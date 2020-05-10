// Using GOb to Encode Data

package main

import (
	"fmt"
	"bytes"
	"encoding/gob"
	"io"
	"log"
)

// client side user struct
type UserClient struct {
	ID string
	Name string
}

// client side transation struct
type TxClient struct {
	ID string 
  	User *UserClient 
 	AccountFrom string 
  	AccountTo string 
  	Amount float64
}

// server side user model struct
type UserServer struct {
	ID string
}

// Server side transation struct
type TxServer struct {
	ID string 
  	User UserClient 
 	AccountFrom string 
  	AccountTo string 
  	Amount *float64
}

func main() {
	// random dummy network which is a buffer from bytes package
	var net bytes.Buffer

	// dummy data for the clientside struct
	clientTx := &TxClient{ 
		ID: "123456789", 
		User: &UserClient{ 
		  ID: "ABCDEF", 
		  Name: "James", 
		}, 
		AccountFrom: "Bob", 
		AccountTo: "Jane", 
		Amount: 9.99, 
	  }

	  // encode the data with our dummy network
	  enc := gob.NewEncoder(&net)

	  // Check for errors and exit if any is found
	  if err := enc.Encode(clientTx); err != nil {
		log.Fatal("error encoding: ", err)
	  }
	  
	  // send the data to the server and check for errors
	  serverTx, err := sendToServer(&net)
	  if err != nil {
		log.Fatal("server error: ", err)
	  }

	  fmt.Printf("%#v\n", serverTx)

}

func sendToServer(net io.Reader) (*TxServer, error){
	tx := &TxServer{}

	// create a decoder with the network as source
	dec := gob.NewDecoder(net)

	// capture errors
	err := dec.Decode(tx)

	return tx, err
}