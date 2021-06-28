package main

import (
	"net"
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

//var datos[][]int

func main(){
	//Con qué nodo se desea establecer la comunicación
	conn, _ := net.Dial("tcp", "localhost:8000") 
	//Establece la conexión hacia el nodo servidor a traves del puerto 8000

	defer conn.Close()
	
	//Este reader captura por consolo el mensaje que se va a enviar
	rin := bufio.NewReader(os.Stdin)
	//Este reader establece conexión para enviar el mensaje
	r := bufio.NewReader(conn)
	for{
		fmt.Print("Ingrese un mensaje: ")
		msg, _ := rin.ReadString('\n')
		fmt.Fprint(conn, msg)//Envío

		resp, _ := r.ReadString('\n')
		fmt.Printf("Respuesta del server: %s", resp)

		procesamientoString(resp)
		distanciaEjemplo(dat[0][0], dat[0][1])
		
		resp2, _ := r.ReadString('\n')
		fmt.Printf("Respuesta 2 del server: %s", resp2)
	}
}

var dat [][]int
var centroids []int

func procesamientoString(s string){
	t := strings.ReplaceAll(s, "[", "")
	t = strings.ReplaceAll(t, "]", "")

	arrString := strings.Split(t, " ")

	for i := 0; i < len(arrString); i=i+2{
		fmt.Println(arrString[i], arrString[i+1])
		num1, _ := strconv.Atoi(arrString[i])
		num2, _ := strconv.Atoi(arrString[i+1])
		dat = append(dat, []int{num1, num2})
	}
}

func distanciaEjemplo(num1, num2 int){
	centroids = append(centroids, num1+1)
	centroids = append(centroids, num2+1)
}