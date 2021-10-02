package Qreport

import (
	"database/sql"
	"fmt"
	"github.com/Mikyas1/SCADA_Go-local-sql/utils/dateTime"
	"github.com/fatih/color"
)

const (
	queryInsertQReport = "INSERT INTO %s (machine_id, process_date, gtem400, gtem350, gtem300, gtem250, gtem200, gtem150, gtem100, gtem050, value, gte050, gte100, gte150, gte200, gte250, gte300, gte350, gte400, sum, m200x200, diff, start_point, accuracy, cylinder_type, final_datetime) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);"
	queryGetLatestQReport = "SELECT MAX(final_datetime) as finalDatetime FROM %s;"
)


type QReportLocalRepositoryDb struct {
	client *sql.DB
}

func (r QReportLocalRepositoryDb) Save(w QReport, branchIndex int) *error {
	query := fmt.Sprintf(queryInsertQReport, fmt.Sprintf("QReport_%v", branchIndex))
	stmt, err := r.client.Prepare(query)
	if err != nil {
		color.Red(fmt.Sprintf("SQL ERROR: error when trying to prepare save QReport statement for index `%v`", branchIndex))
		return &err
	}
	defer stmt.Close()
	_, saveErr := stmt.Exec(w.MachineId, w.ProcessDate, w.Gtem400, w.Gtem350, w.Gtem300, w.Gtem250, w.Gtem200, w.Gtem150, w.Gtem100, w.Gtem050, w.Value, w.Gte050, w.Gte100, w.Gte150, w.Gte200, w.Gte250, w.Gte300, w.Gte350, w.Gte400, w.Sum, w.M200X200, w.Diff, w.StartPoint, w.Accuracy, w.CylinderType, w.FinalDateTime)
	if saveErr != nil {
		color.Red(fmt.Sprintf("QL ERROR: error when trying to run save QReport statement for index `%v`", branchIndex))
		return &saveErr
	}
	return nil
}

func (r QReportLocalRepositoryDb) GetLatestQReport(branchIndex int) (*QReport, *error) {
	query := fmt.Sprintf(queryGetLatestQReport, fmt.Sprintf("QReport_%v", branchIndex))
	stmt, err := r.client.Prepare(query)
	if err != nil {
		color.Red(fmt.Sprintf("SQL ERROR: %s", err))
		return nil, &err
	}
	defer stmt.Close()
	qReport := QReport{}
	result := stmt.QueryRow()
	var finalDatetime string
	if getErr := result.Scan(&finalDatetime); getErr != nil {
		return nil, &getErr
	}
	tempDateTime, timeParseErr := dateTime.ParseDateTimeFromString(dateTime.Layout1, finalDatetime)
	if timeParseErr != nil {
		color.Red("error when trying to parse string datetime to time.Time object")
		return nil, timeParseErr
	}
	qReport.FinalDateTime = *tempDateTime
	return &qReport, nil
}

func NewQReportLocalRepositoryDb(client *sql.DB) QReportLocalRepositoryDb {
	return QReportLocalRepositoryDb{
		client: client,
	}
}
