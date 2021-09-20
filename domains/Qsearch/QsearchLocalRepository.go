package Qsearch

import (
	"database/sql"
	"fmt"
	"github.com/Mikyas1/SCADA_Go-local-sql/utils/dateTime"
	"github.com/fatih/color"
)

const (
	queryInsertQSearch    = "INSERT INTO %s (count, sort_out, plat_input, process_id, namem, count_by_machine_1, count_by_machine_2, count_by_machine_3, process_time, cyl_type) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?);"
	queryGetLatestQSearch = "SELECT MAX(process_time) as processDateTime, count FROM %s;"
)

type qSearchLocalRepositoryDb struct {
	client *sql.DB
}

func (r qSearchLocalRepositoryDb) Save(s QSearch, branchIndex int) *error {
	query := fmt.Sprintf(queryInsertQSearch, fmt.Sprintf("QSearch_%v", branchIndex))
	stmt, err := r.client.Prepare(query)
	if err != nil {
		color.Red(fmt.Sprintf("SQL ERROR: error when trying to prepare save QSearch statement for index `%v`", branchIndex))
		return &err
	}
	defer stmt.Close()
	_, saveErr := stmt.Exec(s.Count, s.SortOut, s.PlatInput, s.getProcessId(), s.Namem, s.CountByMachine1, s.CountByMachine2, s.CountByMachine3, s.ProcessTime, s.getCylinderType())
	if saveErr != nil {
		fmt.Println(saveErr)
		color.Red(fmt.Sprintf("SQL ERROR: error when trying to run save Qsearch statement for index `%v`", branchIndex))
		return &saveErr
	}
	return nil
}

func (r qSearchLocalRepositoryDb) GetLatestQSearch(branchIndex int) (*QSearch, *error) {
	query := fmt.Sprintf(queryGetLatestQSearch, fmt.Sprintf("QSearch_%v", branchIndex))
	stmt, err := r.client.Prepare(query)
	if err != nil {
		color.Red(fmt.Sprintf("SQL ERROR: %s", err))
		return nil, &err
	}
	defer stmt.Close()
	qSearch := QSearch{}
	result := stmt.QueryRow()
	var finalDatetime string
	if getErr := result.Scan(&finalDatetime, &qSearch.Count); getErr != nil {
		return nil, &getErr
	}
	tempDateTime, timeParseErr := dateTime.ParseDateTimeFromString(dateTime.Layout1, finalDatetime)
	if timeParseErr != nil {
		color.Red("error when trying to parse string datetime to time.Time object")
		return nil, timeParseErr
	}
	qSearch.ProcessTime = *tempDateTime
	return &qSearch, nil
}

func NewQSearchLocalRepositoryDb(client *sql.DB) qSearchLocalRepositoryDb {
	return qSearchLocalRepositoryDb{
		client: client,
	}
}