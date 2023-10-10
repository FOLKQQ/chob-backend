package billingController

import (
	"backend/models/billingModel"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func CreateBilling(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	billing := billingModel.Billing{}
	json.NewDecoder(r.Body).Decode(&billing)
	// สร้าง timestamp ในรูปแบบของ datetime
	// บันทึกข้อมูลผู้ใช้ในฐานข้อมูล
	_, err := db.Exec("INSERT INTO tbbilling (client_id, invoice, invoice_date, date_due, date_paid, discount, price, status) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		billing.Client_id, billing.Invoice, billing.Invoice_date, billing.Date_due, billing.Date_paid, billing.Discount, billing.Price, billing.Status)
	if err != nil {
		fmt.Println(err)
	}

	// ส่งข้อความว่า "บันทึกข้อมูลเรียบร้อยแล้ว"
	fmt.Fprintf(w, "บันทึกข้อมูลเรียบร้อยแล้ว")
}

func ListBilling(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// getbody company_id
	// สร้างตัวแปรเพื่อเก็บข้อมูลผู้ใช้ทั้งหมด
	listbilling := []billingModel.Billing{}
	// ค้นหาข้อมูลผู้ใช้ทั้งหมดจากฐานข้อมูล
	rows, err := db.Query("SELECT * FROM tbbilling ")
	if err != nil {
		fmt.Println(err)
	}

	// วนลูปเพื่อเก็บข้อมูลผู้ใช้ทั้งหมดลงในตัวแปร listuser
	for rows.Next() {
		billing := billingModel.Billing{}
		err := rows.Scan(&billing.Id, &billing.Client_id, &billing.Invoice, &billing.Invoice_date, &billing.Date_due, &billing.Date_paid, &billing.Discount, &billing.Price, &billing.Status, &billing.Timestamps)
		if err != nil {
			fmt.Println(err)
		}
		listbilling = append(listbilling, billing)
	}

	// แสดงข้อมูลผู้ใช้ทั้งหมดในรูปแบบ JSON
	json.NewEncoder(w).Encode(listbilling)
}

func GetBillingById(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// getbody company_id
	// สร้างตัวแปรเพื่อเก็บข้อมูลผู้ใช้ทั้งหมด
	billing := billingModel.Billing{}
	// ค้นหาข้อมูลผู้ใช้ทั้งหมดจากฐานข้อมูล
	err := db.QueryRow("SELECT * FROM tbbilling WHERE id = ?", r.FormValue("id")).Scan(&billing.Id, &billing.Client_id, &billing.Invoice, &billing.Invoice_date, &billing.Date_due, &billing.Date_paid, &billing.Discount, &billing.Price, &billing.Status, &billing.Timestamps)
	if err != nil {
		fmt.Println(err)
	}

	// แสดงข้อมูลผู้ใช้ทั้งหมดในรูปแบบ JSON
	json.NewEncoder(w).Encode(billing)
}

func UpdateBilling(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	billing := billingModel.Billing{}
	json.NewDecoder(r.Body).Decode(&billing)
	// สร้าง timestamp ในรูปแบบของ datetime
	// อัพเดทข้อมูลผู้ใช้ในฐานข้อมูล
	_, err := db.Exec("UPDATE tbbilling SET client_id = ?, invoice = ?, invoice_date = ?, date_due = ?, date_paid = ?, discount = ?, price = ?, status = ? WHERE id = ?",
		billing.Client_id, billing.Invoice, billing.Invoice_date, billing.Date_due, billing.Date_paid, billing.Discount, billing.Price, billing.Status, billing.Id)
	if err != nil {
		fmt.Println(err)
	}

	// ส่งข้อความว่า "อัพเดทข้อมูลเรียบร้อยแล้ว"
	fmt.Fprintf(w, "อัพเดทข้อมูลเรียบร้อยแล้ว")
}

func DeleteBilling(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	billing := billingModel.Billing{}
	json.NewDecoder(r.Body).Decode(&billing)
	// ลบข้อมูลผู้ใช้ในฐานข้อมูล
	_, err := db.Exec("DELETE FROM tbbilling WHERE id = ?", billing.Id)
	if err != nil {
		fmt.Println(err)
	}

	// ส่งข้อความว่า "ลบข้อมูลเรียบร้อยแล้ว"
	fmt.Fprintf(w, "ลบข้อมูลเรียบร้อยแล้ว")
}

func CreateBilling_tax(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	billing := billingModel.Billing_tax{}
	json.NewDecoder(r.Body).Decode(&billing)
	// สร้าง timestamp ในรูปแบบของ datetime
	// บันทึกข้อมูลผู้ใช้ในฐานข้อมูล
	_, err := db.Exec("INSERT INTO tbbilling_tax (task_id, client_id, invoice, invoice_date, date_due, date_paid, discount, price, status) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)",
		billing.Task_id, billing.Client_id, billing.Invoice, billing.Invoice_date, billing.Date_due, billing.Date_paid, billing.Discount, billing.Price, billing.Status)
	if err != nil {
		fmt.Println(err)
	}

	// ส่งข้อความว่า "บันทึกข้อมูลเรียบร้อยแล้ว"
	fmt.Fprintf(w, "บันทึกข้อมูลเรียบร้อยแล้ว")
}

func ListBilling_tax(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// getbody company_id
	// สร้างตัวแปรเพื่อเก็บข้อมูลผู้ใช้ทั้งหมด
	listbilling := []billingModel.Billing_tax{}
	// ค้นหาข้อมูลผู้ใช้ทั้งหมดจากฐานข้อมูล
	rows, err := db.Query("SELECT * FROM tbbilling_tax ")
	if err != nil {
		fmt.Println(err)
	}

	// วนลูปเพื่อเก็บข้อมูลผู้ใช้ทั้งหมดลงในตัวแปร listuser
	for rows.Next() {
		billing := billingModel.Billing_tax{}
		err := rows.Scan(&billing.Id, &billing.Task_id, &billing.Client_id, &billing.Invoice, &billing.Invoice_date, &billing.Date_due, &billing.Date_paid, &billing.Discount, &billing.Price, &billing.Status, &billing.Timestamps)
		if err != nil {
			fmt.Println(err)
		}
		listbilling = append(listbilling, billing)
	}

	// แสดงข้อมูลผู้ใช้ทั้งหมดในรูปแบบ JSON
	json.NewEncoder(w).Encode(listbilling)
}

func GetBilling_taxById(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// getbody company_id
	// สร้างตัวแปรเพื่อเก็บข้อมูลผู้ใช้ทั้งหมด
	billing := billingModel.Billing_tax{}
	// ค้นหาข้อมูลผู้ใช้ทั้งหมดจากฐานข้อมูล
	err := db.QueryRow("SELECT * FROM tbbilling_tax WHERE id = ?", r.FormValue("id")).Scan(&billing.Id, &billing.Task_id, &billing.Client_id, &billing.Invoice, &billing.Invoice_date, &billing.Date_due, &billing.Date_paid, &billing.Discount, &billing.Price, &billing.Status, &billing.Timestamps)
	if err != nil {
		fmt.Println(err)
	}

	// แสดงข้อมูลผู้ใช้ทั้งหมดในรูปแบบ JSON
	json.NewEncoder(w).Encode(billing)
}

func UpdateBilling_tax(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	billing := billingModel.Billing_tax{}
	json.NewDecoder(r.Body).Decode(&billing)
	// สร้าง timestamp ในรูปแบบของ datetime
	// อัพเดทข้อมูลผู้ใช้ในฐานข้อมูล
	_, err := db.Exec("UPDATE tbbilling_tax SET task_id = ?, client_id = ?, invoice = ?, invoice_date = ?, date_due = ?, date_paid = ?, discount = ?, price = ?, status = ? WHERE id = ?",
		billing.Task_id, billing.Client_id, billing.Invoice, billing.Invoice_date, billing.Date_due, billing.Date_paid, billing.Discount, billing.Price, billing.Status, billing.Id)
	if err != nil {
		fmt.Println(err)
	}

	// ส่งข้อความว่า "อัพเดทข้อมูลเรียบร้อยแล้ว"
	fmt.Fprintf(w, "อัพเดทข้อมูลเรียบร้อยแล้ว")
}

func DeleteBilling_tax(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	billing := billingModel.Billing_tax{}
	json.NewDecoder(r.Body).Decode(&billing)
	// ลบข้อมูลผู้ใช้ในฐานข้อมูล
	_, err := db.Exec("DELETE FROM tbbilling_tax WHERE id = ?", billing.Id)
	if err != nil {
		fmt.Println(err)
	}

	// ส่งข้อความว่า "ลบข้อมูลเรียบร้อยแล้ว"
	fmt.Fprintf(w, "ลบข้อมูลเรียบร้อยแล้ว")
}
