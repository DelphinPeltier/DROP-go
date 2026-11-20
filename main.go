package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)
func writejson(w http.ResponseWriter,status int,payload any){
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("dropzone-version", "1.0.0")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}


type Error struct{
	Messege string `json:"message"`
	error   string `json:"error"`
}
func writeError(w http.ResponseWriter,status int,msg string){
	writejson(w,status,Error{Messege: msg,error: http.StatusText(status)})
}

func getDrops(w http.ResponseWriter, r *http.Request) {
	drops := []string{"Rabat", "Fes", "Casablanca"}
	writejson(w,http.StatusOK,drops)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /drops", getDrops)

	fmt.Println("Serveur lancé sur le port 8080")
	http.ListenAndServe(":8080", mux)

}
