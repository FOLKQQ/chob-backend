package dashboardtaskcontroller

import (
	adminModel "backend/models/adminModel"
	"backend/models/companyModel"
	servicemodel "backend/models/serviceModel"
	"database/sql"
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func DashboardListTask(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	//get id from url path
	getid := r.URL.Path[len("/dashboardadmins/taskdue/"):]
	var taskduetoday int
	var taskdueweek int
	var taskduenextweek int
	var taskdueover int

	//count task.date_due = today from database tbtas join tbtaskassignees.task_id = task.id where id = getid
	reslut, err := db.Query("SELECT COUNT(tbtask.date_due) FROM tbtask JOIN tbtaskassignees ON tbtask.id = tbtaskassignees.task_id WHERE tbtask.date_due = CURDATE() AND tbtaskassignees.user_id = (SELECT id FROM tbadmins WHERE id = ?)", getid)
	if err != nil {
		panic(err.Error())
	}
	defer reslut.Close()

	for reslut.Next() {
		err := reslut.Scan(&taskduetoday)
		if err != nil {
			panic(err.Error())
		}
	}

	//count task.date_due = week from database tbtas join tbtaskassignees.task_id = task.id where id = getid
	reslut2, err := db.Query("SELECT COUNT(tbtask.date_due) FROM tbtask JOIN tbtaskassignees ON tbtask.id = tbtaskassignees.task_id WHERE tbtask.date_due BETWEEN DATE_SUB(CURDATE(), INTERVAL WEEKDAY(CURDATE()) DAY) AND DATE_ADD(CURDATE(), INTERVAL 6-WEEKDAY(CURDATE()) DAY) AND tbtaskassignees.user_id = (SELECT id FROM tbadmins WHERE id = ?)", getid)
	if err != nil {
		panic(err.Error())
	}
	defer reslut2.Close()

	for reslut2.Next() {
		err := reslut2.Scan(&taskdueweek)
		if err != nil {
			panic(err.Error())
		}
	}

	//count task.date_due = nextweek from database tbtas join tbtaskassignees.task_id = task.id where id = getid
	reslut3, err := db.Query("SELECT COUNT(tbtask.date_due) FROM tbtask JOIN tbtaskassignees ON tbtask.id = tbtaskassignees.task_id WHERE tbtask.date_due BETWEEN DATE_ADD(CURDATE(), INTERVAL 7-WEEKDAY(CURDATE()) DAY) AND DATE_ADD(CURDATE(), INTERVAL 13-WEEKDAY(CURDATE()) DAY) AND tbtaskassignees.user_id = (SELECT id FROM tbadmins WHERE id = ?)", getid)
	if err != nil {
		panic(err.Error())
	}
	defer reslut3.Close()

	for reslut3.Next() {
		err := reslut3.Scan(&taskduenextweek)
		if err != nil {
			panic(err.Error())
		}
	}

	//count task.date_due < today from database tbtas join tbtaskassignees.task_id = task.id where id = getid
	reslut4, err := db.Query("SELECT COUNT(tbtask.date_due) FROM tbtask JOIN tbtaskassignees ON tbtask.id = tbtaskassignees.task_id WHERE tbtask.date_due < CURDATE() AND tbtaskassignees.user_id = (SELECT id FROM tbadmins WHERE id = ?)", getid)
	if err != nil {
		panic(err.Error())
	}
	defer reslut4.Close()

	for reslut4.Next() {
		err := reslut4.Scan(&taskdueover)
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
	//query data from database and tbtaskassignees and tbcompany and tbtask one to many tbtag by id and count data from database to tasks variable
	result, err := db.Query("SELECT tbtask.id , tbcompany.company_name, tbtask.title, tbtaskassignees.user_id,tbadmins.image, tbtask.date_start, tbtask.date_due FROM tbtask JOIN tbtaskassignees ON tbtask.id = tbtaskassignees.task_id JOIN tbcompany ON tbtask.company_id = tbcompany.id JOIN tbadmins ON tbtaskassignees.user_id = tbadmins.id WHERE tbtaskassignees.user_id = (SELECT id FROM tbadmins WHERE id = ?) GROUP BY tbtask.id", getid)
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
		result, err := db.Query("SELECT tbtask.id , tbcompany.company_name, tbtask.title, tbtaskassignees.user_id,tbadmins.image, tbtask.date_start, tbtask.date_due FROM tbtask JOIN tbtaskassignees ON tbtask.id = tbtaskassignees.task_id JOIN tbcompany ON tbtask.company_id = tbcompany.id JOIN tbadmins ON tbtaskassignees.user_id = tbadmins.id WHERE tbtask.date_due = CURDATE() AND tbtaskassignees.user_id = (SELECT id FROM tbadmins WHERE id = ?) GROUP BY tbtask.id", getid)
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
		result, err := db.Query("SELECT tbtask.id , tbcompany.company_name, tbtask.title, tbtaskassignees.user_id,tbadmins.image, tbtask.date_start, tbtask.date_due FROM tbtask JOIN tbtaskassignees ON tbtask.id = tbtaskassignees.task_id JOIN tbcompany ON tbtask.company_id = tbcompany.id JOIN tbadmins ON tbtaskassignees.user_id = tbadmins.id WHERE tbtask.date_due BETWEEN DATE_SUB(CURDATE(), INTERVAL WEEKDAY(CURDATE()) DAY) AND DATE_ADD(CURDATE(), INTERVAL 6-WEEKDAY(CURDATE()) DAY) AND tbtaskassignees.user_id = (SELECT id FROM tbadmins WHERE id = ?) GROUP BY tbtask.id", getid)
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
		result, err := db.Query("SELECT tbtask.id , tbcompany.company_name, tbtask.title, tbtaskassignees.user_id,tbadmins.image, tbtask.date_start, tbtask.date_due FROM tbtask JOIN tbtaskassignees ON tbtask.id = tbtaskassignees.task_id JOIN tbcompany ON tbtask.company_id = tbcompany.id JOIN tbadmins ON tbtaskassignees.user_id = tbadmins.id WHERE tbtask.date_due BETWEEN DATE_ADD(CURDATE(), INTERVAL 7-WEEKDAY(CURDATE()) DAY) AND DATE_ADD(CURDATE(), INTERVAL 13-WEEKDAY(CURDATE()) DAY) AND tbtaskassignees.user_id = (SELECT id FROM tbadmins WHERE id = ?) GROUP BY tbtask.id", getid)
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
		result, err := db.Query("SELECT tbtask.id , tbcompany.company_name, tbtask.title, tbtaskassignees.user_id,tbadmins.image, tbtask.date_start, tbtask.date_due FROM tbtask JOIN tbtaskassignees ON tbtask.id = tbtaskassignees.task_id JOIN tbcompany ON tbtask.company_id = tbcompany.id JOIN tbadmins ON tbtaskassignees.user_id = tbadmins.id WHERE tbtask.date_due < CURDATE() AND tbtaskassignees.user_id = (SELECT id FROM tbadmins WHERE id = ?) GROUP BY tbtask.id", getid)
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

func Dashboardsubtask(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	//get id from url path
	getid := r.URL.Path[len("/dashboardadmins/subtask/"):]

	type assignees struct {
		Id      int `json:"id"`
		User_id int `json:"user_id"`
	}
	type Subtask struct {
		Id             int    `json:"id"`
		Title          string `json:"title"`
		Subtask_status string `json:"subtask_status"`
		Date_start     string `json:"date_start"`
		Date_due       string `json:"date_due"`
		Assignees      []assignees
	}

	var subtasks []Subtask
	//query data from database tbsubtask by id
	result, err := db.Query("SELECT tbsubtask.id, tbsubtask.title, tbsubtask.subtask_status ,tbsubtask.date_start , tbsubtask.date_due FROM tbsubtask WHERE tbsubtask.task_id = ?", getid)
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()

	for result.Next() {
		var subtask Subtask
		err := result.Scan(&subtask.Id, &subtask.Title, &subtask.Subtask_status, &subtask.Date_start, &subtask.Date_due)
		if err != nil {
			panic(err.Error())
		}
		subtasks = append(subtasks, subtask)
	}

	for i := 0; i < len(subtasks); i++ {
		result2, err := db.Query("SELECT tbtaskassignees.id, tbtaskassignees.user_id FROM tbtaskassignees WHERE tbtaskassignees.task_id = ?", subtasks[i].Id)
		if err != nil {
			panic(err.Error())
		}
		defer result2.Close()

		for result2.Next() {
			var assignee assignees
			err := result2.Scan(&assignee.Id, &assignee.User_id)
			if err != nil {
				panic(err.Error())
			}
			subtasks[i].Assignees = append(subtasks[i].Assignees, assignee)
		}
	}

	//encode data to json
	json.NewEncoder(w).Encode(subtasks)

}

func Newproject(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	//query tbrole and join tbadmins tbrole.title = 'worker'
	rows, err := db.Query("SELECT tbadmins.id ,tbadmins.firstname ,tbadmins.lastname FROM tbadmins JOIN tbrole ON tbadmins.role_id = tbrole.id WHERE tbrole.title = 'worker' ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// สร้าง slice เพื่อเก็บข้อมูลที่ค้นพบ
	results := []adminModel.Admin{}
	for rows.Next() {
		// สร้างตัวแปรเพื่อเก็บข้อมูลที่ query ค้นพบ
		admin := adminModel.Admin{}
		// สั่งสแกนข้อมูลจาก query ไปเก็บใน struct ตามชื่อฟิลด์
		err := rows.Scan(
			&admin.ID,
			&admin.Firstname,
			&admin.Lastname,
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// เก็บ struct ลงใน slice
		results = append(results, admin)
	}

	// ตรวจสอบข้อผิดพลาดหลังจากวนลูป
	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//query tbrole and join tbadmins tbrole.title = 'chacker'
	rows2, err := db.Query("SELECT tbadmins.id ,tbadmins.firstname ,tbadmins.lastname FROM tbadmins JOIN tbrole ON tbadmins.role_id = tbrole.id WHERE tbrole.title = 'chacker' ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows2.Close()

	// สร้าง slice เพื่อเก็บข้อมูลที่ค้นพบ
	results2 := []adminModel.Admin{}
	for rows2.Next() {
		// สร้างตัวแปรเพื่อเก็บข้อมูลที่ query ค้นพบ
		admin := adminModel.Admin{}
		// สั่งสแกนข้อมูลจาก query ไปเก็บใน struct ตามชื่อฟิลด์
		err := rows2.Scan(
			&admin.ID,
			&admin.Firstname,
			&admin.Lastname,
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// เก็บ struct ลงใน slice
		results2 = append(results2, admin)
	}

	// ตรวจสอบข้อผิดพลาดหลังจากวนลูป
	if err := rows2.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//query tbcompany
	rows3, err := db.Query("SELECT * FROM tbcompany ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows3.Close()

	// สร้าง slice เพื่อเก็บข้อมูลที่ค้นพบ
	results3 := []companyModel.Company{}
	for rows3.Next() {
		// สร้างตัวแปรเพื่อเก็บข้อมูลที่ query ค้นพบ
		company := companyModel.Company{}
		// สั่งสแกนข้อมูลจาก query ไปเก็บใน struct ตามชื่อฟิลด์
		err := rows3.Scan(
			&company.ID,
			&company.TypeCompany,
			&company.CompanyName,
			&company.CodeIdentification,
			&company.Signature,
			&company.StartComDate,
			&company.StartVatDate,
			&company.BusinessType,
			&company.Id_dbd,
			&company.Pass_dbd,
			&company.Id_filing,
			&company.Pass_filing,
			&company.Email,
			&company.Tal,
			&company.Address,
			&company.Addressextra,
			&company.Subdistrict,
			&company.District,
			&company.Province,
			&company.Zipcode,
			&company.Status,
			&company.Timestamps,
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// เก็บ struct ลงใน slice
		results3 = append(results3, company)
	}

	// ตรวจสอบข้อผิดพลาดหลังจากวนลูป
	if err := rows3.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//query tbservice_type
	rows4, err := db.Query("SELECT * FROM tbservice_type ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows4.Close()

	// สร้าง slice เพื่อเก็บข้อมูลที่ค้นพบ
	results4 := []servicemodel.Servicetype{}
	for rows4.Next() {
		// สร้างตัวแปรเพื่อเก็บข้อมูลที่ query ค้นพบ
		servicetype := servicemodel.Servicetype{}
		// สั่งสแกนข้อมูลจาก query ไปเก็บใน struct ตามชื่อฟิลด์
		err := rows4.Scan(
			&servicetype.Id,
			&servicetype.Service_name,
			&servicetype.Detail,
			&servicetype.Price,
			&servicetype.Timestamps,
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// เก็บ struct ลงใน slice
		results4 = append(results4, servicetype)
	}

	//create map for store data
	data := map[string]interface{}{
		"worker":  results,
		"chacker": results2,
		"company": results3,
		"service": results4,
	}

	//encode data to json
	json.NewEncoder(w).Encode(data)

}
