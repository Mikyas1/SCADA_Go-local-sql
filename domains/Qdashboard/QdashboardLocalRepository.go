package Qdashboard

import (
	"database/sql"
	"fmt"
	"github.com/Mikyas1/SCADA_Go-local-sql/utils/dateTime"
	"github.com/fatih/color"
)

const (
	queryInsertQDashboard = "INSERT INTO %s (residual, check_net, count, final_datetime) VALUES(?, ?, ?, ?);"
	queryGetLatestQDashboard = "SELECT MAX(final_datetime) as finalDatetime, count FROM %s;"
)

type QDashboardLocalRepositoryDb struct {
	client *sql.DB
}

func (r QDashboardLocalRepositoryDb) Save(q QDashboard, branchIndex int) *error {
	query := fmt.Sprintf(queryInsertQDashboard, fmt.Sprintf("QDashboard_%v", branchIndex))
	stmt, err := r.client.Prepare(query)
	if err != nil {
		color.Red(fmt.Sprintf("SQL ERROR: error when trying to prepare save QDashboard statement for index `%v`", branchIndex))
		return &err
	}
	defer stmt.Close()
	_, saveErr := stmt.Exec(q.Residual, q.CheckNet, q.Count, q.FinalDateTime)
	if saveErr != nil {
		color.Red(fmt.Sprintf("QL ERROR: error when trying to run save QDashboard statement for index `%v`", branchIndex))
		return &saveErr
	}
	return nil
}

func (r QDashboardLocalRepositoryDb) GetLatestWeekly(branchIndex int) (*QDashboard, *error) {
	query := fmt.Sprintf(queryGetLatestQDashboard, fmt.Sprintf("QDashboard_%v", branchIndex))
	stmt, err := r.client.Prepare(query)
	if err != nil {
		color.Red(fmt.Sprintf("SQL ERROR: %s", err))
		return nil, &err
	}
	defer stmt.Close()
	qDashboard := QDashboard{}
	result := stmt.QueryRow()
	var finalDatetime string
	if getErr := result.Scan(&finalDatetime, &qDashboard.Count); getErr != nil {
		return nil, &getErr
	}
	tempDateTime, timeParseErr := dateTime.ParseDateTimeFromString(dateTime.Layout1, finalDatetime)
	if timeParseErr != nil {
		color.Red("error when trying to parse string datetime to time.Time object")
		return nil, timeParseErr
	}
	qDashboard.FinalDateTime = *tempDateTime
	return &qDashboard, nil
}

func NewQDashboardLocalRepositoryDb(client *sql.DB) QDashboardLocalRepositoryDb {
	return QDashboardLocalRepositoryDb{
		client: client,
	}
}
