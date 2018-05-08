package weekly

import "github.com/google/go-github/github"

func (w Worker) buildWeekly(issue github.Issue) error {
	// cmd := exec.Cmd("weekly-gen", "--repo", fmt.Sprintf("%s/%s", w.config.Owner, w.config.Repo), "--issue", fmt.Sprintf("%d", *issue.Number))
	return nil
}
