package authcontrollers

import (
	adminModel "backend/models/adminModel"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

const SecretKey = "chob-backend-2023"

func GeneratePasswordHash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func LoginAdmins(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// get json body from request and decode to user struct fmt.Println( username, password)))
	admin := adminModel.Admin{}
	err := json.NewDecoder(r.Body).Decode(&admin)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//fmt.Println(admin.Username, admin.Password)

	// ดำเนินการค้นหาข้อมูลจากฐานข้อมูล
	row := db.QueryRow("SELECT * FROM tbadmins WHERE email = ?", admin.Email)
	// สร้างตัวแปรเพื่อเก็บข้อมูลที่ query ค้นพบ
	adminDB := adminModel.Admin{}
	// สั่งสแกนข้อมูลจาก query ไปเก็บใน struct ตามชื่อฟิลด์
	err = row.Scan(
		&adminDB.Id,
		&adminDB.Role_id,
		&adminDB.Pstag_id,
		&adminDB.Team_id,
		&adminDB.Username,
		&adminDB.Password,
		&adminDB.Fistname,
		&adminDB.Lastname,
		&adminDB.Email,
		&adminDB.Tal,
		&adminDB.Image,
		&adminDB.Status,
		&adminDB.Timestamps,
		&adminDB.User_id,
		&adminDB.Token_link,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// ตรวจสอบว่ามีข้อมูลผู้ใช้งานหรือไม่
	if adminDB.Id == 0 {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	// ตรวจสอบรหัสผ่าน
	err = bcrypt.CompareHashAndPassword([]byte(adminDB.Password), []byte(admin.Password))
	if err != nil {
		http.Error(w, "invalid password", http.StatusUnauthorized)
		return
	}

	// สร้าง token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = adminDB.Username
	claims["role_id"] = adminDB.Role_id
	claims["pstag_id"] = adminDB.Pstag_id
	claims["team_id"] = adminDB.Team_id
	claims["exp"] = time.Now().Add(time.Hour * 5).Unix()

	// สร้าง access token
	accessToken, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// return token to client side
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"accessToken": accessToken,
		"id":          fmt.Sprintf("%d", adminDB.Id),
		"Username":    adminDB.Username,
		"Role":        fmt.Sprintf("%d", adminDB.Role_id),
		"Team":        fmt.Sprintf("%d", adminDB.Team_id),
		"Pstag":       fmt.Sprintf("%d", adminDB.Pstag_id),
		"Status":      adminDB.Status,
		"Email":       adminDB.Email,
		"Tal":         adminDB.Tal,
		"Firstname":   adminDB.Fistname,
		"Lastname":    adminDB.Lastname,
		"Image":       adminDB.Image,
		"Token_link":  adminDB.Token_link,
		"User_id":     adminDB.User_id,
	})
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

type Case struct {
	Company_id   int    `json:"company_id"`
	Service_name string `json:"service_name"`
	Company_name string `json:"company_name"`
	Type_company string `json:"type_company"`
	Date_start   string `json:"date_start"`
	Date_end     string `json:"date_end"`
	Sbt_tax      []Sbt_tax
}
type Sbt_tax struct {
	E_tax_name string `json:"e_tax_name"`
	Status     string `json:"status"`
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
	rows, err := db.Query("SELECT tbcompany.id ,tbcompany.company_name ,tbcompany.type_company ,tbservice_type.service_name,tbservice.date_start,tbservice.date_end,tbsbt_tax.e_tax_name,tbsbt_tax.status FROM tbcase JOIN tbsbt_tax ON tbcase.id = tbsbt_tax.case_id JOIN tbservice ON tbservice.id = tbcase.service_id JOIN tbservice_type ON tbservice_type.id = tbservice.servicetype_id JOIN tbcompany ON tbcompany.id = tbservice.company_id WHERE tbcase.admin_id = ?", user.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	rowsMap := make(map[int]*Case)

	for rows.Next() {
		p := struct {
			Company_id   int
			Service_name string
			Company_name string
			Type_company string
			Date_start   string
			Date_end     string
			E_tax_name   string
			Status       string
		}{}
		if err := rows.Scan(&p.Service_name, &p.Company_id, &p.Company_name, &p.Type_company, &p.Date_start, &p.Date_end, &p.E_tax_name, &p.Status); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if _, ok := rowsMap[p.Company_id]; !ok {
			rowsMap[p.Company_id] = &Case{Service_name: p.Service_name, Company_id: p.Company_id, Company_name: p.Company_name, Type_company: p.Type_company, Date_start: p.Date_start, Date_end: p.Date_end}

		}
		sbt_tax := Sbt_tax{E_tax_name: p.E_tax_name, Status: p.Status}

		rowsMap[p.Company_id].Sbt_tax = append(rowsMap[p.Company_id].Sbt_tax, sbt_tax)

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
