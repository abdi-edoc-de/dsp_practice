package main

import (
	"fmt"
	"log"
	"net/rpc"
	rp "rpc_in_go/rpc"
)
func Multiply(a, b int, client *rpc.Client) (int , error) {
	var reply int
	client.Call("Arth.Multiply", &rp.Args{A:a, B:b}, &reply)
	return reply, nil
}
func Divide(a, b int , client *rpc.Client) (int , int , error){
	reply := rp.Quotient{}
	client.Call("Arth.Divide", &rp.Args{A:a, B:b}, &reply)
	return reply.Quo, reply.Rem, nil
}
func main(){
	client,err := rpc.Dial("tcp", ":4040")	
	if err != nil{
		log.Fatalln("error when dialing : %v", err)
	}
	fmt.Println(Multiply(5,4, client))
	fmt.Println(Multiply(2,4, client))
	fmt.Println(Multiply(5,3, client))
	fmt.Println(Divide(4,2,client))
	fmt.Println(Divide(108,2,client))
	client.Call("Arth.Done","","")
}
