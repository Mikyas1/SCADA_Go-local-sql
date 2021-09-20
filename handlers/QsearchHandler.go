package handlers

import (
	"fmt"
	"github.com/Mikyas1/SCADA_Go-local-sql/service"
	"github.com/Mikyas1/SCADA_Go-local-sql/utils/dateTime"
	"github.com/fatih/color"
	"strings"
	"time"
)

type QSearchHandler struct {
	Service service.DefaultQSearchService
}

func (h *QSearchHandler) FetchAndSaveQSearchFromRemoteToLocal(branchIndex int, dtFrom, dtTo *time.Time) *error {
	var startDateTime *time.Time

	lastQSearch, err := h.Service.GetLatestQSearch(branchIndex)
	if err != nil {
		if strings.Contains(fmt.Sprintf("%s", *err), "converting NULL to string is unsupported") {

			color.Red(fmt.Sprintf("X Error when trying to scan latest QSearch, LOCAL DB might be empty for `%v` Branch Index", branchIndex))
			color.Blue("--> Setting start date and time to the BEGINNING OF THE YEAR")

			// TODO set start date to beginning
			//startDateTime, _ = dateTime.ParseDateTimeFromString(dateTime.Layout1, "2021-01-01 00:00:00")
			startDateTime, _ = dateTime.GetOnlyDateFormDateTime(time.Now().AddDate(0, 0, -2)) // this two days before today
		} else {
			return err
		}
	} else {
		startDateTime = &lastQSearch.ProcessTime
	}

	startDateTime, _ = dateTime.ChangeDateTimeMinToFactorWrapper(startDateTime, interval, true)
	dtTo, _ = dateTime.ChangeDateTimeMinToFactorWrapper(dtTo, interval, false)

	err = h.Service.GetQSearchAndSave(*startDateTime, *dtTo, interval, branchIndex)
	if err != nil {
		return err
	}
	return nil
}

