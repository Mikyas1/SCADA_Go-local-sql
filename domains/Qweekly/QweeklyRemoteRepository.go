package Qweekly

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/Mikyas1/SCADA_Go-local-sql/utils/dateTime"
	"github.com/fatih/color"
)

type RemoteRepositoryDb struct {
	client   *sql.DB
	qWeeklies []QWeekly
}

const (
	str  = "SELECT q.process_id, COUNT(q.process_id) AS count, a.comment as namem, q.cyl_type FROM event_log as q JOIN system_process_relations AS a ON q.process_id = a.process_id"
	str1 = "GROUP BY q.process_id, DATE_FORMAT(process_date, '%Y-%m-%d') ORDER BY DATE_FORMAT(process_date, '%Y-%m-%d'), q.process_id"
)

func (s RemoteRepositoryDb) FindByTimeInterval(branchIndex int, dtFrom, dtTo time.Time) (*QWeekly, *error) {
	var query string

	query = fmt.Sprintf("%s WHERE process_date > '%s' AND process_date <= '%s' AND q.is_sortout = 0 %s ",
		str, dtFrom.Format(dateTime.Layout1), dtTo.Format(dateTime.Layout1), str1)

	selDB, err := s.client.Query(query)
	if err != nil {
		return nil, &err
	}

	qWeekly := QWeekly{ProcessTime: dtTo}
	for selDB.Next() {
		var processId int
		var count int
		var namem string
		var cylinderType int
		if err := selDB.Scan(&processId, &count, &namem, &cylinderType); err != nil {
			color.Red("error when trying to scan remote QWeekly")
			return nil, &err
		}
		qWeekly.ProcessId = processId
		qWeekly.Count = count
		qWeekly.Namem = namem
		qWeekly.CylinderType = cylinderType
	}

	return &qWeekly, nil
}

func NewQWeeklyRemoteRepositoryDb(client *sql.DB) RemoteRepositoryDb {
	return RemoteRepositoryDb{
		client: client,
	}
}
