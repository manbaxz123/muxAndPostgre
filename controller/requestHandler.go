package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"muxAndPostgre/dao"
	"muxAndPostgre/model"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"muxAndPostgre/config"
)

func GetWithID(res http.ResponseWriter, req *http.Request) {
	configure := config.GetConfig()
	db, err := dao.CreateConnection(configure.Username, configure.DBname, configure.Password)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	params := mux.Vars(req)
	id, _ := strconv.Atoi(params["id"])
	name, err := dao.SelectWithId(id, db)
	for _, value := range name {
		fmt.Fprintln(res, value)
	}
	// fmt.Fprintln(res, name)
}

func GetAll(res http.ResponseWriter, req *http.Request) {
	var (
		params   = mux.Vars(req)
		username = params["username"]
		password = params["password"]
	)

	configure := config.GetConfig()
	db, err := dao.CreateConnection(configure.Username, configure.DBname, configure.Password)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	name, err := dao.SelectAll(db, username, password)
	// for _, value := range name {
	// 	fmt.Fprintln(res, value)
	// }
	encoder := json.NewEncoder(res)
	encoder.Encode(name)
	// fmt.Fprintln(res, name)
}

func GetSelectUser(res http.ResponseWriter, req *http.Request) {
	var (
		params   = mux.Vars(req)
		username = params["username"]
	)

	configure := config.GetConfig()
	db, err := dao.CreateConnection(configure.Username, configure.DBname, configure.Password)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	name, err := dao.SelectAccount(db, username)
	encoder := json.NewEncoder(res)
	encoder.Encode(name)
}

func GetGroupedUser(res http.ResponseWriter, req *http.Request) {
	var (
		params  = mux.Vars(req)
		groupID = params["groupID"]
	)
	realGroup, _ := strconv.Atoi(groupID)
	configure := config.GetConfig()
	db, err := dao.CreateConnection(configure.Username, configure.DBname, configure.Password)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	name, err := dao.GetGroupUser(db, realGroup)
	encoder := json.NewEncoder(res)
	encoder.Encode(name)
}

func GetAllTask(res http.ResponseWriter, req *http.Request) {
	var (
		params     = mux.Vars(req)
		assignedTo = params["assigned_to"]
	)

	configure := config.GetConfig()
	db, err := dao.CreateConnection(configure.Username, configure.DBname, configure.Password)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	name, err := dao.SelectAllUserTask(db, assignedTo)
	encoder := json.NewEncoder(res)
	encoder.Encode(name)
}

func GetAllTasks(res http.ResponseWriter, req *http.Request) {
	var (
		params     = mux.Vars(req)
		assignedTo = params["assigned_to"]
	)

	configure := config.GetConfig()
	db, err := dao.CreateConnection(configure.Username, configure.DBname, configure.Password)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	name, err := dao.SelectAllUserTasks(db, assignedTo)
	encoder := json.NewEncoder(res)
	encoder.Encode(name)
}

func GetGroupTask(res http.ResponseWriter, req *http.Request) {
	var (
		params     = mux.Vars(req)
		groupID, _ = strconv.Atoi(params["groupID"])
	)

	configure := config.GetConfig()
	db, err := dao.CreateConnection(configure.Username, configure.DBname, configure.Password)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	name, err := dao.GetGroupTask(db, groupID)
	encoder := json.NewEncoder(res)
	encoder.Encode(name)
}

func GetGroupTaskApprove(res http.ResponseWriter, req *http.Request) {
	var (
		params     = mux.Vars(req)
		groupID, _ = strconv.Atoi(params["groupID"])
	)

	configure := config.GetConfig()
	db, err := dao.CreateConnection(configure.Username, configure.DBname, configure.Password)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	name, err := dao.GetGroupTaskApproved(db, groupID)
	encoder := json.NewEncoder(res)
	encoder.Encode(name)
}

func GetGroup(res http.ResponseWriter, req *http.Request) {
	configure := config.GetConfig()
	db, err := dao.CreateConnection(configure.Username, configure.DBname, configure.Password)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	list, err := dao.GetGroupList(db)
	encoder := json.NewEncoder(res)
	encoder.Encode(list)
}

func GetUser(res http.ResponseWriter, req *http.Request) {
	// var (
	// 	params = mux.Vars(req)
	// 	role   = params["role"]
	// )

	configure := config.GetConfig()
	db, err := dao.CreateConnection(configure.Username, configure.DBname, configure.Password)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	name, err := dao.GetUser(db)
	encoder := json.NewEncoder(res)
	encoder.Encode(name)
}

func InsertNewTaskUser(res http.ResponseWriter, req *http.Request) {

	configure := config.GetConfig()
	db, err := dao.CreateConnection(configure.Username, configure.DBname, configure.Password)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var info model.Task
	err = json.NewDecoder(req.Body).Decode(&info)
	if err != nil {
		log.Println(err.Error())
	}

	errr := dao.CreateTaskForUser(db, info.Title, info.Desc, info.CreatedBy, info.AssignedTo, info.GroupID, info.Start, info.Finish)
	if errr != nil {
		log.Println(errr.Error())
	} else {
		fmt.Fprintln(res, "ok no error here")
	}

}

func UpdateTask(res http.ResponseWriter, req *http.Request) {

	configure := config.GetConfig()
	db, err := dao.CreateConnection(configure.Username, configure.DBname, configure.Password)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var info model.Task
	err = json.NewDecoder(req.Body).Decode(&info)
	if err != nil {
		log.Println(err.Error())
	}

	errr := dao.UpdateTask(db, info.TaskID, info.Title, info.Desc, info.AssignedTo, info.Status, info.Start, info.Finish)
	if errr != nil {
		log.Println(errr.Error())
	} else {
		fmt.Fprintln(res, "ok no error here")
	}

}

func InsertNewGroup(res http.ResponseWriter, req *http.Request) {
	configure := config.GetConfig()
	db, err := dao.CreateConnection(configure.Username, configure.DBname, configure.Password)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var info model.Group
	err = json.NewDecoder(req.Body).Decode(&info)
	if err != nil {
		log.Println(err.Error())
	}

	errr := dao.InsertGroup(db, info.GroupName)
	if errr != nil {
		log.Println(errr.Error())
	} else {
		fmt.Fprintln(res, "ok no error here")
	}
}

func InsertNewUser(res http.ResponseWriter, req *http.Request) {
	configure := config.GetConfig()
	db, err := dao.CreateConnection(configure.Username, configure.DBname, configure.Password)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var info model.Account
	err = json.NewDecoder(req.Body).Decode(&info)
	if err != nil {
		log.Println(err.Error())
	}

	errr := dao.CreateUser(db, info.Username, info.Password, info.Name, info.Role, info.GroupID)
	if errr != nil {
		log.Println(errr.Error())
	} else {
		fmt.Fprintln(res, "ok no error here")
	}
}

func UpdateAccount(res http.ResponseWriter, req *http.Request) {
	configure := config.GetConfig()
	db, err := dao.CreateConnection(configure.Username, configure.DBname, configure.Password)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var info model.Account
	err = json.NewDecoder(req.Body).Decode(&info)
	if err != nil {
		log.Println(err.Error())
	}

	errr := dao.UpdateUser(db, info.Username, info.Password, info.Name, info.Role, info.GroupID)
	if errr != nil {
		log.Println(errr.Error())
	} else {
		fmt.Fprintln(res, "ok no error here")
	}
}

func ApproveTask(res http.ResponseWriter, req *http.Request) {
	configure := config.GetConfig()
	db, err := dao.CreateConnection(configure.Username, configure.DBname, configure.Password)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var info model.Task
	err = json.NewDecoder(req.Body).Decode(&info)
	if err != nil {
		log.Println(err.Error())
	}

	errr := dao.ApproveTask(db, info.TaskID, info.Title, info.Desc, info.AssignedTo, info.Status, info.Start, info.Finish)
	if errr != nil {
		log.Println(errr.Error())
	} else {
		fmt.Fprintln(res, "ok no error here")
	}
}

func DeleteTask(res http.ResponseWriter, req *http.Request) {
	configure := config.GetConfig()
	db, err := dao.CreateConnection(configure.Username, configure.DBname, configure.Password)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var info model.Task
	err = json.NewDecoder(req.Body).Decode(&info)
	if err != nil {
		log.Println(err.Error())
	}

	errr := dao.DeclineTask(db, info.TaskID)
	if errr != nil {
		log.Println(errr.Error())
	} else {
		fmt.Fprintln(res, "ok no error here")
	}
}

func ArchiveAccount(res http.ResponseWriter, req *http.Request) {
	configure := config.GetConfig()
	db, err := dao.CreateConnection(configure.Username, configure.DBname, configure.Password)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var info model.Account
	err = json.NewDecoder(req.Body).Decode(&info)
	if err != nil {
		log.Println(err.Error())
	}

	errr := dao.ArchiveUser(db, info.Username)
	if errr != nil {
		log.Println(errr.Error())
	} else {
		fmt.Fprintln(res, "ok no error here")
	}
}
