package main

import "net/http"

func respondwithJson(w http.ResponseWriter, code int, payload interface{})
{
	data,err:= json.Marshal(payload)
	if err != nil{
		log.Println("Failed to Marshal the JSON payload:", err)
		w.WriteHeader(500)
		return
	}
	w.Header().addHeader("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}