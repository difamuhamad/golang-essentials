package controller

import (
	"database/sql"
	"html/template"
	"net/http"
	"path/filepath"
)

func HandleCreateEmployee(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {

			// spread form data from (application/x-www-form-urlencoded)
			r.ParseForm()
			name := r.Form["name"][0]
			npwp := r.Form["npwp"][0]
			address := r.Form["address"][0]

			// db.Exec( query string (SQL Command), args)
			_, err := db.Exec("INSERT INTO employee (name, npwp, address) VALUES ($1, $2, $3)", name, npwp, address)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}
			http.Redirect(w, r, "/employee", http.StatusSeeOther)
			return
		}
		fp := filepath.Join("views", "create.html")
		tmpl, err := template.ParseFiles(fp)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}

		err = tmpl.Execute(w, nil)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
	}
}
