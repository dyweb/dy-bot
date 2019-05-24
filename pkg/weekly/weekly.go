package weekly

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"path/filepath"

	"github.com/google/go-github/github"

	"github.com/dyweb/dy-bot/pkg/util/weeklyutil"
)

func (w Worker) buildWeekly(issue github.Issue) error {
	fn, err := weeklyutil.GetFileNameFromIssue(issue)
	if err != nil {
		return fmt.Errorf("failed to generate weekly: %v", err)
	}
	cmd := exec.Command("weekly-gen", "--repo", fmt.Sprintf("%s/%s", w.config.Owner, w.config.Repo), "--issue", fmt.Sprintf("%d", *issue.Number))
	log.Info(cmd.Args)
	cmd.Dir = w.config.WeeklyDir
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to generate weekly: %s, %v", string(output), err)
	}
	err = ioutil.WriteFile(filepath.Join(w.config.WeeklyDir, fn), output, 0644)
	log.Infof("Output to %s: %s", filepath.Join(w.config.WeeklyDir, fn), string(output))
	if err != nil {
		return fmt.Errorf("failed to generate weekly: %v", err)
	}

	cmd = exec.Command("bash", "./scripts/build.sh")
	cmd.Dir = w.config.WeeklyDir
	output, err = cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to build weekly static page: %s, %v", string(output), err)
	}
	return nil
}
