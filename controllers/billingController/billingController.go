package billingcontroller

import (
	"backend/models/billingModel"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func ListBilling(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var billings []billingModel.Billing
	result, err := db.Query("SELECT * FROM tbbilling")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {
		var billing billingModel.Billing
		err := result.Scan(&billing.Id, &billing.Client_id, &billing.Invoice, &billing.Invoice_date, &billing.Date_due, &billing.Date_paid, &billing.Totol, &billing.Status, &billing.Timestamps)
		if err != nil {
			panic(err.Error())
		}
		billings = append(billings, billing)
	}
	// แสดงข้อมูลผู้ใช้ทั้งหมดในรูปแบบ JSON
	json.NewEncoder(w).Encode(billings)

}

func CreateBilling(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	billing := billingModel.Billing{}
	json.NewDecoder(r.Body).Decode(&billing)
	// บันทึกข้อมูลผู้ใช้ในฐานข้อมูล
	_, err := db.Exec("INSERT INTO tbbilling (client_id, invoice, invoice_date, date_due, date_paid, totol, status, timestamps) VALUES(?, ?, ?, ?, ?, ?, ?, ?)", billing.Client_id, billing.Invoice, billing.Invoice_date, billing.Date_due, billing.Date_paid, billing.Totol, billing.Status, billing.Timestamps)
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "New billing was created")
}

func UpdateBilling(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	billing := billingModel.Billing{}
	json.NewDecoder(r.Body).Decode(&billing)
	// บันทึกข้อมูลผู้ใช้ในฐานข้อมูล
	_, err := db.Exec("UPDATE tbbilling SET client_id=?, invoice=?, invoice_date=?, date_due=?, date_paid=?, totol=?, status=?, timestamps=? WHERE id=?", billing.Client_id, billing.Invoice, billing.Invoice_date, billing.Date_due, billing.Date_paid, billing.Totol, billing.Status, billing.Timestamps, billing.Id)
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "Billing was updated")
}

func DeleteBilling(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	billing := billingModel.Billing{}
	json.NewDecoder(r.Body).Decode(&billing)
	// ลบข้อมูลผู้ใช้ในฐานข้อมูล
	_, err := db.Exec("DELETE FROM tbbilling WHERE id=?", billing.Id)
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "Billing was deleted")
}

func ListBilling_tax(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var billing_taxs []billingModel.Billing_tax
	result, err := db.Query("SELECT * FROM tbbilling_tax")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {
		var billing_tax billingModel.Billing_tax
		err := result.Scan(&billing_tax.Id, &billing_tax.Task_id, &billing_tax.Price)
		if err != nil {
			panic(err.Error())
		}
		billing_taxs = append(billing_taxs, billing_tax)
	}
	// แสดงข้อมูลผู้ใช้ทั้งหมดในรูปแบบ JSON
	json.NewEncoder(w).Encode(billing_taxs)

}

func CreateBilling_tax(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	billing_tax := billingModel.Billing_tax{}
	json.NewDecoder(r.Body).Decode(&billing_tax)
	// บันทึกข้อมูลผู้ใช้ในฐานข้อมูล
	_, err := db.Exec("INSERT INTO tbbilling_tax (task_id, price) VALUES(?, ?)", billing_tax.Task_id, billing_tax.Price)
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "New billing_tax was created")
}

func UpdateBilling_tax(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	billing_tax := billingModel.Billing_tax{}
	json.NewDecoder(r.Body).Decode(&billing_tax)
	// บันทึกข้อมูลผู้ใช้ในฐานข้อมูล
	_, err := db.Exec("UPDATE tbbilling_tax SET task_id=?, price=? WHERE id=?", billing_tax.Task_id, billing_tax.Price, billing_tax.Id)
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "Billing_tax was updated")
}
