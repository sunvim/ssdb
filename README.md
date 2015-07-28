##instruction
	this is a simple ssdb client, welcome to use.
	From the official client deriver from the client, supports the connection pool.
##  **usage**
	package main

	import (
	"github.com/sunvim/ssdb"
	"log"
	)

	func main() {
		//init thread connect pool
		//default config: Host="127.0.0.1" Port="7530" Min=5 Max=20
		err := ssdb.NewPool(ssdb.Conn{Host: "127.0.0.1", Port: "7530"})
		if err != nil {
			log.Fatal(err)
			return
		}
		defer ssdb.ClosePool()
		client := new(ssdb.Client)
		//get a connect
		client.GetConn()
		defer client.Close()
		//busi operation
		client.Set("keys", "hello")
		data, err := client.Get("keys")
		log.Println("data=>", data.(string), "|error=>", err)
	}


