package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/minmax", func(w http.ResponseWriter, r *http.Request) {
		getBody := json.NewDecoder(r.Body)

		var body []int

		err := getBody.Decode(&body)
		if err != nil {
			log.Print(err.Error())
			return
		}

		min := 99

		max := 0

		for _, value := range body {
			if value < min {
				min = value
			}
			if value > max {
				max = value
			}
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(fmt.Sprintf("%s%d\n%s%d", "min:", min, "max:", max)))

	}).Methods("POST")

	r.HandleFunc("/string/{condition}/{msg}", func(w http.ResponseWriter, r *http.Request) {
		condition := mux.Vars(r)["condition"]
		message := mux.Vars(r)["msg"]

		conditionLength := len(condition)

		resultFlag := "0"

		for index := 0; index < len(message)-conditionLength-1; index++ {
			log.Print(condition, "---", message[index:index+conditionLength])
			if condition == message[index:index+conditionLength] {
				resultFlag = "1"
				break
			}
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(resultFlag))

	}).Methods("GET")

	r.HandleFunc("/bracket/{msg}", func(w http.ResponseWriter, r *http.Request) {
		message := mux.Vars(r)["msg"]

		openCount := 0

		closeCount := 0

		resultFlag := "1"

		for index := 0; index < len(message); index++ {
			if message[index:index+1] == "(" {
				openCount++
			}
			if message[index:index+1] == ")" && openCount < (closeCount+1) {
				resultFlag = "0"
				break
			} else if message[index:index+1] == ")" {
				closeCount++
			}
		}

		if openCount != closeCount {
			resultFlag = "0"
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(resultFlag))

	}).Methods("GET")

	log.Print("running on :10800")

	http.ListenAndServe(":10800", r)
}
