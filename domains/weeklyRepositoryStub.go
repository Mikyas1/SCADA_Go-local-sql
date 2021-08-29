package domains

import "time"

type WeeklyRemoteRepositoryStub struct {
	weeklies []Weekly
}

type WeeklyLocalRepositoryStub struct {

}

func (s WeeklyRemoteRepositoryStub) FindByTimeInterval(branchIndex int, dtFrom, dtTo time.Time) (*Weekly, *error) {
	res := Weekly{
		Count: 320,
		FinalDateTime: dtTo,
	}
	return &res, nil
}

func (s WeeklyLocalRepositoryStub) Save(Weekly, branchIndex int) *error {
	return nil
}

func (r WeeklyLocalRepositoryStub) GetLatestWeekly(branchIndex int) (*Weekly, *error) {
	return nil, nil
}

func NewWeeklyRemoteRepositoryStub() WeeklyRemoteRepositoryStub {
	return WeeklyRemoteRepositoryStub{}
}

func NewWeeklyLocalRepositoryStub() WeeklyLocalRepositoryStub {
	return WeeklyLocalRepositoryStub{}
}