package Qsearch

import "time"

type QSearch struct {
	Count           int
	SortOut         int
	PlatInput       int
	ProcessId       int
	Namem           string
	CountByMachine1 int
	CountByMachine2 int
	CountByMachine3 int
	ProcessTime     time.Time
	CylinderType    int
}

type QCountAndCylTypeDto struct {
	Count        int
	CylinderType int
}

func (d QSearch) getProcessId() *int {
	processId := &d.ProcessId
	if d.ProcessId == 0 {
		processId = nil
	}
	return processId
}

func (d QSearch) getCylinderType() *int {
	cylinderType := &d.CylinderType
	if d.CylinderType == 0 {
		cylinderType = nil
	}
	return cylinderType
}

type LocalRepository interface {
	Save(QSearch, int) *error
	GetLatestQSearch(int) (*QSearch, *error)
}

type RemoteRepository interface {
	FindByTimeInterval(branchIndex int, dtFrom, dtTo time.Time) ([]QSearch, *error)
	FindNamemByProcessId(int) (string, *error)
	FindCountByIdCylTypeDate(pId int, dtFrom, dtTo, tableName string) ([]QCountAndCylTypeDto, *error)
	FindSortOut(pId, cylType int, dtFrom, dtTo, tableName string, use23 bool) (int, *error)
	FindPlatInput(pId, cylType int, dtFrom, dtTo, tableName string, use23 bool) (int, *error)
	FindCountBYMachine(pId, cylType int, dtFrom, dtTo, tableName string, use23 bool) ([3]int, *error)
}
