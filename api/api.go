package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//Declare a global array of Credentials
//See credentials.go

/*YOUR CODE HERE*/
var database []Credentials = []Credentials{}

func RegisterRoutes(router *mux.Router) error {

	/*

		Fill out the appropriate get methods for each of the requests, based on the nature of the request.

		Think about whether you're reading, writing, or updating for each request


	*/

	router.HandleFunc("/api/getCookie", getCookie).Methods(http.MethodGet)
	router.HandleFunc("/api/getQuery", getQuery).Methods(http.MethodGet)
	router.HandleFunc("/api/getJSON", getJSON).Methods(http.MethodGet)

	router.HandleFunc("/api/signup", signup).Methods(http.MethodPost)
	router.HandleFunc("/api/getIndex", getIndex).Methods(http.MethodGet)
	router.HandleFunc("/api/getpw", getPassword).Methods(http.MethodGet)
	router.HandleFunc("/api/updatepw", updatePassword).Methods(http.MethodPut)
	router.HandleFunc("/api/deleteuser", deleteUser).Methods(http.MethodDelete)

	return nil
}

func getCookie(response http.ResponseWriter, request *http.Request) {

	/*
		Obtain the "access_token" cookie's value and write it to the response

		If there is no such cookie, write an empty string to the response
	*/

	/*YOUR CODE HERE*/
	// Requesting cookie, which holds the access_token
	cookie, err := request.Cookie("access_token")
	if err != nil {
		fmt.Fprintf(response, "")
		return
	}
	fmt.Fprintf(response, cookie.Value)

}

func getQuery(response http.ResponseWriter, request *http.Request) {

	/*
		Obtain the "userID" query paramter and write it to the response
		If there is no such query parameter, write an empty string to the response
	*/

	/*YOUR CODE HERE*/
	userid := request.URL.Query().Get("userID")
	fmt.Fprint(response, userid)
}

func getJSON(response http.ResponseWriter, request *http.Request) {

	/*
		Our JSON file will look like this:

		{
			"username" : <username>,
			"password" : <password>
		}

		Decode this json file into an instance of Credentials.

		Then, write the username and password to the response, separated by a newline.

		Make sure to error check! If there are any errors, call http.Error(), and pass in a "http.StatusBadRequest" What kind of errors can we expect here?
	*/

	/*YOUR CODE HERE*/
	userInfo := Credentials{}
	err := json.NewDecoder(request.Body).Decode(&userInfo)
	// Error may be in the request or in the credentials given.
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	userName, userPass := userInfo.Username, userInfo.Password

	if len(userName) == 0 || len(userPass) == 0 {
		http.Error(response, errors.New("bad credentials").Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(response, userName+"\n")
	fmt.Fprintf(response, userPass)
}

func signup(response http.ResponseWriter, request *http.Request) {

	/*
		Our JSON file will look like this:

		{
			"username" : <username>,
			"password" : <password>
		}

		Decode this json file into an instance of Credentials.

		Then store it ("append" it) to the global array of Credentials.

		Make sure to error check! If there are any errors, call http.Error(), and pass in a "http.StatusBadRequest" What kind of errors can we expect here?
	*/

	/*YOUR CODE HERE*/
	userInfo := Credentials{}
	err := json.NewDecoder(request.Body).Decode(&userInfo)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	userName, userPass := userInfo.Username, userInfo.Password

	if len(userName) == 0 || len(userPass) == 0 {
		http.Error(response, errors.New("bad credentials").Error(), http.StatusBadRequest)
		return
	}

	database = append(database, userInfo)

	response.WriteHeader(201)
	return
}

func getIndex(response http.ResponseWriter, request *http.Request) {

	/*
		Our JSON file will look like this:

		{
			"username" : <username>
		}


		Decode this json file into an instance of Credentials. (What happens when we don't have all the fields? Does it matter in this case?)

		Return the array index of the Credentials object in the global Credentials array

		The index will be of type integer, but we can only write strings to the response. What library and function was used to get around this?

		Make sure to error check! If there are any errors, call http.Error(), and pass in a "http.StatusBadRequest" What kind of errors can we expect here?
	*/

	/*YOUR CODE HERE*/
	userInfo := Credentials{}
	err := json.NewDecoder(request.Body).Decode(&userInfo)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	userName := userInfo.Username

	if len(userName) == 0 {
		http.Error(response, errors.New("bad credentials").Error(), http.StatusBadRequest)
		return
	}

	for i := 0; i < len(database); i++ {
		if database[i].Username == userName {
			fmt.Fprintf(response, strconv.Itoa(i))
			return
		}
	}
	// If we exit the loop the the user was NOT in our database
	http.Error(response, errors.New("user doesn't exist").Error(), http.StatusBadRequest)
	return
}

func getPassword(response http.ResponseWriter, request *http.Request) {

	/*
		Our JSON file will look like this:

		{
			"username" : <username>
		}


		Decode this json file into an instance of Credentials. (What happens when we don't have all the fields? Does it matter in this case?)

		Write the password of the specific user to the response

		Make sure to error check! If there are any errors, call http.Error(), and pass in a "http.StatusBadRequest" What kind of errors can we expect here?
	*/

	/*YOUR CODE HERE*/
	userInfo := Credentials{}
	err := json.NewDecoder(request.Body).Decode(&userInfo)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	userName := userInfo.Username

	if len(userName) == 0 {
		http.Error(response, errors.New("bad credentials").Error(), http.StatusBadRequest)
		return
	}

	for i := 0; i < len(database); i++ {
		if database[i].Username == userName {
			fmt.Fprintf(response, database[i].Password)
			return
		}
	}
	// If we exit the loop the the user was NOT in our database
	http.Error(response, errors.New("user doesn't exist").Error(), http.StatusBadRequest)
	return
}

func updatePassword(response http.ResponseWriter, request *http.Request) {

	/*
		Our JSON file will look like this:

		{
			"username" : <username>,
			"password" : <password,
		}


		Decode this json file into an instance of Credentials.

		The password in the JSON file is the new password they want to replace the old password with.

		You don't need to return anything in this.

		Make sure to error check! If there are any errors, call http.Error(), and pass in a "http.StatusBadRequest" What kind of errors can we expect here?
	*/

	/*YOUR CODE HERE*/
	userInfo := Credentials{}
	err := json.NewDecoder(request.Body).Decode(&userInfo)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	userName, userPass := userInfo.Username, userInfo.Password

	if len(userName) == 0 || len(userPass) == 0 {
		http.Error(response, errors.New("bad credentials").Error(), http.StatusBadRequest)
		return
	}

	for i := 0; i < len(database); i++ {
		if database[i].Username == userName {
			database[i].Password = userPass
			return
		}
	}
	// If we exit the loop the the user was NOT in our database
	http.Error(response, errors.New("user doesn't exist").Error(), http.StatusBadRequest)
	return
}

func deleteUser(response http.ResponseWriter, request *http.Request) {

	/*
		Our JSON file will look like this:

		{
			"username" : <username>,
			"password" : <password,
		}


		Decode this json file into an instance of Credentials.

		Remove this user from the array. Preserve the original order. You may want to create a helper function.

		This wasn't covered in lecture, so you may want to read the following:
			- https://gobyexample.com/slices
			- https://www.delftstack.com/howto/go/how-to-delete-an-element-from-a-slice-in-golang/

		Make sure to error check! If there are any errors, call http.Error(), and pass in a "http.StatusBadRequest" What kind of errors can we expect here?
	*/
	/*YOUR CODE HERE*/
	userInfo := Credentials{}
	err := json.NewDecoder(request.Body).Decode(&userInfo)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	userName, userPass := userInfo.Username, userInfo.Password

	if len(userName) == 0 || len(userPass) == 0 {
		http.Error(response, errors.New("bad credentials").Error(), http.StatusBadRequest)
		return
	}

	for i := 0; i < len(database); i++ {
		if database[i].Username == userName && database[i].Password == userPass {
			// Database is now database[0 ... i-1] U database[i+1 ...]
			database = append(database[:i], database[i+1:]...)
			return
		}
	}
	// If we exit the loop the the user was NOT in our database
	http.Error(response, errors.New("user doesn't exist").Error(), http.StatusBadRequest)
	return
}
