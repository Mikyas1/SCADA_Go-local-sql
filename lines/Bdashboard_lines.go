package lines

import (
	"database/sql"
	"fmt"
	"github.com/Mikyas1/SCADA_Go-local-sql/datasources/mysql/local"
	"github.com/Mikyas1/SCADA_Go-local-sql/datasources/mysql/remote"
	"github.com/Mikyas1/SCADA_Go-local-sql/domains/weekly"
	"github.com/Mikyas1/SCADA_Go-local-sql/handlers"
	"github.com/Mikyas1/SCADA_Go-local-sql/service"
	"github.com/Mikyas1/SCADA_Go-local-sql/utils/dateTime"
	"github.com/fatih/color"
	"time"
)

type BDashboardLine struct {
	handler     handlers.WeeklyHandler
	branchIndex int
	remoteDb	*sql.DB
	localDB		*sql.DB
}

func (l BDashboardLine) CloseAllDb() {
	if l.remoteDb != nil {
		l.remoteDb.Close()
	}
	l.localDB.Close()
}

func (l BDashboardLine) RunLine(toDt time.Time) *error {
	defer l.CloseAllDb()
	err := l.handler.FetchAndSaveWeekliesFromRemoteToLocal(l.branchIndex, nil, &toDt)
	if err != nil {
		return err
	}
	color.Green(fmt.Sprintf("Successfully fetched date for branch index `%v` to final time `%s`", l.branchIndex, toDt))
	return nil
}

func NewBDashboardLine(index int) (*BDashboardLine, *error) {

	//var remoteDB *sql.DB = nil
	remoteDB, err := remote.Open(index)
	if err != nil {
		color.Red("error connecting to remote DB")
		return nil, err
	}

	localDb, err := local.Open()
	if err != nil {
		return nil, err
	}

	h := handlers.WeeklyHandler{
		Service: service.NewWeeklyService(
			weekly.NewWeeklyLocalRepositoryDb(localDb),
			weekly.NewWeeklyRemoteRepositoryDb(remoteDB),
			//domains.NewWeeklyRemoteRepositoryStub(),
			),
	}
	return &BDashboardLine{
		handler: h,
		branchIndex: index,
		remoteDb: remoteDB,
		localDB: localDb,
	}, nil
}

func RunBDashboardLine(index int) {
	ln, err := NewBDashboardLine(index)
	if err != nil {
		color.Red(fmt.Sprintf("LINE ERROR: error creating communication line for remoteDB Branch index `%v` and localDB", index))
		return
	}

	toDt, _ := dateTime.GetYesterday()
	err = ln.RunLine(*toDt)
	if err != nil {
		color.Red(fmt.Sprintf("LINE ERROR: error running communication line for remoteDB Branch index `%v` and localDB", index))
		color.Red(fmt.Sprintf("LINE ERROR: %s", *err))
	}
}

func RunConcurAllBDashboardBranches(totalBranches int) {
	for i := 0; i < totalBranches; i++ {
		color.White(fmt.Sprintf("--> Task created for BDashboard Branch id %d", i))
		go RunBDashboardLine(totalBranches)
	}
	fmt.Scanln()
}