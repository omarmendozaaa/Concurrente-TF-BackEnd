package main

import (
	"net"
	"fmt"
	"os"
	"bufio"
)

var dat [][]int

var centroids []int

func handleConnections(conn net.Conn, d [][]int, c []int){
	defer conn.Close()

	//Recuperar los datos que se envían
	r := bufio.NewReader(conn)

	for{
		datos, _ := r.ReadString('\n')
		fmt.Println(datos)

		//Responder al cliente
		fmt.Fprintln(conn, d)
		fmt.Fprintln(conn, c)
	}
}

func main(){

	dat = append(dat, []int{1, 2})
	dat = append(dat, []int{3, 4})
	dat = append(dat, []int{5, 6})

	dat = append(dat, []int{7, 8})
	dat = append(dat, []int{9, 10})
	dat = append(dat, []int{11, 12})

	dat = append(dat, []int{13, 14})
	dat = append(dat, []int{15, 16})
	dat = append(dat, []int{17, 18})

	dat = append(dat, []int{19, 20})
	dat = append(dat, []int{21, 22})
	dat = append(dat, []int{23, 24})

	dat = append(dat, []int{25, 26})
	dat = append(dat, []int{27, 28})
	dat = append(dat, []int{29, 30})

	centroids = append(centroids, 0)
	centroids = append(centroids, 0)

	//Prueba
	fmt.Println(dat)

	//que escucha
	//aquí debemos agregar la parte del ip
	ln, err := net.Listen("tcp", "localhost:8000")

	if err != nil {
		fmt.Println("Falla en la comunicación", err.Error())
		os.Exit(1)
	}
	defer ln.Close()
	
	for{
		//Acepta los mensajes que llega de esa comunicación
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Falla en la conexión", err.Error())
		}

		//Conexión concurrente a varios clientes
		go handleConnections(conn, dat, centroids)
	}
}