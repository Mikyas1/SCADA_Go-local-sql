package service

import (
	"fmt"
	"github.com/Mikyas1/SCADA_Go-local-sql/domains/Qreport"
	"github.com/fatih/color"
	"time"
)

type QReportService interface {
	SaveQReport(QReport Qreport.QReport, branchIndex int) *error
	GetLatestQReport(branchIndex int) (*Qreport.QReport, *error)
	GetQReportAndSave(time.Time, time.Time, int, int) *error
}

type DefaultQReportService struct {
	localRepo  Qreport.LocalRepository
	remoteRepo Qreport.RemoteRepository
}


func (s DefaultQReportService) SaveQReport(qReport Qreport.QReport, branchIndex int) *error {
	err := s.localRepo.Save(qReport, branchIndex)
	if err != nil {
		return err
	}
	return nil
}

func (s DefaultQReportService) GetQReportAndSave(dtFrom, dtTo time.Time, interval, branchIndex int) *error {
	//branch index hard coded
	tempFrom := dtFrom

	for tempFrom.Before(dtTo) {
		tempFormAfterInterval := tempFrom.Add(time.Minute * time.Duration(interval))
		qReport, err := s.remoteRepo.FindByTimeInterval(branchIndex, tempFrom, tempFormAfterInterval)
		if err != nil {
			color.Red(fmt.Sprintf("SERVICE ERROR: error happend when getting QReport from REMOTE DB for \n -> `%v` branch index, \n -> `%v` from time \n -> `%v` to time ", branchIndex, tempFrom, tempFormAfterInterval))
			return err
		}

		err = s.SaveQReport(*qReport, branchIndex)
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

func NewQReportService(localRepo Qreport.LocalRepository, remoteRepo Qreport.RemoteRepository) QReportService {
	return DefaultQReportService{
		localRepo: localRepo,
		remoteRepo: remoteRepo,
	}
}