package errlog

/*
#cgo LDFLAGS: -lerrlog
#include <fcntl.h>
#include <sys/errlog.h>
*/
import "C"

import (
	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/elastic/beats/winlogbeat/checkpoint"

	"fmt"
	"strings"
	"time"
)

type ErrlogReader struct {
	handle C.errlog_handle_t
	cache  *errTemplateCache
}

func (entry *C.errlog_entry_t) toEvent(detail_data ...[2]string) beat.Event {
	timestamp := time.Unix(int64(entry.el_timestamp), 0)
	errlog := common.MapStr{
		"magic":     uint(entry.el_magic),
		"sequence":  uint(entry.el_sequence),
		"label":     strings.TrimSpace(C.GoString(&entry.el_label[0])),
		"crcid":     fmt.Sprintf("%08X", uint(entry.el_crcid)),
		"errdiag":   uint(entry.el_errdiag),
		"machineid": strings.TrimSpace(C.GoString(&entry.el_machineid[0])),
		"nodeid":    strings.TrimSpace(C.GoString(&entry.el_nodeid[0])),
		"class":     strings.TrimSpace(C.GoString(&entry.el_class[0])),
		"type":      strings.TrimSpace(C.GoString(&entry.el_type[0])),
		"resource":  strings.TrimSpace(C.GoString(&entry.el_resource[0])),
		"rclass":    strings.TrimSpace(C.GoString(&entry.el_rclass[0])),
		"rtype":     strings.TrimSpace(C.GoString(&entry.el_rtype[0])),
		"vpd_ibm":   strings.TrimSpace(C.GoString(&entry.el_vpd_ibm[0])),
		"vpd_user":  strings.TrimSpace(C.GoString(&entry.el_vpd_user[0])),
		"in":        strings.TrimSpace(C.GoString(&entry.el_in[0])),
		"connwhere": strings.TrimSpace(C.GoString(&entry.el_connwhere[0])),
		"flags": common.MapStr{
			"err64":   entry.el_flags&C.LE_FLAG_ERR64 != 0,
			"errdup":  entry.el_flags&C.LE_FLAG_ERRDUP != 0,
			"errwpar": entry.el_flags&C.LE_FLAG_ERRWPAR != 0,
		},
		"wparid": strings.TrimSpace(C.GoString(&entry.el_wparid[0])),
	}

	if entry.el_flags&C.LE_FLAG_ERRDUP != 0 {
		errlog["errdup"] = common.MapStr{
			"dupcount": uint(entry.el_errdup.ed_dupcount),
			"time1":    time.Unix(int64(entry.el_errdup.ed_time1), 0),
			"time2":    time.Unix(int64(entry.el_errdup.ed_time2), 0),
		}
	}

	if len(detail_data) > 0 {
		detail := common.MapStr{}
		for _, d := range detail_data {
			detail[d[0]] = d[1]
		}
		errlog["detail"] = detail
	}

	return beat.Event{
		Timestamp: timestamp,
		Fields: common.MapStr{
			"errlog": errlog,
		},
		Private: checkpoint.EventLogState{
			Name:         "errlog",
			RecordNumber: uint64(entry.el_sequence),
			Timestamp:    timestamp,
		},
	}
}

func NewErrlogReader() (*ErrlogReader, error) {
	var handle C.errlog_handle_t
	var rc C.int

	rc = C.errlog_open(nil, C.O_RDONLY, C.LE_MAGIC, &handle)
	if rc != 0 {
		return nil, fmt.Errorf("call to errlog_open failed with return code %d", rc)
	}

	rc = C.errlog_set_direction(handle, C.LE_FORWARD)
	if rc != 0 {
		return nil, fmt.Errorf("call to errlog_set_direction failed with return code %d", rc)
	}

	cache := newErrTemplateCache()

	return &ErrlogReader{
		handle: handle,
		cache:  cache,
	}, nil
}

func (r *ErrlogReader) FindSequence(s uint64) error {
	var entry C.errlog_entry_t

	rc := C.errlog_find_sequence(r.handle, C.int(s), &entry)
	if rc == C.LE_ERR_DONE {
		logp.Warn("Recorded sequence %d could not be found", s)
		rc = C.errlog_set_direction(r.handle, C.LE_FORWARD)
		if rc != 0 {
			return fmt.Errorf("call to errlog_set_direction failed with return code %d", rc)
		}
	} else if rc != 0 {
		return fmt.Errorf("call to errlog_find_sequence failed with return code %d", rc)
	}
	return nil
}

func (r *ErrlogReader) Read() ([]beat.Event, error) {
	entries := make([]*C.errlog_entry_t, 0, 1)
	var entry C.errlog_entry_t

	rc := C.errlog_find_next(r.handle, &entry)
	if rc == C.LE_ERR_DONE {
		return nil, nil
	} else if rc == 0 {
		entries = append(entries, &entry)
		return r.cache.entriesToEvents(entries)
	} else {
		return nil, fmt.Errorf("call to errlog_find_next failed with return code %d", rc)
	}
}
