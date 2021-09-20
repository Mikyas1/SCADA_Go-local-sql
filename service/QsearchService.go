package service

import (
	"fmt"
	"github.com/Mikyas1/SCADA_Go-local-sql/datasources/mysql/remote"
	"github.com/Mikyas1/SCADA_Go-local-sql/domains/Qsearch"
	"github.com/Mikyas1/SCADA_Go-local-sql/utils/dateTime"
	"github.com/fatih/color"
	"time"
)

type QSearchService interface {
	SaveQSearch(qSearch Qsearch.QSearch, branchIndex int) *error
	GetLatestQSearch(branchIndex int) (*Qsearch.QSearch, *error)
	GetQSearchAndSave(time.Time, time.Time, int, int) *error
}

type DefaultQSearchService struct {
	localRepo  Qsearch.LocalRepository
	remoteRepo Qsearch.RemoteRepository
	namem      map[int]string
}

func (s DefaultQSearchService) SaveQSearch(qSearch Qsearch.QSearch, branchIndex int) *error {
	err := s.localRepo.Save(qSearch, branchIndex)
	if err != nil {
		return err
	}
	return nil
}

func (s DefaultQSearchService) GetQSearchAndSave(dtFrom, dtTo time.Time, interval, branchIndex int) *error {
	//branch index hard coded
	tempFrom := dtFrom

	for tempFrom.Before(dtTo) {
		tempFormAfterInterval := tempFrom.Add(time.Minute * time.Duration(interval))
		qSearches, err := s.GetQSearch(branchIndex, tempFrom, tempFormAfterInterval)
		if err != nil {
			color.Red(fmt.Sprintf("SERVICE ERROR: error happend when getting Qsearch from REMOTE DB for \n -> `%v` branch index, \n -> `%v` from time \n -> `%v` to time ", branchIndex, tempFrom, tempFormAfterInterval))
			return err
		}

		for _, qSearch := range qSearches {

			err = s.SaveQSearch(qSearch, branchIndex)
			if err != nil {
				color.Red(fmt.Sprintf("SERVICE ERROR: error happend when saving QSearch to LOCAL DB for \n -> `%v` branch index, \n -> `%v` from time \n -> `%v` to time ", branchIndex, tempFrom, tempFormAfterInterval))
				return err
			}

			color.Green(fmt.Sprintf("Successfully copied QSearch data from REMOTE DB to LOCAL DB for \n -> `%v` branch index, \n -> `%v` from time \n -> `%v` to time ", branchIndex, tempFrom, tempFormAfterInterval))
		}

		tempFrom = tempFormAfterInterval
	}

	return nil
}

func (s DefaultQSearchService) GetLatestQSearch(branchIndex int) (*Qsearch.QSearch, *error) {
	qSearch, err := s.localRepo.GetLatestQSearch(branchIndex)
	if err != nil {
		return nil, err
	}
	return qSearch, nil
}

func (s DefaultQSearchService) GetNamem(processId int) (string, *error) {
	return s.remoteRepo.FindNamemByProcessId(processId)
}

func (s DefaultQSearchService) GetPIds(branchIndex int) []int {
	isUse23 := remote.Use23(branchIndex)
	if isUse23 {
		return []int{7, 5, 23, 24, 10, 9}
	} else {
		return []int{7, 5, 3, 4, 10, 9}
	}
}

func (s DefaultQSearchService) GetIsUse23(branchIndex int) bool {
	return remote.Use23(branchIndex)
}

func (s DefaultQSearchService) GetCountWithCylinderType(pId int, dtFrom, dtTo, tableName string) ([]Qsearch.QCountAndCylTypeDto, *error) {
	return s.remoteRepo.FindCountByIdCylTypeDate(pId, dtFrom, dtTo, tableName)
}

func (s DefaultQSearchService) GetSortOut(pId, cylType int, dtFrom, dtTo, tableName string, use23 bool) (int, *error) {
	return s.remoteRepo.FindSortOut(pId, cylType, dtFrom, dtTo, tableName, use23)
}

func (s DefaultQSearchService) GetPlatInput(pId, cylType int, dtFrom, dtTo, tableName string, use23 bool) (int, *error) {
	return s.remoteRepo.FindPlatInput(pId, cylType, dtFrom, dtTo, tableName, use23)
}

func (s DefaultQSearchService) GetCountByMachine(pId, cylType int, dtFrom, dtTo, tableName string, use23 bool) ([3]int, *error) {
	return s.remoteRepo.FindCountBYMachine(pId, cylType, dtFrom, dtTo, tableName, use23)
}

func (s DefaultQSearchService) MemoizedNamemService(pId int) (string, *error) {
	if s.namem == nil {s.namem = make(map[int]string)}

	if s.namem[pId] == "" {
		val, err := s.GetNamem(pId)
		if err != nil {return "", err}
		s.namem[pId] = val
		return val, nil
	} else {
		return s.namem[pId], nil
	}
}

func (s DefaultQSearchService) GetQSearch(branchIndex int, dtFrom, dtTo time.Time) ([]Qsearch.QSearch, *error) {

	var res []Qsearch.QSearch
	dtFromStr := dtFrom.Format(dateTime.Layout1)
	dtToStr := dtTo.Format(dateTime.Layout1)
	isUse23 := s.GetIsUse23(branchIndex)

	for _, pId := range s.GetPIds(branchIndex) {
		var tableName = "event_log"
		if pId == 5 {
			tableName = "production_log"
		}

		namem, err := s.MemoizedNamemService(pId)
		if err != nil {
			color.Red(fmt.Sprintf("SERVICE ERROR: error happend when getting namem from REMOTE DB for \n -> `%v` branch index, \n -> `%v` from time \n -> `%v` to time, \n -> `%v` Process id ", branchIndex, dtFrom, dtTo, pId))
			color.Red(fmt.Sprintf("%s", *err))
			return nil, err
		}

		countWithCylTypes, err := s.GetCountWithCylinderType(pId, dtFromStr, dtToStr, tableName)
		if err != nil {
			color.Red(fmt.Sprintf("SERVICE ERROR: error happend when getting Count With Cylinder Type from REMOTE DB for \n -> `%v` branch index, \n -> `%v` from time \n -> `%v` to time, \n -> `%v` Process id ", branchIndex, dtFrom, dtTo, pId))
			color.Red(fmt.Sprintf("%s", *err))
			return nil, err
		}

		for _, countWithCylType := range countWithCylTypes {

			sortOut, err := s.GetSortOut(pId, countWithCylType.CylinderType, dtFromStr, dtToStr, tableName, isUse23)
			if err != nil {
				color.Red(fmt.Sprintf("SERVICE ERROR: error happend when getting Sort Out from REMOTE DB for \n -> `%v` branch index, \n -> `%v` from time \n -> `%v` to time, \n -> `%v` Process id ", branchIndex, dtFrom, dtTo, pId))
				color.Red(fmt.Sprintf("%s", *err))
				return nil, err
			}

			platInput, err := s.GetPlatInput(pId, countWithCylType.CylinderType, dtFromStr, dtToStr, tableName, isUse23)
			if err != nil {
				color.Red(fmt.Sprintf("SERVICE ERROR: error happend when getting Plat Input from REMOTE DB for \n -> `%v` branch index, \n -> `%v` from time \n -> `%v` to time, \n -> `%v` Process id ", branchIndex, dtFrom, dtTo, pId))
				color.Red(fmt.Sprintf("%s", *err))
				return nil, err
			}

			countByMachine, err := s.GetCountByMachine(pId, countWithCylType.CylinderType, dtFromStr, dtToStr, tableName, isUse23)
			if err != nil {
				color.Red(fmt.Sprintf("SERVICE ERROR: error happend when getting Count By Machine from REMOTE DB for \n -> `%v` branch index, \n -> `%v` from time \n -> `%v` to time, \n -> `%v` Process id ", branchIndex, dtFrom, dtTo, pId))
				color.Red(fmt.Sprintf("%s", *err))
				return nil, err
			}

			res = append(res, Qsearch.QSearch{
				Count:           countWithCylType.Count,
				SortOut:         sortOut,
				PlatInput:       platInput,
				ProcessId:       pId,
				Namem:           namem,
				CountByMachine1: countByMachine[0],
				CountByMachine2: countByMachine[1],
				CountByMachine3: countByMachine[2],
				ProcessTime:     dtTo,
				CylinderType:    countWithCylType.CylinderType,
			})
		}

	}

	return res, nil
}

func NewQSearchService(localRepo Qsearch.LocalRepository, remoteRepo Qsearch.RemoteRepository) DefaultQSearchService {
	return DefaultQSearchService{
		localRepo:  localRepo,
		remoteRepo: remoteRepo,
	}
}
