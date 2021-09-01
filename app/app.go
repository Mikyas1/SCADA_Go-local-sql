package app

import (
	"fmt"
	"github.com/Mikyas1/SCADA_Go-local-sql/lines"
	"github.com/fatih/color"
	"time"

	"github.com/Mikyas1/SCADA_Go-local-sql/utils/dateTime"
)

func Start() {

	go X(1)
	go X(2)

	fmt.Scanln()

}


func X(index int) {
	ln, err := lines.NewBDashboardLine(index)
	if err != nil {
		color.Red(fmt.Sprintf("LINE ERROR: error creating communication line for remoteDB Branch index `%v` and localDB", index))
		return
	}

	dt, _ := dateTime.GetYesterday()
	err = ln.RunLine(*dt)
	if err != nil {
		color.Red(fmt.Sprintf("LINE ERROR: error running communication line for remoteDB Branch index `%v` and localDB", index))
		color.Red(fmt.Sprintf("LINE ERROR: %s", *err))
	}
}

func StartTwo() {
	//dt, _ := dateTime.ParseDateTimeFromString(dateTime.Layout1, "2021-01-01 00:00:00")
	dt := time.Now()
	newDt, _ := dateTime.ChangeDateTimeMinToFactorWrapper(&dt, 5, true)
	fmt.Println(newDt)
}
