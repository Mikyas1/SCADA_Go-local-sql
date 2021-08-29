package domains

import "time"

type Weekly struct {
	Count         int       `json:"count"`
	FinalDateTime time.Time `json:"final_date_time"`
}

type WeeklyLocalRepository interface {
	Save(weekly Weekly, branchIndex int) *error
	GetLatestWeekly(branchIndex int) (*Weekly, *error)
}

type WeeklyRemoteRepository interface {
	FindByTimeInterval(branchIndex int, dtFrom, dtTo time.Time) (*Weekly, *error)
}
