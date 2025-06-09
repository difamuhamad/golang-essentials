package controller

import (
	"database/sql"
	"html/template"
	"net/http"
	"path/filepath"
)

func HandleUpdateEmployee(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// get id from query param
		id := r.URL.Query().Get("id")

		if r.Method == "POST" {
			r.ParseForm()
			name := r.Form["name"][0]
			npwp := r.Form["npwp"][0]
			address := r.Form["address"][0]

			// db.Exec( query string (SQL Command), args)
			_, err := db.Exec("UPDATE employee SET name= $1, npwp= $2, address= $3 WHERE id= $4", name, npwp, address, id)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}
			http.Redirect(w, r, "/employee", http.StatusSeeOther)
			return
		} else if r.Method == "GET" {
			row := db.QueryRow("SELECT name, npwp, address FROM employee WHERE id= $1 ", id)

			if row.Err() != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(row.Err().Error()))
			}

			// fill employee struct
			var employee Employee
			err := row.Scan(
				&employee.Name,
				&employee.Npwp,
				&employee.Address,
			)
			employee.Id = id

			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte(err.Error()))
			}

			fp := filepath.Join("views", "update.html")
			tmpl, err := template.ParseFiles(fp)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
			}

			data := make(map[string]any)
			data["employee"] = employee
			err = tmpl.Execute(w, data)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
			}
		}
	}
}
