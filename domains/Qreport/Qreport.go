package Qreport

import "time"

type QReport struct {
	MachineId     string `json:"machine_id"`
	ProcessDate   string
	Gtem400       int     `json:"gtem400"`
	Gtem350       int     `json:"gtem350"`
	Gtem300       int     `json:"gtem300"`
	Gtem250       int     `json:"gtem250"`
	Gtem200       int     `json:"gtem200"`
	Gtem150       int     `json:"gtem150"`
	Gtem100       int     `json:"gtem100"`
	Gtem050       int     `json:"gtem050"`
	Value         int     `json:"value"`
	Gte050        int     `json:"gte050"`
	Gte100        int     `json:"gte100"`
	Gte150        int     `json:"gte150"`
	Gte200        int     `json:"gte200"`
	Gte250        int     `json:"gte250"`
	Gte300        int     `json:"gte300"`
	Gte350        int     `json:"gte350"`
	Gte400        int     `json:"gte400"`
	Sum           int     `json:"sum"`
	M200X200      int     `json:"m200X200"`
	Diff          int     `json:"diff"`
	StartPoint    float64 `json:"startpoint"`
	Accuracy      float64 `json:"accuracy"`
	CylinderType  int
	FinalDateTime time.Time `json:"final_datetime"`
}

type LocalRepository interface {
	Save(QReport, int) *error
	GetLatestQReport(branchIndex int) (*QReport, *error)
}

type RemoteRepository interface {
	FindByTimeInterval(branchIndex int, dtFrom, dtTo time.Time) (*QReport, *error)
	GetMachineId(value, processId int, dtFrom, dtTo time.Time) ([]string, *error)
}
