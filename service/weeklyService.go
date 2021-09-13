package service

import (
	"fmt"
	"github.com/Mikyas1/SCADA_Go-local-sql/domains/weekly"
	"github.com/fatih/color"
	"time"
)

type WeeklyService interface {
	SaveWeekly(weekly weekly.Weekly, branchIndex int) *error
	GetLatestWeekly(branchIndex int) (*weekly.Weekly, *error)
	GetWeekliesAndSave(time.Time, time.Time, int, int) *error
}


type DefaultWeeklyService struct {
	localRepo  weekly.WeeklyLocalRepository
	remoteRepo weekly.WeeklyRemoteRepository
}

func (s DefaultWeeklyService) SaveWeekly(weekly weekly.Weekly, branchIndex int) *error {
	err := s.localRepo.Save(weekly, branchIndex)
	if err != nil {
		return err
	}
	return nil
}

func (s DefaultWeeklyService) GetWeekliesAndSave(dtFrom, dtTo time.Time, interval, branchIndex int) *error {
	//branch index hard coded
	tempFrom := dtFrom

	for tempFrom.Before(dtTo) {
		tempFormAfterInterval := tempFrom.Add(time.Minute * time.Duration(interval))
		weekly, err := s.remoteRepo.FindByTimeInterval(branchIndex, tempFrom, tempFormAfterInterval)
		if err != nil {
			color.Red(fmt.Sprintf("SERVICE ERROR: error happend when getting weekly from REMOTE DB for \n -> `%v` branch index, \n -> `%v` from time \n -> `%v` to time ", branchIndex, tempFrom, tempFormAfterInterval))
			return err
		}

		err = s.SaveWeekly(*weekly, branchIndex)
		if err != nil {
			color.Red(fmt.Sprintf("SERVICE ERROR: error happend when saving weekly to LOCAL DB for \n -> `%v` branch index, \n -> `%v` from time \n -> `%v` to time ", branchIndex, tempFrom, tempFormAfterInterval))
			return err
		}

		color.Green(fmt.Sprintf("Successfully copied BDashbord data from REMOTE DB to LOCAL DB for \n -> `%v` branch index, \n -> `%v` from time \n -> `%v` to time ", branchIndex, tempFrom, tempFormAfterInterval))

		tempFrom = tempFormAfterInterval
	}

	return nil
}

func (s DefaultWeeklyService) GetLatestWeekly(branchIndex int) (*weekly.Weekly, *error) {
	weekly, err := s.localRepo.GetLatestWeekly(branchIndex)
	if err != nil {
		return nil, err
	}
	return weekly, nil
}

func NewWeeklyService(localRepo weekly.WeeklyLocalRepository, remoteRepo weekly.WeeklyRemoteRepository) DefaultWeeklyService {
	return DefaultWeeklyService{
		localRepo: localRepo,
		remoteRepo: remoteRepo,
	}
}