package sbttaxcontroller

import (
	"backend/models/sbttaxModel"
	"database/sql"
	"encoding/json"
	"net/http"
)

func Addsbttax(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// get json body from request and decode to user struct
	sbttax := sbttaxmodel.Sbttax{}
	err := json.NewDecoder(r.Body).Decode(&sbttax)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// insert data to database
	stmt, err := db.Prepare("INSERT INTO tbsbt_tax (case_id, e_tax_name, sales_total, sales_tax, purchase_total, purchase_tax, excess_tax_paid, excess_tax_paid_from_previous_month, excess_tax_not_to_be_paid, net_tax_payable , status) VALUES (?,?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// ปิดการเชื่อมต่อกับฐานข้อมูล
	defer stmt.Close()

	// สั่งให้เริ่มทำการเพิ่มข้อมูลลงในฐานข้อมูล
	_, err = stmt.Exec(sbttax.Case_id, sbttax.E_tax_name, sbttax.Sales_total, sbttax.Sales_tax, sbttax.Purchase_total, sbttax.Purchase_tax, sbttax.Excess_tax_paid, sbttax.Excess_tax_paid_from_previous_month, sbttax.Excess_tax_not_to_be_paid, sbttax.Net_tax_payable, sbttax.Status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// return data to client
	Listsbttax(w, r, db)

}

func Listsbttax(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// ดำเนินการค้นหาข้อมูลจากฐานข้อมูล
	rows, err := db.Query("SELECT * FROM tbsbt_tax")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// สร้างตัวแปรเพื่อเก็บข้อมูลที่ query ค้นพบ
	sbttaxresults := []sbttaxmodel.Sbttax{}
	// เก็บข้อมูลที่ query ค้นพบลงในตัวแปร
	for rows.Next() {
		// สร้างตัวแปรเพื่อเก็บข้อมูลที่ query ค้นพบ
		sbttax := sbttaxmodel.Sbttax{}
		// สั่งสแกนข้อมูลจาก query ไปเก็บใน struct ตามชื่อฟิลด์
		err := rows.Scan(
			&sbttax.Id,
			&sbttax.Case_id,
			&sbttax.E_tax_name,
			&sbttax.Sales_total,
			&sbttax.Sales_tax,
			&sbttax.Purchase_total,
			&sbttax.Purchase_tax,
			&sbttax.Excess_tax_paid,
			&sbttax.Excess_tax_paid_from_previous_month,
			&sbttax.Excess_tax_not_to_be_paid,
			&sbttax.Net_tax_payable,
			&sbttax.Timestamps,
			&sbttax.Status,
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// เก็บ struct ลงใน slice
		sbttaxresults = append(sbttaxresults, sbttax)
	}

	// ตรวจสอบข้อผิดพลาดหลังจากวนลูป
	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// return data to client
	json.NewEncoder(w).Encode(sbttaxresults)
}

func Updatesbttax(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// get json body from request and decode to user struct
	sbttax := sbttaxmodel.Sbttax{}
	err := json.NewDecoder(r.Body).Decode(&sbttax)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// update data to database
	stmt, err := db.Prepare("UPDATE tbsbt_tax SET case_id=?, e_tax_name=?, sales_total=?, sales_tax=?, purchase_total=?, purchase_tax=?, excess_tax_paid=?, excess_tax_paid_from_previous_month=?, excess_tax_not_to_be_paid=?, net_tax_payable=? , status=? WHERE id=?")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// ปิดการเชื่อมต่อกับฐานข้อมูล
	defer stmt.Close()

	// สั่งให้เริ่มทำการแก้ไขข้อมูลลงในฐานข้อมูล
	_, err = stmt.Exec(sbttax.Case_id, sbttax.E_tax_name, sbttax.Sales_total, sbttax.Sales_tax, sbttax.Purchase_total, sbttax.Purchase_tax, sbttax.Excess_tax_paid, sbttax.Excess_tax_paid_from_previous_month, sbttax.Excess_tax_not_to_be_paid, sbttax.Net_tax_payable, sbttax.Status, sbttax.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// return data to client
	Listsbttax(w, r, db)
}

func Deletesbttax(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// get json body from request and decode to user struct
	sbttax := sbttaxmodel.Sbttax{}
	err := json.NewDecoder(r.Body).Decode(&sbttax)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// delete data to database
	stmt, err := db.Prepare("DELETE FROM tbsbt_tax WHERE id=?")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// ปิดการเชื่อมต่อกับฐานข้อมูล
	defer stmt.Close()

	// สั่งให้เริ่มทำการลบข้อมูลลงในฐานข้อมูล
	_, err = stmt.Exec(sbttax.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// return data fuce Listsbttax to client
	Listsbttax(w, r, db)
}
