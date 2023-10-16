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

	type Tag struct {
		Name  string `json:"name"`
		Color string `json:"color"`
	}

	type Task struct {
		Id           int    `json:"id"`
		Company_name string `json:"company_name"`
		Title        string `json:"title"`
		UserId       string `json:"userId"`
		Image        string `json:"image"`
		DateStart    string `json:"dateStart"`
		DateDue      string `json:"dateDue"`
		Tags         []Tag
	}

	var tasks []Task
	//query data from database tbtask and tbtaskdue and tbtaskassignees and tbcompany and tbadmins select image where id = getid and tbtask one to many tbtag by id and count data from database to tasks variable
	result5, err := db.Query("SELECT tbtask.id, tbcompany.company_name, tbtask.title, tbtaskassignees.user_id, tbadmins.image, tbtaskdue.date_start, tbtaskdue.date_due FROM tbtask JOIN tbtaskdue ON tbtask.id = tbtaskdue.task_id JOIN tbtaskassignees ON tbtask.id = tbtaskassignees.task_id JOIN tbcompany ON tbtask.company_id = tbcompany.id JOIN tbadmins ON tbtaskassignees.user_id = tbadmins.id WHERE tbtaskassignees.user_id = (SELECT id FROM tbadmins WHERE id = ?) GROUP BY tbtask.id", getid)
	if err != nil {
		panic(err.Error())
	}
	defer result5.Close()

	for result5.Next() {
		var task Task
		err := result5.Scan(&task.Id, &task.Company_name, &task.Title, &task.UserId, &task.Image, &task.DateStart, &task.DateDue)
		if err != nil {
			panic(err.Error())
		}
		tasks = append(tasks, task)
	}

	for i := 0; i < len(tasks); i++ {
		result6, err := db.Query("SELECT tbtag.name, tbtag.color FROM tbtag WHERE tbtag.task_id = ?", tasks[i].Id)
		if err != nil {
			panic(err.Error())
		}
		defer result6.Close()

		for result6.Next() {
			var tag Tag
			err := result6.Scan(&tag.Name, &tag.Color)
			if err != nil {
				panic(err.Error())
			}
			tasks[i].Tags = append(tasks[i].Tags, tag)
		}
	}

	//create map for store data
	data := map[string]interface{}{
		"taskduetoday":    taskduetoday,
		"taskdueweek":     taskdueweek,
		"taskduenextweek": taskduenextweek,
		"taskdueover":     taskdueover,
		"tasks":           tasks,
	}

	//encode data to json
	json.NewEncoder(w).Encode(data)

}

func DashboardListSelectResDue(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	//get url path
	parts := r.URL.Path[len("/dashboardadmins/selectresdue/")-1:]
	//get id,name from url path
	getid := parts[0:1]
	getname := parts[1:]
	//getname delete / from string
	getname = getname[1:]

	type Tag struct {
		Name  string `json:"name"`
		Color string `json:"color"`
	}

	type Task struct {
		Id           int    `json:"id"`
		Company_name string `json:"company_name"`
		Title        string `json:"title"`
		UserId       string `json:"userId"`
		Image        string `json:"image"`
		DateStart    string `json:"dateStart"`
		DateDue      string `json:"dateDue"`
		Tags         []Tag
	}
	if getname == "resduetoday" {
		var tasks []Task
		//query data from database tbtask and tbtaskdue and tbtaskassignees and tbcompany and tbtask one to many tbtag by id and count data from database to tasks variable
		result, err := db.Query("SELECT tbtask.id , tbcompany.company_name, tbtask.title, tbtaskassignees.user_id,tbadmins.image, tbtaskdue.date_start, tbtaskdue.date_due FROM tbtask JOIN tbtaskdue ON tbtask.id = tbtaskdue.task_id JOIN tbtaskassignees ON tbtask.id = tbtaskassignees.task_id JOIN tbcompany ON tbtask.company_id = tbcompany.id JOIN tbadmins ON tbtaskassignees.user_id = tbadmins.id WHERE tbtaskdue.date_due = CURDATE() AND tbtaskassignees.user_id = (SELECT id FROM tbadmins WHERE id = ?) GROUP BY tbtask.id", getid)
		if err != nil {
			panic(err.Error())
		}
		defer result.Close()

		for result.Next() {
			var task Task
			err := result.Scan(&task.Id, &task.Company_name, &task.Title, &task.UserId, &task.Image, &task.DateStart, &task.DateDue)
			if err != nil {
				panic(err.Error())
			}
			tasks = append(tasks, task)
		}

		for i := 0; i < len(tasks); i++ {
			result2, err := db.Query("SELECT tbtag.name, tbtag.color FROM tbtag WHERE tbtag.task_id = ?", tasks[i].Id)
			if err != nil {
				panic(err.Error())
			}
			defer result2.Close()

			for result2.Next() {
				var tag Tag
				err := result2.Scan(&tag.Name, &tag.Color)
				if err != nil {
					panic(err.Error())
				}
				tasks[i].Tags = append(tasks[i].Tags, tag)
			}
		}

		//create map for store data
		data := map[string]interface{}{
			"tasks": tasks,
		}

		//encode data to json
		json.NewEncoder(w).Encode(data)
	} else if getname == "resdueweek" {
		var tasks []Task
		//query data from database tbtask and tbtaskdue and tbtaskassignees and tbcompany and tbtask one to many tbtag by id and count data from database to tasks variable
		result, err := db.Query("SELECT tbtask.id , tbcompany.company_name, tbtask.title, tbtaskassignees.user_id,tbadmins.image, tbtaskdue.date_start, tbtaskdue.date_due FROM tbtask JOIN tbtaskdue ON tbtask.id = tbtaskdue.task_id JOIN tbtaskassignees ON tbtask.id = tbtaskassignees.task_id JOIN tbcompany ON tbtask.company_id = tbcompany.id JOIN tbadmins ON tbtaskassignees.user_id = tbadmins.id WHERE tbtaskdue.date_due BETWEEN DATE_SUB(CURDATE(), INTERVAL WEEKDAY(CURDATE()) DAY) AND DATE_ADD(CURDATE(), INTERVAL 6-WEEKDAY(CURDATE()) DAY) AND tbtaskassignees.user_id = (SELECT id FROM tbadmins WHERE id = ?) GROUP BY tbtask.id", getid)
		if err != nil {
			panic(err.Error())
		}
		defer result.Close()

		for result.Next() {
			var task Task
			err := result.Scan(&task.Id, &task.Company_name, &task.Title, &task.UserId, &task.Image, &task.DateStart, &task.DateDue)
			if err != nil {
				panic(err.Error())
			}
			tasks = append(tasks, task)
		}

		for i := 0; i < len(tasks); i++ {
			result2, err := db.Query("SELECT tbtag.name, tbtag.color FROM tbtag WHERE tbtag.task_id = ?", tasks[i].Id)
			if err != nil {
				panic(err.Error())
			}
			defer result2.Close()

			for result2.Next() {
				var tag Tag
				err := result2.Scan(&tag.Name, &tag.Color)
				if err != nil {
					panic(err.Error())
				}
				tasks[i].Tags = append(tasks[i].Tags, tag)
			}
		}
		//create map for store data
		data := map[string]interface{}{
			"tasks": tasks,
		}

		//encode data to json
		json.NewEncoder(w).Encode(data)
	} else if getname == "resduenextweek" {
		var tasks []Task
		//query data from database tbtask and tbtaskdue and tbtaskassignees and tbcompany and tbtask one to many tbtag by id and count data from database to tasks variable
		result, err := db.Query("SELECT tbtask.id , tbcompany.company_name, tbtask.title, tbtaskassignees.user_id,tbadmins.image, tbtaskdue.date_start, tbtaskdue.date_due FROM tbtask JOIN tbtaskdue ON tbtask.id = tbtaskdue.task_id JOIN tbtaskassignees ON tbtask.id = tbtaskassignees.task_id JOIN tbcompany ON tbtask.company_id = tbcompany.id JOIN tbadmins ON tbtaskassignees.user_id = tbadmins.id WHERE tbtaskdue.date_due BETWEEN DATE_ADD(CURDATE(), INTERVAL 7-WEEKDAY(CURDATE()) DAY) AND DATE_ADD(CURDATE(), INTERVAL 13-WEEKDAY(CURDATE()) DAY) AND tbtaskassignees.user_id = (SELECT id FROM tbadmins WHERE id = ?) GROUP BY tbtask.id", getid)
		if err != nil {
			panic(err.Error())
		}
		defer result.Close()

		for result.Next() {
			var task Task
			err := result.Scan(&task.Id, &task.Company_name, &task.Title, &task.UserId, &task.Image, &task.DateStart, &task.DateDue)
			if err != nil {
				panic(err.Error())
			}
			tasks = append(tasks, task)
		}

		for i := 0; i < len(tasks); i++ {
			result2, err := db.Query("SELECT tbtag.name, tbtag.color FROM tbtag WHERE tbtag.task_id = ?", tasks[i].Id)
			if err != nil {
				panic(err.Error())
			}
			defer result2.Close()

			for result2.Next() {
				var tag Tag
				err := result2.Scan(&tag.Name, &tag.Color)
				if err != nil {
					panic(err.Error())
				}
				tasks[i].Tags = append(tasks[i].Tags, tag)
			}
		}
		//create map for store data
		data := map[string]interface{}{
			"tasks": tasks,
		}

		//encode data to json
		json.NewEncoder(w).Encode(data)
	} else if getname == "resdueover" {
		var tasks []Task
		//query data from database tbtask and tbtaskdue and tbtaskassignees and tbcompany and tbtask one to many tbtag by id and count data from database to tasks variable
		result, err := db.Query("SELECT tbtask.id , tbcompany.company_name, tbtask.title, tbtaskassignees.user_id,tbadmins.image, tbtaskdue.date_start, tbtaskdue.date_due FROM tbtask JOIN tbtaskdue ON tbtask.id = tbtaskdue.task_id JOIN tbtaskassignees ON tbtask.id = tbtaskassignees.task_id JOIN tbcompany ON tbtask.company_id = tbcompany.id JOIN tbadmins ON tbtaskassignees.user_id = tbadmins.id WHERE tbtaskdue.date_due < CURDATE() AND tbtaskassignees.user_id = (SELECT id FROM tbadmins WHERE id = ?) GROUP BY tbtask.id", getid)
		if err != nil {
			panic(err.Error())
		}
		defer result.Close()

		for result.Next() {
			var task Task
			err := result.Scan(&task.Id, &task.Company_name, &task.Title, &task.UserId, &task.Image, &task.DateStart, &task.DateDue)
			if err != nil {
				panic(err.Error())
			}
			tasks = append(tasks, task)
		}

		for i := 0; i < len(tasks); i++ {
			result2, err := db.Query("SELECT tbtag.name, tbtag.color FROM tbtag WHERE tbtag.task_id = ?", tasks[i].Id)
			if err != nil {
				panic(err.Error())
			}
			defer result2.Close()

			for result2.Next() {
				var tag Tag
				err := result2.Scan(&tag.Name, &tag.Color)
				if err != nil {
					panic(err.Error())
				}
				tasks[i].Tags = append(tasks[i].Tags, tag)
			}
		}
		//create map for store data
		data := map[string]interface{}{
			"tasks": tasks,
		}

		//encode data to json
		json.NewEncoder(w).Encode(data)
	}

}
