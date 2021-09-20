package Qsearch

import (
	"database/sql"
	"fmt"
	"github.com/fatih/color"
	"time"
)

const (
	queryGetNamemByProcessId = "SELECT comment as namem from system_process_relations WHERE process_id = ?"
)

type RemoteRepositoryDb struct {
	client    *sql.DB
}

func (db RemoteRepositoryDb) FindByTimeInterval(branchIndex int, dtFrom, dtTo time.Time) ([]*QSearch, *error) {
	return nil, nil
}

func (db RemoteRepositoryDb) FindNamemByProcessId(processId int) (string, error) {
	stmt, err := db.client.Prepare(queryGetNamemByProcessId)
	if err != nil {
		color.Red(fmt.Sprintf("SQL ERROR: %s", err))
		return "", err
	}
	defer stmt.Close()
	var namem string
	result := stmt.QueryRow()
	if getErr := result.Scan(&namem); getErr != nil {
		return "", getErr
	}
	return namem, nil
}
