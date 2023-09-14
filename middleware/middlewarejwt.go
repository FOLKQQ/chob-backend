package middlewarejwt

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

const SecretKey = "chob-backend-2023"

// check Authorization header and validate token if exist in header and valid set user in context and call next handler else return 401 status code and error message to client

func ValidateToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// ตรวจสอบว่ามีการส่ง access token มาหรือไม่
		if r.Header.Get("Authorization") == "" {
			http.Error(w, "Authorization header required", http.StatusUnauthorized)
			return
		}
		// ดึงค่า token จาก header ของ request
		tokenString := r.Header.Get("Authorization")[7:]
		// สร้างตัวแปรเพื่อเก็บผลลัพธ์ที่ได้จากการตรวจสอบ token
		claims := jwt.MapClaims{}
		// สร้างตัวแปรเพื่อเก็บผลลัพธ์ที่ได้จากการตรวจสอบ token
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(SecretKey), nil
		})
		// ตรวจสอบว่า token มีข้อผิดพลาดหรือไม่
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		// ตรวจสอบว่า token ถูก validate และไม่มีข้อผิดพลาด
		if !token.Valid {
			http.Error(w, "token is invalid", http.StatusUnauthorized)
			return
		}

		// call next handler
		next.ServeHTTP(w, r)
	})
}
