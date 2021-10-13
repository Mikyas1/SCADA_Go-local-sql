package service

import (
	"fmt"
	"time"

	"github.com/Mikyas1/SCADA_Go-local-sql/domains/Qdashboard"
	"github.com/fatih/color"
)

type QDashboardService interface {
	SaveQDashboard(qDashboard Qdashboard.QDashboard, branchIndex int) *error
	GetLatestQDashboard(branchIndex int) (*Qdashboard.QDashboard, *error)
	GetQDashboardsAndSave(time.Time, time.Time, int, int) *error
}

type DefaultQDashboardService struct {
	localRepo  Qdashboard.LocalRepository
	remoteRepo Qdashboard.RemoteRepository
}

func (s DefaultQDashboardService) SaveQDashboard(qDashboard Qdashboard.QDashboard, branchIndex int) *error {
	err := s.localRepo.Save(qDashboard, branchIndex)
	if err != nil {
		return err
	}
	return nil
}

func (s DefaultQDashboardService) GetQDashboardsAndSave(dtFrom, dtTo time.Time, interval, branchIndex int) *error {
	//branch index hard coded
	tempFrom := dtFrom

	for tempFrom.Before(dtTo) {
		tempFormAfterInterval := tempFrom.Add(time.Minute * time.Duration(interval))
		qDashboard, err := s.remoteRepo.FindByTimeInterval(branchIndex, tempFrom, tempFormAfterInterval)
		if err != nil {
			color.Red(fmt.Sprintf("SERVICE ERROR: error happend when getting qDashboard from REMOTE DB for \n -> `%v` branch index, \n -> `%v` from time \n -> `%v` to time ", branchIndex, tempFrom, tempFormAfterInterval))
			return err
		}

		if qDashboard.CheckNet == 0 && qDashboard.Count == 0 && qDashboard.Residual == 0 && !dtFrom.Equal(tempFrom) {
			color.Green(fmt.Sprintf("Data is 0 so skipping \n -> `%v` branch index, \n -> `%v` from time \n -> `%v` to time ", branchIndex, tempFrom, tempFormAfterInterval))
			tempFrom = tempFormAfterInterval
			continue
		}

		err = s.SaveQDashboard(*qDashboard, branchIndex)
		if err != nil {
			color.Red(fmt.Sprintf("SERVICE ERROR: error happend when saving qDashboard to LOCAL DB for \n -> `%v` branch index, \n -> `%v` from time \n -> `%v` to time ", branchIndex, tempFrom, tempFormAfterInterval))
			return err
		}

		color.Green(fmt.Sprintf("Successfully copied qDashboard data from REMOTE DB to LOCAL DB for \n -> `%v` branch index, \n -> `%v` from time \n -> `%v` to time ", branchIndex, tempFrom, tempFormAfterInterval))

		tempFrom = tempFormAfterInterval
	}

	return nil
}

func (s DefaultQDashboardService) GetLatestQDashboard(branchIndex int) (*Qdashboard.QDashboard, *error) {
	qDashboard, err := s.localRepo.GetLatestWeekly(branchIndex)
	if err != nil {
		return nil, err
	}
	return qDashboard, nil
}

func NewQDashboardService(localRepo Qdashboard.LocalRepository, remoteRepo Qdashboard.RemoteRepository) DefaultQDashboardService {
	return DefaultQDashboardService{
		localRepo:  localRepo,
		remoteRepo: remoteRepo,
	}
}
