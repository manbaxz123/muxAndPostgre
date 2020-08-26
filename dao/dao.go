package dao

import (
	"database/sql"
	"fmt"
	"log"
	"muxAndPostgre/model"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func CreateConnection(user, dbname, pass string) (*sql.DB, error) {
	connStr := fmt.Sprintf("root:123456@tcp(localhost:3306)/finalDB?parseTime=true")
	db, err := sql.Open("mysql", connStr)
	log.Println(db)
	return db, err
}
func SelectAll(db *sql.DB, username string, password string) (model.Account, error) {
	account := new(model.Account)
	statement := fmt.Sprintf("SELECT * FROM account where accountID = '%s' AND password = '%s'", username, password)
	rows, err := db.Query(statement)
	log.Println(rows)

	for rows.Next() {
		rows.Scan(&account.Username, &account.Password, &account.Name, &account.Role, &account.GroupID, &account.Status)

	}
	return *account, err
}

func SelectAccount(db *sql.DB, username string) (model.Account, error) {
	account := new(model.Account)
	statement := fmt.Sprintf("SELECT * FROM account where accountID = '%s'", username)
	rows, err := db.Query(statement)
	log.Println(rows)

	for rows.Next() {
		rows.Scan(&account.Username, &account.Password, &account.Name, &account.Role, &account.GroupID, &account.Status)

	}
	return *account, err
}

func SelectAllUserTask(db *sql.DB, assigned_to string) ([]model.Task, error) {
	taskList := make([]model.Task, 0)
	statement := fmt.Sprintf("SELECT taskID, title, description, created_by, assigned_to, approved, status, groupID,start, finish FROM task WHERE assigned_to = '%s' AND (status ='in progress' OR status ='not started') ", assigned_to)
	rows, err := db.Query(statement)
	for rows.Next() {
		task := new(model.Task)
		var start, finish time.Time
		rows.Scan(&task.TaskID, &task.Title, &task.Desc, &task.CreatedBy, &task.AssignedTo, &task.Approved, &task.Status, &task.GroupID, &start, &finish)

		// rows.Scan(&start, &finish)
		task.Start = start.Format("2006-01-02")
		task.Finish = finish.Format("2006-01-02")
		// log.Println(start)
		taskList = append(taskList, *task)
	}
	return taskList, err
}

func SelectAllUserTasks(db *sql.DB, assigned_to string) ([]model.Task, error) {
	taskList := make([]model.Task, 0)
	statement := fmt.Sprintf("SELECT taskID, title, description, created_by, assigned_to, approved, status, groupID,start, finish FROM task WHERE assigned_to = '%s' ", assigned_to)
	rows, err := db.Query(statement)
	for rows.Next() {
		task := new(model.Task)
		var start, finish time.Time
		rows.Scan(&task.TaskID, &task.Title, &task.Desc, &task.CreatedBy, &task.AssignedTo, &task.Approved, &task.Status, &task.GroupID, &start, &finish)

		// rows.Scan(&start, &finish)
		task.Start = start.Format("2006-01-02")
		task.Finish = finish.Format("2006-01-02")
		// log.Println(start)
		taskList = append(taskList, *task)
	}
	return taskList, err
}

// func CreateTaskForUser(db *sql.DB, taskName string, description string, createdBy string, assignedTo string, groupID int) error {
// 	statement := fmt.Sprintf("INSERT INTO task (title, description, created_by, assigned_to, approved, status, groupID) VALUES ('%s', '%s', '%s', '%s', 0, 'not started', '%d')", taskName, description, createdBy, assignedTo, groupID)
// 	_, err := db.Exec(statement)
// 	log.Println(statement)
// 	return err
// }
func CreateTaskForUser(db *sql.DB, taskName string, description string, createdBy string, assignedTo string, groupID int, start string, finish string) error {
	currentTime := time.Now()
	currentTimeDB := currentTime.Format("2006-01-02 15:04:05")
	log.Println(currentTimeDB)
	statement := fmt.Sprintf("INSERT INTO task (title, description, created_by, assigned_to, approved, status, groupID, createdAt, start, finish) VALUES ('%s', '%s', '%s', '%s', 0, 'not started', '%d','%s', '%s', '%s')", taskName, description, createdBy, assignedTo, groupID, currentTimeDB, start, finish)
	_, err := db.Exec(statement)
	log.Println(statement)
	return err
}

func UpdateTask(db *sql.DB, taskID int, title string, desc string, assigned_to string, status string, start string, finish string) error {
	statement := fmt.Sprintf("UPDATE task SET title = '%s', description = '%s', assigned_to = '%s',status = '%s', start = '%s', finish = '%s' WHERE taskID = %d", title, desc, assigned_to, status, start, finish, taskID)
	_, error := db.Exec(statement)
	return error
}

func SelectWithId(id int, db *sql.DB) ([]model.FacebookAudienceSegment, error) {
	nameList := make([]model.FacebookAudienceSegment, 0)
	statement := "SELECT * FROM facebook_audience_segment WHERE id = $1"
	rows, err := db.Query(statement, id)
	facebook := new(model.FacebookAudienceSegment)
	for rows.Next() {

		rows.Scan(&facebook.ID, &facebook.InstanceID, &facebook.AdvertiserID, &facebook.FbAccountID, &facebook.Typed,
			&facebook.CreatedTime, &facebook.ModifiedTime, &facebook.Archived, &facebook.DmpAudience,
			&facebook.ServiceSegment, &facebook.EventSource)
		nameList = append(nameList, *facebook)
	}
	return nameList, err
}

func InsertGroup(db *sql.DB, group string) error {
	statement := fmt.Sprintf("Insert into `group` (name) values ('%s')", group)
	_, err := db.Exec(statement)
	return err
}

func CreateUser(db *sql.DB, username, password, name, role string, groupID int) error {
	statement := fmt.Sprintf("INSERT INTO account VALUES('%s','%s','%s','%s',%d,%d)", username, password, name, role, groupID, 1)
	_, err := db.Exec(statement)
	return err
}
func GetGroupList(db *sql.DB) ([]model.Group, error) {
	statement := "SELECT * from finaldb.group"
	groupList := make([]model.Group, 0)
	rows, err := db.Query(statement)
	for rows.Next() {
		temp := new(model.Group)
		rows.Scan(&temp.GroupID, &temp.GroupName)
		groupList = append(groupList, *temp)
	}

	return groupList, err
}

func GetUser(db *sql.DB) ([]model.Account, error) {
	statement := fmt.Sprintf("SELECT * from account WHERE status = 1")
	accountList := make([]model.Account, 0)
	log.Println(statement)
	rows, err := db.Query(statement)
	for rows.Next() {
		temp := new(model.Account)
		rows.Scan(&temp.Username, &temp.Password, &temp.Name, &temp.Role, &temp.GroupID, &temp.Status)
		accountList = append(accountList, *temp)
	}
	return accountList, err
}

func GetGroupUser(db *sql.DB, groupID int) ([]model.Account, error) {
	statement := fmt.Sprintf("SELECT accountID from account where groupID = %d", groupID)
	accountList := make([]model.Account, 0)
	log.Println(statement)
	rows, err := db.Query(statement)
	for rows.Next() {
		temp := new(model.Account)
		rows.Scan(&temp.Username)
		accountList = append(accountList, *temp)
	}
	return accountList, err
}

func GetGroupTask(db *sql.DB, groupID int) ([]model.Task, error) {
	taskList := make([]model.Task, 0)
	statement := fmt.Sprintf("SELECT taskID, title, description, created_by, assigned_to, approved, status, groupID,start, finish FROM task WHERE groupID = %d AND approved = 0", groupID)
	rows, err := db.Query(statement)
	for rows.Next() {
		task := new(model.Task)
		var start, finish time.Time
		rows.Scan(&task.TaskID, &task.Title, &task.Desc, &task.CreatedBy, &task.AssignedTo, &task.Approved, &task.Status, &task.GroupID, &start, &finish)

		// rows.Scan(&start, &finish)
		task.Start = start.Format("2006-01-02")
		task.Finish = finish.Format("2006-01-02")
		// log.Println(start)
		taskList = append(taskList, *task)
	}
	return taskList, err
}

func GetGroupTaskApproved(db *sql.DB, groupID int) ([]model.Task, error) {
	taskList := make([]model.Task, 0)
	statement := fmt.Sprintf("SELECT taskID, title, description, created_by, assigned_to, approved, status, groupID,start, finish FROM task WHERE groupID = %d AND approved = 1", groupID)
	rows, err := db.Query(statement)
	for rows.Next() {
		task := new(model.Task)
		var start, finish time.Time
		rows.Scan(&task.TaskID, &task.Title, &task.Desc, &task.CreatedBy, &task.AssignedTo, &task.Approved, &task.Status, &task.GroupID, &start, &finish)

		// rows.Scan(&start, &finish)
		task.Start = start.Format("2006-01-02")
		task.Finish = finish.Format("2006-01-02")
		// log.Println(start)
		taskList = append(taskList, *task)
	}
	return taskList, err
}

func UpdateUser(db *sql.DB, username, password, name, role string, groupID int) error {
	statement := fmt.Sprintf("UPDATE account SET password = '%s', name = '%s', role = '%s', groupID = %d WHERE accountID = '%s'", password, name, role, groupID, username)
	_, error := db.Exec(statement)
	return error
}

func ApproveTask(db *sql.DB, taskID int, title string, desc string, assigned_to string, status string, start string, finish string) error {
	statement := fmt.Sprintf("UPDATE task SET approved = 1, title = '%s', description = '%s', assigned_to = '%s',status = '%s', start = '%s', finish = '%s' WHERE taskID = %d", title, desc, assigned_to, status, start, finish, taskID)
	_, error := db.Exec(statement)
	return error
}

// func UpdateTaskStatus(db *sql.DB, taskID int, status string) error {
// 	statement := fmt.Sprintf("UPDATE task SET status = '%s' WHERE taskID = %d", status, taskID)
// 	_, error := db.Exec(statement)
// 	return error
// }

func DeclineTask(db *sql.DB, taskID int) error {
	statement := fmt.Sprintf("DELETE FROM task WHERE taskID = %d", taskID)
	log.Println(statement)
	_, error := db.Exec(statement)
	return error
}

func ArchiveUser(db *sql.DB, username string) error {
	statement := fmt.Sprintf("UPDATE account SET status = 0 WHERE accountID = '%s'", username)
	_, error := db.Exec(statement)
	return error
}

// func GetTaskUser(db *sql.DB, username string) error {
// 	statement := "Select * FROM task WHERE assigned_to = $1"
// 	taskList := make([]model.Task, 0)
// 	task := new(model.Task)
// 	rows, err := db.Query(statement, username)
// 	for rows.Next() {
// 		rows.Scan(&task.TaskID, &task.Title, &task.Desc, &task.CreatedBy, &task.AssignedTo, &task.Approved, &task.Status, &task.gr)
// 	}
// }

// func SelectDB(id, instId, adverId string, db *sql.DB) ([]model.FacebookAudienceSegment, error) {
// 	tempList := make([]model.FacebookAudienceSegment, 0)

// 	statement := "SELECT * FROM facebook_audience_segment WHERE id LIKE $1 AND instance_id LIKE $2" +
// 		" AND advertiser_id LIKE $3"
// 	rows, err := db.Query(statement, id, instId, adverId)
// 	for rows.Next() {
// 		a := ""
// 		facebook := new(model.FacebookAudienceSegment)
// 		rows.Scan(&facebook.ID, &facebook.InstanceID, &facebook.AdvertiserID, &facebook.FbAccountID, &facebook.Typed,
// 			&facebook.CreatedTime, &facebook.ModifiedTime, &facebook.Archived, &facebook.DmpAudience,
// 			&facebook.ServiceSegment, &facebook.EventSource)
// 		fmt.Println(a)
// 		tempList = append(tempList, *facebook)
// 	}
// 	return tempList, err
// }

// func InsertIntoDB(id, adverId, fbID, types string, db *sql.DB) error {
// 	statement := "INSERT into facebook_audience_segment values($1,NULL,$2,$3,$4,NULL,NULL,NULL,NULL,NULL,NULL)"
// 	_, err := db.Exec(statement, id, adverId, fbID, types)
// 	return err

// }

// func DeleteFromDB(id string, db *sql.DB) error {
// 	statement := "DELETE From facebook_audience_segment WHERE id LIKE $1"
// 	_, err := db.Exec(statement, id)
// 	return err

// }

// func UpdateDB(id, instId, adverId string, db *sql.DB) error {
// 	statement := "UPDATE facebook_audience_segment SET advertiser_id = $1 WHERE id LIKE $2 AND instance_id LIKE $3"
// 	_, err := db.Exec(statement, adverId, id, instId)
// 	return err
// }
