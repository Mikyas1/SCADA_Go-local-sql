package local

import (
	"database/sql"
	"fmt"
	"github.com/Mikyas1/SCADA_Go-local-sql/datasources/mysql/remote"
	"github.com/fatih/color"
)

const (
	createBDashboardTable = "create table IF NOT EXISTS BDashboard_%d (count int default 0 not null, final_datetime datetime not null, id int auto_increment, constraint BDashboard_2_pk primary key (id));"
	createBDashboardIndex = "create unique index IF NOT EXISTS BDashboard_%d_final_datetime_uindex on BDashboard_%d (final_datetime);"
	createQWeeklyTable = "create table IF NOT EXISTS QWeekly_%d (id int auto_increment, process_time datetime not null, namem varchar(50) not null, process_id int null, count int not null, cyl_type int null, constraint QWeekly_%d_pk primary key (id));"
	createQSearchTable = "create table IF NOT EXISTS QSearch_%d (id int auto_increment, count int not null, sort_out int not null, plat_input int not null, process_id int not null, namem varchar(50) not null, count_by_machine_1 int not null, count_by_machine_2 int not null, count_by_machine_3 int not null, process_time datetime not null, cyl_type int null, constraint QSearch_%d_pk primary key (id));"
)

func SetUpTables() error {
	branches := remote.TotalBranches()
	db, openErr := Open()
	defer db.Close()
	if openErr != nil {
		return *openErr
	}

	for i := 0; i <= branches; i++ {

		err := CreateBDashboardTable(db, i)
		if err != nil {
			color.Red(err.Error())
			return err
		}
		err = CreateQWeeklyTable(db, i)
		if err != nil {
			color.Red(err.Error())
			return err
		}
		err = CreateQSearchTable(db, i)
		if err != nil {
			color.Red(err.Error())
			return err
		}
	}

	return nil
}

func CreateBDashboardTable(db *sql.DB, i int) error {
	_, err := db.Exec(fmt.Sprintf(createBDashboardTable, i))
	if err != nil {
		return err
	}
	color.Green(fmt.Sprintf("Created BDashboard table for branch id: %d", i))
	_, err = db.Exec(fmt.Sprintf(createBDashboardIndex, i, i))
	if err != nil {
		return err
	}
	color.Green(fmt.Sprintf("Created BDashboard index for branch id: %d", i))
	return nil
}

func CreateQWeeklyTable(db *sql.DB, i int) error {
	_, err := db.Exec(fmt.Sprintf(createQWeeklyTable, i, i))
	if err != nil {
		return err
	}
	color.Green(fmt.Sprintf("Created QWeekly table for branch id: %d", i))
	return nil
}

func CreateQSearchTable(db *sql.DB, i int) error {
	_, err := db.Exec(fmt.Sprintf(createQSearchTable, i, i))
	if err != nil {
		return err
	}
	color.Green(fmt.Sprintf("Created QSearch table for branch id: %d", i))
	return nil
}