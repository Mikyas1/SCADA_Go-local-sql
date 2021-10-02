package handlers

import (
	"fmt"
	"github.com/Mikyas1/SCADA_Go-local-sql/service"
	"github.com/Mikyas1/SCADA_Go-local-sql/utils/dateTime"
	"github.com/fatih/color"
	"strings"
	"time"
)

type QDashboardHandler struct {
	Service service.QDashboardService
}

func (h *QDashboardHandler) FetchAndSaveQDashboardFromRemoteToLocal(branchIndex int, dtFrom, dtTo *time.Time) *error {
	var startDateTime *time.Time

	lastWeekly, err := h.Service.GetLatestQDashboard(branchIndex)
	if err != nil {
		if strings.Contains(fmt.Sprintf("%s", *err), "converting NULL to string is unsupported") {

			color.Red(fmt.Sprintf("X Error when trying to scan latest QDashoard, LOCAL DB might be empty for `%v` Branch Index", branchIndex))
			color.Blue("--> Setting start date and time to the BEGINNING OF THE YEAR")

			// TODO set start date to beginning
			//startDateTime, _ = dateTime.ParseDateTimeFromString(dateTime.Layout1, startDate)
			startDateTime, _ = dateTime.GetOnlyDateFormDateTime(time.Now().AddDate(0, 0, -4)) // this two days before today
		} else {
			return err
		}
	} else {
		startDateTime = &lastWeekly.FinalDateTime
	}

	startDateTime, _ = dateTime.ChangeDateTimeMinToFactorWrapper(startDateTime, interval, true)
	dtTo, _ = dateTime.ChangeDateTimeMinToFactorWrapper(dtTo, interval, false)

	err = h.Service.GetQDashboardsAndSave(*startDateTime, *dtTo, interval, branchIndex)
	if err != nil {
		return err
	}
	return nil
}
