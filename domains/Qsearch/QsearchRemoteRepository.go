package Qsearch

import (
	"database/sql"
	"fmt"
	"github.com/fatih/color"
	"time"
)

const (
	queryGetNamemByProcessId   = "SELECT comment as namem from system_process_relations WHERE process_id = ?"
	queryGetCountWithDateRange = "SELECT COUNT(*) as count, q.cyl_type as cylinder_type from %s as q WHERE q.process_id=%d AND (q.process_date BETWEEN '%s' AND '%s' ) GROUP BY q.cyl_type"
	querySortOut               = "SELECT COUNT(*) as count from %s as q WHERE q.process_id=%d AND (q.process_date BETWEEN '%s' AND '%s' ) AND q.process_status = 16 AND q.cyl_type=%d"
	queryPlatInput             = "SELECT COUNT(*) as count from %s as q WHERE q.process_id=%d AND (q.process_date BETWEEN '%s' AND '%s' ) AND q.machine_id = 1 AND q.cyl_type=%d"
	queryCountByMachine        = "SELECT COUNT(*) as count from %s as q WHERE q.process_id=%d AND (q.process_date BETWEEN '%s' AND '%s' ) AND q.machine_id = %d AND q.cyl_type=%d"
)

type RemoteRepositoryDb struct {
	client *sql.DB
}

func (db RemoteRepositoryDb) FindByTimeInterval(branchIndex int, dtFrom, dtTo time.Time) ([]QSearch, *error) {
	return nil, nil
}

func (db RemoteRepositoryDb) FindNamemByProcessId(processId int) (string, *error) {
	stmt, err := db.client.Prepare(queryGetNamemByProcessId)
	if err != nil {
		color.Red(fmt.Sprintf("SQL ERROR: %s", err))
		return "", &err
	}
	defer stmt.Close()
	var namem string
	result := stmt.QueryRow()
	if getErr := result.Scan(&namem); getErr != nil {
		color.Red(fmt.Sprintf("SQL ERROR: %s", err))
		return "", &getErr
	}
	return namem, nil
}

func (db RemoteRepositoryDb) FindCountByIdCylTypeDate(pId int, dtFrom, dtTo, tableName string) ([]QCountAndCylTypeDto, *error) {
	query := fmt.Sprintf(queryGetCountWithDateRange, tableName, pId, dtFrom, dtTo)

	selDB, err := db.client.Query(query)
	if err != nil {
		color.Red(fmt.Sprintf("SQL ERROR: %s", err))
		return nil, &err
	}

	var res []QCountAndCylTypeDto
	for selDB.Next() {
		countWithCylType := QCountAndCylTypeDto{}
		err := selDB.Scan(&countWithCylType.Count, &countWithCylType.CylinderType)
		if err != nil {
			color.Red(fmt.Sprintf("SQL ERROR: %s", err))
			return nil, &err
		}
		res = append(res, countWithCylType)
	}
	return res, nil
}

func (db RemoteRepositoryDb) FindSortOut(pId, cylType int, dtFrom, dtTo, tableName string, use23 bool) (int, *error) {
	if use23 == false && (pId == 3 || pId == 4) || use23 == true && (pId == 23 || pId == 24) {
		query := fmt.Sprintf(querySortOut, tableName, pId, dtFrom, dtTo, cylType)
		selDB, _ := db.client.Query(query)
		var sortOut int
		for selDB.Next() {
			err := selDB.Scan(&sortOut)
			if err != nil {
				color.Red(fmt.Sprintf("SQL ERROR: %s", err))
				return -1, &err
			}
		}
		return sortOut, nil
	} else {
		return -1, nil
	}
}

func (db RemoteRepositoryDb) FindPlatInput(pId, cylType int, dtFrom, dtTo, tableName string, use23 bool) (int, *error) {
	if pId == 7 {
		query := fmt.Sprintf(queryPlatInput, tableName, pId, dtFrom, dtTo, cylType)
		selDB, _ := db.client.Query(query)
		var platInput int
		for selDB.Next() {
			err := selDB.Scan(&platInput)
			if err != nil {
				color.Red(fmt.Sprintf("SQL ERROR: %s", err))
				return -1, &err
			}
		}
		return platInput, nil
	} else {
		return -1, nil
	}
}

func (db RemoteRepositoryDb) FindCountBYMachine(pId, cylType int, dtFrom, dtTo, tableName string, use23 bool) ([3]int, *error) {
	if use23 == true && (pId == 9 || pId == 10) {

		mIds := []int{1, 2, 3}
		countByM := [3]int{}

		for index, mID := range mIds {
			query := fmt.Sprintf(queryCountByMachine, tableName, pId, dtFrom, dtTo, mID, cylType)
			selDB, _ := db.client.Query(query)
			for selDB.Next() {
				err := selDB.Scan(&countByM[index])
				if err != nil {
					color.Red(fmt.Sprintf("SQL ERROR: %s", err))
					return [3]int{-1,-1,-1}, &err
				}
			}

		}
		return countByM, nil

	} else {
		return [3]int{-1,-1,-1}, nil
	}
}
