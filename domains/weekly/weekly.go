package weekly

import "time"

type Weekly struct {
	Count         int       `json:"count"`
	FinalDateTime time.Time `json:"final_datetime"`
}

type LocalRepository interface {
	Save(weekly Weekly, branchIndex int) *error
	GetLatestWeekly(branchIndex int) (*Weekly, *error)
}

type RemoteRepository interface {
	FindByTimeInterval(branchIndex int, dtFrom, dtTo time.Time) (*Weekly, *error)
}
