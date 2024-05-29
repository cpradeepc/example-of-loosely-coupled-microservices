package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

const (
	hostName string = "127.0.0.1"
	portSelf string = "8083"
)

type employee struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
type account struct {
	AccType string  `json:"acctype"`
	Salary  float64 `json:"salary"`
}

var emp []employee
var acc []account

func getDataFromSrv1(w http.ResponseWriter, r *http.Request) {
	log.Printf("len<emps>: %d, len<accs> : %d\n", len(emp), len(acc))
	w.Header().Set("Content-Type", "application")
	var e interface{}
	var ac interface{}
	if emp == nil {
		log.Println("employee data is empty")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("employee data nil"))
		//return
	}
	if acc == nil {
		log.Println("account data is empty")
		//w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("account data nil"))
		//return
	}

	e = emp
	ac = acc
	data := make(map[string]interface{})
	data["employees"] = e
	data["accounts"] = ac
	eJson, _ := json.Marshal(data)
	log.Println("emp acc data :", data)
	//w.WriteHeader(http.StatusOK)
	w.Write(eJson)
}

func setCompData(w http.ResponseWriter, r *http.Request) {
	// r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	name := r.FormValue("name")
	age := r.FormValue("age")
	accType := r.FormValue("acctype")
	salary := r.FormValue("salary")

	log.Println("name, age, accType, salary,r.FormValue(name),r.FormValue(age),r.FormValue(salary),r.FormValue(acctype) >> ", name, age, accType, salary, r.FormValue("name"), r.FormValue("age"), r.FormValue("acctype"), r.FormValue("salary"))

	log.Printf("r.URL.String() : %v, r.URL.Host : %v , r.URL.Path : %v, r.URL.RawQuery  %v>> ", r.URL.String(), r.URL.Host, r.URL.Path, r.URL.RawQuery)
	empData := employee{}
	accData := account{}
	if name != "" || age != "" {
		empData.Name = name
		ageInt, _ := strconv.Atoi(age)
		empData.Age = ageInt
		emp = append(emp, empData)
		log.Println("emp : ", empData)
		w.WriteHeader(http.StatusOK)
		log.Printf("len<emps>: %d, len<accs> : %d\n", len(emp), len(acc))
		return
	}
	if accType != "" || salary != "" {
		accData.AccType = accType
		salInt, _ := strconv.Atoi(salary)
		accData.Salary = float64(salInt)
		acc = append(acc, accData)
		log.Println("emp : ", accData)
		w.WriteHeader(http.StatusOK)
		log.Printf("len<emps>: %d, len<accs> : %d\n", len(emp), len(acc))
		return
	}
	w.WriteHeader(http.StatusBadRequest)
	custResp := make(map[string]any)
	custResp["error"] = "wrong data send"
	err, _ := json.Marshal(custResp)
	log.Printf("len<emps>: %d, len<accs> : %d\n", len(emp), len(acc))
	w.Write(err)
}

func main() {
	fmt.Println("service companee")
	mux := http.NewServeMux()
	mux.HandleFunc("GET /com/get", getDataFromSrv1)
	mux.HandleFunc("POST /com/send", setCompData)
	log.Println("company server is running...")
	log.Fatalln("comp server err: ", http.ListenAndServe("127.0.0.1:8083", mux))
}
