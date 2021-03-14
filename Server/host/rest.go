package host

import (
	"fmt"
	"log"
	"encoding/json"
	"net/http"
	"io/ioutil"
	"strconv"
	"strings"
	"github.com/gorilla/mux"	
	data "../data/stores"
	products "../data/products"
	"../reports"
)

var MainVector *data.Vector
var MainInventory *products.Inventarios

func Request() {
	fmt.Println("Listening And Serving ...")
	myrouter := mux.NewRouter().StrictSlash(true)
	myrouter.HandleFunc("/cargartienda", setStores).Methods("POST")
	myrouter.HandleFunc("/getArreglo", getArreglo).Methods("GET")
	myrouter.HandleFunc("/TiendaEspecifica", searchByName).Methods("POST")
	myrouter.HandleFunc("/id/{id}", searchByPosition).Methods("GET")
	myrouter.HandleFunc("/Eliminar", deleteStore).Methods("POST")
	myrouter.HandleFunc("/setInventarios", setInventarios).Methods("POST")
	log.Fatal(http.ListenAndServe(":3000", myrouter))
}

//Ingrsa las tiendas POST
func setStores(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var response data.Data
	err := json.Unmarshal(body, &response)
	Error(err)
	if err == nil {
		// if len(MainVector.Vector) == 0 {
		// 	MainVector.GetVector(response)
		// } else {
		// 	auxVector := data.NewVector()
		// 	auxVector.GetVector(response)
	
		// 	//reports.SaveVector(*auxVector)
	
		// 	MainVector = data.JoinVectors(*MainVector, *auxVector)
		// }

		MainVector.GetVector(response)
		reports.SaveVector(*MainVector)
		fmt.Fprintf(w, "Seted")
		fmt.Println("Seted")
	}
}

func getArreglo(w http.ResponseWriter, r *http.Request) {
	reports.GetComplete(MainVector)
	fmt.Println("Archivo creado")
	fmt.Fprintf(w, "Archivo creado")
}

func searchByName(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	
	response := data.NewVstore()
	err := json.Unmarshal(body, &response)
	Error(err)
	if err == nil {
		result := reports.GetSearchByStore(response, MainVector)

		fmt.Fprintf(w, result)
		fmt.Println(result)
	}
}

func searchByPosition(w http.ResponseWriter, r *http.Request) {
	url := []byte(r.URL.Path)	
	idURL := string(url[4:])
	id, err := strconv.Atoi(idURL)
	Error(err)
	if err == nil {
		result := reports.GetSearchByPosition(id, *MainVector)

		fmt.Println(result)
		fmt.Fprintf(w, result)
	}
}

func deleteStore(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	auxBody := strings.ReplaceAll(string(body), "\"Categoria\":", "\"Departamento\":")
	
	response := data.NewVstore()
	err := json.Unmarshal([]byte(auxBody), &response)
	Error(err)

	if err == nil {
		result := reports.DeleteStore(response, MainVector)
		reports.SaveVector(*MainVector)
		fmt.Fprintf(w, result)
		fmt.Println(result)
	}	
}

func setInventarios(w http.ResponseWriter, r *http.Request) {
	body ,_ := ioutil.ReadAll(r.Body)
	response := products.NewInventarios()
	err := json.Unmarshal(body, &response)

	if !Error(err) {
		MainInventory = response

		fmt.Println("Seted")
		fmt.Fprintf(w, "Seted")
	}
}

func Error(err error) bool {
	if err != nil {
		fmt.Println("error:", err)
		return true
	}
	return false
}