package pstagcontrollers

import (
	pstagmodel "backend/models/pstagModel"
	"database/sql"
	"encoding/json"
	"net/http"
)

func Listpstag(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// ดำเนินการค้นหาข้อมูลทั้งหมดจากฐานข้อมูล
	rows, err := db.Query("SELECT * FROM tbpstag")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// สร้าง slice เพื่อเก็บข้อมูลที่ค้นพบ
	results := []pstagmodel.Pstag{}
	//
	for rows.Next() {
		// สร้างตัวแปรเพื่อเก็บข้อมูลที่ query ค้นพบ
		pstag := pstagmodel.Pstag{}
		// สั่งสแกนข้อมูลจาก query ไปเก็บใน struct ตามชื่อฟิลด์
		err := rows.Scan(
			&pstag.Id,
			&pstag.Name,
			&pstag.Permis_tag,
			&pstag.Status,
			&pstag.Timestamps,
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// เก็บ struct ลงใน slice
		results = append(results, pstag)
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

func Addpstag(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// get body from client request and decode to user struct
	pstag := pstagmodel.Pstag{}
	err := json.NewDecoder(r.Body).Decode(&pstag)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// ดำเนินการเพิ่มข้อมูลลงในฐานข้อมูล
	stmt, err := db.Prepare("INSERT INTO tbpstag (name,permis_tag,status) VALUES (?,?,?)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// ปิดการเชื่อมต่อกับฐานข้อมูล
	defer stmt.Close()

	// สั่งให้เริ่มทำการเพิ่มข้อมูลลงในฐานข้อมูล
	_, err = stmt.Exec(pstag.Name, pstag.Permis_tag, pstag.Status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// return data to client
	Listpstag(w, r, db)

}

func Updatepstag(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// get body from client request and decode to user struct
	pstag := pstagmodel.Pstag{}
	err := json.NewDecoder(r.Body).Decode(&pstag)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// ดำเนินการแก้ไขข้อมูลลงในฐานข้อมูล
	stmt, err := db.Prepare("UPDATE tbpstag SET name=?,permis_tag=?,status=? WHERE id=?")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// ปิดการเชื่อมต่อกับฐานข้อมูล
	defer stmt.Close()

	// สั่งให้เริ่มทำการแก้ไขข้อมูลลงในฐานข้อมูล
	_, err = stmt.Exec(pstag.Name, pstag.Permis_tag, pstag.Status, pstag.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// return data to client
	Listpstag(w, r, db)

}

func Deletepstag(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// get body from client request and decode to user struct
	pstag := pstagmodel.Pstag{}
	err := json.NewDecoder(r.Body).Decode(&pstag)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// delete data to database
	stmt, err := db.Prepare("DELETE FROM tbpermis_tag WHERE id=?")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// ปิดการเชื่อมต่อกับฐานข้อมูล
	defer stmt.Close()

	// สั่งให้เริ่มทำการลบข้อมูลลงในฐานข้อมูล
	_, err = stmt.Exec(pstag.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// return data to client
	Listpstag(w, r, db)
}
