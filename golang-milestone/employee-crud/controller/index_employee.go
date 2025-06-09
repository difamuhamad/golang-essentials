package controller

import (
	"database/sql"
	"html/template"
	"net/http"
	"path/filepath"
)

type Employee struct {
	Id      string // for safety
	Name    string
	Npwp    string
	Address string
}

func NewIndexEmployee(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		// query to database
		rows, err := db.Query("SELECT id, name, npwp, address FROM employee")
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(err.Error()))
		}

		//	Slice of employee
		var employees []Employee

		for rows.Next() {
			var employee Employee

			rows.Scan(
				&employee.Id,
				&employee.Name,
				&employee.Npwp,
				&employee.Address,
			)
			employees = append(employees, employee)
		}

		fp := filepath.Join("views", "index.html")
		tmpl, err := template.ParseFiles(fp)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		data := make(map[string]any)
		data["employees"] = employees
		// fmt.Println("Employee data:", data)

		// Execute( writer, data)
		err = tmpl.Execute(w, data)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
	}
}
