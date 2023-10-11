package dashboardcontroller

import (
	"database/sql"
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func DashboardListTask(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	//get id from url path
	getid := r.URL.Path[len("/dashboardadmins/taskdue/"):]
	var taskduetoday int
	result, err := db.Query("SELECT COUNT(*) FROM tbtask JOIN tbtaskdue ON tbtask.id = tbtaskdue.task_id JOIN tbtaskassignees ON tbtask.id = tbtaskassignees.task_id WHERE tbtaskdue.date_due = CURDATE() AND tbtaskassignees.user_id = (SELECT id FROM tbadmins WHERE id = ?)", getid)
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()

	for result.Next() {
		err := result.Scan(&taskduetoday)
		if err != nil {
			panic(err.Error())
		}
	}

	var taskdueweek int
	//query data from database tbtask and tbtaskdue and tbtaskassignees by id and date due this week in monday to sunday and count data from database to taskdueweek variable
	result2, err := db.Query("SELECT COUNT(*) FROM tbtask JOIN tbtaskdue ON tbtask.id = tbtaskdue.task_id JOIN tbtaskassignees ON tbtask.id = tbtaskassignees.task_id WHERE tbtaskdue.date_due BETWEEN DATE_SUB(CURDATE(), INTERVAL WEEKDAY(CURDATE()) DAY) AND DATE_ADD(CURDATE(), INTERVAL 6-WEEKDAY(CURDATE()) DAY) AND tbtaskassignees.user_id = (SELECT id FROM tbadmins WHERE id = ?)", getid)
	if err != nil {
		panic(err.Error())
	}
	defer result2.Close()

	for result2.Next() {
		err := result2.Scan(&taskdueweek)
		if err != nil {
			panic(err.Error())
		}
	}

	var taskduenextweek int
	//query data from database tbtask and tbtaskdue and tbtaskassignees by id and date due next week in monday to sunday and count data from database to taskduenextweek variable
	result3, err := db.Query("SELECT COUNT(*) FROM tbtask JOIN tbtaskdue ON tbtask.id = tbtaskdue.task_id JOIN tbtaskassignees ON tbtask.id = tbtaskassignees.task_id WHERE tbtaskdue.date_due BETWEEN DATE_ADD(CURDATE(), INTERVAL 7-WEEKDAY(CURDATE()) DAY) AND DATE_ADD(CURDATE(), INTERVAL 13-WEEKDAY(CURDATE()) DAY) AND tbtaskassignees.user_id = (SELECT id FROM tbadmins WHERE id = ?)", getid)
	if err != nil {
		panic(err.Error())
	}
	defer result3.Close()

	for result3.Next() {
		err := result3.Scan(&taskduenextweek)
		if err != nil {
			panic(err.Error())
		}
	}

	var taskdueover int
	//query data from database tbtask and tbtaskdue and tbtaskassignees by id and date due over in monday to sunday and count data from database to taskdueover variable
	result4, err := db.Query("SELECT COUNT(*) FROM tbtask JOIN tbtaskdue ON tbtask.id = tbtaskdue.task_id JOIN tbtaskassignees ON tbtask.id = tbtaskassignees.task_id WHERE tbtaskdue.date_due < CURDATE() AND tbtaskassignees.user_id = (SELECT id FROM tbadmins WHERE id = ?)", getid)
	if err != nil {
		panic(err.Error())
	}
	defer result4.Close()

	for result4.Next() {
		err := result4.Scan(&taskdueover)
		if err != nil {
			panic(err.Error())
		}
	}

	//return data to json
	json.NewEncoder(w).Encode(map[string]int{"taskduetoday": taskduetoday, "taskdueweek": taskdueweek, "taskduenextweek": taskduenextweek, "taskdueover": taskdueover})

}
