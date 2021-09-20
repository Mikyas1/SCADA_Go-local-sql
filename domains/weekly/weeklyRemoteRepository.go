package weekly

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/Mikyas1/SCADA_Go-local-sql/utils/dateTime"
	"github.com/fatih/color"
)

type RemoteRepositoryDb struct {
	client   *sql.DB
}

const (
	weeklyStr  = "SELECT q.process_id, COUNT(q.process_id) AS count FROM event_log as q JOIN system_process_relations AS a ON q.process_id = a.process_id"
	weeklyStr1 = "GROUP BY q.process_id, DATE_FORMAT(process_date, '%Y-%m-%d') ORDER BY DATE_FORMAT(process_date, '%Y-%m-%d'), q.process_id"
)

func (s RemoteRepositoryDb) FindByTimeInterval(branchIndex int, dtFrom, dtTo time.Time) (*Weekly, *error) {
	var query string

	query = fmt.Sprintf("%s WHERE process_date > '%s' AND process_date <= '%s' AND q.is_sortout = 0 %s AND q.process_id = 9 ",
		weeklyStr, dtFrom.Format(dateTime.Layout1), dtTo.Format(dateTime.Layout1), weeklyStr1)

	selDB, err := s.client.Query(query)
	if err != nil {
		return nil, &err
	}

	weekly := Weekly{FinalDateTime: dtTo}
	for selDB.Next() {
		var processId int
		var count int
		if err := selDB.Scan(&processId, &count); err != nil {
			color.Red("error when trying to scan remote weekly")
			return nil, &err
		}
		if processId == 9 {
			weekly.Count += count
		}
	}

	return &weekly, nil
}

func NewWeeklyRemoteRepositoryDb(client *sql.DB) RemoteRepositoryDb {
	return RemoteRepositoryDb{
		client: client,
	}
}
