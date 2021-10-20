package lines

import (
	"database/sql"
	"fmt"
	"sync"
	"time"

	"github.com/Mikyas1/SCADA_Go-local-sql/datasources/mysql/local"
	"github.com/Mikyas1/SCADA_Go-local-sql/datasources/mysql/remote"
	"github.com/Mikyas1/SCADA_Go-local-sql/domains/Qdashboard"
	"github.com/Mikyas1/SCADA_Go-local-sql/handlers"
	"github.com/Mikyas1/SCADA_Go-local-sql/service"
	"github.com/Mikyas1/SCADA_Go-local-sql/utils/dateTime"
	"github.com/fatih/color"
)

type QDashboardLine struct {
	handler     handlers.QDashboardHandler
	branchIndex int
	remoteDb    *sql.DB
	localDB     *sql.DB
}

func (l QDashboardLine) CloseAllDb() {
	if l.remoteDb != nil {
		l.remoteDb.Close()
	}
	l.localDB.Close()
}

func (l QDashboardLine) RunLine(toDt time.Time) *error {
	defer l.CloseAllDb()
	err := l.handler.FetchAndSaveQDashboardFromRemoteToLocal(l.branchIndex, nil, &toDt)
	if err != nil {
		return err
	}
	color.Green(fmt.Sprintf("Successfully fetched date for branch index `%v` to final time `%s`", l.branchIndex, toDt))
	return nil
}

func NewQDashboardLine(index int) (*QDashboardLine, *error) {

	// var remoteDB *sql.DB = nil
	remoteDB, err := remote.Open(index)
	if err != nil {
		color.Red("error connecting to remote DB")
		return nil, err
	}

	localDb, err := local.Open()
	if err != nil {
		return nil, err
	}

	h := handlers.QDashboardHandler{
		Service: service.NewQDashboardService(
			Qdashboard.NewQDashboardLocalRepositoryDb(localDb),
			// Qdashboard.NewQDashboardRemoteRepositoryStub(),
			Qdashboard.NewQDashboardRemoteRepositoryDb(remoteDB),
		),
	}
	return &QDashboardLine{
		handler:     h,
		branchIndex: index,
		remoteDb:    remoteDB,
		localDB:     localDb,
	}, nil
}

func RunQDashboardLine(index int, wg *sync.WaitGroup) {
	ln, err := NewQDashboardLine(index)
	if err != nil {
		color.Red(fmt.Sprintf("LINE ERROR: error creating QDashboard communication line for remoteDB Branch index `%v` and localDB", index))
		//return *err
	}

	toDt, _ := dateTime.GetYesterday()
	err = ln.RunLine(*toDt)
	if err != nil {
		color.Red(fmt.Sprintf("LINE ERROR: error running QDashboard communication line for remoteDB Branch index `%v` and localDB", index))
		color.Red(fmt.Sprintf("LINE ERROR: %s", *err))
		//return *err
	}
	wg.Done()
	//return nil
}

func RunConcurAllQDashboardBranches(totalBranches int) {
	var wg sync.WaitGroup
	for i := 0; i < totalBranches; i++ {
		wg.Add(1)
		color.White(fmt.Sprintf("--> Task created for QDashboard Branch id %d", i))
		go RunQDashboardLine(totalBranches, &wg)
	}
	wg.Wait()
	//RunQDashboardLine(3)
}
