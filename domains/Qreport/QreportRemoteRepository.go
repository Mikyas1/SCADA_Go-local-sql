package Qreport

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
	str = "SELECT distinct Machine_id from production_log where check_net >= %d AND check_net <= %d AND process_date BETWEEN '%s' AND '%s' AND process_id = %d AND cyl_type = 1 GROUP BY machine_id "
)

func (s RemoteRepositoryDb) GetMachineId(value, processId int, dtFrom, dtTo time.Time) ([]string, *error) {
	var query string
	query = fmt.Sprintf(str, value-400, value+400, dtFrom.Format(dateTime.Layout1), dtTo.Format(dateTime.Layout1), processId)

	selDB, err := s.client.Query(query)
	if err != nil {
		return nil, &err
	}
	var res []string
	for selDB.Next() {
		var MId string
		if err := selDB.Scan(&MId); err != nil {
			color.Red("error when trying to scan remote Machine id.")
			return nil, &err
		}
		res = append(res, MId)
	}
	return res, nil
}
