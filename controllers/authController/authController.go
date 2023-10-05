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
		&adminDB.ID,
		&adminDB.RoleID,
		&adminDB.TeamID,
		&adminDB.UserID,
		&adminDB.Username,
		&adminDB.Password,
		&adminDB.Firstname,
		&adminDB.Lastname,
		&adminDB.Email,
		&adminDB.Image,
		&adminDB.Tal,
		&adminDB.Token_link,
		&adminDB.Status,
		&adminDB.Timestamps,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// ตรวจสอบว่ามีข้อมูลผู้ใช้งานหรือไม่
	if adminDB.ID == 0 {
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
	claims["role_id"] = adminDB.RoleID
	claims["team_id"] = adminDB.TeamID
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
		"access_token": accessToken,
		"id":           fmt.Sprintf("%d", adminDB.ID),
		"role_id":      fmt.Sprintf("%d", adminDB.RoleID),
		"team_id":      fmt.Sprintf("%d", adminDB.TeamID),
		"user_id":      admin.UserID,
		"username":     adminDB.Username,
		"firstname":    adminDB.Firstname,
		"lastname":     adminDB.Lastname,
		"email":        adminDB.Email,
		"image":        adminDB.Image,
		"tal":          adminDB.Tal,
		"token_link":   adminDB.Token_link,
		"status":       adminDB.Status,
	})

}
