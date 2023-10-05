package servicecontroller

import (
	"backend/models/serviceModel"
	"database/sql"
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func Addservicetype(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	// get json body from request and decode to user struct
	servicetype := servicemodel.Servicetype{}
	err := json.NewDecoder(r.Body).Decode(&servicetype)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// insert data to database
	stmt, err := db.Prepare("INSERT INTO tbservice_type (service_name, detail, price) VALUES (?,?,?)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// ปิดการเชื่อมต่อกับฐานข้อมูล
	defer stmt.Close()

	// สั่งให้เริ่มทำการเพิ่มข้อมูลลงในฐานข้อมูล
	_, err = stmt.Exec(servicetype.Service_name, servicetype.Detail, servicetype.Price)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// return data to client
	Listservicetype(w, r, db)

}

func Listservicetype(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// ดำเนินการค้นหาข้อมูลจากฐานข้อมูล
	rows, err := db.Query("SELECT * FROM tbservice_type")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// สร้างตัวแปรเพื่อเก็บข้อมูลที่ query ค้นพบ
	servicetyperesults := []servicemodel.Servicetype{}
	// เก็บข้อมูลที่ query ค้นพบลงในตัวแปร
	for rows.Next() {
		// สร้างตัวแปรเพื่อเก็บข้อมูลที่ query ค้นพบ
		servicetype := servicemodel.Servicetype{}
		// สั่งสแกนข้อมูลจาก query ไปเก็บใน struct ตามชื่อฟิลด์
		err := rows.Scan(
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
		servicetyperesults = append(servicetyperesults, servicetype)
	}

	// ตรวจสอบข้อผิดพลาดหลังจากวนลูป
	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// return data to client
	json.NewEncoder(w).Encode(servicetyperesults)
}

func Updateservicetype(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// get json body from request and decode to user struct
	servicetype := servicemodel.Servicetype{}
	err := json.NewDecoder(r.Body).Decode(&servicetype)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// update data to database
	stmt, err := db.Prepare("UPDATE tbservice_type SET service_name=?, detail=?, price=? WHERE id=?")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// ปิดการเชื่อมต่อกับฐานข้อมูล
	defer stmt.Close()

	// สั่งให้เริ่มทำการแก้ไขข้อมูลลงในฐานข้อมูล
	_, err = stmt.Exec(servicetype.Service_name, servicetype.Detail, servicetype.Price, servicetype.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// return data to client
	Listservicetype(w, r, db)
}

func Deleteservicetype(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// get json body from request and decode to user struct
	servicetype := servicemodel.Servicetype{}
	err := json.NewDecoder(r.Body).Decode(&servicetype)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// delete data to database
	stmt, err := db.Prepare("DELETE FROM tbservice_type WHERE id=?")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// ปิดการเชื่อมต่อกับฐานข้อมูล
	defer stmt.Close()

	// สั่งให้เริ่มทำการลบข้อมูลลงในฐานข้อมูล
	_, err = stmt.Exec(servicetype.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// return data fuce Listservicetype to client
	Listservicetype(w, r, db)
}

func Listservice(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// ดำเนินการค้นหาข้อมูลจากฐานข้อมูล
	rows, err := db.Query("SELECT * FROM tbservice")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// สร้างตัวแปรเพื่อเก็บข้อมูลที่ query ค้นพบ
	serviceresults := []servicemodel.Service{}
	// วนลูปอ่านข้อมูลทีละแถว
	for rows.Next() {
		// สร้างตัวแปรเพื่อเก็บข้อมูลที่ query ค้นพบ
		service := servicemodel.Service{}
		// สั่งสแกนข้อมูลจาก query ไปเก็บใน struct ตามชื่อฟิลด์
		err := rows.Scan(
			&service.Id,
			&service.Servicetype_id,
			&service.Company_id,
			&service.Date_start,
			&service.Date_due,
			&service.Timestamps,
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// เก็บ struct ลงใน slice
		serviceresults = append(serviceresults, service)
	}

	// ตรวจสอบข้อผิดพลาดหลังจากวนลูป
	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// return data to client
	json.NewEncoder(w).Encode(serviceresults)
}

func Addservice(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// get json body from request and decode to user struct
	service := servicemodel.Service{}
	err := json.NewDecoder(r.Body).Decode(&service)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// insert data to database
	_, err = db.Exec("INSERT INTO tbservice ( servicetype_id,company_id, date_start, date_dule) VALUES (?,?,?,?)", service.Servicetype_id, service.Company_id, service.Date_start, service.Date_due)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// สั่งให้แสดงผลลัพธ์ที่ได้จากการเพิ่มข้อมูล

	// return data func Listservice to client
	Listservice(w, r, db)

}

func Updateservice(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// get json body from request and decode to user struct
	service := servicemodel.Service{}
	err := json.NewDecoder(r.Body).Decode(&service)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// update data to database
	stmt, err := db.Prepare("UPDATE tbservice SET company_id=?, servicetype_id=?, date_start=?, date_dule=? WHERE id=?")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// ปิดการเชื่อมต่อกับฐานข้อมูล
	defer stmt.Close()

	// สั่งให้เริ่มทำการแก้ไขข้อมูลลงในฐานข้อมูล
	_, err = stmt.Exec(service.Company_id, service.Servicetype_id, service.Date_start, service.Date_due, service.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// return data to client
	Listservice(w, r, db)

}

func Deleteservice(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// get json body from request and decode to user struct
	service := servicemodel.Service{}
	err := json.NewDecoder(r.Body).Decode(&service)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// delete data to database
	stmt, err := db.Prepare("DELETE FROM tbservice WHERE id=?")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// ปิดการเชื่อมต่อกับฐานข้อมูล
	defer stmt.Close()

	// สั่งให้เริ่มทำการลบข้อมูลลงในฐานข้อมูล
	_, err = stmt.Exec(service.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// return data fuce Listservice to client
	Listservice(w, r, db)

}
