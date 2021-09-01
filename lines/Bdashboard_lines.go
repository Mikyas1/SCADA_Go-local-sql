package lines

import (
	"database/sql"
	"fmt"
	"github.com/Mikyas1/SCADA_Go-local-sql/datasources/mysql/local"
	"github.com/Mikyas1/SCADA_Go-local-sql/datasources/mysql/remote"
	"github.com/Mikyas1/SCADA_Go-local-sql/domains"
	"github.com/Mikyas1/SCADA_Go-local-sql/handlers"
	"github.com/Mikyas1/SCADA_Go-local-sql/service"
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
	err := l.handler.FetchAndSaveFromRemoteToLocal(l.branchIndex, nil, &toDt)
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
		Service: service.NewCustomerService(
			domains.NewWeeklyLocalRepositoryDb(localDb),
			domains.NewWeeklyRemoteRepositoryDb(remoteDB),
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