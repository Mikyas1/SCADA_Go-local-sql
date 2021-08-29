package domains

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/Mikyas1/SCADA_Go-local-sql/utils/dateTime"
	"github.com/fatih/color"
)

type WeeklyRemoteRepositoryDb struct {
	client   *sql.DB
	weeklies []Weekly
}

const (
	queryGetRemoteWeeklies = `SELECT q.process_id, COUNT(q.process_id) AS count,
								DATE_FORMAT(process_date, '%Y-%m-%d') AS period, 
								a.comment as namem FROM event_log as q JOIN system_process_relations 
								AS a ON q.process_id = a.process_id WHERE process_date > ? 
								AND process_date <= ? AND q.is_sortout = 0
								GROUP BY q.process_id, DATE_FORMAT(process_date, '%Y-%m-%d')
								ORDER BY DATE_FORMAT(process_date, '%Y-%m-%d'), q.process_id;`

	//must be tested
	betterQueryGetRemoteWeeklies = `SELECT process_id, COUNT(process_id) as count FROM event_log 
									WHERE process_date > ? AND process_date <= ? AND is_sortout = 0 AND process_id = ?;`
)

func (s WeeklyRemoteRepositoryDb) FindByTimeInterval(branchIndex int, dtFrom, dtTo time.Time) (*Weekly, *error) {
	// stmt, err := s.client.Prepare(queryGetRemoteWeeklies)
	// if err != nil {
	// 	color.Red("error when trying to prepare get remote weeklies statement")
	// 	return nil, &err
	// }

	var query string

	str := "SELECT q.process_id, COUNT(q.process_id) AS count FROM event_log as q JOIN system_process_relations AS a ON q.process_id = a.process_id"
	str1 := "GROUP BY q.process_id, DATE_FORMAT(process_date, '%Y-%m-%d') ORDER BY DATE_FORMAT(process_date, '%Y-%m-%d'), q.process_id"
	query = fmt.Sprintf("%s WHERE process_date > '%s' AND process_date <= '%s' AND q.is_sortout = 0 %s AND q.process_id = 9 ",
		str, dtFrom.Format(dateTime.Layout1), dtTo.Format(dateTime.Layout1), str1)

	selDB, err := s.client.Query(query)
	if err != nil {
		log.Panic(err)
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

	// defer stmt.Close()
	// weekly := Weekly{FinalDateTime: dtTo}
	// results, err := stmt.Query(dtFrom.Format(dateTime.Layout1), dtTo.Format(dateTime.Layout1), 9)
	// if err != nil {
	// 	color.Red("error when trying to run get remote weeklies statement")
	// 	return nil, &err
	// }
	// for results.Next() {
	// 	var processId int
	// 	var count int

	// 	if err := results.Scan(&processId, &count); err != nil {
	// 		color.Red("error when trying to scan remote weekly")
	// 		return nil, &err
	// 	}
	// 	if processId == 9 {
	// 		weekly.Count += count
	// 	}
	// }
	return &weekly, nil
}

func NewWeeklyRemoteRepositoryDb(client *sql.DB) WeeklyRemoteRepositoryDb {
	return WeeklyRemoteRepositoryDb{
		client: client,
	}
}
