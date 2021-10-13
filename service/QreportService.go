package service

import (
	"fmt"
	"time"

	"github.com/Mikyas1/SCADA_Go-local-sql/datasources/mysql/remote"
	"github.com/Mikyas1/SCADA_Go-local-sql/domains/Qreport"
	"github.com/fatih/color"
)

type QReportService interface {
	SaveQReports(QReports []Qreport.QReport, branchIndex int) *error
	GetLatestQReport(branchIndex int) (*Qreport.QReport, *error)
	GetQReportAndSave(time.Time, time.Time, int, int) *error
}

type DefaultQReportService struct {
	localRepo  Qreport.LocalRepository
	remoteRepo Qreport.RemoteRepository
}

func (s DefaultQReportService) SaveQReports(qReports []Qreport.QReport, branchIndex int) *error {
	for _, qReport := range qReports {
		err := s.localRepo.Save(qReport, branchIndex)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s DefaultQReportService) GetQReportAndSave(dtFrom, dtTo time.Time, interval, branchIndex int) *error {
	//branch index hard coded
	tempFrom := dtFrom

	for tempFrom.Before(dtTo) {
		tempFormAfterInterval := tempFrom.AddDate(0, 0, interval)

		machineIdes, err := s.GetMachineIds(branchIndex, tempFrom, tempFormAfterInterval)
		if err != nil {
			color.Red(fmt.Sprintf("SERVICE ERROR: error happend when getting Machine Ids from REMOTE DB for \n -> `%v` branch index, \n -> `%v` from time \n -> `%v` to time ", branchIndex, tempFrom, tempFormAfterInterval))
			return nil
		}

		var qReports []Qreport.QReport
		fmt.Println("len of machine ids")
		fmt.Println(len(machineIdes))
		for _, machineId := range machineIdes {
			color.Red("================================= got herer ======================")
			color.Red(fmt.Sprintf("caled get Report for machineId: %s", machineId))
			color.Red("time: %s", tempFrom)

			ress, err := s.GetQReports(branchIndex, machineId, tempFrom, tempFormAfterInterval)
			if err != nil {
				color.Red(fmt.Sprintf("SERVICE ERROR: error happend when getting QReport from REMOTE DB for \n -> `%v` branch index, \n -> `%v` from time \n -> `%v` to time ", branchIndex, tempFrom, tempFormAfterInterval))
			}
			for _, re := range ress {
				qReports = append(qReports, re)
			}
		}

		for _, d := range qReports {
			color.Blue("qreports from remote")
			color.Blue(fmt.Sprintf("%v", d.ProcessDate))
		}

		err = s.SaveQReports(qReports, branchIndex)
		if err != nil {
			color.Red(fmt.Sprintf("SERVICE ERROR: error happend when saving QReport to LOCAL DB for \n -> `%v` branch index, \n -> `%v` from time \n -> `%v` to time ", branchIndex, tempFrom, tempFormAfterInterval))
			return err
		}

		color.Green(fmt.Sprintf("Successfully copied QReport data from REMOTE DB to LOCAL DB for \n -> `%v` branch index, \n -> `%v` from time \n -> `%v` to time ", branchIndex, tempFrom, tempFormAfterInterval))

		tempFrom = tempFormAfterInterval
	}

	return nil
}

func (s DefaultQReportService) GetLatestQReport(branchIndex int) (*Qreport.QReport, *error) {
	qReport, err := s.localRepo.GetLatestQReport(branchIndex)
	if err != nil {
		return nil, err
	}
	return qReport, nil
}

func (s DefaultQReportService) GetMachineIds(branchIndex int, dtFrom, dtTo time.Time) ([]string, *error) {
	isUse23 := remote.Use23(branchIndex)
	processId := 2
	if isUse23 {
		processId = 22
	}
	return s.remoteRepo.GetMachineId(processId, dtFrom, dtTo)
}

func (s DefaultQReportService) GetQReports(branchIndex int, machineId string, dtFrom, dtTo time.Time) ([]Qreport.QReport, *error) {
	isUse23 := remote.Use23(branchIndex)
	processId := 2
	if isUse23 {
		processId = 22
	}
	filling := remote.GetFilling(branchIndex)
	return s.remoteRepo.FindByTimeInterval(branchIndex, processId, filling, machineId, dtFrom, dtTo)
}

func NewQReportService(localRepo Qreport.LocalRepository, remoteRepo Qreport.RemoteRepository) QReportService {
	return DefaultQReportService{
		localRepo:  localRepo,
		remoteRepo: remoteRepo,
	}
}
