#A basic port scanner that scans only one port

package main

import(
	"fmt"
	"net"
)

func main (){
	_, := net.Dial("tcp","URL:Port")
	if err == nil{
		fmt.Println("Connection successful")
	}
}