package Qreport

import "time"

type QReportRemoteRepositoryStub struct {
}

type QReportLocalRepositoryStub struct {
}

func (s QReportRemoteRepositoryStub) FindByTimeInterval(branchIndex, processId, filling int, machineId string, dtFrom, dtTo time.Time) ([]QReport, *error) {
	results := []QReport{
		{
			MachineId:     "12",
			ProcessDate:   "somedate",
			Gtem400:       12,
			Gtem350:       11,
			Gtem300:       11,
			Gtem250:       11,
			Gtem200:       11,
			Gtem150:       11,
			Gtem100:       11,
			Gtem050:       11,
			Value:         45,
			Gte050:        4,
			Gte100:        4,
			Gte150:        4,
			Gte200:        4,
			Gte250:        4,
			Gte300:        4,
			Gte350:        4,
			Gte400:        4,
			Sum:           300,
			M200X200:      98,
			Diff:          43,
			StartPoint:    2,
			Accuracy:      4,
			CylinderType:  1,
			FinalDateTime: dtTo,
		},
		{
			MachineId:     "2",
			ProcessDate:   "somedate",
			Gtem400:       112,
			Gtem350:       111,
			Gtem300:       111,
			Gtem250:       111,
			Gtem200:       111,
			Gtem150:       111,
			Gtem100:       111,
			Gtem050:       11,
			Value:         45,
			Gte050:        14,
			Gte100:        14,
			Gte150:        14,
			Gte200:        14,
			Gte250:        14,
			Gte300:        14,
			Gte350:        14,
			Gte400:        14,
			Sum:           300,
			M200X200:      198,
			Diff:          143,
			StartPoint:    12,
			Accuracy:      14,
			CylinderType:  2,
			FinalDateTime: dtTo,
		},
	}
	return results, nil
}

func (s QReportRemoteRepositoryStub) GetMachineId(processId int, dtFrom, dtTo time.Time) ([]string, *error) {
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
