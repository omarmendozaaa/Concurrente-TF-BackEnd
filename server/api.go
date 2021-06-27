package server

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type api struct {
	router http.Handler
}

type Server interface {
	Router() http.Handler
}

var Centroids []Node

func middlewareCors(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, req *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
			next.ServeHTTP(w, req)
		})
}
func enableCORS(router *mux.Router) {
	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
	}).Methods(http.MethodOptions)
	router.Use(middlewareCors)
}

var DataSetNodes []Node

func leerCSVdesdeURL(url string) ([][]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	reader := csv.NewReader(resp.Body)
	reader.Comma = ','
	data, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func New() Server {

	a := &api{}
	r := mux.NewRouter()
	//Habilitamos los CORS
	enableCORS(r)
	url := "https://raw.githubusercontent.com/omarmendozaaa/Concurrente-TA2-BackEnd/master/server/N_DataSetFeminicidio.csv"

	fmt.Println("Se cargó satisfactoriamente el DataSet")

	DataSetLines, err := leerCSVdesdeURL(url)
	if err != nil {
		panic(err)
	}
	if err != nil {
		fmt.Println(err)
	}
	for _, line := range DataSetLines {

		param1, _ := strconv.ParseFloat(line[1], 64)
		param2, _ := strconv.ParseFloat(line[5], 64)
		param3, _ := strconv.ParseFloat(line[6], 64)
		param4, _ := strconv.ParseFloat(line[8], 64)
		param6, _ := strconv.ParseFloat(line[10], 64)
		param7, _ := strconv.ParseFloat(line[11], 64)
		param8, _ := strconv.ParseFloat(line[12], 64)
		param9, _ := strconv.ParseFloat(line[13], 64)
		param10, _ := strconv.ParseFloat(line[14], 64)
		param11, _ := strconv.ParseFloat(line[15], 64)
		param12, _ := strconv.ParseFloat(line[16], 64)
		param13, _ := strconv.ParseFloat(line[17], 64)
		param14, _ := strconv.ParseFloat(line[18], 64)
		param15, _ := strconv.ParseFloat(line[19], 64)

		var datita Node = Node{
			float64(param1),
			float64(param2),
			float64(param3),
			float64(param4),
			float64(param6),
			float64(param7),
			float64(param8),
			float64(param9),
			float64(param10),
			float64(param11),
			float64(param12),
			float64(param13),
			float64(param14),
			float64(param15)}
		DataSetNodes = append(DataSetNodes, datita)
	}

	//Train( data, clusters, iteraciones para definir centroide)
	_, Centroids = Train(DataSetNodes, 2, 148)

	r.HandleFunc("/gokmeans/predict", PredictKmeans).Methods("GET", "OPTIONS")
	r.HandleFunc("/gokmeans/centroids/graph", GetCentroidsGraph).Methods("GET", "OPTIONS")
	r.HandleFunc("/gokmeans/centroids", GetCentroids).Methods("GET", "OPTONS")
	//Iniciar Servidor
	a.router = r
	return a
}

func (a *api) Router() http.Handler {
	return a.router
}
