package service

import (
	"fmt"
	"github.com/Mikyas1/SCADA_Go-local-sql/domains/Qweekly"
	"github.com/fatih/color"
	"time"
)

type QWeeklyService interface {
	SaveQWeekly(qWeekly Qweekly.QWeekly, branchIndex int) *error
	GetLatestQWeekly(branchIndex int) (*Qweekly.QWeekly, *error)
	GetQWeekliesAndSave(time.Time, time.Time, int, int) *error
}


type DefaultQWeeklyService struct {
	localRepo  Qweekly.LocalRepository
	remoteRepo Qweekly.RemoteRepository
}

func (s DefaultQWeeklyService) SaveQWeekly(qWeekly Qweekly.QWeekly, branchIndex int) *error {
	err := s.localRepo.Save(qWeekly, branchIndex)
	if err != nil {
		return err
	}
	return nil
}

func (s DefaultQWeeklyService) GetQWeekliesAndSave(dtFrom, dtTo time.Time, interval, branchIndex int) *error {
	//branch index hard coded
	tempFrom := dtFrom

	for tempFrom.Before(dtTo) {
		tempFormAfterInterval := tempFrom.Add(time.Minute * time.Duration(interval))
		weekly, err := s.remoteRepo.FindByTimeInterval(branchIndex, tempFrom, tempFormAfterInterval)
		if err != nil {
			color.Red(fmt.Sprintf("SERVICE ERROR: error happend when getting Qweekly from REMOTE DB for \n -> `%v` branch index, \n -> `%v` from time \n -> `%v` to time ", branchIndex, tempFrom, tempFormAfterInterval))
			return err
		}

		err = s.SaveQWeekly(*weekly, branchIndex)
		if err != nil {
			color.Red(fmt.Sprintf("SERVICE ERROR: error happend when saving QWeekly to LOCAL DB for \n -> `%v` branch index, \n -> `%v` from time \n -> `%v` to time ", branchIndex, tempFrom, tempFormAfterInterval))
			return err
		}

		color.Green(fmt.Sprintf("Successfully copied QWeekly data from REMOTE DB to LOCAL DB for \n -> `%v` branch index, \n -> `%v` from time \n -> `%v` to time ", branchIndex, tempFrom, tempFormAfterInterval))

		tempFrom = tempFormAfterInterval
	}

	return nil
}

func (s DefaultQWeeklyService) GetLatestQWeekly(branchIndex int) (*Qweekly.QWeekly, *error) {
	qWeekly, err := s.localRepo.GetLatestQWeekly(branchIndex)
	if err != nil {
		return nil, err
	}
	return qWeekly, nil
}

func NewQWeeklyService(localRepo Qweekly.LocalRepository, remoteRepo Qweekly.RemoteRepository) DefaultQWeeklyService {
	return DefaultQWeeklyService{
		localRepo: localRepo,
		remoteRepo: remoteRepo,
	}
}