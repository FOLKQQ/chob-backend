package chatcontroller

import (
	"backend/models/chatModel"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func CreateChat_Team(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	chat := chatModel.Chat_team{}
	json.NewDecoder(r.Body).Decode(&chat)
	// สร้าง timestamp ในรูปแบบของ datetime
	// บันทึกข้อมูลผู้ใช้ในฐานข้อมูล
	_, err := db.Exec("INSERT INTO tbchat_team (user_id, company_id, comment) VALUES (?, ?, ?)",
		chat.User_id, chat.Company_id, chat.Comment)
	if err != nil {
		fmt.Println(err)
	}

	// ส่งข้อความว่า "บันทึกข้อมูลเรียบร้อยแล้ว"
	fmt.Fprintf(w, "บันทึกข้อมูลเรียบร้อยแล้ว")
}

func ListChat_Team(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// getbody company_id
	// สร้างตัวแปรเพื่อเก็บข้อมูลผู้ใช้ทั้งหมด
	listchat := []chatModel.Chat_team{}
	// ค้นหาข้อมูลผู้ใช้ทั้งหมดจากฐานข้อมูล
	rows, err := db.Query("SELECT * FROM tbchat_team ")
	if err != nil {
		fmt.Println(err)
	}

	// วนลูปเพื่อเก็บข้อมูลผู้ใช้ทั้งหมดลงในตัวแปร listuser
	for rows.Next() {
		chat := chatModel.Chat_team{}
		err := rows.Scan(&chat.Id, &chat.User_id, &chat.Company_id, &chat.Comment, &chat.Timestamps)
		if err != nil {
			fmt.Println(err)
		}
		listchat = append(listchat, chat)
	}

	// แสดงข้อมูลผู้ใช้ทั้งหมดในรูปแบบ JSON
	json.NewEncoder(w).Encode(listchat)
}

func UpdateChat_Team(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	chat := chatModel.Chat_team{}
	json.NewDecoder(r.Body).Decode(&chat)
	// สร้าง timestamp ในรูปแบบของ datetime
	// อัพเดทข้อมูลผู้ใช้ในฐานข้อมูล
	_, err := db.Exec("UPDATE tbchat_team SET user_id = ?, company_id = ?, comment = ? WHERE id = ?",
		chat.User_id, chat.Company_id, chat.Comment, chat.Id)
	if err != nil {
		fmt.Println(err)
	}

	// ส่งข้อความว่า "อัพเดทข้อมูลเรียบร้อยแล้ว"
	fmt.Fprintf(w, "อัพเดทข้อมูลเรียบร้อยแล้ว")
}

func DeleteChat_Team(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	chat := chatModel.Chat_team{}
	json.NewDecoder(r.Body).Decode(&chat)
	// ลบข้อมูลผู้ใช้ในฐานข้อมูล
	_, err := db.Exec("DELETE FROM tbchat_team WHERE id = ?", chat.Id)
	if err != nil {
		fmt.Println(err)
	}

	// ส่งข้อความว่า "ลบข้อมูลเรียบร้อยแล้ว"
	fmt.Fprintf(w, "ลบข้อมูลเรียบร้อยแล้ว")
}

func CreateChat_Task(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	//getbody request
	chat := chatModel.Chat_task_input{}
	json.NewDecoder(r.Body).Decode(&chat)
	fmt.Println(chat)
	// สร้าง timestamp ในรูปแบบของ datetime
	// บันทึกข้อมูลผู้ใช้ในฐานข้อมูล
	_, err := db.Exec("INSERT INTO tbchat_task (task_id, user_id, comment) VALUES (?, ?, ?)",
		chat.Task_id, chat.User_id, chat.Comment)
	if err != nil {
		fmt.Println(err)
	}

	// ส่งข้อความว่า "บันทึกข้อมูลเรียบร้อยแล้ว"
	fmt.Fprintf(w, "บันทึกข้อมูลเรียบร้อยแล้ว")

}

func ListChat_Task(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// สร้างตัวแปรเพื่อเก็บข้อมูลผู้ใช้ทั้งหมด
	listchat := []chatModel.Chat_task{}
	// ค้นหาข้อมูลผู้ใช้ทั้งหมดจากฐานข้อมูล
	rows, err := db.Query("SELECT * FROM tbchat_task")
	if err != nil {
		fmt.Println(err)
	}

	// วนลูปเพื่อเก็บข้อมูลผู้ใช้ทั้งหมดลงในตัวแปร listuser
	for rows.Next() {
		chat := chatModel.Chat_task{}
		err := rows.Scan(&chat.Id, &chat.Task_id, &chat.User_id, &chat.Comment, &chat.Timestamps)
		if err != nil {
			fmt.Println(err)
		}
		listchat = append(listchat, chat)
	}

	// แสดงข้อมูลผู้ใช้ทั้งหมดในรูปแบบ JSON
	json.NewEncoder(w).Encode(listchat)
}

func UpdateChat_Task(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	chat := chatModel.Chat_task{}
	json.NewDecoder(r.Body).Decode(&chat)
	// สร้าง timestamp ในรูปแบบของ datetime
	// อัพเดทข้อมูลผู้ใช้ในฐานข้อมูล
	_, err := db.Exec("UPDATE tbchat_task SET task_id = ?, user_id = ?, comment = ? WHERE id = ?",
		chat.Task_id, chat.User_id, chat.Comment, chat.Id)
	if err != nil {
		fmt.Println(err)
	}

	// ส่งข้อความว่า "อัพเดทข้อมูลเรียบร้อยแล้ว"
	fmt.Fprintf(w, "อัพเดทข้อมูลเรียบร้อยแล้ว")
}

func DeleteChat_Task(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	chat := chatModel.Chat_task{}
	json.NewDecoder(r.Body).Decode(&chat)
	// ลบข้อมูลผู้ใช้ในฐานข้อมูล
	_, err := db.Exec("DELETE FROM tbchat_task WHERE id = ?", chat.Id)
	if err != nil {
		fmt.Println(err)
	}

	// ส่งข้อความว่า "ลบข้อมูลเรียบร้อยแล้ว"
	fmt.Fprintf(w, "ลบข้อมูลเรียบร้อยแล้ว")
}

func ChatByID(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// Parse the request body to get the company_id
	type RequestBody struct {
		Company_id int `json:"company_id"`
	}
	var reqBody RequestBody
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Query the database for the chat teams with the given company_id
	rows, err := db.Query("SELECT * FROM tbchat_team WHERE company_id = ?", reqBody.Company_id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Create a slice to hold the chat teams
	var chatTeams []chatModel.Chat_team

	// Iterate over the rows and append the chat teams to the slice
	for rows.Next() {
		var chatTeam chatModel.Chat_team
		err := rows.Scan(&chatTeam.Id, &chatTeam.User_id, &chatTeam.Company_id, &chatTeam.Comment, &chatTeam.Timestamps)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		chatTeams = append(chatTeams, chatTeam)
	}
	err = rows.Err()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Encode the chat teams slice as JSON and send it in the response
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(chatTeams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func Chat_TaskByID(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	//getid from url path
	getid := r.URL.Path[len("/chattask/"):]
	// Query the database for the chat tasks with the given task_id
	rows, err := db.Query("SELECT tbchat_task.id ,tbchat_task.user_id, tbchat_task.comment, tbchat_task.timestamp ,tbadmins.username FROM tbchat_task JOIN tbadmins ON tbadmins.id = tbchat_task.user_id  WHERE tbchat_task.task_id = ?", getid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Create a slice to hold the chat tasks
	type ChatTask struct {
		Id         int    `json:"id"`
		User_id    int    `json:"user_id"`
		Comment    string `json:"comment"`
		Timestamps string `json:"timestamps"`
		Username   string `json:"username"`
	}
	var chatTasks []ChatTask

	// Iterate over the rows and append the chat tasks to the slice
	for rows.Next() {
		var chatTask ChatTask
		err := rows.Scan(&chatTask.Id, &chatTask.User_id, &chatTask.Comment, &chatTask.Timestamps, &chatTask.Username)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		chatTasks = append(chatTasks, chatTask)
	}
	err = rows.Err()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Encode the chat tasks slice as JSON and send it in the response
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(chatTasks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
