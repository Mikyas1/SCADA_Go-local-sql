package Qdashboard

import "time"

type QDashboardRemoteRepositoryStub struct {
	weeklies []QDashboard
}

type QDashboardLocalRepositoryStub struct {

}

func (s QDashboardRemoteRepositoryStub) FindByTimeInterval(branchIndex int, dtFrom, dtTo time.Time) (*QDashboard, *error) {
	res := QDashboard{
		Residual: 34,
		CheckNet: 43,
		Count: 320,
		FinalDateTime: dtTo,
	}
	return &res, nil
}

func (s QDashboardLocalRepositoryStub) Save(q QDashboard, branchIndex int) *error {
	return nil
}

func (r QDashboardLocalRepositoryStub) GetLatestWeekly(branchIndex int) (*QDashboard, *error) {
	return nil, nil
}

func NewQDashboardRemoteRepositoryStub() QDashboardRemoteRepositoryStub {
	return QDashboardRemoteRepositoryStub{}
}

func NewWDashboardLocalRepositoryStub() QDashboardLocalRepositoryStub {
	return QDashboardLocalRepositoryStub{}
}