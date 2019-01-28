package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	sw "test_simple_app_vue_go/back/go"
	"testing"
)

func logErrFatal (place string,err error) {
	if err != nil {
		log.Fatalf("[0;31m Error on %s: %s :[39m", place, err)
	}
}

func clearUsers()   {
	result, err := sw.DB.Exec("DELETE from users")

	logErrFatal("ClearUsers", err)

	count, err := result.RowsAffected()

	logErrFatal("ClearUsers", err)

	fmt.Printf("Count delete %s \n", string(count))
}

func TestCircle(t *testing.T) {

	userName := "testor"

	err := sw.ENV.Init()

	if (err != nil) {
		log.Fatalf("[0;31m Error on initializing envirement: %s :[39m", err)
	}

	err = sw.DatabaseInit()

	if (err != nil) {
		log.Fatalf("[0;31m Error on initializing database connection: %s :[39m", err)
	}

	clearUsers();

	client := &http.Client{}
	router := sw.NewRouter()
	ts := httptest.NewServer(router)
	defer ts.Close()

	Convey("App run", t, func() {

		req, err := http.NewRequest("GET", ts.URL+ "/v2/", nil)
		So(err, ShouldBeNil)

		resp, err := client.Do(req)
		So(err, ShouldBeNil)
		So(resp.StatusCode, ShouldEqual, http.StatusOK)

		body, err := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		bodyString := string(body)
		So(bodyString, ShouldEqual, "Hello World!")
	})

	Convey("Created user", t, func() {

		user := &sw.User{
			Username: userName,
			Email: "test@gom.ua",
			Password:"123",
		}

		data, err := json.Marshal(user)
		So(err, ShouldBeNil)

		buf := bytes.NewBuffer(data)

		req, err := http.NewRequest("POST", ts.URL+ "/v2/user", buf)
		So(err, ShouldBeNil)

		resp, err := client.Do(req)

		So(err, ShouldBeNil)
		defer resp.Body.Close()
		So(resp.StatusCode, ShouldEqual, http.StatusOK)

		actual := &sw.ApiResponse{}
		expect := &sw.ApiResponse{Code:http.StatusOK, Message:sw.MessageOk}
		json.NewDecoder(resp.Body).Decode(actual)

		So(actual.Code, ShouldEqual, expect.Code)
		So(actual.Message, ShouldEqual, expect.Message)
	})

	Convey("Get user", t, func() {

		req, _ := http.NewRequest("GET", ts.URL+ "/v2/user/" + userName, nil)
		// NOTE this !!
		req.Close = true

		req.Header.Add("Accept", "application/json")
		resp, err := client.Do(req)

		So(err, ShouldBeNil)
		So(resp.StatusCode, ShouldEqual, http.StatusOK)

		defer resp.Body.Close()

		userResponse := sw.User{}
		actual := &sw.ApiResponse{Data:&userResponse}
		expect := &sw.ApiResponse{Code:http.StatusOK, Message:sw.MessageOk}

		json.NewDecoder(resp.Body).Decode(actual)

		So(actual.Code, ShouldEqual, expect.Code)
		So(actual.Message, ShouldEqual, expect.Message)
		So(userResponse.Username, ShouldEqual, userName)
	})

	Convey("Get delete user", t, func() {
		req, _ := http.NewRequest("DELETE", ts.URL+ "/v2/user/" + userName, nil)
		// NOTE this !!
		req.Close = true

		req.Header.Add("Accept", "application/json")
		resp, err := client.Do(req)

		So(err, ShouldBeNil)
		So(resp.StatusCode, ShouldEqual, http.StatusOK)

		defer resp.Body.Close()

		actual := &sw.ApiResponse{}
		expect := &sw.ApiResponse{Code:http.StatusOK, Message:sw.MessageOk}

		json.NewDecoder(resp.Body).Decode(actual)

		So(actual.Code, ShouldEqual, expect.Code)
		So(actual.Message, ShouldEqual, expect.Message)
	});
}
