package event

import (
	"fmt"

	"github.com/dyweb/dy-bot/pkg/event/issue"
	"github.com/dyweb/gommon/util/logutil"
)

var log = logutil.NewPackageLogger()

// Processor is the type for processor.
type Processor interface {
	// Process processes item automan gets, and then execute operations torwards items on GitHub
	Process(data []byte) error
}

// Manager contains several specific processors.
type Manager struct {
	issueProcessor issue.Processor
}

func NewManager() *Manager {
	return &Manager{}
}

// HandleEvent processes an event received from github.
func (m Manager) HandleEvent(eventType string, data []byte) error {
	switch eventType {
	case "ping":
		log.Info("Got ping from GitHub!")
	case "issues":
		log.Info("Got issue events.")
		if err := m.issueProcessor.Process(data); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unknown event type %s", eventType)
	}
	return nil
}
