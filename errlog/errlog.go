package errlog

/*
#cgo LDFLAGS: -lerrlog
#include <fcntl.h>
#include <sys/errlog.h>
*/
import "C"

import (
	"fmt"
	"github.com/elastic/beats/libbeat/logp"
	"os"
	"os/exec"
	"os/signal"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"syscall"
)

type ErrlogReader struct {
	handle C.errlog_handle_t
	cache  *ErrTemplateCache
}

type ErrlogEntry struct {
	Magic     uint
	Sequence  uint
	Label     string
	Timestamp uint
	Crcid     uint
	Errdiag   uint
	Machineid string
	Nodeid    string
	Class     string
	Type      string
	Resource  string
	Rclass    string
	Rtype     string
	Vpd_ibm   string
	Vpd_user  string
	In        string
	Connwhere string
	Flags     uint
	Detail    string
	Symptom   string
}

type ErrTemplate struct {
	detail_message []string
	detail_length  []string
	detail_encode  []string
}

type ErrTemplateCache struct {
	cache             map[C.uint]ErrTemplate
	detail_message_re *regexp.Regexp
	detail_length_re  *regexp.Regexp
	detail_encode_re  *regexp.Regexp
}

func ErrlogSender(seq int, c chan *ErrlogEntry) {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	er, _ := NewErrlogReader(seq)

	trigger := make(chan os.Signal, 1)
	signal.Notify(trigger, syscall.SIGUSR1)

	for {
		if el, err := er.GetNext(); err == nil {
			if el != nil {
				c <- el
			} else {
				logp.Debug("errlogsender", "Going to sleep.")
				select {
				case <-trigger:
					logp.Debug("errlogsender", "Awakened by signal.")
				}
			}
		} else {
			close(c)
			return
		}
	}
}

func NewErrlogReader(seq int) (*ErrlogReader, error) {
	var handle C.errlog_handle_t
	var entry C.errlog_entry_t
	var rc C.int

	rc = C.errlog_open(nil, C.O_RDONLY, C.LE_MAGIC, &handle)
	if rc != 0 {
		return nil, fmt.Errorf("call to errlog_open failed with return code %d", rc)
	}

	rc = C.errlog_set_direction(handle, C.LE_FORWARD)
	if rc != 0 {
		return nil, fmt.Errorf("call to errlog_set_direction failed with return code %d", rc)
	}

	if seq != -1 {
		logp.Debug("errlogreader", "Searching for sequence %d.", seq)
		rc = C.errlog_find_sequence(handle, C.int(seq), &entry)
		if rc == C.LE_ERR_DONE {
			logp.Warn("Recorded sequence %d could not be found", seq)
			rc = C.errlog_set_direction(handle, C.LE_FORWARD)
			if rc != 0 {
				return nil, fmt.Errorf("call to errlog_set_direction failed with return code %d", rc)
			}
		} else if rc != 0 {
			return nil, fmt.Errorf("call to errlog_find_sequence failed with return code %d", rc)
		}
	}

	cache := NewErrTemplateCache()

	return &ErrlogReader{
		handle: handle,
		cache:  cache,
	}, nil
}

func (r *ErrlogReader) GetNext() (*ErrlogEntry, error) {
	var entry C.errlog_entry_t
	var rc C.int

	logp.Debug("errlogreader", "calling errlog_find_next: %v", r.handle)
	rc = C.errlog_find_next(r.handle, &entry)
	logp.Debug("errlogreader", "errlog_find_next returned %d", rc)
	if rc == C.LE_ERR_DONE {
		return nil, nil
	} else if rc == 0 {
		sequence := uint(entry.el_sequence)

		r.cache.Lookup(&entry)

		return &ErrlogEntry{
			Magic:     uint(entry.el_magic),
			Sequence:  sequence,
			Label:     C.GoString(&entry.el_label[0]),
			Timestamp: uint(entry.el_timestamp),
			Crcid:     uint(entry.el_crcid),
			Errdiag:   uint(entry.el_errdiag),
			Machineid: C.GoString(&entry.el_machineid[0]),
			Nodeid:    C.GoString(&entry.el_nodeid[0]),
			Class:     C.GoString(&entry.el_class[0]),
			Type:      C.GoString(&entry.el_type[0]),
			Resource:  C.GoString(&entry.el_resource[0]),
			Rclass:    C.GoString(&entry.el_rclass[0]),
			Rtype:     C.GoString(&entry.el_rtype[0]),
			Vpd_ibm:   C.GoString(&entry.el_vpd_ibm[0]),
			Vpd_user:  C.GoString(&entry.el_vpd_user[0]),
			In:        C.GoString(&entry.el_in[0]),
			Connwhere: C.GoString(&entry.el_connwhere[0]),
			Flags:     uint(entry.el_flags),
			Detail:    C.GoStringN(&entry.el_detail_data[0], C.int(entry.el_detail_length)),
			Symptom:   C.GoStringN(&entry.el_symptom_data[0], C.int(entry.el_symptom_length)),
		}, nil
	} else {
		return nil, fmt.Errorf("call to errlog_find_next failed with return code %d", rc)
	}
}

func NewErrTemplateCache() *ErrTemplateCache {
	return &ErrTemplateCache{
		cache:             make(map[C.uint]ErrTemplate),
		detail_message_re: regexp.MustCompile("(?s)\\nDetail Data\\n(.*)\\n"),
		detail_length_re:  regexp.MustCompile("(?m)^et_detail_length\\s+((?:0x[[:xdigit:]]{4},)*(?:0x[[:xdigit:]]{4})?)"),
		detail_encode_re:  regexp.MustCompile("(?m)^et_detail_encode\\s+((?:[ADHILX],)*(?:[ADHILX])?)"),
	}
}

func (c *ErrTemplateCache) add_template(entry *C.errlog_entry_t) error {
	out, err := exec.Command("/usr/bin/errpt", "-t", "-A", "-j", strconv.FormatUint(uint64(entry.el_crcid), 16)).Output()
	if err != nil {
		return err
	}
	var match []string

	var messages []string
	var lengths []string
	var encodes []string

	if match = c.detail_message_re.FindStringSubmatch(string(out[:])); len(match) == 2 {
		messages = strings.Split(match[1], "\n")
	}
	logp.Debug("errtemplatecache", "errpt output: %q", messages)

	out, err = exec.Command("/usr/bin/errpt", "-g", "-l", strconv.FormatUint(uint64(entry.el_sequence), 10)).Output()
	if err != nil {
		return err
	}
	out_str := string(out[:])
	if match = c.detail_length_re.FindStringSubmatch(string(out_str)); len(match) == 2 {
		lengths = strings.Split(match[1], ",")
	}
	logp.Debug("errtemplatecache", "errpt output: %q", lengths)

	if match = c.detail_encode_re.FindStringSubmatch(string(out_str)); len(match) == 2 {
		encodes = strings.Split(match[1], ",")
	}
	logp.Debug("errtemplatecache", "errpt output: %q", encodes)
	c.cache[entry.el_crcid] = ErrTemplate{
		detail_message: messages,
		detail_length:  lengths,
		detail_encode:  encodes,
	}
	return nil
}

func (c *ErrTemplateCache) Lookup(entry *C.errlog_entry_t) error {
	if _, ok := c.cache[entry.el_crcid]; !ok {
		c.add_template(entry)
		logp.Debug("errtemplatecache", "id %d missing in cache", entry.el_crcid)
	} else {
		logp.Debug("errtemplatecache", "id %d found in cache", entry.el_crcid)
	}
	return nil
}
