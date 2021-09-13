package handlers

import (
	"fmt"
	"strings"
	"time"

	"github.com/Mikyas1/SCADA_Go-local-sql/service"
	"github.com/Mikyas1/SCADA_Go-local-sql/utils/dateTime"
	"github.com/fatih/color"
)

type QWeeklyHandler struct {
	Service service.DefaultQWeeklyService
}

func (h *QWeeklyHandler) FetchAndSaveQWeekliesFromRemoteToLocal(branchIndex int, dtFrom, dtTo *time.Time) *error {
	var startDateTime *time.Time

	lastQWeekly, err := h.Service.GetLatestQWeekly(branchIndex)
	if err != nil {
		if strings.Contains(fmt.Sprintf("%s", *err), "converting NULL to string is unsupported") {

			color.Red(fmt.Sprintf("X Error when trying to scan latest Qweekly, LOCAL DB might be empty for `%v` Branch Index", branchIndex))
			color.Blue("--> Setting start date and time to the BEGINNING OF THE YEAR")

			// TODO set start date to beginning
			//startDateTime, _ = dateTime.ParseDateTimeFromString(dateTime.Layout1, "2021-01-01 00:00:00")
			startDateTime, _ = dateTime.GetOnlyDateFormDateTime(time.Now().AddDate(0, 0, -2)) // this two days before today
		} else {
			return err
		}
	} else {
		startDateTime = &lastQWeekly.ProcessTime
	}

	startDateTime, _ = dateTime.ChangeDateTimeMinToFactorWrapper(startDateTime, interval, true)
	dtTo, _ = dateTime.ChangeDateTimeMinToFactorWrapper(dtTo, interval, false)

	err = h.Service.GetQWeekliesAndSave(*startDateTime, *dtTo, interval, branchIndex)
	if err != nil {
		return err
	}
	return nil
}
