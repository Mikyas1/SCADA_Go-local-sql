package app

import (
	"fmt"
	"time"

	"github.com/Mikyas1/SCADA_Go-local-sql/lines"

	"github.com/Mikyas1/SCADA_Go-local-sql/utils/dateTime"
)

func Start() {

	go lines.RunQWeeklyLine(3)
	// go lines.RunBDashboardLine(3)
	// go lines.RunQSearchLine(1)

	fmt.Scanln()

	// if err := local.SetUpTables(); err != nil {
	// 	panic(err)
	// }

}

func StartTwo() {
	//dt, _ := dateTime.ParseDateTimeFromString(dateTime.Layout1, "2021-01-01 00:00:00")
	dt := time.Now()
	newDt, _ := dateTime.ChangeDateTimeMinToFactorWrapper(&dt, 5, true)
	fmt.Println(newDt)
}
