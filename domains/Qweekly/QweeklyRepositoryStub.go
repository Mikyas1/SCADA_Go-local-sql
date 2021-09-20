package Qweekly

import "time"

type QWeeklyRemoteRepositoryStub struct {
	qWeeklies []QWeekly
}

type QWeeklyLocalRepositoryStub struct {

}

func (s QWeeklyRemoteRepositoryStub) FindByTimeInterval(branchIndex int, dtFrom, dtTo time.Time) ([]*QWeekly, *error) {
	qWeeklies := []*QWeekly{
		{
			Count: 320,
			ProcessTime: dtTo,
			Namem: "Leak Gas",
			ProcessId: 3,
			CylinderType: 1,
		},
		{
			Count: 220,
			ProcessTime: dtTo,
			Namem: "Leak Gas",
			ProcessId: 3,
			CylinderType: 2,
		},
		{
			Count: 120,
			ProcessTime: dtTo,
			Namem: "Leak Pressure",
			ProcessId: 4,
			CylinderType: 1,
		},{
			Count: 320,
			ProcessTime: dtTo,
			Namem: "Leak Pressure",
			ProcessId: 4,
			CylinderType: 2,
		},
	}
	return qWeeklies, nil
}

func (s QWeeklyLocalRepositoryStub) Save(qWeekly QWeekly, branchIndex int) *error {
	return nil
}

func (r QWeeklyLocalRepositoryStub) GetLatestQWeekly(branchIndex int) (*QWeekly, *error) {
	return nil, nil
}

func NewQWeeklyRemoteRepositoryStub() QWeeklyRemoteRepositoryStub {
	return QWeeklyRemoteRepositoryStub{}
}

func NewQWeeklyLocalRepositoryStub() QWeeklyLocalRepositoryStub {
	return QWeeklyLocalRepositoryStub{}
}