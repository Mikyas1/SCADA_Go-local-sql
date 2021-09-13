package app

import (
	"fmt"
	"github.com/Mikyas1/SCADA_Go-local-sql/lines"
	"time"

	"github.com/Mikyas1/SCADA_Go-local-sql/utils/dateTime"
)

func Start() {

	go lines.RunQWeeklyLine(1)
	//go lines.RunBDashboardLine(2)

	fmt.Scanln()

}


func StartTwo() {
	//dt, _ := dateTime.ParseDateTimeFromString(dateTime.Layout1, "2021-01-01 00:00:00")
	dt := time.Now()
	newDt, _ := dateTime.ChangeDateTimeMinToFactorWrapper(&dt, 5, true)
	fmt.Println(newDt)
}
