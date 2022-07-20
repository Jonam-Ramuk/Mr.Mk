package main

import (
   "fmt"
   "net/http"
   "time"
   "github.com/gorilla/mux"
)

var Chargingstations []Chargingstation

var StartChargingstations []StartChargingstation

type Chargingstation struct {
  
	Id     string
	Energy string
    Typec  string
}

type StartChargingstation struct {
	Id string
	Vbc string
	Cvc string
	CST time.Time
}

func (w StartChargingstation) string() {
	fmt.Println(w.CST.String())
	w.CST = time.Now()
	fmt.Println(w.CST.String())
}

func handleReqs() {
	w := mux.NewRouter()
	w.HandleFunc("/all", returnAllChargings)
	w.HandleFunc("/post", addNewChargingStation)
	w.HandleFunc("/post", addNewStartChargingstation)
	http.ListenAndServe("127.0.0.1:5555", w)
}