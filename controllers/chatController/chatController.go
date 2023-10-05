package chatcontroller

import (
	"backend/models/chatModel"
	"backend/models/companyModel"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func CreateChat_Team(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	chat := chatModel.Chat_team{}
	json.NewDecoder(r.Body).Decode(&chat)
	// สร้าง timestamp ในรูปแบบของ datetime
	timestamps := time.Now().Format("2006-01-02 15:04:05")
	// บันทึกข้อมูลผู้ใช้ในฐานข้อมูล
	_, err := db.Exec("INSERT INTO tbchat_team (user_id, company_id, comment, timestamps) VALUES (?, ?, ?, ?)",
		chat.User_id, chat.Company_id, chat.Comment, timestamps)
	if err != nil {
		fmt.Println(err)
	}

	// ส่งข้อความว่า "บันทึกข้อมูลเรียบร้อยแล้ว"
	fmt.Fprintf(w, "บันทึกข้อมูลเรียบร้อยแล้ว")
}

func ListChat_Team(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// getbody company_id
	company := companyModel.Company{}
	// สร้างตัวแปรเพื่อเก็บข้อมูลผู้ใช้ทั้งหมด
	listchat := []chatModel.Chat_team{}
	// ค้นหาข้อมูลผู้ใช้ทั้งหมดจากฐานข้อมูล
	rows, err := db.Query("SELECT * FROM tbchat_team wehere company_id = ?", company.ID)
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
	timestamps := time.Now().Format("2006-01-02 15:04:05")
	// อัพเดทข้อมูลผู้ใช้ในฐานข้อมูล
	_, err := db.Exec("UPDATE tbchat_team SET user_id = ?, company_id = ?, comment = ?, timestamps = ? WHERE id = ?",
		chat.User_id, chat.Company_id, chat.Comment, timestamps, chat.Id)
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
	chat := chatModel.Chat_task{}
	json.NewDecoder(r.Body).Decode(&chat)
	// สร้าง timestamp ในรูปแบบของ datetime
	timestamps := time.Now().Format("2006-01-02 15:04:05")
	// บันทึกข้อมูลผู้ใช้ในฐานข้อมูล
	_, err := db.Exec("INSERT INTO tbchat_task (task_id, user_id, comment, timestamps) VALUES (?, ?, ?, ?)",
		chat.Task_id, chat.User_id, chat.Comment, timestamps)
	if err != nil {
		fmt.Println(err)
	}

	// ส่งข้อความว่า "บันทึกข้อมูลเรียบร้อยแล้ว"
	fmt.Fprintf(w, "บันทึกข้อมูลเรียบร้อยแล้ว")
}

func ListChat_Task(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// getbody task_id
	// สร้างตัวแปรเพื่อเก็บข้อมูลผู้ใช้ทั้งหมด
	listchat := []chatModel.Chat_task{}
	// ค้นหาข้อมูลผู้ใช้ทั้งหมดจากฐานข้อมูล
	rows, err := db.Query("SELECT * FROM tbchat_task wehere task_id = ?")
	if err != nil {
		fmt.Println(err)
	}

	// วนลูปเพื่อเก็บข้อมูลผู้ใช้ทั้งหมดลงในตัวแปร listuser
	for rows.Next() {
		chat := chatModel.Chat_task{}
		err := rows.Scan(&chat.Task_id, &chat.User_id, &chat.Comment, &chat.Timestamps)
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
	timestamps := time.Now().Format("2006-01-02 15:04:05")
	// อัพเดทข้อมูลผู้ใช้ในฐานข้อมูล
	_, err := db.Exec("UPDATE tbchat_task SET task_id = ?, user_id = ?, comment = ?, timestamps = ? WHERE id = ?",
		chat.Task_id, chat.User_id, chat.Comment, timestamps, chat.Id)
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
