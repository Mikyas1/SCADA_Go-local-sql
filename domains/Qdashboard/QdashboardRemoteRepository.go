package Qdashboard

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/Mikyas1/SCADA_Go-local-sql/utils/dateTime"
	"github.com/fatih/color"
)

type RemoteRepositoryDb struct {
	client *sql.DB
}

const (
	qDashboardStr = "SELECT COALESCE(SUM(residual), 0) as residual, COALESCE(SUM(check_net), 0) as check_net, COUNT(process_id) as count from production_log where process_date > '%s' AND process_date <= '%s'"
)

func (s RemoteRepositoryDb) FindByTimeInterval(branchIndex int, dtFrom, dtTo time.Time) (*QDashboard, *error) {
	var query string

	query = fmt.Sprintf(qDashboardStr, dtFrom.Format(dateTime.Layout1), dtTo.Format(dateTime.Layout1))

	selDB, err := s.client.Query(query)
	if err != nil {
		return nil, &err
	}

	qDashboard := QDashboard{FinalDateTime: dtTo}
	for selDB.Next() {
		var residual, checkNet, count int
		if err := selDB.Scan(&residual, &checkNet, &count); err != nil {
			color.Red("error when trying to scan remote QDashboard")
			return nil, &err
		}
		qDashboard.Residual = residual
		qDashboard.CheckNet = checkNet
		qDashboard.Count = count
	}

	return &qDashboard, nil
}

func NewQDashboardRemoteRepositoryDb(client *sql.DB) RemoteRepositoryDb {
	return RemoteRepositoryDb{
		client: client,
	}
}
