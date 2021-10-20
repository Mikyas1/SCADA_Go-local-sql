package lines

import (
	"database/sql"
	"fmt"
	"sync"
	"time"

	"github.com/Mikyas1/SCADA_Go-local-sql/datasources/mysql/local"
	"github.com/Mikyas1/SCADA_Go-local-sql/datasources/mysql/remote"
	"github.com/Mikyas1/SCADA_Go-local-sql/domains/Qweekly"
	"github.com/Mikyas1/SCADA_Go-local-sql/handlers"
	"github.com/Mikyas1/SCADA_Go-local-sql/service"
	"github.com/Mikyas1/SCADA_Go-local-sql/utils/dateTime"
	"github.com/fatih/color"
)

type QWeeklyLine struct {
	handler     handlers.QWeeklyHandler
	branchIndex int
	remoteDb    *sql.DB
	localDB     *sql.DB
}

func (l QWeeklyLine) CloseAllDb() {
	if l.remoteDb != nil {
		l.remoteDb.Close()
	}
	l.localDB.Close()
}

func (l QWeeklyLine) RunLine(toDt time.Time) *error {
	defer l.CloseAllDb()
	err := l.handler.FetchAndSaveQWeekliesFromRemoteToLocal(l.branchIndex, nil, &toDt)
	if err != nil {
		return err
	}
	color.Green(fmt.Sprintf("Successfully fetched date for branch index `%v` to final time `%s`", l.branchIndex, toDt))
	return nil
}

func NewQWeeklyLine(index int) (*QWeeklyLine, *error) {

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

	h := handlers.QWeeklyHandler{
		Service: service.NewQWeeklyService(
			Qweekly.NewQWeeklyLocalRepositoryDb(localDb),
			// Qweekly.NewQWeeklyRemoteRepositoryStub(),
			Qweekly.NewQWeeklyRemoteRepositoryDb(remoteDB),
		),
	}
	return &QWeeklyLine{
		handler:     h,
		branchIndex: index,
		remoteDb:    remoteDB,
		localDB:     localDb,
	}, nil
}

func RunQWeeklyLine(index int, wg *sync.WaitGroup) {
	ln, err := NewQWeeklyLine(index)
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
	wg.Done()
}

func RunConcurAllQWeeklyBranches(totalBranches int) {
	var wg sync.WaitGroup
	for i := 0; i < totalBranches; i++ {
		wg.Add(1)
		color.White(fmt.Sprintf("--> Task created for QWeekly Branch id %d", i))
		go RunQWeeklyLine(totalBranches, &wg)
	}
	wg.Wait()
}