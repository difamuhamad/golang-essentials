package controller

import (
	"database/sql"
	"net/http"
)

func HandleDeleteEmployee(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// get id from query param
		id := r.URL.Query().Get("id")

		// result, err := db.Exec
		_, err := db.Exec("DELETE FROM employee WHERE id= $1", id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		http.Redirect(w, r, "/employee", http.StatusSeeOther)
	}
}
