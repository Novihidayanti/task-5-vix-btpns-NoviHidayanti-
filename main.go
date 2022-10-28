package main

import (
	"encoding/json"
	"net/http"
)

type Profile struct {
	Nama   string `json:"nama"`
	TTL    string `json:"ttl"`
	Alamat string `json:"alamat"`
	Foto   image
}

var data = []Profile{
	{"Asep", "Bandung", "Bandung"},
	{"Yayat", "Bandung", "Bandung"},
}

func main() {
	http.HandleFunc("/profile", PostProfile)

	http.ListenAndServe(":8000", nill)
}

func PostProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	if r.Method == "POST" {
		res, err := json.Marshal(data)

		if err != nill {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(res)
		return
	}

}
