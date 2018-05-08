package weeklyutil

import (
	"fmt"
	"strconv"
	"time"

	"github.com/google/go-github/github"
)

func GetWeeklyNumber(issue github.Issue) (int, error) {
	title := *issue.Title
	num, err := strconv.Atoi(title[7:len(title)])
	return num, err
}

func GenerateTimeFromNumber(number int) time.Time {
	secondsEastOfUTC := int((8 * time.Hour).Seconds())
	beijing := time.FixedZone("Beijing Time", secondsEastOfUTC)
	// 74 is 2018.05.07
	dateFor74 := time.Date(2018, time.May, 7, 12, 0, 0, 0, beijing)

	weeksToAdd := number - 74
	dateFor74.AddDate(0, 0, 7*weeksToAdd)
	return dateFor74.AddDate(0, 0, 7*weeksToAdd)
}

func GetFileNameFromTime(date time.Time) string {
	return fmt.Sprintf("%d/%d-%d-%d-weekly.md", date.Year(), date.Year(), int(date.Month()), date.Day())
}

func GetFileNameFromIssue(issue github.Issue) (string, error) {
	num, err := GetWeeklyNumber(issue)
	if err != nil {
		return "", err
	}
	return GetFileNameFromTime(GenerateTimeFromNumber(num)), nil
}