package taxcontroller

import (
	"backend/models/taxModel"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func ListTax(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var taxs []taxModel.Tax_from
	result, err := db.Query("SELECT * FROM tbtax")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {
		var tax taxModel.Tax_from
		err := result.Scan(&tax.Task_id, &tax.Payment_date, &tax.Name, &tax.Tax_id, &tax.Income_type, &tax.Amount, &tax.Witholding_tax, &tax.Net_income, &tax.Number, &tax.Village, &tax.Sub_district, &tax.District, &tax.Province, &tax.Zipcode, &tax.View, &tax.Status, &tax.Timestamps)
		if err != nil {
			panic(err.Error())
		}
		taxs = append(taxs, tax)
	}

	json.NewEncoder(w).Encode(taxs)
}

func CreateTax(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	tax := taxModel.Tax_from{}
	json.NewDecoder(r.Body).Decode(&tax)
	// บันทึกข้อมูลผู้ใช้ในฐานข้อมูล
	_, err := db.Exec("INSERT INTO tbtax (task_id, payment_date, name, tax_id, income_type, amount, witholding_tax, net_income, number, village, sub_district, district, province, zipcode, view, status, timestamps) VALUES(?, ?,?, ?, ?, ?, ?,?, ?, ?, ?,?, ?, ?, ?,?, ?, ?)", tax.Task_id, tax.Payment_date, tax.Name, tax.Tax_id, tax.Income_type, tax.Amount, tax.Witholding_tax, tax.Net_income, tax.Number, tax.Village, tax.Sub_district, tax.District, tax.Province, tax.Zipcode, tax.View, tax.Status, tax.Timestamps)
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "New tax was created")
}

func UpdateTax(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	tax := taxModel.Tax_from{}
	json.NewDecoder(r.Body).Decode(&tax)
	// บันทึกข้อมูลผู้ใช้ในฐานข้อมูล
	_, err := db.Exec("UPDATE tbtax SET task_id=?, payment_date=?, name=?, tax_id=?, income_type=?, amount=?, witholding_tax=?, net_income=?, number=?, village=?, sub_district=?, district=?, province=?, zipcode=?, view=?, status=?, timestamps=? WHERE id=?", tax.Task_id, tax.Payment_date, tax.Name, tax.Tax_id, tax.Income_type, tax.Amount, tax.Witholding_tax, tax.Net_income, tax.Number, tax.Village, tax.Sub_district, tax.District, tax.Province, tax.Zipcode, tax.View, tax.Status, tax.Timestamps, tax.Task_id)
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "Tax was updated")
}

func DeleteTax(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	tax := taxModel.Tax_from{}
	json.NewDecoder(r.Body).Decode(&tax)
	// ลบข้อมูลผู้ใช้ในฐานข้อมูล
	_, err := db.Exec("DELETE FROM tbtax WHERE id=?", tax.Task_id)
	if err != nil {
		panic(err.Error())
	}
}

func ListTax30(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var taxs []taxModel.Tax_30
	result, err := db.Query("SELECT * FROM tbtax30")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {
		var tax taxModel.Tax_30
		err := result.Scan(&tax.Task_id, &tax.Excess_tax_paid, &tax.Excess_tax_paid_from_previous_month, &tax.Excess_tax_not_to_be_paid, &tax.Net_tax_payable, &tax.Sales_total, &tax.Sales_tax, &tax.Purchase_total, &tax.Purchase_tax, &tax.Timestamp)
		if err != nil {
			panic(err.Error())
		}
		taxs = append(taxs, tax)
	}

	json.NewEncoder(w).Encode(taxs)
}

func CreateTax30(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	tax := taxModel.Tax_30{}
	json.NewDecoder(r.Body).Decode(&tax)
	// บันทึกข้อมูลผู้ใช้ในฐานข้อมูล
	_, err := db.Exec("INSERT INTO tbtax30 (task_id, excess_tax_paid, excess_tax_paid_from_previous_month, excess_tax_not_to_be_paid, net_tax_payable, sales_total, sales_tax, purchase_total, purchase_tax, timestamp) VALUES(?, ?,?, ?, ?, ?, ?,?, ?, ?)", tax.Task_id, tax.Excess_tax_paid, tax.Excess_tax_paid_from_previous_month, tax.Excess_tax_not_to_be_paid, tax.Net_tax_payable, tax.Sales_total, tax.Sales_tax, tax.Purchase_total, tax.Purchase_tax, tax.Timestamp)
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "New tax was created")
}

func UpdateTax30(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	tax := taxModel.Tax_30{}
	json.NewDecoder(r.Body).Decode(&tax)
	// บันทึกข้อมูลผู้ใช้ในฐานข้อมูล
	_, err := db.Exec("UPDATE tbtax30 SET excess_tax_paid=?, excess_tax_paid_from_previous_month=?, excess_tax_not_to_be_paid=?, net_tax_payable=?, sales_total=?, sales_tax=?, purchase_total=?, purchase_tax=?, timestamp=? WHERE id=?", tax.Excess_tax_paid, tax.Excess_tax_paid_from_previous_month, tax.Excess_tax_not_to_be_paid, tax.Net_tax_payable, tax.Sales_total, tax.Sales_tax, tax.Purchase_total, tax.Purchase_tax, tax.Timestamp, tax.Task_id)
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "Tax was updated")
}

func DeleteTax30(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	tax := taxModel.Tax_30{}
	json.NewDecoder(r.Body).Decode(&tax)
	// ลบข้อมูลผู้ใช้ในฐานข้อมูล
	_, err := db.Exec("DELETE FROM tbtax30 WHERE id=?", tax.Task_id)
	if err != nil {
		panic(err.Error())
	}
}
