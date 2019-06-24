package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"rest-api-crud-2/helpers"

	"github.com/gorilla/mux"
)

func returnErrorResponse(response http.ResponseWriter, request *http.Request) {
	jsonResponse, err := json.Marshal("ERROR")
	if err != nil {
		panic(err)
	}
	fmt.Println(jsonResponse)
}

// GetUsers ...
func GetUsers(response http.ResponseWriter, request *http.Request) {
	// var httpError = ErrorResponse{
	// 	Code: http.StatusInternalServerError, Message: "It's not you it's me.",
	// }

	jsonResponse := helpers.GetUsersFromDB()

	if jsonResponse == nil {
		// returnErrorResponse(response, request, httpError)
		returnErrorResponse(response, request)
	} else {
		response.Header().Set("Content-Type", "application/json")
		response.Write(jsonResponse)
	}
}

// GetUser ...
func GetUser(response http.ResponseWriter, request *http.Request) {
	// var httpError = ErrorResponse{
	// 	Code: http.StatusInternalServerError, Message: "It's not you it's me.",
	// }

	// userID := mux.Vars(request)["id"]
	params := mux.Vars(request) // Gets params
	userID := params["id"]
	// fmt.Println(params["id"])

	if userID == "" {
		// httpError.Message = "User id can't be empty"
		returnErrorResponse(response, request)
	} else {
		// string to int
		id, err := strconv.Atoi(userID)
		if err != nil {
			// handle error
			fmt.Println(err)
			// os.Exit(2)
		}
		jsonResponse := helpers.GetUserFromDB(id)

		if jsonResponse == nil {
			// returnErrorResponse(response, request, httpError)
			returnErrorResponse(response, request)
		} else {
			response.Header().Set("Content-Type", "application/json")
			response.Write(jsonResponse)
		}
	}
}

// AddUser ...
func AddUser(response http.ResponseWriter, request *http.Request) {
	// var httpError = ErrorResponse{
	// 	Code: http.StatusInternalServerError, Message: "It's not you it's me.",
	// }

	var userDetails helpers.User
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&userDetails)
	defer request.Body.Close()
	// t := strconv.Itoa(123) // int to string

	if err != nil {
		returnErrorResponse(response, request)
	} else {
		// httpError.Code = http.StatusBadRequest
		if userDetails.Name == "" {
			// httpError.Message = "Name can't be empty"
			returnErrorResponse(response, request)
		} else if userDetails.Email == "" {
			// httpError.Message = "Email can't be empty"
			returnErrorResponse(response, request)
		} else if strconv.Itoa(userDetails.Status) == "" {
			// httpError.Message = "Status can't be empty"
			returnErrorResponse(response, request)
		} else {
			isInserted := helpers.InsertUserInDB(userDetails)
			if isInserted {
				GetUsers(response, request)
			} else {
				returnErrorResponse(response, request)
			}
		}
	}
}

// UpdateUser ...
func UpdateUser(response http.ResponseWriter, request *http.Request) {
	// var httpError = ErrorResponse{
	// 	Code: http.StatusInternalServerError, Message: "It's not you it's me.",
	// }
	var userDetails helpers.User
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&userDetails)
	defer request.Body.Close()

	userID := mux.Vars(request)["id"]
	id, err := strconv.Atoi(userID)
	userDetails.ID = id // important, note the change := and =

	if err != nil {
		returnErrorResponse(response, request)
	} else {

		// fmt.Println(userDetails)

		// httpError.Code = http.StatusBadRequest
		if userDetails.Name == "" {
			// httpError.Message = "Name can't be empty"
			returnErrorResponse(response, request)
		} else if userDetails.Email == "" {
			// httpError.Message = "Email can't be empty"
			returnErrorResponse(response, request)
		} else if strconv.Itoa(userDetails.Status) == "" {
			// httpError.Message = "Status can't be empty"
			returnErrorResponse(response, request)
		} else {
			isUpdated := helpers.UpdateUserInDB(userDetails)
			if isUpdated {
				GetUsers(response, request)
			} else {
				returnErrorResponse(response, request)
			}
		}
	}
}

// DeleteUser ...
func DeleteUser(response http.ResponseWriter, request *http.Request) {
	// var httpError = ErrorResponse{
	// 	Code: http.StatusInternalServerError, Message: "It's not you it's me.",
	// }

	userID := mux.Vars(request)["id"]
	if userID == "" {
		// httpError.Message = "User id can't be empty"
		returnErrorResponse(response, request)
	} else {
		// string to int
		id, err := strconv.Atoi(userID)
		if err != nil {
			// handle error
			fmt.Println(err)
			// os.Exit(2)
		}
		isdeleted := helpers.DeleteUserFromDB(id)
		if isdeleted {
			GetUsers(response, request)
		} else {
			returnErrorResponse(response, request)
		}
	}
}
