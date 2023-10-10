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
	result, err := db.Query("SELECT * FROM tbtax_from")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {
		var tax taxModel.Tax_from
		err := result.Scan(&tax.Id, &tax.Task_id, &tax.Payment_date, &tax.Name, &tax.Tax_id, &tax.Income_type, &tax.Amount, &tax.Withholding_tax, &tax.Net_income, &tax.Number, &tax.Village, &tax.Sub_district, &tax.District, &tax.Province, &tax.Zipcode, &tax.View, &tax.Status, &tax.Timestamps)
		if err != nil {
			panic(err.Error())
		}
		taxs = append(taxs, tax)
	}

	json.NewEncoder(w).Encode(taxs)
}

func GetTaxById(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	tax := taxModel.Tax_from{}
	id := r.URL.Query().Get("id")
	result, err := db.Query("SELECT * FROM tbtax_from WHERE id=?", id)
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {
		err := result.Scan(&tax.Id, &tax.Task_id, &tax.Payment_date, &tax.Name, &tax.Tax_id, &tax.Income_type, &tax.Amount, &tax.Withholding_tax, &tax.Net_income, &tax.Number, &tax.Village, &tax.Sub_district, &tax.District, &tax.Province, &tax.Zipcode, &tax.View, &tax.Status, &tax.Timestamps)
		if err != nil {
			panic(err.Error())
		}
	}
	json.NewEncoder(w).Encode(tax)
}

func CreateTax(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	tax := taxModel.Tax_from{}
	json.NewDecoder(r.Body).Decode(&tax)
	fmt.Println(tax.Id)
	// บันทึกข้อมูลผู้ใช้ในฐานข้อมูล
	_, err := db.Exec("INSERT INTO tbtax_from (task_id, payment_date, name, tax_id, income_type, amount, withholding_tax, net_income, number, village, sub_district, district, province, zipcode, view, status ) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)", tax.Task_id, tax.Payment_date, tax.Name, tax.Tax_id, tax.Income_type, tax.Amount, tax.Withholding_tax, tax.Net_income, tax.Number, tax.Village, tax.Sub_district, tax.District, tax.Province, tax.Zipcode, tax.View, tax.Status)
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "New tax was created")
}

func UpdateTax(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	tax := taxModel.Tax_from{}
	json.NewDecoder(r.Body).Decode(&tax)
	// บันทึกข้อมูลผู้ใช้ในฐานข้อมูล
	_, err := db.Exec("UPDATE tbtax_from SET task_id=?, payment_date=?, name=?, tax_id=?, income_type=?, amount=?, withholding_tax=?, net_income=?, number=?, village=?, sub_district=?, district=?, province=?, zipcode=?, view=?, status=?  WHERE id=?", tax.Task_id, tax.Payment_date, tax.Name, tax.Tax_id, tax.Income_type, tax.Amount, tax.Withholding_tax, tax.Net_income, tax.Number, tax.Village, tax.Sub_district, tax.District, tax.Province, tax.Zipcode, tax.View, tax.Status, tax.Id)
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "Tax was updated")
}

func DeleteTax(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	tax := taxModel.Tax_from{}
	json.NewDecoder(r.Body).Decode(&tax)
	// ลบข้อมูลผู้ใช้ในฐานข้อมูล
	_, err := db.Exec("DELETE FROM tbtax_from WHERE id=?", tax.Id)
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "Tax was deleted")
}

func ListTax30(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var taxs []taxModel.Tax_30
	result, err := db.Query("SELECT * FROM tbtax_30")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {
		var tax taxModel.Tax_30
		err := result.Scan(&tax.Id, &tax.Task_id, &tax.Excess_tax_paid, &tax.Excess_tax_paid_from_previous_month, &tax.Excess_tax_not_to_be_paid, &tax.Net_tax_payable, &tax.Sales_total, &tax.Sales_tax, &tax.Purchase_total, &tax.Purchase_tax, &tax.Timestamp)
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
	_, err := db.Exec("INSERT INTO tbtax_30 (task_id, excess_tax_paid, excess_tax_paid_from_previous_month, excess_tax_not_to_be_paid, net_tax_payable, sales_total, sales_tax, purchase_total, purchase_tax, timestamp) VALUES(?, ?,?, ?, ?, ?, ?,?, ?, ?)", tax.Task_id, tax.Excess_tax_paid, tax.Excess_tax_paid_from_previous_month, tax.Excess_tax_not_to_be_paid, tax.Net_tax_payable, tax.Sales_total, tax.Sales_tax, tax.Purchase_total, tax.Purchase_tax, tax.Timestamp)
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "New tax was created")
}

func UpdateTax30(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	tax := taxModel.Tax_30{}
	json.NewDecoder(r.Body).Decode(&tax)
	// บันทึกข้อมูลผู้ใช้ในฐานข้อมูล
	_, err := db.Exec("UPDATE tbtax_30 SET task_id=?,excess_tax_paid=?, excess_tax_paid_from_previous_month=?, excess_tax_not_to_be_paid=?, net_tax_payable=?, sales_total=?, sales_tax=?, purchase_total=?, purchase_tax=? WHERE id=?", tax.Task_id, tax.Excess_tax_paid, tax.Excess_tax_paid_from_previous_month, tax.Excess_tax_not_to_be_paid, tax.Net_tax_payable, tax.Sales_total, tax.Sales_tax, tax.Purchase_total, tax.Purchase_tax, tax.Id)
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "Tax was updated")
}

func DeleteTax30(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	tax := taxModel.Tax_30{}
	json.NewDecoder(r.Body).Decode(&tax)
	// ลบข้อมูลผู้ใช้ในฐานข้อมูล
	_, err := db.Exec("DELETE FROM tbtax_30 WHERE id=?", tax.Id)
	if err != nil {
		panic(err.Error())
	}

	fmt.Fprintf(w, "Tax was deleted")
}

func GetTax30ById(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	Tax_id := r.URL.Query().Get("id")
	// ดึงข้อมูลผู้ใช้จากฐานข้อมูล
	rows, err := db.Query("SELECT * FROM tbtax_30 where id=?", Tax_id)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	// สร้าง slice ของ Tax_30 เพื่อเก็บข้อมูลทั้งหมด
	taxes := []taxModel.Tax_30{}

	// วนลูปเพื่อเก็บข้อมูลทุกแถว
	for rows.Next() {
		var tax taxModel.Tax_30
		err := rows.Scan(&tax.Id, &tax.Task_id, &tax.Excess_tax_paid, &tax.Excess_tax_paid_from_previous_month, &tax.Excess_tax_not_to_be_paid, &tax.Net_tax_payable, &tax.Sales_total, &tax.Sales_tax, &tax.Purchase_total, &tax.Purchase_tax, &tax.Timestamp)
		if err != nil {
			panic(err.Error())
		}
		taxes = append(taxes, tax)
	}

	// แปลง slice ของ Tax_30 เป็น JSON และส่งคืนไปยัง client
	json.NewEncoder(w).Encode(taxes)
}
