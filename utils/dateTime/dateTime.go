package dateTime

import (
	"time"
)

const (
	Layout1 string = "2006-01-02 15:04:05"
	Layout2 string = "2006-01-02 15:04"
	Layout3 string = "2006-01-02"
)

func ParseDateTimeFromString(layout, str string) (*time.Time, *error) {
	t, err := time.Parse(layout, str)
	if err != nil {
		return nil, &err
	}
	return &t, nil
}

func removeSecondsFromDateTime(dt time.Time) (*time.Time, *error) {
	temp := dt.Format(Layout2)
	resDt, err := ParseDateTimeFromString(Layout2, temp)
	if err != nil {
		return nil, err
	}
	return resDt, nil
}

func GetOnlyDateFormDateTime(dt time.Time) (*time.Time, *error) {
	temp := dt.Format(Layout3)
	resDt, err := ParseDateTimeFromString(Layout3, temp)
	if err != nil {
		return nil, err
	}
	return resDt, nil
}

func getMin(dt time.Time) int {
	_, min, _ := dt.Clock()
	return min
}

// ChangeDateTimeMinToFactor return a new datetime that has its minute set to
// the factor of the factor param.
func changeDateTimeMinToFactor(dt *time.Time, factor int, up bool) (*time.Time, *error) {
	min := getMin(*dt)
	if min % factor == 0 {
		// base case
		return dt, nil
	}
	if up {
		*dt = dt.Add(time.Minute)
		return changeDateTimeMinToFactor(dt, factor, up)
	}

	*dt = dt.Add(-time.Minute)
	return changeDateTimeMinToFactor(dt, factor, up)
}

func ChangeDateTimeMinToFactorWrapper(dt *time.Time, factor int, up bool) (*time.Time, *error) {
	dt, err := removeSecondsFromDateTime(*dt)
	if err != nil {
		return nil, err
	}
	return changeDateTimeMinToFactor(dt, factor, up)
}

func GetYesterday() (*time.Time, *error) {
	//now := time.Now().AddDate(0, 0, -1)
	now := time.Now()
	dt, err := GetOnlyDateFormDateTime(now)
	if err != nil {
		return nil, err
	}
	return dt, nil
}