package beater

import (
	"fmt"
	"time"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"

	"github.com/WuerthIT/errlogbeat/config"
	"github.com/WuerthIT/errlogbeat/errlog"

	"os"
	"os/signal"
	"syscall"

	"github.com/elastic/beats/winlogbeat/checkpoint"
)

type Errlogbeat struct {
	done       chan struct{}
	config     config.Config
	client     beat.Client
	checkpoint *checkpoint.Checkpoint
}

// Creates beater
func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {
	config := config.DefaultConfig
	if err := cfg.Unpack(&config); err != nil {
		return nil, fmt.Errorf("Error reading config file: %v", err)
	}

	bt := &Errlogbeat{
		done:   make(chan struct{}),
		config: config,
	}
	return bt, nil
}

func (bt *Errlogbeat) Run(b *beat.Beat) error {
	logp.Info("errlogbeat is running! Hit CTRL-C to stop it.")

	errlogfile := "errlog"

	var err error
	bt.checkpoint, err = checkpoint.NewCheckpoint(bt.config.RegistryFile, 10, 5*time.Second)
	if err != nil {
		return err
	}
	defer bt.checkpoint.Shutdown()

	err = b.Publisher.SetACKHandler(beat.PipelineACKHandler{
		ACKLastEvents: func(data []interface{}) {
			logp.Debug("errlogbeat", "Call to ACKLastEvents.")
			for _, datum := range data {
				if st, ok := datum.(checkpoint.EventLogState); ok {
					bt.checkpoint.PersistState(st)
				}
			}
		},
	})
	if err != nil {
		return err
	}

	bt.client, err = b.Publisher.ConnectWith(beat.ClientConfig{PublishMode: beat.GuaranteedSend})
	if err != nil {
		return err
	}

	persistedState := bt.checkpoint.States()

	var start_sequence int
	if state, ok := persistedState[errlogfile]; ok {
		start_sequence = int(state.RecordNumber)
	} else {
		start_sequence = -1
	}

	er, err := errlog.NewErrlogReader(start_sequence)
	if err != nil {
		return err
	}

	trigger := make(chan os.Signal, 1)
	signal.Notify(trigger, syscall.SIGUSR1)

	for {
		select {
		case <-bt.done:
			return nil
		default:
		}

		if el, err := er.GetNext(); err == nil {
			if el != nil {
				timestamp := time.Unix(int64(el.Timestamp), 0)
				event := beat.Event{
					Timestamp: timestamp,
					Fields: common.MapStr{
						"system": common.MapStr{
							"errlog": el,
						},
					},
					Private: checkpoint.EventLogState{
						Name:         "errlog",
						RecordNumber: uint64(el.Sequence),
						Timestamp:    timestamp,
					},
				}
				bt.client.Publish(event)
				logp.Info("Event sent")
			} else {
				logp.Debug("errlogbeat", "Going to sleep.")
				select {
				case <-bt.done:
					return nil
				case <-trigger:
					logp.Debug("errlogbeat", "Awakened by signal.")
				}
			}
		} else {
			return err
		}
	}
}

func (bt *Errlogbeat) Stop() {
	bt.client.Close()
	close(bt.done)
}
