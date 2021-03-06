package handlers

import (
	"fmt"
	"github.com/Mikyas1/SCADA_Go-local-sql/service"
	"github.com/Mikyas1/SCADA_Go-local-sql/utils/dateTime"
	"github.com/fatih/color"
	"strings"
	"time"
)

type QReportHandler struct {
	Service service.QReportService
}

func (h *QReportHandler) FetchAndSaveQReportsFromRemoteToLocal(branchIndex int, dtFrom, dtTo *time.Time) *error {
	var startDateTime *time.Time

	lastWeekly, err := h.Service.GetLatestQReport(branchIndex)
	if err != nil {
		if strings.Contains(fmt.Sprintf("%s", *err), "converting NULL to string is unsupported") {

			color.Red(fmt.Sprintf("X Error when trying to scan latest QReport, LOCAL DB might be empty for `%v` Branch Index", branchIndex))
			color.Blue("--> Setting start date and time to the BEGINNING OF THE YEAR")

			// TODO set start date to beginning
			//startDateTime, _ = dateTime.ParseDateTimeFromString(dateTime.Layout1, startDate)
			startDateTime, _ = dateTime.GetOnlyDateFormDateTime(time.Now().AddDate(0, 0, -5)) // this two days before today
		} else {
			return err
		}
	} else {
		startDateTime = &lastWeekly.FinalDateTime
	}

	//startDateTime, _ = dateTime.ChangeDateTimeMinToFactorWrapper(startDateTime, interval, true)
	dtTo, _ = dateTime.GetOnlyDateFormDateTime(*dtTo)

	err = h.Service.GetQReportAndSave(*startDateTime, *dtTo, dayInterval, branchIndex)
	if err != nil {
		return err
	}
	return nil
}
