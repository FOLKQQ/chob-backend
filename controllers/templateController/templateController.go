package templatecontroller

import (
	templatecontroller "backend/models/templateModel"
	"database/sql"
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func Listtemplates(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// ดำเนินการค้นหาข้อมูลทั้งหมดจากฐานข้อมูล
	rows, err := db.Query("SELECT * FROM tbtmp")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// สร้าง slice เพื่อเก็บข้อมูลที่ค้นพบ
	results := []templatecontroller.Template{}
	//
	for rows.Next() {
		// สร้างตัวแปรเพื่อเก็บข้อมูลที่ query ค้นพบ
		template := templatecontroller.Template{}
		// สั่งสแกนข้อมูลจาก query ไปเก็บใน struct ตามชื่อฟิลด์
		err := rows.Scan(
			&template.Id,
			&template.Title,
			&template.Description,
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// เก็บ struct ลงใน slice
		results = append(results, template)
	}

	// ตรวจสอบข้อผิดพลาดหลั
	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// แสดงข้อมูลในรูปแบบ JSON
	jsonData, err := json.Marshal(results)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}

func Addtemplates(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// ดึงข้อมูลจากฟอร์ม
	template := templatecontroller.Template{}
	template.Title = r.FormValue("title")
	template.Description = r.FormValue("description")

	// ดำเนินการเพิ่มข้อมูลลงในฐานข้อมูล
	_, err := db.Exec("INSERT INTO tbtmp (title, description) VALUES (?, ?)",
		template.Title,
		template.Description,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("เพิ่มข้อมูลเรียบร้อยแล้ว"))
}

func Gettemplate(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	//get url parameter id
	id := r.URL.Query().Get("id")

	// ดึงข้อมูลจากฟอร์ม
	template := templatecontroller.Template{}

	// ดำเนินการค้นหาข้อมูลจากฐานข้อมูล
	row := db.QueryRow("SELECT * FROM tbtmp WHERE id = ?", id)
	err := row.Scan(
		&template.Id,
		&template.Title,
		&template.Description,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// แสดงข้อมูลในรูปแบบ JSON
	jsonData, err := json.Marshal(template)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}

func Updatetemplate(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// รับข้อมูลจากฟอร์ม json และแปลงเป็น struct
	template := templatecontroller.Template{}
	err := json.NewDecoder(r.Body).Decode(&template)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return

	}

	// ดำเนินการแก้ไขข้อมูลลงในฐานข้อมูล
	_, err = db.Exec("UPDATE tbtmp SET title = ?, description = ? WHERE id = ?", template.Title, template.Description, template.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("แก้ไขข้อมูลเรียบร้อยแล้ว"))
}

func Deletetemplate(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	//get url parameter id
	id := r.URL.Query().Get("id")

	// ดำเนินการลบข้อมูลลงในฐานข้อมูล
	_, err := db.Exec("DELETE FROM tbtmp WHERE id = ?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("ลบข้อมูลเรียบร้อยแล้ว"))
}

func Listtemplatetasks(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// ดำเนินการค้นหาข้อมูลทั้งหมดจากฐานข้อมูล
	rows, err := db.Query("SELECT * FROM tbtmptask")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// สร้าง slice เพื่อเก็บข้อมูลที่ค้นพบ
	results := []templatecontroller.Templatetask{}
	//
	for rows.Next() {
		// สร้างตัวแปรเพื่อเก็บข้อมูลที่ query ค้นพบ
		templatetask := templatecontroller.Templatetask{}
		// สั่งสแกนข้อมูลจาก query ไปเก็บใน struct ตามชื่อฟิลด์
		err := rows.Scan(
			&templatetask.Id,
			&templatetask.Tmp_id,
			&templatetask.Title,
			&templatetask.Tax_status,
			&templatetask.Tasklist_status,
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// เก็บ struct ลงใน slice
		results = append(results, templatetask)
	}

	// ตรวจสอบข้อผิดพลาดหลังจากวนลูป
	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// แสดงข้อมูลในรูปแบบ JSON
	jsonData, err := json.Marshal(results)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}

func Addtemplatetasks(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// ดึงข้อมูลจากฟอร์ม
	templatetask := templatecontroller.Templatetask{}
	err := json.NewDecoder(r.Body).Decode(&templatetask)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// ดำเนินการเพิ่มข้อมูลลงในฐานข้อมูล
	_, err = db.Exec("INSERT INTO tbtmptask (tmp_id, title, tax_status, tasklist_status) VALUES (?, ?, ?, ?)",
		templatetask.Tmp_id,
		templatetask.Title,
		templatetask.Tax_status,
		templatetask.Tasklist_status,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("เพิ่มข้อมูลเรียบร้อยแล้ว"))
}

func Gettemplatetask(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	//get url parameter id
	id := r.URL.Query().Get("id")

	// ดึงข้อมูลจากฟอร์ม
	templatetask := templatecontroller.Templatetask{}

	// ดำเนินการค้นหาข้อมูลจากฐานข้อมูล
	row := db.QueryRow("SELECT * FROM tbtmptask WHERE id = ?", id)
	err := row.Scan(
		&templatetask.Id,
		&templatetask.Tmp_id,
		&templatetask.Title,
		&templatetask.Tax_status,
		&templatetask.Tasklist_status,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// แสดงข้อมูลในรูปแบบ JSON
	jsonData, err := json.Marshal(templatetask)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}

func Updatetemplatetask(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// รับข้อมูลจากฟอร์ม json และแปลงเป็น struct
	templatetask := templatecontroller.Templatetask{}
	err := json.NewDecoder(r.Body).Decode(&templatetask)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return

	}

	// ดำเนินการแก้ไขข้อมูลลงในฐานข้อมูล
	_, err = db.Exec("UPDATE tbtmptask SET tmp_id = ?, title = ?, tax_status = ?, tasklist_status = ? WHERE id = ?", templatetask.Tmp_id, templatetask.Title, templatetask.Tax_status, templatetask.Tasklist_status, templatetask.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("แก้ไขข้อมูลเรียบร้อยแล้ว"))
}

func Deletetemplatetask(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	//get url parameter id
	id := r.URL.Query().Get("id")

	// ดำเนินการลบข้อมูลลงในฐานข้อมูล
	_, err := db.Exec("DELETE FROM tbtmptask WHERE id = ?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("ลบข้อมูลเรียบร้อยแล้ว"))
}

func Listtemplatesubtasks(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// ดำเนินการค้นหาข้อมูลทั้งหมดจากฐานข้อมูล
	rows, err := db.Query("SELECT * FROM tbtmpsubtask")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// สร้าง slice เพื่อเก็บข้อมูลที่ค้นพบ
	results := []templatecontroller.Templatesubtask{}
	//
	for rows.Next() {
		// สร้างตัวแปรเพื่อเก็บข้อมูลที่ query ค้นพบ
		templatesubtask := templatecontroller.Templatesubtask{}
		// สั่งสแกนข้อมูลจาก query ไปเก็บใน struct ตามชื่อฟิลด์
		err := rows.Scan(
			&templatesubtask.Id,
			&templatesubtask.Tmptask_id,
			&templatesubtask.Title,
			&templatesubtask.Detail,
			&templatesubtask.Subtask_status,
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// เก็บ struct ลงใน slice
		results = append(results, templatesubtask)
	}

	// ตรวจสอบข้อผิดพลาดหลังจากวนลูป
	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// แสดงข้อมูลในรูปแบบ JSON
	jsonData, err := json.Marshal(results)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}

func Addtemplatesubtasks(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// ดึงข้อมูลจากฟอร์ม
	templatesubtask := templatecontroller.Templatesubtask{}
	err := json.NewDecoder(r.Body).Decode(&templatesubtask)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// ดำเนินการเพิ่มข้อมูลลงในฐานข้อมูล
	_, err = db.Exec("INSERT INTO tbtmpsubtask (tmptask_id, title, detail, subtask_status) VALUES (?, ?, ?, ?)",
		templatesubtask.Tmptask_id,
		templatesubtask.Title,
		templatesubtask.Detail,
		templatesubtask.Subtask_status,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("เพิ่มข้อมูลเรียบร้อยแล้ว"))
}

func Gettemplatesubtask(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	//get url parameter id
	id := r.URL.Query().Get("id")

	// ดึงข้อมูลจากฟอร์ม
	templatesubtask := templatecontroller.Templatesubtask{}

	// ดำเนินการค้นหาข้อมูลจากฐานข้อมูล
	row := db.QueryRow("SELECT * FROM tbtmpsubtask WHERE id = ?", id)
	err := row.Scan(
		&templatesubtask.Id,
		&templatesubtask.Tmptask_id,
		&templatesubtask.Title,
		&templatesubtask.Detail,
		&templatesubtask.Subtask_status,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// แสดงข้อมูลในรูปแบบ JSON
	jsonData, err := json.Marshal(templatesubtask)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}

func Updatetemplatesubtask(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// รับข้อมูลจากฟอร์ม json และแปลงเป็น struct
	templatesubtask := templatecontroller.Templatesubtask{}
	err := json.NewDecoder(r.Body).Decode(&templatesubtask)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return

	}

	// ดำเนินการแก้ไขข้อมูลลงในฐานข้อมูล
	_, err = db.Exec("UPDATE tbtmpsubtask SET tmptask_id = ?, title = ?, detail = ?, subtask_status = ? WHERE id = ?", templatesubtask.Tmptask_id, templatesubtask.Title, templatesubtask.Detail, templatesubtask.Subtask_status, templatesubtask.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("แก้ไขข้อมูลเรียบร้อยแล้ว"))
}

func Deletetemplatesubtask(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	//get url parameter id
	id := r.URL.Query().Get("id")

	// ดำเนินการลบข้อมูลลงในฐานข้อมูล
	_, err := db.Exec("DELETE FROM tbtmpsubtask WHERE id = ?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("ลบข้อมูลเรียบร้อยแล้ว"))
}

func Listtemplatesubtasklists(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// ดำเนินการค้นหาข้อมูลทั้งหมดจากฐานข้อมูล
	rows, err := db.Query("SELECT * FROM tbtmpsubtasklist")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// สร้าง slice เพื่อเก็บข้อมูลที่ค้นพบ
	results := []templatecontroller.Templatesubtasklist{}
	//
	for rows.Next() {
		// สร้างตัวแปรเพื่อเก็บข้อมูลที่ query ค้นพบ
		templatesubtasklist := templatecontroller.Templatesubtasklist{}
		// สั่งสแกนข้อมูลจาก query ไปเก็บใน struct ตามชื่อฟิลด์
		err := rows.Scan(
			&templatesubtasklist.Id,
			&templatesubtasklist.Tmpsubtask_id,
			&templatesubtasklist.Title,
			&templatesubtasklist.Subtasklist_status,
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// เก็บ struct ลงใน slice
		results = append(results, templatesubtasklist)
	}

	// ตรวจสอบข้อผิดพลาดหลังจากวนลูป
	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// แสดงข้อมูลในรูปแบบ JSON
	jsonData, err := json.Marshal(results)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}

func Addtemplatesubtasklists(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// ดึงข้อมูลจากฟอร์ม
	templatesubtasklist := templatecontroller.Templatesubtasklist{}
	err := json.NewDecoder(r.Body).Decode(&templatesubtasklist)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// ดำเนินการเพิ่มข้อมูลลงในฐานข้อมูล
	_, err = db.Exec("INSERT INTO tbtmpsubtasklist (tmpsubtask_id, title, subtasklist_status) VALUES (?, ?, ?)",
		templatesubtasklist.Tmpsubtask_id,
		templatesubtasklist.Title,
		templatesubtasklist.Subtasklist_status,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("เพิ่มข้อมูลเรียบร้อยแล้ว"))
}

func Gettemplatesubtasklist(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	//get url parameter id
	id := r.URL.Query().Get("id")

	// ดึงข้อมูลจากฟอร์ม
	templatesubtasklist := templatecontroller.Templatesubtasklist{}

	// ดำเนินการค้นหาข้อมูลจากฐานข้อมูล
	row := db.QueryRow("SELECT * FROM tbtmpsubtasklist WHERE id = ?", id)
	err := row.Scan(
		&templatesubtasklist.Id,
		&templatesubtasklist.Tmpsubtask_id,
		&templatesubtasklist.Title,
		&templatesubtasklist.Subtasklist_status,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// แสดงข้อมูลในรูปแบบ JSON
	jsonData, err := json.Marshal(templatesubtasklist)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}

func Updatetemplatesubtasklist(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// รับข้อมูลจากฟอร์ม json และแปลงเป็น struct
	templatesubtasklist := templatecontroller.Templatesubtasklist{}
	err := json.NewDecoder(r.Body).Decode(&templatesubtasklist)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return

	}

	// ดำเนินการแก้ไขข้อมูลลงในฐานข้อมูล
	_, err = db.Exec("UPDATE tbtmpsubtasklist SET tmpsubtask_id = ?, title = ?, subtasklist_status = ? WHERE id = ?", templatesubtasklist.Tmpsubtask_id, templatesubtasklist.Title, templatesubtasklist.Subtasklist_status, templatesubtasklist.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("แก้ไขข้อมูลเรียบร้อยแล้ว"))
}

func Deletetemplatesubtasklist(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	//get url parameter id
	id := r.URL.Query().Get("id")

	// ดำเนินการลบข้อมูลลงในฐานข้อมูล
	_, err := db.Exec("DELETE FROM tbtmpsubtasklist WHERE id = ?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("ลบข้อมูลเรียบร้อยแล้ว"))
}
