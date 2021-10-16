package Qreport

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"github.com/Mikyas1/SCADA_Go-local-sql/utils/dateTime"
	"github.com/fatih/color"
)

type RemoteRepositoryDb struct {
	client *sql.DB
}

const (
	str    = "SELECT distinct Machine_id from production_log where check_net >= %d AND check_net <= %d AND process_date > '%s' AND process_date <= '%s' AND process_id = %d AND cyl_type = 1 GROUP BY machine_id "
	query1 = "SELECT count(machine_id), cyl_type from production_log where machine_id  = '%v' AND  check_net <= %d AND process_date > '%s' AND process_date <= '%s' AND process_id = %d AND (process_status = 0 OR process_status = 1040 OR process_status = 2064) GROUP BY machine_id, cyl_type  "
	query2 = "SELECT count(machine_id), cyl_type from production_log where machine_id  = '%v' AND  check_net >= %d AND process_date > '%s' AND process_date <= '%s' AND process_id = %d AND (process_status = 0 OR process_status = 1040 OR process_status = 2064) GROUP BY machine_id, cyl_type  "
	query3 = "SELECT count(machine_id), cyl_type from production_log where machine_id = '%v' AND  check_net = %d AND process_date > '%s' AND process_date <= '%s' AND process_id = %d AND (process_status = 0 OR process_status = 1040 OR process_status = 2064) GROUP BY machine_id, cyl_type  "
)

func (s RemoteRepositoryDb) GetMachineId(processId int, dtFrom, dtTo time.Time) ([]string, *error) {
	var query string
	Value := 11000
	query = fmt.Sprintf(str, Value-400, Value+400, dtFrom.Format(dateTime.Layout1), dtTo.Format(dateTime.Layout1), processId)

	selDB, err := s.client.Query(query)
	if err != nil {
		return nil, &err
	}
	var res []string
	for selDB.Next() {
		var MId string
		if err := selDB.Scan(&MId); err != nil {
			color.Red("error when trying to scan remote Machine id.")
			return nil, &err
		}
		res = append(res, MId)
	}
	return res, nil
}

func (s RemoteRepositoryDb) FindByTimeInterval(branchIndex, processId, filling int, machineId string, dtFrom, dtTo time.Time) ([]QReport, *error) {
	var results []QReport
	var query string
	checkNets := []int{10600, 10650, 10700, 10750, 10800, 10850, 10900, 10950, 11000, 11050, 11100, 11150, 11200, 11250, 11300, 11350, 11400}
	resCyl := map[int][17]int{}
	for i := 0; i < 17; i++ {
		color.Green("========================================== got here ====================")
		color.Green("running the 17 loops")
		if i == 0 {
			query = fmt.Sprintf(query1, machineId, checkNets[i], dtFrom.Format(dateTime.Layout1), dtTo.Format(dateTime.Layout1), processId)
		} else if i == 16 {
			query = fmt.Sprintf(query2, machineId, checkNets[i], dtFrom.Format(dateTime.Layout1), dtTo.Format(dateTime.Layout1), processId)
		} else {
			query = fmt.Sprintf(query3, machineId, checkNets[i], dtFrom.Format(dateTime.Layout1), dtTo.Format(dateTime.Layout1), processId)
		}
		res, err := s.client.Query(query)
		if err != nil {
			color.Red(fmt.Sprintf("SQL ERROR: %s.", err.Error()))
			return nil, &err
		}
		for res.Next() {
			var count int
			var cylType int
			//err := res.Scan(&n[i])
			err := res.Scan(&count, &cylType)
			if err != nil {
				fmt.Println(fmt.Sprintf("%s", err.Error()))
				color.Red("SQL ERROR: error when trying to scan remote QReport.")
				return nil, &err
			}
			// get if exists cylType int array from map
			// assign the i's element to the result count from db
			// if not exist create an int array with i's element to count from db
			mutateCylTypeQReports(resCyl, count, i, cylType)
		}
	}

	for cylType, n := range resCyl {
		gte400 := n[0]
		gte350 := n[1]
		gte300 := n[2]
		gte250 := n[3]
		gte200 := n[4]
		gte150 := n[5]
		gte100 := n[6]
		gte050 := n[7]
		Cvalue := n[8]
		gtem050 := n[9]
		gtem100 := n[10]
		gtem150 := n[11]
		gtem200 := n[12]
		gtem250 := n[13]
		gtem300 := n[14]
		gtem350 := n[15]
		gtem400 := n[16]
		dateString := fmt.Sprintf("From: %s -  To:  %s", dtFrom.Format(dateTime.Layout1), dtTo.Format(dateTime.Layout1))
		var j, sum int
		for j = 0; j < 17; j++ {
			sum += n[j]
		}
		m200X200 := n[16] + n[15] + n[14] + n[13] + n[3] + n[2] + n[1] + n[0]
		diff := sum - m200X200

		results = append(results, QReport{
			MachineId:     machineId,
			ProcessDate:   dateString,
			Gtem400:       gtem400,
			Gtem350:       gtem350,
			Gtem300:       gtem300,
			Gtem250:       gtem250,
			Gtem200:       gtem200,
			Gtem150:       gtem150,
			Gtem100:       gtem100,
			Gtem050:       gtem050,
			Value:         Cvalue,
			Gte050:        gte050,
			Gte100:        gte100,
			Gte150:        gte150,
			Gte200:        gte200,
			Gte250:        gte250,
			Gte300:        gte300,
			Gte350:        gte350,
			Gte400:        gte400,
			Sum:           sum,
			M200X200:      m200X200,
			Diff:          diff,
			StartPoint:    0.0,
			Accuracy:      0.0,
			CylinderType:  cylType,
			FinalDateTime: dtTo,
		})
	}

	var sumOfSum float64
	for k := 0; k < len(results); k++ {
		sumOfSum += float64(results[k].Sum)
	}
	for l := 0; l < len(results); l++ {
		var temp1, temp2, perScale float64
		perScale = sumOfSum / float64(filling)
		temp1 = (float64(results[l].Sum) / perScale) * 100
		results[l].StartPoint, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", temp1), 64)

		temp2 = (float64(results[l].Diff) / float64(results[l].Sum)) * 100
		results[l].Accuracy, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", temp2), 64)
	}
	return results, nil
}

// get if exists cylType int array from map
// assign the i's element to the result count from db
// if not exist create an int array with i's element to count from db
func mutateCylTypeQReports(data map[int][17]int, count, i, cylType int) {
	val, ok := data[cylType]
	if ok {
		color.Green(fmt.Sprintf("val %v", val[i]))
		color.Green(fmt.Sprintf("val + count %v", val[i]+count))
		val[i] = val[i] + count
	} else {
		var inits [17]int
		inits[i] = count
		color.Green(fmt.Sprintf("array %d", inits))
		data[cylType] = inits
	}
}

func NewRemoteRepositoryDb(client *sql.DB) RemoteRepository {
	return RemoteRepositoryDb{
		client: client,
	}
}
