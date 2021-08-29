package handlers

import (
	"fmt"
	"strings"
	"time"

	"github.com/Mikyas1/SCADA_Go-local-sql/service"
	"github.com/Mikyas1/SCADA_Go-local-sql/utils/dateTime"
	"github.com/fatih/color"
)

const (
	interval  int    = 5
	startDate string = "2021-01-01 00:00:00"
)

type WeeklyHandler struct {
	Service service.DefaultCustomerService
}

func (h *WeeklyHandler) FetchAndSaveFromRemoteToLocal(branchIndex int, dtFrom, dtTo *time.Time) *error {
	var startDateTime *time.Time

	lastWeekly, err := h.Service.GetLatestWeekly(branchIndex)
	if err != nil {
		if strings.Contains(fmt.Sprintf("%s", *err), "converting NULL to string is unsupported") {

			color.Red(fmt.Sprintf("X Error when trying to scan latest weekly, LOCAL DB might be empty for `%v` Branch Index", branchIndex))
			color.Blue("--> Setting start date and time to the BEGINNING OF THE YEAR")

			// TODO set start date to beginning
			//startDateTime, _ = dateTime.ParseDateTimeFromString(dateTime.Layout1, "2021-01-01 00:00:00")
			tempTime := time.Now().Add(-time.Hour * time.Duration(5*24)) // this tow days before today
			startDateTime = &tempTime
		} else {
			return err
		}
	} else {
		startDateTime = &lastWeekly.FinalDateTime
	}

	startDateTime, _ = dateTime.ChangeDateTimeMinToFactorWrapper(startDateTime, interval, true)
	dtTo, _ = dateTime.ChangeDateTimeMinToFactorWrapper(dtTo, interval, false)

	// TODO prepare start and to date time
	err = h.Service.GetWeekliesAndSave(*startDateTime, *dtTo, interval, branchIndex)
	if err != nil {
		return err
	}
	return nil
}
