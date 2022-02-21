package natshelper

import (
	"fmt"
	"sync"
	"time"

	"github.com/civitops/Ecommercify/user/pkg/config"

	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
)

func SetupConnOptions(log *zap.SugaredLogger, wg *sync.WaitGroup) []nats.Option {
	opts := make([]nats.Option, 0)
	// Buffering Messages During Reconnect Attempts
	opts = append(opts, nats.ReconnectBufSize(5*1024*1024),
		nats.ReconnectWait(config.NatsReconnectDelay),                                    //Set Reconnect Interval
		nats.MaxReconnects(int(config.NatsReconnectTotalWait/config.NatsReconnectDelay)), // Set max reconnects attempts
		nats.ReconnectHandler(func(nc *nats.Conn) {
			log.Infof("Reconnected [%s]", nc.ConnectedUrl())
		}),
		nats.ClosedHandler(func(nc *nats.Conn) {
			// done when nats is closed
			wg.Done()
			log.Infof("Exiting: %v", nc.LastError())
		}),
	)

	return opts
}

func CreateStream(js nats.JetStreamContext, log *zap.SugaredLogger) (err error) {
	stream, _ := js.StreamInfo(config.StreamName)
	if stream == nil {
		subj := fmt.Sprintf("%s.*", config.StreamName)
		log.Debugf("creating stream %q and subjects %q", config.StreamName, subj)

		if _, err = js.AddStream(&nats.StreamConfig{
			Name:        config.StreamName,
			Description: "user stream",
			Subjects:    []string{subj},
			Retention:   nats.WorkQueuePolicy,
			Discard:     nats.DiscardOld,
			MaxAge:      24 * time.Hour,
			Storage:     nats.FileStorage,
		}); err != nil {
			return
		}
	}

	return nil
}
