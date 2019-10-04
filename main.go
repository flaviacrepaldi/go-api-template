package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	"os"
	"github.com/rs/cors"
	"github.com/gorilla/mux"
	"app/src/models"
	"gopkg.in/mgo.v2"
)

var collection *mgo.Collection

func main() {
	session, err := mgo.Dial("mongo:27017")
	
	if err != nil {
		log.Fatalln(err)
		log.Fatalln("Ops, connection to mongo failed")
		os.Exit(1)
	}

	session.SetMode(mgo.Monotonic, true)

	collection = session.DB("my-api").C("my-collection")

	defer session.Close()
	
	r := mux.NewRouter()
	r.HandleFunc("/post-endpoint", postMethod).Methods("POST")
	r.HandleFunc("/get-endpoint", getMethod).Methods("GET")

	http.ListenAndServe(":8080", cors.AllowAll().Handler(r))

	log.Println("I'm ready and listening on port 8080!")
}

func postMethod(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responseError(w, err.Error(), http.StatusBadRequest)
		return
	}

	myModel := &mymodel.Model{}
	err = json.Unmarshal(data, myModel)
	if err != nil {
		responseError(w, err.Error(), http.StatusBadRequest)
		return
	}

	myModel.SecondBeautifulField = time.Now().UTC()

	if err := collection.Insert(myModel); err != nil {
		responseError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	responseJSON(w, myModel)
}

func getMethod(w http.ResponseWriter, r *http.Request) {
	result := []mymodel.Model{}
	if err := collection.Find(nil).Sort("-second_beautiful_field").All(&result); err != nil {
		responseError(w, err.Error(), http.StatusInternalServerError)
	} else {
		responseJSON(w, result)
	}
}

func responseError(w http.ResponseWriter, message string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}

func responseJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
