package lines

import (
	"database/sql"
	"fmt"
	"github.com/Mikyas1/SCADA_Go-local-sql/datasources/mysql/local"
	"github.com/Mikyas1/SCADA_Go-local-sql/domains/Qsearch"
	"github.com/Mikyas1/SCADA_Go-local-sql/handlers"
	"github.com/Mikyas1/SCADA_Go-local-sql/service"
	"github.com/Mikyas1/SCADA_Go-local-sql/utils/dateTime"
	"github.com/fatih/color"
	"time"
)

type QSearchLine struct {
	handler     handlers.QSearchHandler
	branchIndex int
	remoteDb	*sql.DB
	localDB		*sql.DB
}

func (l QSearchLine) CloseAllDb() {
	if l.remoteDb != nil {
		l.remoteDb.Close()
	}
	l.localDB.Close()
}

func (l QSearchLine) RunLine(toDt time.Time) *error {
	defer l.CloseAllDb()
	err := l.handler.FetchAndSaveQSearchFromRemoteToLocal(l.branchIndex, nil, &toDt)
	if err != nil {
		return err
	}
	color.Green(fmt.Sprintf("Successfully fetched date for branch index `%v` to final time `%s`", l.branchIndex, toDt))
	return nil
}

func NewQSearchLine(index int) (*QSearchLine, *error) {

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

	h := handlers.QSearchHandler{
		Service: service.NewQSearchService(
			Qsearch.NewQSearchLocalRepositoryDb(localDb),
			//Qsearch.NewWeeklyRemoteRepositoryDb(remoteDB),
			Qsearch.NewQSearchRemoteRepositoryStub(),
		),
	}
	return &QSearchLine{
		handler: h,
		branchIndex: index,
		remoteDb: remoteDB,
		localDB: localDb,
	}, nil
}

func RunQSearchLine(index int) error {
	ln, err := NewQSearchLine(index)
	if err != nil {
		color.Red(fmt.Sprintf("LINE ERROR: error creating QSearch communication line for remoteDB Branch index `%v` and localDB", index))
		return *err
	}

	toDt, _ := dateTime.GetYesterday()
	err = ln.RunLine(*toDt)
	if err != nil {
		color.Red(fmt.Sprintf("LINE ERROR: error running QSearch communication line for remoteDB Branch index `%v` and localDB", index))
		color.Red(fmt.Sprintf("LINE ERROR: %s", *err))
		return *err
	}
	return nil
}

func RunAllQSearchBranches(totalBranches int) {
	for i := 0; i < totalBranches; i++ {
		color.White(fmt.Sprintf("--> Task created for QSearch Branch id %d", i))
		err := RunQSearchLine(totalBranches)
		if err != nil {}
	}
}