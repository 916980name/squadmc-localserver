package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"flag"
)

var (
	rtss_cli_path = flag.String("clipath", ".\\saku-rtss-cli\\Saku RTSS CLI.exe", "The RTSS cli exe file location")
	port = flag.Int("port", 16900, "The server runs in local port")
)

type Motar struct {
	Angle string `json:"angle"`
	Mil   string `json:"mil"`
}

func bbb(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method == "OPTIONS" {
		fmt.Fprintf(w, "Options %v\n", "ok")
		return
	}

	var list []Motar
	err := json.NewDecoder(r.Body).Decode(&list)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Print the map
	fmt.Println("Request body data:")
	for key, value := range list {
		fmt.Printf("%d: ang:%s, mil: %s\n", key, value.Angle, value.Mil)
	}
	fmt.Fprintf(w, "%v\n", "ok")
	execSakuRTSSCli_text(list)
}

func Run_Simple_Server() {
	execSakuRTSSCli_clean()
	http.HandleFunc("/mil", bbb)

	fmt.Printf("server listen on port %d\n", *port)
	http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	// (*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers, Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	// (*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func main() {
	flag.Parse()
	Run_Simple_Server()
}