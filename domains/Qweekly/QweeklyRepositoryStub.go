package Qweekly

import "time"

type QWeeklyRemoteRepositoryStub struct {
	qWeeklies []QWeekly
}

type QWeeklyLocalRepositoryStub struct {

}

func (s QWeeklyRemoteRepositoryStub) FindByTimeInterval(branchIndex int, dtFrom, dtTo time.Time) (*QWeekly, *error) {
	res := QWeekly{
		Count: 320,
		ProcessTime: dtTo,
		Namem: "Leak Gas",
		ProcessId: 3,
		CylinderType: 1,
	}
	return &res, nil
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