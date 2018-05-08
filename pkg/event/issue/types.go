package issue

import (
	"github.com/dyweb/gommon/util/logutil"

	"github.com/dyweb/dy-bot/pkg/util/githubutil"
)

var log = logutil.NewPackageLogger()

type Processor struct {
}

func (p Processor) Process(data []byte) error {
	// process details
	actionType, err := githubutil.ExtractActionType(data)
	if err != nil {
		return err
	}

	log.Infof("received event type [issues], action type [%s]", actionType)

	issue, err := githubutil.ExactIssue(data)
	if err != nil {
		return err
	}

	switch actionType {
	case "closed":
		if err := p.processEventClosed(issue); err != nil {
			return err
		}
	default:
		log.Errorf("action type %s is not supported yet", actionType)
	}
	return nil
}
