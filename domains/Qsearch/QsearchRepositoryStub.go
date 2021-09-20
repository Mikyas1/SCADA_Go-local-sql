package Qsearch

import "time"

type QSearchRemoteRepositoryStub struct {

}

type QSearchLocalRepositoryStub struct {

}

func (s QSearchRemoteRepositoryStub) FindByTimeInterval(branchIndex int, dtFrom, dtTo time.Time) ([]QSearch, *error) {
	qSearches := []QSearch {
		{
			Count: 320,
			ProcessTime: dtTo,
			Namem: "Leak Gas",
			ProcessId: 3,
			SortOut: 2,
			PlatInput: 5,
			CountByMachine1: 32,
			CountByMachine2: 54,
			CountByMachine3: 11,
			CylinderType: 1,
		},
		{
			Count: 420,
			ProcessTime: dtTo,
			Namem: "Palletizer",
			ProcessId: 4,
			SortOut: 1,
			PlatInput: 3,
			CountByMachine1: 32,
			CountByMachine2: 54,
			CountByMachine3: 11,
			CylinderType: 2,
		},
		{
			Count: 120,
			ProcessTime: dtTo,
			Namem: "Leak Gas",
			ProcessId: 3,
			SortOut: 8,
			PlatInput: 13,
			CountByMachine1: 32,
			CountByMachine2: 54,
			CountByMachine3: 11,
			CylinderType: 1,
		},
		{
			Count: 120,
			ProcessTime: dtTo,
			Namem: "Leak Gas",
			ProcessId: 7,
			SortOut: 8,
			PlatInput: 13,
			CountByMachine1: -1,
			CountByMachine2: -1,
			CountByMachine3: -1,
			CylinderType: 2,
		},
	}
	return qSearches, nil
}

func (s QSearchRemoteRepositoryStub) FindNamemByProcessId(processId int) (string, *error) {
	if processId == 3 {return "Leak Gas", nil}
	if processId == 4 {return "Palletizer", nil}
	if processId == 9 {return "Sealing Machine", nil}
	return "", nil
}

func (s QSearchRemoteRepositoryStub) FindCountByIdCylTypeDate(pId int, dtFrom, dtTo, tableName string) ([]QCountAndCylTypeDto, *error) {
	return []QCountAndCylTypeDto{
		{
			Count: 320,
			CylinderType: 1,
		},
		{
			Count: 420,
			CylinderType: 2,
		},
		{
			Count: 120,
			CylinderType: 1,
		},
	}, nil
}

func (s QSearchRemoteRepositoryStub) FindSortOut(pId, cylType int, dtFrom, dtTo, tableName string, use23 bool) (int, *error) {
	if pId == 3 {return 2, nil}
	if pId == 4 {return 1, nil}
	if pId == 9 {return 8, nil}
	return -1, nil
}

func (db QSearchRemoteRepositoryStub) FindPlatInput(pId, cylType int, dtFrom, dtTo, tableName string, use23 bool) (int, *error) {
	if pId == 3 {return 5, nil}
	if pId == 4 {return 3, nil}
	if pId == 9 {return 13, nil}
	return -1, nil
}

func (db QSearchRemoteRepositoryStub) FindCountBYMachine(pId, cylType int, dtFrom, dtTo, tableName string, use23 bool) ([3]int, *error) {
	if pId == 3 {return [3]int{12, 4, 54}, nil}
	if pId == 4 {return [3]int{10, 20, 30}, nil}
	if pId == 9 {return [3]int{71, 65, 54}, nil}
	return [3]int{-1,-1,-1}, nil
}

func (s QSearchLocalRepositoryStub) Save(qWeekly QSearch, branchIndex int) *error {
	return nil
}

func (r QSearchLocalRepositoryStub) GetLatestQSearch(branchIndex int) (*QSearch, *error) {
	return nil, nil
}

func NewQSearchRemoteRepositoryStub() QSearchRemoteRepositoryStub {
	return QSearchRemoteRepositoryStub{}
}

func NewQSearchLocalRepositoryStub() QSearchLocalRepositoryStub {
	return QSearchLocalRepositoryStub{}
}