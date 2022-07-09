package main

import (
	"log"
	"net"
	// "net/http"
	"net/rpc"
	rp "rpc_in_go/rpc"
)
type Arth struct {
	IsDone bool 
}
func (a *Arth) Multiply(args *rp.Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (a *Arth) Divide(args *rp.Args, quo *rp.Quotient) error{
	if args.B== 0{
		log.Fatalln("number can not be divided by zero")
	}
	quo.Quo = args.A/args.B
	quo.Rem = args.A % args.B
	return nil
}
func (a *Arth) Done(args string, reply *string) error{
	a.IsDone = false
	return nil
}
func main(){
	arth := new(Arth)
	arth.IsDone = true
	rpcs := rpc.NewServer()
	err := rpcs.Register(arth)

	if err != nil{
		log.Fatalln("error registering arth: %v", err)
	}
	// rpcs.HandleHTTP()
	list , err := net.Listen("tcp", ":4040")

	if err != nil{
		log.Fatalln("Listener error: %v", err)
	}
	 func () {
		for arth.IsDone{
			if arth.IsDone == false {
				break
			}
			println(arth.IsDone)
			con , err := list.Accept()
			if err == nil{
				go rpcs.ServeConn(con)
			}else{
				break
			}
		}
		list.Close()
	}()
	println(arth.IsDone)
}