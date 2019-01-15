package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// -----------------------------------------------------
// ----------------- Data Structures -------------------
// -----------------------------------------------------

type app struct {
	Router *mux.Router
}

type ResponseUpload struct {
	URL  string `json:"url_file"`
	Hash string `json:"hash_name"`
}

// -----------------------------------------------------
// ---------------- Server Conections ------------------
// -----------------------------------------------------

func (a *app) initializeRoutes() {
	// AWS Tools API
	a.Router.HandleFunc("/", a.helloworld).Methods("GET", "OPTIONS")
}

func (a *app) serve(addr string) {
	handler := cors.AllowAll().Handler(a.Router)
	http.ListenAndServe(addr, handler)
}

func (a *app) initialize() {

	// Server Mux

	a.Router = mux.NewRouter()
	a.initializeRoutes()
	a.serve(":3535")
}

// -----------------------------------------------------
// ---------------- Server Functions  ------------------
// -----------------------------------------------------

func (a *app) helloworld(w http.ResponseWriter, r *http.Request) {
	log.Print("service helloworld consult by: " + r.RemoteAddr)

	checkError(nil, w)
	json.NewEncoder(w).Encode("Hello World From Dev!")
}

// -----------------------------------------------------
// ----------------- Helper Functions ------------------
// -----------------------------------------------------

// Este metodo permite administrar los errores que puedan
// surgir al momento de generar alguna consulta en el API
// y envia la respuesta en el cuerpo del request
func checkError(err error, w http.ResponseWriter) bool {
	if err != nil {
		log.Print(err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		s := err.Error()
		w.Write([]byte(s))
		return true
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		return false
	}
}

// -----------------------------------------------------
// ------------------ Main Structure -------------------
// -----------------------------------------------------
func main() {

	// cntxt := &daemon.Context{
	// 	PidFileName: "PID_FILE.pid",
	// 	PidFilePerm: 0644,
	// 	LogFileName: "LOG_APP.log",
	// 	LogFilePerm: 0640,
	// 	WorkDir:     "./",
	// 	Umask:       027,
	// 	Args:        []string{"[go-daemon sample]"},
	// }

	// d, err := cntxt.Reborn()
	// if err != nil {
	// 	log.Fatal("Unable to run: ", err)
	// }
	// if d != nil {
	// 	return
	// }
	// defer cntxt.Release()

	log.Print("- - - - - - - - - - - - - - -")
	log.Print("		  Daemon Started		")
	log.Print("- - - - - - - - - - - - - - -")

	// ---------- Start System ----------

	a := app{}

	a.initialize()

	log.Print("- - - - - - - - - - - - - - -")
	log.Print("			Run System			")
	log.Print("- - - - - - - - - - - - - - -")
}
