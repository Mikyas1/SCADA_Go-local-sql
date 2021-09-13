package Qweekly

import "time"

type QWeekly struct {
	Count        int
	ProcessTime  time.Time
	Namem        string
	ProcessId    int
	CylinderType int
}

type LocalRepository interface {
	Save(QWeekly, int) *error
	GetLatestQWeekly(int) (*QWeekly, *error)
}

type RemoteRepository interface {
	FindByTimeInterval(branchIndex int, dtFrom, dtTo time.Time) (*QWeekly, *error)
}
