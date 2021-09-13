package Qweekly

import "time"

type QWeekly struct {
	Count        int
	ProcessTime  time.Time
	Namem        string
	ProcessId    int
	CylinderType int
}

func (d QWeekly) getProcessId() *int {
	processId := &d.ProcessId
	if d.ProcessId == 0 {
		processId = nil
	}
	return processId
}

func (d QWeekly) getCylinderType() *int {
	cylinderType := &d.CylinderType
	if d.CylinderType == 0 {
		cylinderType = nil
	}
	return cylinderType
}

type LocalRepository interface {
	Save(QWeekly, int) *error
	GetLatestQWeekly(int) (*QWeekly, *error)
}

type RemoteRepository interface {
	FindByTimeInterval(branchIndex int, dtFrom, dtTo time.Time) (*QWeekly, *error)
}
