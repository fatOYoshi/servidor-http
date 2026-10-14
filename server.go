
package main

import(
"fmt"
"net"
"log"
)

func main(){

	listener, err := net.Listen("tcp",":9000")
	if err != nil{
		fmt.Println("falha ao conectar na porta 9000", err)
	}else{
		fmt.Println("[SERVIDOR] Escutando na porta :9000")
	}

	defer listener.Close()

	for{
		conn, err := listener.Accept()
		if err != nil{
			log.Fatal(err)
		}
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn){

	defer conn.Close()
	defer fmt.Println("[CONEXÃO] Encerrada", conn.RemoteAddr().String())

	fmt.Println("[CONEXÃO] Nova conexão de", conn.RemoteAddr().String())

	buffer := make([]byte ,1024)

	for{
	n, err := conn.Read(buffer)
	if err != nil{
		fmt.Println("erro na leitura ou client desconectado", err)
		break
	}

	fmt.Printf("Recebido %d byte: %s", n, string(buffer[:n]))

	_, err = conn.Write(buffer[:n])
	if err != nil{
		fmt.Println("Erro na escrita: ", err)
	}
	}

}