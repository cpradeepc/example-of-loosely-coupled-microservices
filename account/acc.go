package main

//server at localhost:8082
import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

const (
	hostName string = "127.0.0.1"
	portNum  string = "8083"
)

type account struct {
	AccType string  `json:"acctype"`
	Salary  float64 `json:"salary"`
}

var acc []account

func getAcc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content_Type", "application/json; charset=utf-8")
	var a interface{}
	if acc == nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("data not found"))
		return
	}
	a = acc
	//log.Println(">> ", u)
	accByte, _ := json.Marshal(a)
	log.Println("acc byte >> ", string(accByte))
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(accByte))
}
func createAcc(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content_Type", "application/json; charset=utf-8")
	q := r.URL.Query()
	acType := q.Get("acctype")
	acSal := q.Get("salary")
	accInt, _ := strconv.Atoi(acSal)
	accVal := account{AccType: acType, Salary: float64(accInt)}
	acc = append(acc, accVal)
	data := fmt.Sprintf("acc create : %s, %d", acType, accInt)
	fmt.Println("acc data:", accVal)
	urlStr := "http://127.0.0.1:8083/com/send"
	params := url.Values{}
	params.Add("acctype", acType)
	params.Add("salary", acSal)
	respNew, err := http.PostForm(urlStr, params)
	if err != nil {
		log.Println("error in new url request by server account: ", err)
		//return
	} else {
		log.Println("create data also in company")
		respNew.Body.Close()
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(data))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /acc/get", getAcc)
	mux.HandleFunc("GET /acc/send", createAcc)
	fmt.Println("account server a is running...")
	log.Fatalln("error in running server: ", http.ListenAndServe("localhost:8082", mux))
}
