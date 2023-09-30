package admincontroller

import (
	adminModel "backend/models/adminModel"
	dashboardmodel "backend/models/dashboardModel"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func GeneratePasswordHash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func ListAdmin(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// ดำเนินการค้นหาข้อมูลทั้งหมดจากฐานข้อมูล
	rows, err := db.Query("SELECT * FROM tbadmins")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// สร้าง slice เพื่อเก็บข้อมูลที่ค้นพบ
	results := []adminModel.Admin{}
	//
	for rows.Next() {
		// สร้างตัวแปรเพื่อเก็บข้อมูลที่ query ค้นพบ
		admin := adminModel.Admin{}
		// สั่งสแกนข้อมูลจาก query ไปเก็บใน struct ตามชื่อฟิลด์
		err := rows.Scan(
			&admin.Id,
			&admin.Role_id,
			&admin.Pstag_id,
			&admin.Team_id,
			&admin.Username,
			&admin.Password,
			&admin.Fistname,
			&admin.Lastname,
			&admin.Email,
			&admin.Tal,
			&admin.Image,
			&admin.Status,
			&admin.Timestamps,
			&admin.User_id,
			&admin.Token_link,
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// เก็บ struct ลงใน slice
		results = append(results, admin)
	}

	// ตรวจสอบข้อผิดพลาดหลังจากวนลูป
	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// แปลงข้อมูลใน slice เป็น JSON
	jsonData, err := json.Marshal(results)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// ตั้งค่า Header และส่ง JSON กลับไปยัง HTTP Response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func AddAdmin(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// สร้างตัวแปรเพื่อเก็บข้อมูลที่รับมาจาก client
	user := adminModel.AddAdmin{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	//check email in database have or not have
	row := db.QueryRow("SELECT email FROM tbadmins WHERE email = ?", user.Email)
	// สร้างตัวแปรเพื่อเก็บข้อมูลที่ query ค้นพบ
	adminDB := adminModel.Admin{}
	// สั่งสแกนข้อมูลจาก query ไปเก็บใน struct ตามชื่อฟิลด์
	err := row.Scan(
		&adminDB.Email,
	)

	// ตรวจสอบว่ามีข้อมูลผู้ใช้งานหรือไม่
	if adminDB.Email != "" {
		http.Error(w, "user have in database", http.StatusNotFound)
		return
	}

	// สร้าง hash password
	hash, err := GeneratePasswordHash(user.Password)
	if err != nil {
		log.Fatal(err)
	}
	// rendom token_link สำหรับยืนยันตัวตน
	token_link := uuid.New().String()

	// get time utc+7 thailand time zone
	//t := time.Now().UTC().Add(time.Hour * 7)
	// บันทึกข้อมูลผู้ใช้ในฐานข้อมูล
	_, err = db.Exec("INSERT INTO tbadmins (username, password, firstname, lastname, email, status, role_id, pstag_id, team_id , token_link ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ? ,?)",
		user.Username, hash, user.Firstname, user.Lastname, user.Email, user.Status, user.RoleID, user.PstagID, user.TeamID, token_link)
	if err != nil {
		log.Fatal(err)
	}

	// ส่งข้อความว่า "บันทึกข้อมูลเรียบร้อยแล้ว"
	fmt.Fprintf(w, "บันทึกข้อมูลเรียบร้อยแล้ว")
}

func UpdateAdmin(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// สร้างตัวแปรเพื่อเก็บข้อมูลที่รับมาจาก client
	user := adminModel.Admin{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// อัพเดทข้อมูลผู้ใช้ในฐานข้อมูล เช็คจาก email
	_, err := db.Exec("UPDATE tbadmins SET username = ?, firstname = ?, lastname = ?, email = ?, status = ?, role_id = ?, pstag_id = ?, team_id = ? WHERE email = ?",
		user.Username, user.Fistname, user.Lastname, user.Email, user.Status, user.Role_id, user.Pstag_id, user.Team_id, user.Email)
	if err != nil {
		log.Fatal(err)
	}

	// ส่งข้อความว่า "อัพเดทข้อมูลเรียบร้อยแล้ว"
	fmt.Fprintf(w, "อัพเดทข้อมูลเรียบร้อยแล้ว")
}

func DashboardAdmin(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// สร้างตัวแปรเพื่อเก็บข้อมูลที่รับมาจาก client
	user := adminModel.Admin{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// ดึงข้อมูลจากฐานข้อมูล
	rows, err := db.Query("SELECT tbcase.id ,tbcompany.type_company,tbcompany.company_name,tbservice_type.service_name,tbservice.date_start,tbservice.date_end ,tbsbt_tax.E_tax_name,tbsbt_tax.Status FROM tbcase JOIN tbsbt_tax ON tbcase.id = tbsbt_tax.case_id JOIN tbservice ON tbservice.id = tbcase.service_id JOIN tbservice_type ON tbservice.servicetype_id = tbservice_type.id JOIN tbcompany ON tbcompany.id = tbservice.company_id = tbcompany.id WHERE tbcase.admin_id = ? AND tbcase.status = 'active' ", user.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	rowsMap := make(map[int]*dashboardmodel.Case)

	for rows.Next() {
		casee := dashboardmodel.Case{}
		sbt_tax := dashboardmodel.Sbt_tax{}
		err := rows.Scan(
			&casee.Id,
			&casee.Type_company,
			&casee.Company_name,
			&casee.Service_name,
			&casee.Date_start,
			&casee.Date_end,
			&sbt_tax.E_tax_name,
			&sbt_tax.Status,
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if _, ok := rowsMap[casee.Id]; !ok {
			rowsMap[casee.Id] = &casee
		}

		rowsMap[casee.Id].Sbt_tax = append(rowsMap[casee.Id].Sbt_tax, sbt_tax)

	}

	// แปลงข้อมูลใน map เป็น JSON
	jsonData, err := json.Marshal(rowsMap)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// ตั้งค่า Header และส่ง JSON กลับไปยัง HTTP Response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)

}

func StatusWork(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	user := adminModel.Admin{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	//statusworks := dashboardmodel.Statusworks{}

	// ดึงข้อมูลจากฐานข้อมูล

}
