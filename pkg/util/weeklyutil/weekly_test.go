package weeklyutil

import (
	"testing"
	"time"
)

func TestGenerateTimeFromNumber(t *testing.T) {
	secondsEastOfUTC := int((8 * time.Hour).Seconds())
	beijing := time.FixedZone("Beijing Time", secondsEastOfUTC)
	testCases := []struct {
		Case int
		Date time.Time
	}{
		{
			Case: 124,
			Date: time.Date(2019, time.May, 27, 12, 0, 0, 0, beijing),
		},
		{
			Case: 125,
			Date: time.Date(2019, time.May, 27, 12, 0, 0, 0, beijing).AddDate(0, 0, 7),
		},
	}

	for _, tc := range testCases {
		if !GenerateTimeFromNumber(tc.Case).Equal(tc.Date) {
			t.Errorf("Expected %s, got %s", tc.Date.String(), GenerateTimeFromNumber(tc.Case).String())
		}
	}
}
