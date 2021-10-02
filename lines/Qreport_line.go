package lines

import (
	"database/sql"
	"fmt"
	"github.com/Mikyas1/SCADA_Go-local-sql/datasources/mysql/local"
	"github.com/Mikyas1/SCADA_Go-local-sql/domains/Qreport"
	"github.com/Mikyas1/SCADA_Go-local-sql/handlers"
	"github.com/Mikyas1/SCADA_Go-local-sql/service"
	"github.com/Mikyas1/SCADA_Go-local-sql/utils/dateTime"
	"github.com/fatih/color"
	"time"
)

type QReportLine struct {
	handler     handlers.QReportHandler
	branchIndex int
	remoteDb    *sql.DB
	localDB     *sql.DB
}

func (l QReportLine) CloseAllDb() {
	if l.remoteDb != nil {
		l.remoteDb.Close()
	}
	l.localDB.Close()
}

func (l QReportLine) RunLine(toDt time.Time) *error {
	defer l.CloseAllDb()
	err := l.handler.FetchAndSaveQReportsFromRemoteToLocal(l.branchIndex, nil, &toDt)
	if err != nil {
		return err
	}
	color.Green(fmt.Sprintf("Successfully fetched date for branch index `%v` to final time `%s`", l.branchIndex, toDt))
	return nil
}

func NewQReportLine(index int) (*QReportLine, *error) {

	var remoteDB *sql.DB = nil
	//remoteDB, err := remote.Open(index)
	//if err != nil {
	//	color.Red("error connecting to remote DB")
	//	return nil, err
	//}

	localDb, err := local.Open()
	if err != nil {
		return nil, err
	}

	h := handlers.QReportHandler{
		Service: service.NewQReportService(
			//Qreport.NewQReportLocalRepositoryStub(),
			Qreport.NewQReportLocalRepositoryDb(localDb),
			Qreport.NewQReportRemoteRepositoryStub(),
		),
	}
	return &QReportLine{
		handler:     h,
		branchIndex: index,
		remoteDb:    remoteDB,
		localDB:     localDb,
	}, nil
}

func RunQReportLine(index int) {
	ln, err := NewQReportLine(index)
	if err != nil {
		color.Red(fmt.Sprintf("LINE ERROR: error creating communication line for remoteDB Branch `QReport` index `%v` and localDB", index))
		return
	}

	toDt, _ := dateTime.GetYesterday()
	err = ln.RunLine(*toDt)
	if err != nil {
		color.Red(fmt.Sprintf("LINE ERROR: error running communication line for remoteDB Branch `QReport` index `%v` and localDB", index))
		color.Red(fmt.Sprintf("LINE ERROR: %s", *err))
	}
}


func RunConcurAllQReportBranches(totalBranches int) {
	//for i := 0; i < totalBranches; i++ {
	//	color.White(fmt.Sprintf("--> Task created for QReport Branch id %d", i))
	//	go RunQReportLine(totalBranches)
	//}
	//fmt.Scanln()
	RunQReportLine(1)
}