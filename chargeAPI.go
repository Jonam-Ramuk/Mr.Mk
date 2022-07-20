func returnAllCharge(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Chargingstations)
}

func addNewChargingStation(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var station Chargingstation
	json.Unmarshal(reqBody, &station)
	Chargingstations = append(Chargingstations, station)
	json.NewEncoder(w).Encode(station)
}

