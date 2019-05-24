package weekly

import (
	"context"
	"fmt"
	"os/exec"
	"strings"
	"time"

	"github.com/dyweb/dy-bot/pkg/gh"
	"github.com/google/go-github/github"
)

// ErrNothingChanged is used when git commit has nothing to commit.
var ErrNothingChanged = fmt.Errorf("nothing to commit")

func generateNewBranch() string {
	timeStr := time.Now().String()
	dateStrSlice := strings.SplitN(timeStr, " ", 2)
	dateStr := dateStrSlice[0]

	return fmt.Sprintf("weekly-%s-%s", dateStr, RandStringRunes(5))
}

func (w Worker) prepareGitEnv(newBranch string) error {
	// sync latest master branch and checkout new branch

	// checkout local master branch
	cmd := exec.Command("git", "checkout", "master")
	cmd.Dir = w.config.WeeklyDir
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to checkout master: %v", err)
	}

	// fetch upstream master to local
	cmd = exec.Command("git", "fetch", "upstream", "master")
	cmd.Dir = w.config.WeeklyDir
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to git fetch upstream master: %v", err)
	}

	// rebase local master on origin/master
	cmd = exec.Command("git", "rebase", "upstream/master")
	cmd.Dir = w.config.WeeklyDir
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to git rebase upstream/master: %v", err)
	}

	// push local master to origin/master
	cmd = exec.Command("git", "push", "-f", "origin", "master")
	cmd.Dir = w.config.WeeklyDir
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to git push -f origin master: %v", err)
	}

	// create a new branch named by input newBranch
	// the following doc generation are all on this new branch
	cmd = exec.Command("git", "checkout", "-b", newBranch)
	cmd.Dir = w.config.WeeklyDir
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to git checkout -b %s: %v", newBranch, err)
	}
	log.Infof("New branch: %s", newBranch)
	return nil
}

func (w Worker) gitCommitAndPush(newBranch string) error {
	// git add all updated files.
	cmd := exec.Command("git", "add", ".")
	cmd.Dir = w.config.WeeklyDir
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to git add .: %v", err)
	}

	// check whether nothing changed.
	cmd = exec.Command("git", "status")
	cmd.Dir = w.config.WeeklyDir
	out, err := cmd.Output()
	if err != nil {
		return err
	}

	// if nothing changes, return nil to quit git procedure.
	if strings.Contains(string(out), "nothing to commit") {
		log.Infof("no changes happened, quit git procedure")
		return ErrNothingChanged
	}

	// git commit all the staged files.
	cmd = exec.Command("git", "commit", "-s", "-m", "Weekly: Add")
	cmd.Dir = w.config.WeeklyDir
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to git commit -s -m : %v", err)
	}

	// git push forcely to origin repo.
	cmd = exec.Command("git", "push", "-f", "origin", newBranch)
	cmd.Dir = w.config.WeeklyDir
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to git push -f origin %s: %v", newBranch, err)
	}

	// git branch -D to delete branch to free resources.
	cmd = exec.Command("git", "checkout", "master")
	cmd.Dir = w.config.WeeklyDir
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to git checkout master before deleting branch %s: %v", newBranch, err)
	}

	// git branch -D to delete branch to free resources.
	cmd = exec.Command("git", "branch", "-D", newBranch)
	cmd.Dir = w.config.WeeklyDir
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to git push branch -D %s: %v", newBranch, err)
	}
	log.Infof("Succeeded to push the change")
	return nil
}

func (w Worker) sumbitPR(branch string, issueNumber int) error {
	title := fmt.Sprintf("Weekly: Add %d", issueNumber)
	head := fmt.Sprintf("gaocegege-bot:%s", branch)
	base := "master"
	body := fmt.Sprintf(`weekly: Generate

gaocegege-bot powered by github.com/dyweb/dy-bot

Ref https://github.com/%s/%s/issues/%d`, w.config.Owner, w.config.Repo, issueNumber)

	newPR := &github.NewPullRequest{
		Title: &title,
		Head:  &head,
		Base:  &base,
		Body:  &body,
	}
	log.Infof("PR: %v", newPR)

	gc := gh.GetGitHubClient()
	ctx := context.Background()
	if _, _, err := gc.PullRequests.Create(ctx, gc.Owner(), gc.Repo(), newPR); err != nil {
		log.Errorf("failed to create pull request: %v", err)
		return err
	}

	return nil
}
