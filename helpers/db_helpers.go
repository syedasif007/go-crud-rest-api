package helpers

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

// User ...
type User struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Status int    `json:"status"`
}

// var users []User

// GetUsersFromDB ...
func GetUsersFromDB() []byte {
	var (
		user  User
		users []User
	)

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test")
	rows, err := db.Query("SELECT * FROM go_users")

	if err != nil {
		fmt.Println(err)
		return nil
	}

	for rows.Next() {
		rows.Scan(&user.ID, &user.Name, &user.Email, &user.Status)
		users = append(users, user)
	}
	defer rows.Close()

	jsonResponse, jsonError := json.Marshal(users)
	if jsonError != nil {
		fmt.Println(jsonError)
		return nil
	}
	return jsonResponse
}

// GetUserFromDB ...
func GetUserFromDB(userID int) []byte {
	var (
		user  User
		users []User
	)

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test")
	rows, err := db.Query("SELECT * FROM go_users WHERE data_id=?", userID)
	// rows, err := db.QueryRow("SELECT * FROM go_users WHERE data_id=?", userID) // limit to 1 row

	if err != nil {
		fmt.Println(err)
		return nil
	}

	for rows.Next() {
		rows.Scan(&user.ID, &user.Name, &user.Email, &user.Status)
		users = append(users, user)
	}
	defer rows.Close()

	jsonResponse, jsonError := json.Marshal(users)
	if jsonError != nil {
		fmt.Println(jsonError)
		return nil
	}
	return jsonResponse
}

// InsertUserInDB ...
func InsertUserInDB(userDetails User) bool {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test")
	stmt, err := db.Prepare("INSERT into go_users SET name=?, email=?, status=?")
	if err != nil {
		fmt.Println(err)
		return false
	}
	_, queryError := stmt.Exec(userDetails.Name, userDetails.Email, userDetails.Status)
	if queryError != nil {
		fmt.Println(queryError)
		return false
	}
	return true
}

// UpdateUserInDB ...
func UpdateUserInDB(userDetails User) bool {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test")
	// stmt, err := db.Prepare("UPDATE users SET status=? WHERE id=?")
	stmt, err := db.Prepare("UPDATE go_users SET name=?, email=?, status=? WHERE data_id=?")
	if err != nil {
		fmt.Println(err)
		return false
	}
	_, queryError := stmt.Exec(userDetails.Name, userDetails.Email, userDetails.Status, userDetails.ID)
	if queryError != nil {
		fmt.Println(queryError)
		return false
	}
	return true
}

// DeleteUserFromDB ...
func DeleteUserFromDB(userID int) bool {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test")
	stmt, err := db.Prepare("DELETE FROM go_users WHERE data_id=?")

	if err != nil {
		fmt.Println(err)
		return false
	}
	_, queryError := stmt.Exec(userID)
	if queryError != nil {
		fmt.Println(queryError)
		return false
	}
	return true
}
