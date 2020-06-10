package datetime

import "time"

// GetCheTime gets "zero" offset: 16:00 MSK
func GetCheTime() time.Time {
	now := time.Now().UTC().Add(11 * time.Hour)
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
}

// GetWeekFirstDate returns truncated time of a first week day by given currentDate
func GetWeekFirstDate(currentDate time.Time) time.Time {
	offset := int(currentDate.Weekday()) - 1
	if offset == -1 {
		offset = 6
	}

	return currentDate.
		Truncate(24*time.Hour).
		AddDate(0, 0, -offset)
}

// GetCurrentWeekFirstDate returns truncated time of a first week day by given currentDate
func GetCurrentWeekFirstDate() time.Time {
	return GetWeekFirstDate(time.Now())
}
