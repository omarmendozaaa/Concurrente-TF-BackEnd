package server

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
)

type Node []float64

type Data struct {
	DEPARTAMENTO            float64 `json:"departamento"`
	VICTIMA_EDAD            float64 `json:"victima_edad"`
	VICTIMA_NRO_HIJOS       float64 `json:"victima_nro_hijos"`
	AGRESOR_EDAD            float64 `json:"agresor_edad"`
	ALCOHOL_DROGAS          float64 `json:"alcohol_drogas"`
	ACUCHILLAMIENTO         float64 `json:"acuchillamiento"`
	GOLPES_DIVERSOS         float64 `json:"golpes_diversos"`
	DISPARO_BALA            float64 `json:"disparo_bala"`
	ENVENENAMIENTO          float64 `json:"envenenamiento"`
	DESBARRANCAMIENTO       float64 `json:"desbarrancamiento"`
	ASFIXIA_ESTRAGULAMIENTO float64 `json:"asfixia_extrangulamiento"`
	ATROPELLAMIENTO         float64 `json:"atropellamiento"`
	QUEMADURA               float64 `json:"quemadura"`
	OTRO                    float64 `json:"otro"`
}
type Cluster struct {
	Index int `json:"index"`
}
type Cent struct {
	Valor  float64 `json:"Claster 1"`
	Valor2 float64 `json:"Claster 2"`
}

func GetCentroidsGraph(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var datita []Cent
	var datitaoficial []Cent
	var temp Cent
	iterable := 0
	for _, value := range Centroids[0] {
		temp.Valor = value
		datita = append(datita, temp)
	}

	for _, value2 := range datita {
		temp = value2
		temp.Valor2 = Centroids[1][iterable]
		iterable += 1
		datitaoficial = append(datitaoficial, temp)
	}

	json.NewEncoder(w).Encode(datitaoficial)
}
func GetCentroids(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Centroids)
}

//Realiza la predicción
func PredictKmeans(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newdata Data
	_ = json.NewDecoder(r.Body).Decode(&newdata)
	var datita Node = Node{
		newdata.DEPARTAMENTO,
		newdata.VICTIMA_EDAD,
		newdata.VICTIMA_NRO_HIJOS,
		newdata.AGRESOR_EDAD,
		newdata.ALCOHOL_DROGAS,
		newdata.ACUCHILLAMIENTO,
		newdata.GOLPES_DIVERSOS,
		newdata.DISPARO_BALA,
		newdata.ENVENENAMIENTO,
		newdata.DESBARRANCAMIENTO,
		newdata.ASFIXIA_ESTRAGULAMIENTO,
		newdata.ATROPELLAMIENTO,
		newdata.QUEMADURA,
		newdata.OTRO}
	cluster := Nearest(datita, Centroids)
	var ResponseCluster Cluster
	ResponseCluster.Index = cluster
	json.NewEncoder(w).Encode(ResponseCluster)
}
func wait(c chan int, values Node) {
	count := len(values)

	<-c
	for respCnt := 1; respCnt < count; respCnt++ {
		<-c
	}
}

func Train(Nodes []Node, clusterCount int, maxRounds int) (bool, []Node) {
	if int(len(Nodes)) < clusterCount {
		return false, nil
	}

	stdLen := 0
	for i, Node := range Nodes {
		curLen := len(Node)

		if i > 0 && len(Node) != stdLen {
			return false, nil
		}

		stdLen = curLen
	}

	centroids := make([]Node, clusterCount)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < clusterCount; i++ {
		srcIndex := r.Intn(len(Nodes))
		srcLen := len(Nodes[srcIndex])
		centroids[i] = make(Node, srcLen)
		copy(centroids[i], Nodes[r.Intn(len(Nodes))])
	}

	return Train2(Nodes, clusterCount, maxRounds, centroids)
}

func Train2(Nodes []Node, clusterCount int, maxRounds int, centroids []Node) (bool, []Node) {
	movement := true
	for i := 0; i < maxRounds && movement; i++ {
		movement = false

		groups := make(map[int][]Node)

		for _, Node := range Nodes {
			near := Nearest(Node, centroids)
			groups[near] = append(groups[near], Node)
		}

		for key, group := range groups {
			newNode := meanNode(group)

			if !equal(centroids[key], newNode) {
				centroids[key] = newNode
				movement = true
			}
		}
	}

	return true, centroids
}

func equal(node1, node2 Node) bool {
	if len(node1) != len(node2) {
		return false
	}

	for i, v := range node1 {
		if v != node2[i] {
			return false
		}
	}

	return true
}

func Nearest(in Node, nodes []Node) int {
	count := len(nodes)

	results := make(Node, count)
	cnt := make(chan int)
	for i, node := range nodes {
		go func(i int, node, cl Node) {
			results[i] = distance(in, node)
			cnt <- 1
		}(i, node, in)
	}

	wait(cnt, results)

	mindex := 0
	curdist := results[0]

	for i, dist := range results {
		if dist < curdist {
			curdist = dist
			mindex = i
		}
	}

	return mindex
}

func distance(node1 Node, node2 Node) float64 {
	length := len(node1)
	squares := make(Node, length)

	cnt := make(chan int)

	for i := range node1 {
		go func(i int) {
			diff := node1[i] - node2[i]
			squares[i] = diff * diff
			cnt <- 1
		}(i)
	}

	wait(cnt, squares)

	sum := 0.0
	for _, val := range squares {
		sum += val
	}

	return sum
}

func meanNode(values []Node) Node {
	newNode := make(Node, len(values[0]))

	for _, value := range values {
		for j := 0; j < len(newNode); j++ {
			newNode[j] += value[j]
		}
	}

	for i, value := range newNode {
		newNode[i] = value / float64(len(values))
	}

	return newNode
}
