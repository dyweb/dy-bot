package weekly

import (
	"fmt"
	"os/exec"

	"github.com/google/go-github/github"

	"github.com/dyweb/dy-bot/pkg/util/weeklyutil"
)

func (w Worker) buildWeekly(issue github.Issue) error {
	fn, err := weeklyutil.GetFileNameFromIssue(issue)
	if err != nil {
		return fmt.Errorf("failed to generate weekly: %v", err)
	}

	cmdInBash := fmt.Sprintf("weekly-gen --repo %s/%s --issue %d > %s", w.config.Owner, w.config.Repo, *issue.Number, fn)

	cmd := exec.Command("bash", "-c", cmdInBash)
	log.Info(cmd.Args)
	cmd.Dir = w.config.WeeklyDir
	if err = cmd.Run(); err != nil {
		return fmt.Errorf("failed to generate weekly: %v", err)
	}
	return nil
}
