/*
 * Fast api desc
 *
 *  This is a sample server Petstore server.  You can find out more about      Swagger at [http://swagger.io](http://swagger.io) or on [irc.freenode.net, #swagger](http://swagger.io/irc/).      For this sample, you can use the api key `special-key` to test the authorization     filters.
 *
 * API version: 1.0.0
 * Contact: verycooleagle@gamail.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

func sendJsonMessage (w http.ResponseWriter, message string, code int) {
	w.WriteHeader(code)

	data, _ :=json.Marshal(ApiResponse{
		Code: code,
		Message: message,
	})

	fmt.Fprintf(w, string(data))
}

func logErr (err error) {
	log.Printf("[0;31m DatabaseError: %s [39m \n", err)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		sendJsonMessage(w, "Bad request", http.StatusBadRequest)
		return
	}

	user := User{}

	err = json.Unmarshal(body, &user)

	if err != nil {
		sendJsonMessage(w, "Bad request json", http.StatusBadRequest)
		return
	}

	_, err = user.Save()

	if err != nil {
		sendJsonMessage(w, "Not user save to database", http.StatusInternalServerError)
		logErr(err)

		return
	}

	sendJsonMessage(w, "Successful operation", http.StatusOK)

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "successful operation")
}

func GetUserByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	vars := mux.Vars(r)
	username := vars["username"] // the book title slug

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello " + username)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func LogoutUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
