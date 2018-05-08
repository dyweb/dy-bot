package weeklyutil

import (
	"strconv"

	"github.com/google/go-github/github"
)

func GetWeeklyNumber(issue github.Issue) (int, error) {
	title := *issue.Title
	num, err := strconv.Atoi(title[7:len(title)])
	return num, err
}
