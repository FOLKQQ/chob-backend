package controllers

import (
	"backend/models"
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

const SecretKey = "chob-backend-2023"

func LoginAdmins(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// get json body from request and decode to user struct fmt.Println( username, password)))
	admin := models.Admin{}
	err := json.NewDecoder(r.Body).Decode(&admin)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// ดำเนินการค้นหาข้อมูลจากฐานข้อมูล
	row := db.QueryRow("SELECT * FROM tbadmins WHERE username = ?", admin.Username)
	// สร้างตัวแปรเพื่อเก็บข้อมูลที่ query ค้นพบ
	adminDB := models.Admin{}
	// สั่งสแกนข้อมูลจาก query ไปเก็บใน struct ตามชื่อฟิลด์
	err = row.Scan(
		&adminDB.Id,
		&adminDB.Username,
		&adminDB.Password,
		&adminDB.Fistname,
		&adminDB.Lastname,
		&adminDB.Email,
		&adminDB.Image,
		&adminDB.Status,
		&adminDB.Timestamps,
		&adminDB.Role_id,
		&adminDB.Pstag_id,
		&adminDB.Team_id,
		&adminDB.Regis_time,
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
	claims["id"] = adminDB.Id
	claims["username"] = adminDB.Username
	claims["exp"] = time.Now().Add(time.Hour * 4).Unix()

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
		"access_token": accessToken,
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
	results := []models.Admin{}
	//
	for rows.Next() {
		// สร้างตัวแปรเพื่อเก็บข้อมูลที่ query ค้นพบ
		admin := models.Admin{}
		// สั่งสแกนข้อมูลจาก query ไปเก็บใน struct ตามชื่อฟิลด์
		err := rows.Scan(
			&admin.Id,
			&admin.Username,
			&admin.Password,
			&admin.Fistname,
			&admin.Lastname,
			&admin.Email,
			&admin.Image,
			&admin.Status,
			&admin.Timestamps,
			&admin.Role_id,
			&admin.Pstag_id,
			&admin.Team_id,
			&admin.Regis_time,
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
