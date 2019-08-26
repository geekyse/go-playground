package routes

import (
	"github.com/gorilla/mux"
	"github.com/sql-queries/controller"
	"net/http"
)

func Routes()  {
	r := mux.NewRouter()
	r.HandleFunc("/user", IsLoggedin(controller.CreateUser)).Methods("POST")
	r.HandleFunc("/update", IsLoggedin(controller.UpdateUser)).Methods("PUT")
	r.HandleFunc("/user/login", IsLoggedin(controller.Login)).Methods("POST")
	r.HandleFunc("/user/deactivate", IsLoggedin(controller.DeactivateUser)).Methods("POST")
	r.HandleFunc("/user/logout", IsLoggedin(controller.Logout)).Methods("DELETE")
	http.ListenAndServe(":8000", r)

}

