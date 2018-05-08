package issue

import (
	"github.com/google/go-github/github"

	"github.com/dyweb/dy-bot/pkg/gh"
)

const (
	labelWorking = "working"
)

func (p Processor) processEventClosed(issue github.Issue) error {
	log.Infof("Processing issue#%d closed event.", *issue.Number)
	for _, label := range issue.Labels {
		if *label.Name == labelWorking {
			log.Info("The working issue is closed.")
			if err := removeWorkingLabel(issue); err != nil {
				return err
			}
		}
	}
	return nil
}

func removeWorkingLabel(issue github.Issue) error {
	gc := gh.GetGitHubClient()
	_, err := gc.Client.Issues.RemoveLabelForIssue(gc.Owner(), gc.Repo(), *issue.Number, labelWorking)
	if err != nil {
		return err
	}
	return nil
}
