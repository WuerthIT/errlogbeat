package errlog

/*
#include <sys/errlog.h>
*/
import "C"

import (
	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/logp"

	"encoding/binary"
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"unsafe"
)

type errDetail struct {
	message string
	length  int
	encode  encoding
}

type errTemplate []errDetail

type errTemplateCache struct {
	cache      map[C.uint]errTemplate
	message_re *regexp.Regexp
	length_re  *regexp.Regexp
	encode_re  *regexp.Regexp
}

func newErrTemplateCache() *errTemplateCache {
	return &errTemplateCache{
		cache:      make(map[C.uint]errTemplate),
		message_re: regexp.MustCompile("(?s)\\nDetail Data\\n(.*)\\n"),
		length_re:  regexp.MustCompile("(?m)^et_detail_length\\s+((?:0x[[:xdigit:]]{4},)*(?:0x[[:xdigit:]]{4})?)"),
		encode_re:  regexp.MustCompile("(?m)^et_detail_encode\\s+((?:[ADHIL],)*(?:[ADHIL])?)"),
	}
}

const (
	enc_ALPHA encoding = iota
	enc_DEC
	enc_HEX
)

type encoding uint

func (c *errTemplateCache) lookup(entry *C.errlog_entry_t) (errTemplate, error) {
	if t, ok := c.cache[entry.el_crcid]; ok {
		return t, nil
	}

	var match []string

	errpt_t, err := exec.Command("/usr/bin/errpt", "-t", "-A", "-j", strconv.FormatUint(uint64(entry.el_crcid), 16)).Output()
	if err != nil {
		return nil, err
	}

	out, err := exec.Command("/usr/bin/errpt", "-g", "-l", strconv.FormatUint(uint64(entry.el_sequence), 10)).Output()
	if err != nil {
		return nil, err
	}
	errpt_g := string(out[:])

	c.cache[entry.el_crcid] = nil

	if match = c.message_re.FindStringSubmatch(string(errpt_t[:])); len(match) != 2 {
		logp.Debug("errtemplatecache", "Detail Data not found")
		return nil, nil
	}
	messages := strings.Split(match[1], "\n")

	if match = c.length_re.FindStringSubmatch(errpt_g); len(match) != 2 {
		logp.Warn("et_detail_length not found")
		return nil, nil
	}
	lengths := strings.Split(match[1], ",")
	if len(messages) > len(lengths) {
		logp.Warn("Number of messages (%d) greater than number of lenghts (%d)", len(messages), len(lengths))
		return nil, nil
	}

	if match = c.encode_re.FindStringSubmatch(errpt_g); len(match) != 2 {
		logp.Warn("et_detail_encode not found")
		return nil, nil
	}
	encodes := strings.Split(match[1], ",")
	if len(messages) > len(encodes) {
		logp.Warn("Number of messages (%d) greater than number of encodings (%d)", len(messages), len(encodes))
		return nil, nil
	}

	d := make([]errDetail, len(messages))

	for n, m := range messages {
		d[n] = errDetail{
			message: m,
		}
		if l, err := strconv.ParseUint(lengths[n], 0, 16); err == nil {
			d[n].length = int(l)
		} else {
			return nil, err
		}
		switch encodes[n] {
		case "A":
			d[n].encode = enc_ALPHA
		case "D", "L":
			d[n].encode = enc_DEC
		case "H":
			d[n].encode = enc_HEX
		case "I":
			d[n].encode = enc_HEX
			d[n].length = d[n].length * 4
		default:
			return nil, fmt.Errorf("Invalid encoding %s", encodes[n])
		}
	}
	c.cache[entry.el_crcid] = d
	return d, nil
}

func encode(d []byte, e encoding) string {
	var s string

	switch e {
	case enc_ALPHA:
		s = strings.TrimSpace(strings.SplitN(string(d[:]), "\x00", 2)[0])
	case enc_DEC:
		b := make([]byte, 8)
		copy(b[8-len(d):], d)
		s = strconv.FormatUint(binary.BigEndian.Uint64(b), 10)
	case enc_HEX:
		for len(d) >= 2 {
			s += fmt.Sprintf("%04X ", binary.BigEndian.Uint16(d[:]))
			d = d[2:]
		}
		if len(d) == 1 {
			s += fmt.Sprintf("%02X ", uint8(d[0]))
		}
		s = strings.TrimSpace(s)
	}
	return s
}

func (c *errTemplateCache) entriesToEvents(entries []*C.errlog_entry_t) ([]beat.Event, error) {
	events := make([]beat.Event, 0, len(entries))

	for _, entry := range entries {
		template, err := c.lookup(entry)
		if err != nil {
			return nil, err
		}
		details := make([][2]string, 0, len(template))
		data := C.GoBytes(unsafe.Pointer(&entry.el_detail_data), C.int(entry.el_detail_length))
		for _, detail := range template {
			if len(data) < detail.length {
				detail.length = len(data)
			}
			details = append(details, [2]string{detail.message, encode(data[:detail.length], detail.encode)})
			data = data[detail.length:]
		}
		events = append(events, entry.toEvent(details...))
	}
	return events, nil
}
