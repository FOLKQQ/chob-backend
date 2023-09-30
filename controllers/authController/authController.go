package authcontrollers

import (
	adminModel "backend/models/adminModel"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	_ "github.com/go-sql-driver/mysql"
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
