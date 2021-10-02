package Qreport

import "time"

type QReportRemoteRepositoryStub struct {

}

type QReportLocalRepositoryStub struct {

}

func (s QReportRemoteRepositoryStub) FindByTimeInterval(branchIndex int, dtFrom, dtTo time.Time) (*QReport, *error) {
	res := QReport{
		MachineId: "12",
		ProcessDate: "somedate",
		Gtem400: 12,
		Gtem350: 11,
		Gtem300: 11,
		Gtem250: 11,
		Gtem200: 11,
		Gtem150: 11,
		Gtem100: 11,
		Gtem050: 11,
		Value: 45,
		Gte050: 4,
		Gte100: 4,
		Gte150: 4,
		Gte200: 4,
		Gte250: 4,
		Gte300: 4,
		Gte350: 4,
		Gte400: 4,
		Sum: 300,
		M200X200: 98,
		Diff: 43,
		StartPoint: 2,
		Accuracy: 4,
		CylinderType: 1,
		FinalDateTime: dtTo,
	}
	return &res, nil
}

func (s QReportRemoteRepositoryStub) GetMachineId(value, processId int, dtFrom, dtTo time.Time) ([]string, *error) {
	return []string{"1", "2", "3", "4"}, nil
}

func (s QReportLocalRepositoryStub) Save(qr QReport, branchIndex int) *error {
	return nil
}

func (r QReportLocalRepositoryStub) GetLatestQReport(branchIndex int) (*QReport, *error) {
	return nil, nil
}

func NewQReportRemoteRepositoryStub() QReportRemoteRepositoryStub {
	return QReportRemoteRepositoryStub{}
}

func NewQReportLocalRepositoryStub() QReportLocalRepositoryStub {
	return QReportLocalRepositoryStub{}
}