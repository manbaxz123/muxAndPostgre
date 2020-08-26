package main

import (
	"muxAndPostgre/controller"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	subPath := r.PathPrefix("/api/v1").Subrouter()
	subPath.HandleFunc("/get/{id}", controller.GetWithID).Methods("GET")
	subPath.HandleFunc("/getTask/{assigned_to}", controller.GetAllTask).Methods("GET")
	subPath.HandleFunc("/getTasks/{assigned_to}", controller.GetAllTasks).Methods("GET")
	subPath.HandleFunc("/get/{username}/{password}", controller.GetAll).Methods("GET")
	subPath.HandleFunc("/newtask", controller.InsertNewTaskUser).Methods("POST")
	subPath.HandleFunc("/newgroup", controller.InsertNewGroup).Methods("POST")
	subPath.HandleFunc("/newuser", controller.InsertNewUser).Methods("POST")
	subPath.HandleFunc("/getGroup", controller.GetGroup).Methods("GET")
	subPath.HandleFunc("/getUser", controller.GetUser).Methods("GET")
	subPath.HandleFunc("/updateUser", controller.UpdateAccount).Methods("PUT")
	subPath.HandleFunc("/archiveUser", controller.ArchiveAccount).Methods("PUT")
	subPath.HandleFunc("/getUserDetail/{username}", controller.GetSelectUser).Methods("GET")
	subPath.HandleFunc("/getGroupUser/{groupID}", controller.GetGroupedUser).Methods("GET")
	subPath.HandleFunc("/getGroupTask/{groupID}", controller.GetGroupTask).Methods("GET")
	subPath.HandleFunc("/getGroupTaskApprove/{groupID}", controller.GetGroupTaskApprove).Methods("GET")
	subPath.HandleFunc("/approveTask", controller.ApproveTask).Methods("PUT")
	subPath.HandleFunc("/deleteTask", controller.DeleteTask).Methods("POST")
	subPath.HandleFunc("/updateTask", controller.UpdateTask).Methods("PUT")

	http.Handle("/", r)
	http.ListenAndServe(":3500", nil)
}

// func testPOST(w http.ResponseWriter, req *http.Request) {
// 	// req.ParseForm()
// 	// log.Println(req.Form)
// 	// for i, value := range req.Form {
// 	// 	log.Println(i)
// 	// 	log.Println(value)

// 	// }
// }
