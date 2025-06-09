package routes

import (
	"database/sql"
	"employee-crud/controller"
	"net/http"
)

func RoutesMap(server *http.ServeMux, db *sql.DB) {

	server.HandleFunc("/", controller.NewIndexEmployee(db))
	server.HandleFunc("/employee/create", controller.HandleCreateEmployee(db))
	server.HandleFunc("/employee/update", controller.HandleUpdateEmployee(db))
	server.HandleFunc("/employee/delete", controller.HandleDeleteEmployee(db))
}
