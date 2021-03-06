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

	bt.checkpoint, err = checkpoint.NewCheckpoint(
		bt.config.RegistryFile,
		bt.config.CheckpointMaxUpdates,
		bt.config.CheckpointInterval)
	if err != nil {
		return err
	}
	defer bt.checkpoint.Shutdown()

	bt.client, err = b.Publisher.ConnectWith(beat.ClientConfig{
		PublishMode: beat.GuaranteedSend,
		WaitClose:   bt.config.WaitClose,
	})
	if err != nil {
		return err
	}
	defer bt.client.Close()

	er, err := errlog.NewErrlogReader()
	if err != nil {
		return err
	}

	persistedState := bt.checkpoint.States()

	if state, ok := persistedState[errlogfile]; ok {
		err = er.FindSequence(state.RecordNumber)
		if err != nil {
			return err
		}
	}

	trigger := make(chan os.Signal, 1)
	signal.Notify(trigger, syscall.SIGUSR1)

	for {
		select {
		case <-bt.done:
			return nil
		default:
		}

		if events, err := er.Read(); err == nil {
			if len(events) > 0 {
				for _, event := range events {
					bt.client.Publish(event)
					logp.Info("Event sent")
				}
			} else {
				if bt.config.Once {
					logp.Debug("errlogbeat", "No more entries.")
					return nil
				}
				select {
				case <-bt.done:
					return nil
				case <-trigger:
					logp.Debug("errlogbeat", "Awakened by signal.")
				case <-time.After(bt.config.PollingInterval):
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
