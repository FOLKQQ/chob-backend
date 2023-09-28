package teamcontroller

import (
	teammodel "backend/models/teamModel"
	"database/sql"
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func Listteams(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// ดำเนินการค้นหาข้อมูลทั้งหมดจากฐานข้อมูล
	rows, err := db.Query("SELECT * FROM tbteam")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// สร้าง slice เพื่อเก็บข้อมูลที่ค้นพบ
	results := []teammodel.Team{}
	//
	for rows.Next() {
		// สร้างตัวแปรเพื่อเก็บข้อมูลที่ query ค้นพบ
		team := teammodel.Team{}
		// สั่งสแกนข้อมูลจาก query ไปเก็บใน struct ตามชื่อฟิลด์
		err := rows.Scan(
			&team.Id,
			&team.Name,
			&team.Status,
			&team.Timestamps,
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// เก็บ struct ลงใน slice
		results = append(results, team)
	}

	// ตรวจสอบข้อผิดพลาดหลังจากวนลูป
	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// แสดงข้อมูลในรูปแบบ JSON
	jsonData, err := json.Marshal(results)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}

func Addteams(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// ดึงข้อมูลจากฟอร์ม
	team := teammodel.Team{}
	err := json.NewDecoder(r.Body).Decode(&team)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// ดำเนินการเพิ่มข้อมูลลงในฐานข้อมูล
	_, err = db.Exec("INSERT INTO tbteam (name,status) VALUES (?,?)",
		team.Name,
		team.Status,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Insert success"))
}

func Updateteams(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// ดึงข้อมูลจากฟอร์ม
	team := teammodel.Team{}
	err := json.NewDecoder(r.Body).Decode(&team)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// ดำเนินการแก้ไขข้อมูลลงในฐานข้อมูล
	_, err = db.Exec("UPDATE tbteam SET name=?,status=? WHERE id=?",
		team.Name,
		team.Status,
		team.Id,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Update success"))
}

func Deleteteams(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// ดึงข้อมูลจากฟอร์ม
	team := teammodel.Team{}
	err := json.NewDecoder(r.Body).Decode(&team)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// ดำเนินการลบข้อมูลออกจากฐานข้อมูล
	_, err = db.Exec("DELETE FROM tbteam WHERE id=?", team.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Delete success"))
}
