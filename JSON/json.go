package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	jsonData := []byte(`
		{
			"id": 2,
			"lname": "wahsington",
			"fname": "Bill",  
    		"IsEnrolled": true,  
			"grades":[100,76,93,50],  
			"class":  
    	{  
      		"coursename": "World Lit",  
      		"coursenum": 101,  
      		"coursehours": 3  
    	}  
	}
	`)

	if !json.Valid(jsonData) {
		fmt.Printf("JSON is not valid: %s", jsonData)
		os.Exit(1)
	  }

	  var i interface{} 

	err := json.Unmarshal(jsonData, &i)
  	if err != nil {
    fmt.Println(err)
    os.Exit(1)
 	}

	  data := i.(map[string]interface{})
	  for k, v := range data {
		  switch value := v.(type){
		  case string:
			fmt.Println("(String): ", k, value)
		  case float64: 
			fmt.Println("(float64):", k, value) 
		  case bool: 
			fmt.Println("(bool):", k, value) 
		  case []interface{}: 
		  fmt.Println("(slice): ", k)
		  for k, v := range value {
			fmt.Println(k, v)
		  }
		default:
			fmt.Println("(Unknown type): ", k, value)
	  }
	}
}