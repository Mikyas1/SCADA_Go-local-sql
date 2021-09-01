package domains

import (
	"database/sql"
	"fmt"
	"github.com/Mikyas1/SCADA_Go-local-sql/utils/dateTime"
	"github.com/fatih/color"
)

const (
	queryInsertWeekly = "INSERT INTO %s (count, final_datetime) VALUES(?, ?);"
	queryGetLatestWeekly = "SELECT MAX(final_datetime) as finalDatetime, count FROM %s;"
)

type weeklyLocalRepositoryDb struct {
	client *sql.DB
}

func (r weeklyLocalRepositoryDb) Save(w Weekly, branchIndex int) *error {
	query := fmt.Sprintf(queryInsertWeekly, fmt.Sprintf("BDashboard_%v", branchIndex))
	stmt, err := r.client.Prepare(query)
	if err != nil {
		color.Red(fmt.Sprintf("SQL ERROR: error when trying to prepare save weekly statement for index `%v`", branchIndex))
		return &err
	}
	defer stmt.Close()
	_, saveErr := stmt.Exec(w.Count, w.FinalDateTime)
	if saveErr != nil {
		color.Red(fmt.Sprintf("QL ERROR: error when trying to run save weekly statement for index `%v`", branchIndex))
		return &saveErr
	}
	return nil
}

func (r weeklyLocalRepositoryDb) GetLatestWeekly(branchIndex int) (*Weekly, *error) {
	query := fmt.Sprintf(queryGetLatestWeekly, fmt.Sprintf("BDashboard_%v", branchIndex))
	stmt, err := r.client.Prepare(query)
	if err != nil {
		color.Red(fmt.Sprintf("SQL ERROR: %s", err))
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
