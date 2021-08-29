package service

import (
	"fmt"
	"github.com/Mikyas1/SCADA_Go-local-sql/domains"
	"github.com/fatih/color"
	"time"
)

type WeeklyService interface {
	SaveWeekly(weekly domains.Weekly, branchIndex int) *error
	GetLatestWeekly(branchIndex int) (*domains.Weekly, *error)
	GetWeekliesAndSave(time.Time, time.Time, int, int) *error
}


type DefaultCustomerService struct {
	localRepo  domains.WeeklyLocalRepository
	remoteRepo domains.WeeklyRemoteRepository
}

func (s DefaultCustomerService) SaveWeekly(weekly domains.Weekly, branchIndex int) *error {
	err := s.localRepo.Save(weekly, branchIndex)
	if err != nil {
		return err
	}
	return nil
}

func (s DefaultCustomerService) GetWeekliesAndSave(dtFrom, dtTo time.Time, interval, branchIndex int) *error {
	//branch index hard coded
	tempFrom := dtFrom

	for tempFrom.Before(dtTo) {
		tempFormAfterInterval := tempFrom.Add(time.Minute * time.Duration(interval))
		weekly, err := s.remoteRepo.FindByTimeInterval(branchIndex, tempFrom, tempFormAfterInterval)
		if err != nil {
			color.Red(fmt.Sprintf("error happend when geting weekly from REMOTE DB for \n -> `%v` branch index, \n -> `%v` from time \n -> `%v` to time ", branchIndex, tempFrom, tempFormAfterInterval))
			return err
		}

		err = s.SaveWeekly(*weekly, branchIndex)
		if err != nil {
			color.Red(fmt.Sprintf("error happend when saving weekly to LOCAL DB for \n -> `%v` branch index, \n -> `%v` from time \n -> `%v` to time ", branchIndex, tempFrom, tempFormAfterInterval))
			return err
		}

		color.Green(fmt.Sprintf("Successfully copied BDashbord data from REMOTE DB to LOCAL DB for \n -> `%v` branch index, \n -> `%v` from time \n -> `%v` to time ", branchIndex, tempFrom, tempFormAfterInterval))

		tempFrom = tempFormAfterInterval
	}

	return nil
}

func (s DefaultCustomerService) GetLatestWeekly(branchIndex int) (*domains.Weekly, *error) {
	weekly, err := s.localRepo.GetLatestWeekly(branchIndex)
	if err != nil {
		return nil, err
	}
	return weekly, nil
}

func NewCustomerService(localRepo domains.WeeklyLocalRepository, remoteRepo domains.WeeklyRemoteRepository) DefaultCustomerService {
	return DefaultCustomerService{
		localRepo: localRepo,
		remoteRepo: remoteRepo,
	}
}