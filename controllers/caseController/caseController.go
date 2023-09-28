package casecontroller

import (
	casemodel "backend/models/caseModel"
	"database/sql"
	"encoding/json"
	"net/http"
)

func Addcase(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// get json body from request and decode to user struct
	casee := casemodel.Case{}
	err := json.NewDecoder(r.Body).Decode(&casee)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// insert data to database
	stmt, err := db.Prepare("INSERT INTO tbcase (service_id, admin_id, payment_date, name, tax_id, income_type, amount, withholding_tax, net_income, number, village, sub_district, district, province, zip_code, view, status, timestamps) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// ปิดการเชื่อมต่อกับฐานข้อมูล
	defer stmt.Close()

	// สั่งให้เริ่มทำการเพิ่มข้อมูลลงในฐานข้อมูล
	_, err = stmt.Exec(casee.Service_id, casee.Admin_id, casee.Payment_date, casee.Name, casee.Tax_id, casee.Income_type, casee.Amount, casee.Withholding_tax, casee.Net_income, casee.Number, casee.Village, casee.Sub_district, casee.District, casee.Province, casee.Zip_code, casee.View, casee.Status, casee.Timestamps)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// return data to client
	Listcase(w, r, db)

}

func Listcase(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// ดำเนินการค้นหาข้อมูลจากฐานข้อมูล
	rows, err := db.Query("SELECT * FROM tbcase")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// สร้างตัวแปรเพื่อเก็บข้อมูลที่ query ค้นพบ
	caseresults := []casemodel.Case{}

	// เก็บข้อมูลที่ query ค้นพบลงในตัวแปร
	for rows.Next() {
		// สร้างตัวแปรเพื่อเก็บข้อมูลที่ query ค้นพบ
		casee := casemodel.Case{}

		// สั่งสแกนข้อมูลจาก query ไปเก็บใน struct ตามชื่อฟิลด์
		err := rows.Scan(
			&casee.Id,
			&casee.Service_id,
			&casee.Admin_id,
			&casee.Payment_date,
			&casee.Name,
			&casee.Tax_id,
			&casee.Income_type,
			&casee.Amount,
			&casee.Withholding_tax,
			&casee.Net_income,
			&casee.Number,
			&casee.Village,
			&casee.Sub_district,
			&casee.District,
			&casee.Province,
			&casee.Zip_code,
			&casee.View,
			&casee.Status,
			&casee.Timestamps,
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// เก็บ struct ลงใน slice
		caseresults = append(caseresults, casee)
	}

	// ตรวจสอบข้อผิดพลาดหลังจากวนลูป
	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// แปลงข้อมูลที่ query ค้นพบเป็น JSON แล้วส่งคืนไปยัง client
	jsonData, err := json.Marshal(caseresults)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}

func Updatecase(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// get json body from request and decode to user struct
	casee := casemodel.Case{}
	err := json.NewDecoder(r.Body).Decode(&casee)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// update data to database
	stmt, err := db.Prepare("UPDATE tbcase SET service_id=?, admin_id=?, payment_date=?, name=?, tax_id=?, income_type=?, amount=?, withholding_tax=?, net_income=?, number=?, village=?, sub_district=?, district=?, province=?, zip_code=?, view=?, status=?, timestamps=? WHERE id=?")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// ปิดการเชื่อมต่อกับฐานข้อมูล
	defer stmt.Close()

	// สั่งให้เริ่มทำการแก้ไขข้อมูลลงในฐานข้อมูล
	_, err = stmt.Exec(casee.Service_id, casee.Admin_id, casee.Payment_date, casee.Name, casee.Tax_id, casee.Income_type, casee.Amount, casee.Withholding_tax, casee.Net_income, casee.Number, casee.Village, casee.Sub_district, casee.District, casee.Province, casee.Zip_code, casee.View, casee.Status, casee.Timestamps, casee.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// return data to client
	Listcase(w, r, db)
}

func Deletecase(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// get json body from request and decode to user struct
	casee := casemodel.Case{}
	err := json.NewDecoder(r.Body).Decode(&casee)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// delete data to database
	stmt, err := db.Prepare("DELETE FROM tbcase WHERE id=?")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// ปิดการเชื่อมต่อกับฐานข้อมูล
	defer stmt.Close()

	// สั่งให้เริ่มทำการลบข้อมูลลงในฐานข้อมูล
	_, err = stmt.Exec(casee.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// return data fuce Listcase to client
	Listcase(w, r, db)
}
