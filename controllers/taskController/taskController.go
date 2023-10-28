package taskController

import (
	"backend/models/taskModel"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func CreateTask(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// ดึงข้อมูลจากฟอร์ม และ แปลงเป็น struct
	task := taskModel.Taskadd{}
	json.NewDecoder(r.Body).Decode(&task)
	//how to format task.date_start and task.date_due to dd/mm/yyyy ใช้ได้แค่ตอนเริ่มต้น ถ้าเปลี่ยนวันที่จะต้องเปลี่ยนในฟังก์ชันนี้ด้วย
	var tasklist_status = "active"
	// บันทึกข้อมูลผู้ใช้ในฐานข้อมูล
	_, err := db.Exec("INSERT INTO tbtask (company_id,title,date_start,date_due,repeatwork,cyclemonth,tasklist_status) VALUES (?, ?, ?, ?, ?)", task.CompanyName, task.ProjectName, task.StartDate, task.EndDate, task.Repeat, task.Cyclemonth, tasklist_status)
	if err != nil {
		panic(err.Error())
	}
	//get id task ที่เพิ่งสร้าง จากนั้นเอาไปใช้ในการสร้าง taskassignees , tbchacker_work ,tbservice  ต่อไป
	var task_id int
	result, err := db.Query("SELECT id FROM tbtask ORDER BY id DESC LIMIT 1")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {
		err := result.Scan(&task_id)
		if err != nil {
			panic(err.Error())
		}
	}
	// สร้าง taskassignees
	for _, worker := range task.Worker {
		_, err := db.Exec("INSERT INTO tbtaskassignees (task_id, user_id) VALUES (?, ?)", task_id, worker)
		if err != nil {
			panic(err.Error())
		}
	}
	// สร้าง tbchacker_work
	for _, checker := range task.Checker {
		_, err := db.Exec("INSERT INTO tbchacker_work (task_id, user_id) VALUES (?, ?)", task_id, checker)
		if err != nil {
			panic(err.Error())
		}
	}
	// สร้าง tbservice
	for _, service := range task.Service {
		_, err := db.Exec("INSERT INTO tbservice (servicetype_id, task_id) VALUES (?, ?)", service, task_id)
		if err != nil {
			panic(err.Error())
		}
	}
	//return status 200
	w.WriteHeader(http.StatusOK)

}

func ListTask(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var tasks []taskModel.Task
	result, err := db.Query("SELECT * FROM tbtask")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {
		var task taskModel.Task
		err := result.Scan(&task.Id, &task.Company_id, &task.Title, &task.Tax_status, &task.Tasklist_status, &task.Timestamps)
		if err != nil {
			panic(err.Error())
		}
		tasks = append(tasks, task)
	}
	json.NewEncoder(w).Encode(tasks)
}

func GetTask(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	task := taskModel.Task{}
	getid := r.URL.Path[len("/task/"):]

	result, err := db.Query("SELECT * FROM tbtask WHERE id=?", getid)
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {
		err := result.Scan(&task.Id, &task.Company_id, &task.Title, &task.Tax_status, &task.Tasklist_status, &task.Timestamps)
		if err != nil {
			panic(err.Error())
		}
	}
	json.NewEncoder(w).Encode(task)
}

func UpdateTask(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	task := taskModel.Task{}
	json.NewDecoder(r.Body).Decode(&task)
	// บันทึกข้อมูลผู้ใช้ในฐานข้อมูล
	_, err := db.Exec("UPDATE tbtask SET company_id=?, title=?, tax_status=?, tasklist_status=? WHERE id=?", task.Company_id, task.Title, task.Tax_status, task.Tasklist_status, task.Id)
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "Task was updated")
}

func DeleteTask(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	task := taskModel.Task{}
	json.NewDecoder(r.Body).Decode(&task)
	// ลบข้อมูลผู้ใช้ในฐานข้อมูล
	_, err := db.Exec("DELETE FROM tbtask WHERE id=?", task.Id)
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "Task was deleted")
}

func CreateSubtask(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	type Addsubtask struct {
		Task_id string `json:"task_id"`
		Title   string `json:"title"`
	}
	subtask := Addsubtask{}
	json.NewDecoder(r.Body).Decode(&subtask)
	fmt.Println(subtask)
	// บันทึกข้อมูลผู้ใช้ในฐานข้อมูล
	_, err := db.Exec("INSERT INTO tbsubtask (task_id, title, subtask_status) VALUES (?, ?, ?)", subtask.Task_id, subtask.Title, "active")
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode(subtask)
}

func ListSubtask(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var subtasks []taskModel.Subtask
	result, err := db.Query("SELECT * FROM tbsubtask")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {
		var subtask taskModel.Subtask
		err := result.Scan(&subtask.Id, &subtask.Task_id, &subtask.Title, &subtask.Subtask_status, &subtask.Timestamps)
		if err != nil {
			panic(err.Error())
		}
		subtasks = append(subtasks, subtask)
	}

	w.WriteHeader(http.StatusOK)

}

func GetSubtask(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	subtask := taskModel.Subtask{}
	//how to get url /subtask/id
	getid := r.URL.Path[len("/subtask/"):]
	fmt.Println(getid)

	result, err := db.Query("SELECT * FROM tbsubtask WHERE id=?", getid)
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {
		err := result.Scan(&subtask.Id, &subtask.Task_id, &subtask.Title, &subtask.Subtask_status, &subtask.Timestamps)
		if err != nil {
			panic(err.Error())
		}
	}
	json.NewEncoder(w).Encode(subtask)
}

func UpdateSubtask(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	subtask := taskModel.Subtask{}
	json.NewDecoder(r.Body).Decode(&subtask)
	// บันทึกข้อมูลผู้ใช้ในฐานข้อมูล
	_, err := db.Exec("UPDATE tbsubtask SET task_id=?, title=?, subtask_status=? WHERE id=?", subtask.Task_id, subtask.Title, subtask.Subtask_status, subtask.Id)
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "Subtask was updated")
}

func DeleteSubtask(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	subtask := taskModel.Subtask{}
	json.NewDecoder(r.Body).Decode(&subtask)
	fmt.Println(subtask.Id)
	// ลบข้อมูลผู้ใช้ในฐานข้อมูล
	_, err := db.Exec("DELETE FROM tbsubtask WHERE id=?", subtask.Id)
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "Subtask was deleted")
}

func CreateTaskassignees(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	taskassignees := taskModel.Taskassignees{}
	json.NewDecoder(r.Body).Decode(&taskassignees)
	// บันทึกข้อมูลผู้ใช้ในฐานข้อมูล
	_, err := db.Exec("INSERT INTO tbtaskassignees (task_id, user_id) VALUES (?, ?, ?)", taskassignees.Task_id, taskassignees.User_id)
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "New taskassignees was created")
}

func ListTaskassignees(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var taskassignees []taskModel.Taskassignees
	result, err := db.Query("SELECT * FROM tbtaskassignees")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {
		var taskassignee taskModel.Taskassignees
		err := result.Scan(&taskassignee.Id, &taskassignee.Task_id, &taskassignee.User_id, &taskassignee.Timestamps)
		if err != nil {
			panic(err.Error())
		}
		taskassignees = append(taskassignees, taskassignee)
	}
	json.NewEncoder(w).Encode(taskassignees)
}

func GetTaskassignees(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	taskassignees := taskModel.Taskassignees{}
	getid := r.URL.Path[len("/taskassignees/"):]

	result, err := db.Query("SELECT * FROM tbtaskassignees WHERE id=?", getid)
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {
		err := result.Scan(&taskassignees.Id, &taskassignees.Task_id, &taskassignees.User_id, &taskassignees.Timestamps)
		if err != nil {
			panic(err.Error())
		}
	}
	json.NewEncoder(w).Encode(taskassignees)
}

func UpdateTaskassignees(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	taskassignees := taskModel.Taskassignees{}
	json.NewDecoder(r.Body).Decode(&taskassignees)
	// บันทึกข้อมูลผู้ใช้ในฐานข้อมูล
	_, err := db.Exec("UPDATE tbtaskassignees SET task_id=?, user_id=? WHERE id=?", taskassignees.Task_id, taskassignees.User_id, taskassignees.Id)
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "Taskassignees was updated")
}

func DeleteTaskassignees(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	taskassignees := taskModel.Taskassignees{}
	json.NewDecoder(r.Body).Decode(&taskassignees)
	// ลบข้อมูลผู้ใช้ในฐานข้อมูล
	_, err := db.Exec("DELETE FROM tbtaskassignees WHERE id=?", taskassignees.Id)
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "Taskassignees was deleted")
}
