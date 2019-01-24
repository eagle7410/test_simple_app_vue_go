package main

import (
	"bytes"
	"encoding/json"
	. "github.com/smartystreets/goconvey/convey"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	sw "back/go"
	"testing"
)


func TestCircle(t *testing.T) {

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
			Email: "test@gom.ua",
			Password:"123",
		}

		data, err := json.Marshal(user)
		So(err, ShouldBeNil)

		buf := bytes.NewBuffer(data)

		t.Log(ts.URL)
		req, err := http.NewRequest("POST", ts.URL+ "/v2/user", buf)
		So(err, ShouldBeNil)

		resp, err := client.Do(req)

		So(err, ShouldBeNil)
		So(resp.StatusCode, ShouldEqual, http.StatusOK)

		actual := &sw.ApiResponse{}
		expect := &sw.ApiResponse{Code:http.StatusOK, Message:"Successful operation"}
		json.NewDecoder(resp.Body).Decode(actual)

		So(actual.Code, ShouldEqual, expect.Code)
		So(actual.Message, ShouldEqual, expect.Message)
	})
}
