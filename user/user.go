package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

type user struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var userData []user

func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content_Type", "application/json; charset=utf-8")
	var u interface{}
	if userData == nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("data not found"))
		return
	}
	//fmt.Println("userdata : ", userData)
	u = userData
	//log.Println(">> ", u)
	userByte, _ := json.Marshal(u)
	log.Println("user byte >> ", string(userByte))
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(userByte))
}
func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content_Type", "application/json; charset=utf-8")
	r.Header.Set("Content-Type", "application/json")
	q := r.URL.Query()
	nm := q.Get("name")
	ag := q.Get("age")
	ageInt, _ := strconv.Atoi(ag)
	us := user{Name: nm, Age: ageInt}
	userData = append(userData, us)
	data := fmt.Sprintf("user create : %s, %s", nm, ag)
	fmt.Println("userdata:", userData)
	urlStr := "http://127.0.0.1:8083/com/send"
	params := url.Values{}
	params.Add("name", nm)
	params.Add("age", ag)
	respNew, err := http.PostForm(urlStr, params)

	if err != nil {
		log.Println("error in new url request by server user: ", err)
		//return
	} else {
		log.Println("create data also in company")
		respNew.Body.Close()
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(data)) //http://127.0.0.1:8081/a/b?name=raja&age=14 =>     your query data are: raja, 14
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /usr/get", getUser)
	mux.HandleFunc("GET /usr/send", createUser)
	fmt.Println("user server  is running...")
	log.Fatalln("error in running server: ", http.ListenAndServe("localhost:8081", mux))
}
