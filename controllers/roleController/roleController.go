package rolecontrollers

import (
	"backend/models/roleModel"
	"database/sql"
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func Listroles(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// ดำเนินการค้นหาข้อมูลทั้งหมดจากฐานข้อมูล
	rows, err := db.Query("SELECT * FROM tbrole")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// สร้าง slice เพื่อเก็บข้อมูลที่ค้นพบ
	results := []roleModel.Role{}
	//
	for rows.Next() {
		// สร้างตัวแปรเพื่อเก็บข้อมูลที่ query ค้นพบ
		role := roleModel.Role{}
		// สั่งสแกนข้อมูลจาก query ไปเก็บใน struct ตามชื่อฟิลด์
		err := rows.Scan(
			&role.ID,
			&role.Title,
			&role.Project,
			&role.Managerrole,
			&role.Addtags,
			&role.Report,
			&role.Timestamps,
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// เก็บ struct ลงใน slice
		results = append(results, role)
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

func Addroles(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// ดึงข้อมูลจากฟอร์ม
	role := roleModel.Role{}
	err := json.NewDecoder(r.Body).Decode(&role)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// สร้างคำสั่ง SQL
	stmt, err := db.Prepare("INSERT INTO tbrole (title, project, manager_role, addtags, report) VALUES (?,?,?,?,?)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	// สั่งเรียกใช้งานคำสั่ง SQL
	result, err := stmt.Exec(role.Title, role.Project, role.Managerrole, role.Addtags, role.Report)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// ตรวจสอบผลลัพธ์
	if _, err := result.LastInsertId(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Insert Success"))
}

func UpdateRoles(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// ดึงข้อมูลจากฟอร์ม
	role := roleModel.Role{}
	err := json.NewDecoder(r.Body).Decode(&role)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// สร้างคำสั่ง SQL
	stmt, err := db.Prepare("UPDATE tbrole SET title=?, project=?, manager_role=?, addtags=?, report=? WHERE id=?")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	// สั่งเรียกใช้งานคำสั่ง SQL
	result, err := stmt.Exec(role.Title, role.Project, role.Managerrole, role.Addtags, role.Report, role.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// ตรวจสอบผลลัพธ์
	if _, err := result.RowsAffected(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Update Success"))
}

func DeleteRoles(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// ดึงข้อมูลจากฟอร์ม
	role := roleModel.Role{}
	err := json.NewDecoder(r.Body).Decode(&role)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// สร้างคำสั่ง SQL
	stmt, err := db.Prepare("DELETE FROM tbrole WHERE id=?")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	// สั่งเรียกใช้งานคำสั่ง SQL
	result, err := stmt.Exec(role.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// ตรวจสอบผลลัพธ์
	if _, err := result.RowsAffected(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Delete Success"))
}
