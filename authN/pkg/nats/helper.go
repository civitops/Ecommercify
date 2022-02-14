package natshelper

import (
	"fmt"
	"sync"
	"time"

	"github.com/civitops/Ecommercify/authN/pkg/config"

	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
)

func SetupConnOptions(log *zap.SugaredLogger, wg *sync.WaitGroup) []nats.Option {
	opts := make([]nats.Option, 0)
	// Buffering Messages During Reconnect Attempts
	opts = append(opts, nats.ReconnectBufSize(5*1024*1024))
	// Set reconnect interval
	opts = append(opts, nats.ReconnectWait(config.NatsReconnectDelay))
	// Set max reconnects attempts
	opts = append(opts, nats.MaxReconnects(int(config.NatsReconnectTotalWait/config.NatsReconnectDelay)))

	opts = append(opts, nats.ReconnectHandler(func(nc *nats.Conn) {
		log.Infof("Reconnected [%s]", nc.ConnectedUrl())
	}))

	opts = append(opts, nats.ClosedHandler(func(nc *nats.Conn) {
		// done when nats is closed
		wg.Done()
		log.Infof("Exiting: %v", nc.LastError())
	}))

	return opts
}

func CreateStream(js nats.JetStreamContext, log *zap.SugaredLogger) (err error) {
	stream, _ := js.StreamInfo(config.UserStreamName)
	if stream == nil {
		subj := fmt.Sprintf("%s.*", config.UserStreamName)
		log.Debugf("creating stream %q and subjects %q", config.UserStreamName, subj)

		if _, err = js.AddStream(&nats.StreamConfig{
			Name:        config.UserStreamName,
			Description: "notification stream",
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
