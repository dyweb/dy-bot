package issue

import (
	"github.com/dyweb/dy-bot/pkg/weekly"
	"github.com/google/go-github/github"
)

const (
	labelWorking = "working"
)

func (p Processor) processEventClosed(issue github.Issue) error {
	log.Infof("Processing issue#%d closed event.", *issue.Number)
	for _, label := range issue.Labels {
		if *label.Name == labelWorking {
			log.Info("The working issue is closed.")
			worker := weekly.NewWorker(p.config)
			if err := worker.HandleWeekly(issue); err != nil {
				return err
			}
		}
	}
	return nil
}
