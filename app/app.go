package app

import (
	"fmt"
	"github.com/Mikyas1/SCADA_Go-local-sql/datasources/mysql/local"
	"github.com/Mikyas1/SCADA_Go-local-sql/datasources/mysql/remote"
	"github.com/Mikyas1/SCADA_Go-local-sql/domains"
	"github.com/Mikyas1/SCADA_Go-local-sql/handlers"
	"github.com/Mikyas1/SCADA_Go-local-sql/service"
	"github.com/Mikyas1/SCADA_Go-local-sql/utils/dateTime"
	"time"
)

func Start() {

	remoteDB, err := remote.Open(1)
	if err != nil {
		panic("couldn't connect to remote server")
	}
	defer remoteDB.Close()

	//wh := handlers.WeeklyHandler{
	//	Service: service.NewCustomerService(
	//		domains.NewWeeklyLocalRepositoryDb(local.Client),
	//		domains.NewWeeklyRemoteRepositoryDb(remoteDB)),
	//}

	wh := handlers.WeeklyHandler{
		Service: service.NewCustomerService(
			domains.NewWeeklyLocalRepositoryDb(local.Client),
			domains.NewWeeklyRemoteRepositoryStub()),
	}

	now := time.Now()
	// app should not be concerned with errors
	err = wh.FetchAndSaveFromRemoteToLocal(1, nil, &now)
	if err != nil {
		fmt.Println(err)
	}
}

func StartTwo() {
	//dt, _ := dateTime.ParseDateTimeFromString(dateTime.Layout1, "2021-01-01 00:00:00")
	dt := time.Now()
	newDt, _ := dateTime.ChangeDateTimeMinToFactorWrapper(&dt, 5, true)
	fmt.Println(newDt)
}