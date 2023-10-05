package tagcontroller

import (
	"backend/models/tagModel"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func CreateTag(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	tag := tagModel.Tag{}
	json.NewDecoder(r.Body).Decode(&tag)
	// บันทึกข้อมูลผู้ใช้ในฐานข้อมูล
	_, err := db.Exec("INSERT INTO tbtag (company_id, name, color) VALUES(?, ?, ?)", tag.Company_id, tag.Name, tag.Color)
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "New tag was created")
}

func ListTag(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var tags []tagModel.Tag
	result, err := db.Query("SELECT * FROM tbtag")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {
		var tag tagModel.Tag
		err := result.Scan(&tag.Id, &tag.Company_id, &tag.Name, &tag.Color)
		if err != nil {
			panic(err.Error())
		}
		tags = append(tags, tag)
	}

	json.NewEncoder(w).Encode(tags)
}

func UpdateTag(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	tag := tagModel.Tag{}
	json.NewDecoder(r.Body).Decode(&tag)
	// บันทึกข้อมูลผู้ใช้ในฐานข้อมูล
	_, err := db.Exec("UPDATE tbtag SET company_id=?, name=?, color=? WHERE id=?", tag.Company_id, tag.Name, tag.Color, tag.Id)
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "Tag was updated")
}

func DeleteTag(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	tag := tagModel.Tag{}
	json.NewDecoder(r.Body).Decode(&tag)
	// ลบข้อมูลผู้ใช้ในฐานข้อมูล
	_, err := db.Exec("DELETE FROM tbtag WHERE id=?", tag.Id)
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "Tag was deleted")
}
