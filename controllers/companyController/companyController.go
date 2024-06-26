package companycontroller

import (
	"backend/models/companyModel"
	"database/sql"
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func ListCompany(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// ดำเนินการค้นหาข้อมูลจากฐานข้อมูล
	rows, err := db.Query("SELECT * FROM tbcompany")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// สร้างตัวแปรเพื่อเก็บข้อมูลที่ query ค้นพบ
	companys := []companyModel.Company{}
	// วนลูป query ข้อมูลที่ค้นพบแล้วเก็บลงใน struct
	for rows.Next() {
		// สร้างตัวแปรเพื่อเก็บข้อมูลที่ query ค้นพบ
		company := companyModel.Company{}
		// สั่งสแกนข้อมูลจาก query ไปเก็บใน struct ตามชื่อฟิลด์
		err := rows.Scan(
			&company.ID,
			&company.TypeCompany,
			&company.CompanyName,
			&company.CodeIdentification,
			&company.Signature,
			&company.StartComDate,
			&company.StartVatDate,
			&company.BusinessType,
			&company.Id_dbd,
			&company.Pass_dbd,
			&company.Id_filing,
			&company.Pass_filing,
			&company.Email,
			&company.Tal,
			&company.Address,
			&company.Addressextra,
			&company.Subdistrict,
			&company.District,
			&company.Province,
			&company.Zipcode,
			&company.Status,
			&company.Timestamps,
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// เก็บ struct ลงใน slice
		companys = append(companys, company)
	}
	// ตรวจสอบข้อผิดพลาดหลังจากวนลูป
	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// แปลงข้อมูลใน slice เป็น JSON
	jsonData, err := json.Marshal(companys)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// ส่ง JSON กลับไปเป็น response
	w.Write(jsonData)
}

func AddCompany(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// ดึงข้อมูลจาก form-data
	// และแปลงเป็น struct ของ Company
	company := companyModel.Company{}
	json.NewDecoder(r.Body).Decode(&company)
	var status = "active"
	// ดำเนินการเพิ่มข้อมูลลงในฐานข้อมูล
	_, err := db.Exec("INSERT INTO tbcompany (type_company,company_name,code_identification,signature,regis_company,regis_vat,business_type,id_dbd,pass_dbd,id_filing,pass_filing,email,tal,addresses,addressesextra,subdistrict,district,province,zipcode,status) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)",
		company.TypeCompany,
		company.CompanyName,
		company.CodeIdentification,
		company.Signature,
		company.StartComDate,
		company.StartVatDate,
		company.BusinessType,
		company.Id_dbd,
		company.Pass_dbd,
		company.Id_filing,
		company.Pass_filing,
		company.Email,
		company.Tal,
		company.Address,
		company.Addressextra,
		company.Subdistrict,
		company.District,
		company.Province,
		company.Zipcode,
		status,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// ส่ง response ค่า 200 และข้อความกลับไป
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Insert data successfully."))
}

func UpdateCompany(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// ดึงข้อมูลจาก form-data
	// และแปลงเป็น struct ของ Company
	company := companyModel.Company{}
	json.NewDecoder(r.Body).Decode(&company)
	// ดำเนินการแก้ไขข้อมูลลงในฐานข้อมูล
	_, err := db.Exec("UPDATE tbcompany SET type_company=?,company_name=?,code_identification=?,signature=?,regis_company=?,regis_vat=?,business_type=?,id_dbd=?,pass_dbd=?,id_filing=?,pass_filing=?,email=?,tal=?,status=? WHERE id=?",
		company.TypeCompany,
		company.CompanyName,
		company.CodeIdentification,
		company.Signature,
		company.StartComDate,
		company.StartVatDate,
		company.BusinessType,
		company.Id_dbd,
		company.Pass_dbd,
		company.Id_filing,
		company.Pass_filing,
		company.Email,
		company.Tal,
		company.Address,
		company.Addressextra,
		company.Subdistrict,
		company.District,
		company.Status,
		company.ID,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// ส่ง response ค่า 200 และข้อความกลับไป
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Update data successfully."))
}

func DeleteCompany(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// ดึงข้อมูลจาก form-data
	// และแปลงเป็น struct ของ Company
	company := companyModel.Company{}
	json.NewDecoder(r.Body).Decode(&company)
	// ดำเนินการลบข้อมูลออกจากฐานข้อมูล
	_, err := db.Exec("DELETE FROM tbcompany WHERE id=?", company.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// ส่ง response ค่า 200 และข้อความกลับไป
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Delete data successfully."))
}
