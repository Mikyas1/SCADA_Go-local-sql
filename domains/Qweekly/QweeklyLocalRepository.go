package Qweekly

import (
	"database/sql"
	"fmt"

	"github.com/Mikyas1/SCADA_Go-local-sql/utils/dateTime"
	"github.com/fatih/color"
)

const (
	queryInsertQWeekly    = "INSERT INTO %s (count, process_time, namem, process_id, cyl_type) VALUES(?, ?, ?, ?, ?);"
	queryGetLatestQWeekly = "SELECT MAX(process_time) as processDateTime, count FROM %s;"
)

type qWeeklyLocalRepositoryDb struct {
	client *sql.DB
}

func (r qWeeklyLocalRepositoryDb) Save(w QWeekly, branchIndex int) *error {
	query := fmt.Sprintf(queryInsertQWeekly, fmt.Sprintf("QWeekly_%v", branchIndex))
	stmt, err := r.client.Prepare(query)
	if err != nil {
		color.Red(fmt.Sprintf("SQL ERROR: error when trying to prepare save QWeekly statement for index `%v`", branchIndex))
		return &err
	}
	defer stmt.Close()
	_, saveErr := stmt.Exec(w.Count, w.ProcessTime, w.Namem, w.getProcessId(), w.getCylinderType())
	if saveErr != nil {
		color.Red(fmt.Sprintf("SQL ERROR: error when trying to run save QWeekly statement for index `%v`", branchIndex))
		return &saveErr
	}
	return nil
}

func (r qWeeklyLocalRepositoryDb) GetLatestQWeekly(branchIndex int) (*QWeekly, *error) {
	query := fmt.Sprintf(queryGetLatestQWeekly, fmt.Sprintf("QWeekly_%v", branchIndex))
	stmt, err := r.client.Prepare(query)
	if err != nil {
		color.Red(fmt.Sprintf("SQL ERROR: %s", err))
		return nil, &err
	}
	defer stmt.Close()
	qWeekly := QWeekly{}
	result := stmt.QueryRow()
	var finalDatetime string
	if getErr := result.Scan(&finalDatetime, &qWeekly.Count); getErr != nil {
		return nil, &getErr
	}
	tempDateTime, timeParseErr := dateTime.ParseDateTimeFromString(dateTime.Layout1, finalDatetime)
	if timeParseErr != nil {
		color.Red("error when trying to parse string datetime to time.Time object")
		return nil, timeParseErr
	}
	qWeekly.ProcessTime = *tempDateTime
	return &qWeekly, nil
}

func NewQWeeklyLocalRepositoryDb(client *sql.DB) qWeeklyLocalRepositoryDb {
	return qWeeklyLocalRepositoryDb{
		client: client,
	}
}
