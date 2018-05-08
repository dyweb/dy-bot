package issue

import "github.com/google/go-github/github"

func (p Processor) processEventClosed(issue github.Issue) error {
	log.Infof("Processing issue#%d closed event.", *issue.Number)
	return nil
}
