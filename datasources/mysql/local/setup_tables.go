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
	createQWeeklyTable    = "create table IF NOT EXISTS QWeekly_%d (id int auto_increment, process_time datetime not null, namem varchar(50) not null, process_id int null, count int not null, cyl_type int null, constraint QWeekly_%d_pk primary key (id));"
	createQSearchTable    = "create table IF NOT EXISTS QSearch_%d (id int auto_increment, count int not null, sort_out int not null, plat_input int not null, process_id int not null, namem varchar(50) not null, count_by_machine_1 int not null, count_by_machine_2 int not null, count_by_machine_3 int not null, process_time datetime not null, cyl_type int null, constraint QSearch_%d_pk primary key (id));"
	createQDashboardTable = "create table IF NOT EXISTS QDashboard_%d (residual int not null, check_net int not null, count int not null, final_datetime datetime not null);"
	createQDashboardIndex = "create unique index IF NOT EXISTS QDashboard_%d_final_datetime_uindex on QDashboard_%d (final_datetime);"
	createQReportTable    = "create table IF NOT EXISTS QReport_%d (machine_id varchar(20) not null, process_date varchar(25) not null, gtem400 int null, gtem350 int null, gtem300 int null, gtem250 int null, gtem200 int null, gtem150 int null, gtem100 int null, gtem050 int null, value int null, gte050 int null, gte100 int null, gte150 int null, gte200 int null, gte250 int null, gte300 int null, gte350 int null, gte400 int null, sum int null, m200x200 int null, diff int null, start_point float null, accuracy float null, cylinder_type int null, final_datetime datetime null);"
)

func SetUpTables() error {
	branches := remote.TotalBranches()
	db, openErr := Open()
	if openErr != nil {
		return *openErr
	}
	defer db.Close()

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
		err = CreateQDashboardTable(db, i)
		if err != nil {
			color.Red(err.Error())
			return err
		}
		//err = CreateQReportTable(db, i)
		//if err != nil {
		//	color.Red(err.Error())
		//	return err
		//}
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

func CreateQDashboardTable(db *sql.DB, i int) error {
	_, err := db.Exec(fmt.Sprintf(createQDashboardTable, i))
	if err != nil {
		return err
	}
	color.Green(fmt.Sprintf("Created QDashboard table for branch id: %d", i))
	_, err = db.Exec(fmt.Sprintf(createQDashboardIndex, i, i))
	if err != nil {
		return err
	}
	color.Green(fmt.Sprintf("Created QDashboard index for branch id: %d", i))
	return nil
}


func CreateQReportTable(db *sql.DB, i int) error {
	_, err := db.Exec(fmt.Sprintf(createQReportTable, i))
	if err != nil {
		return err
	}
	color.Green(fmt.Sprintf("Created QReport table for branch id: %d", i))
	return nil
}