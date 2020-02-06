package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"../db"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

var client = db.GetClient()

func getBlock(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	vars := mux.Vars(r)
	blockID, err := strconv.ParseInt(vars["blockID"], 10, 64)
	if err != nil {
		log.Println("couldnt covert blockID to int", err)
		return
	}

	block := db.GetBlockFromDB(client, bson.M{"blockid": blockID})
	j, err := json.Marshal(block)
	if err != nil {
		log.Fatalln("unable to parse json ", err)
	}
	w.Write([]byte(j))
}

func getCount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write([]byte(strconv.FormatInt(db.GetCount(client), 10)))
}
func Handler() {
	r := mux.NewRouter()
	r.HandleFunc("/api/getBlock/{blockID}", getBlock)
	r.HandleFunc("/api/getCount", getCount)
	http.ListenAndServe(":8080", r)
}
