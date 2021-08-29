package domains

import (
	"database/sql"
	"github.com/Mikyas1/SCADA_Go-local-sql/utils/dateTime"
	"github.com/fatih/color"
)

const (
	queryInsertWeekly = "INSERT INTO table_name (count, final_datetime) VALUES(?, ?);"
	queryGetLatestWeekly = "SELECT MAX(final_datetime) as final_datetime, count FROM table_name;"
)

type weeklyLocalRepositoryDb struct {
	client *sql.DB
}

func (r weeklyLocalRepositoryDb) Save(w Weekly, branchIndex int) *error {
	stmt, err := r.client.Prepare(queryInsertWeekly)
	if err != nil {
		color.Red("error when trying to prepare save weekly statement")
		return &err
	}
	defer stmt.Close()
	_, saveErr := stmt.Exec(w.Count, w.FinalDateTime)
	if saveErr != nil {
		color.Red("error when trying to run save weekly statement")
		return &saveErr
	}
	return nil
}

func (r weeklyLocalRepositoryDb) GetLatestWeekly(branchIndex int) (*Weekly, *error) {
	stmt, err := r.client.Prepare(queryGetLatestWeekly)
	if err != nil {
		color.Red("error when trying to prepare get latest weekly statement")
		return nil, &err
	}
	defer stmt.Close()
	weekly := Weekly{}
	result := stmt.QueryRow()
	var finalDatetime string
	if getErr := result.Scan(&finalDatetime, &weekly.Count); getErr != nil {
		return nil, &getErr
	}
	tempDateTime, timeParseErr := dateTime.ParseDateTimeFromString(dateTime.Layout1, finalDatetime)
	if timeParseErr != nil {
		color.Red("error when trying to parse string datetime to time.Time object")
		return nil, timeParseErr
	}
	weekly.FinalDateTime = *tempDateTime
	return &weekly, nil
}

func NewWeeklyLocalRepositoryDb(client *sql.DB) weeklyLocalRepositoryDb {
	return weeklyLocalRepositoryDb{
		client: client,
	}
}
