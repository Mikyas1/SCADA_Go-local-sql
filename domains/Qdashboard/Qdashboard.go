package Qdashboard

import "time"

type QDashboard struct {
	Residual      int       `json:"residual"`
	CheckNet      int       `json:"checkNet"`
	Count         int       `json:"count"`
	FinalDateTime time.Time `json:"final_datetime"`
}

type LocalRepository interface {
	Save(data QDashboard, branchIndex int) *error
	GetLatestWeekly(branchIndex int) (*QDashboard, *error)
}

type RemoteRepository interface {
	FindByTimeInterval(branchIndex int, dtFrom, dtTo time.Time) (*QDashboard, *error)
}
