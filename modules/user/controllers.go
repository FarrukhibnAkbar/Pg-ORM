package user

import(
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)


func GET_ID (w http.ResponseWriter, r *http.Request){
	
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)

	// user := User {}

	encoder := json.NewEncoder(w)

	u, _ := GetById(vars["Id"])

	encoder.Encode(u)
}

// func POST (w http.ResponseWriter, r *http.Request) {
	
// 	body, err := ioutil.ReadAll(r.Body)

// 	if err != nil{
// 		log.Fatalf("%v", err)
// 	}

// 	u := User {}

// 	json.Unmarshal(body, &u)

// 	encoder.Encode(GetMany())

// }