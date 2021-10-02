package Qdashboard

import (
	"database/sql"
	"fmt"
	"github.com/Mikyas1/SCADA_Go-local-sql/utils/dateTime"
	"github.com/fatih/color"
	"time"
)

type RemoteRepositoryDb struct {
	client *sql.DB
}

const (
	qDashboardStr = "SELECT SUM(residual) as residual, SUM(check_net) as check_net, COUNT(process_id) as count from production_log where process_date BETWEEN '%s' AND '%s'"
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
