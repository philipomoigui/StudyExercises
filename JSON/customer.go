package main

import (
	"fmt"
	"os"
	"encoding/json"
)

type address struct {
	Street string `json:"street"`
	City string `json:"city"`
	State string `json:"state"`
	Zipcode int `json:"zipcode"`
}

type item struct {
	Name string `json:"itemname"`
	Description string `json:"desc,omitempty"`
	Quantity int `json:"qty"`
	Price int `json:"price"`
}

type order struct {
	TotalPrice int `json:"total"`
	IsPaid bool `json:"paid"`
	Fragile bool `json:",omitempty"`
	OrderDetail []item `json:"Orderdetail"`
}

type customer struct {
	UserName string `json:"username"`
	Password string `json:"-"`
	Token string `json:"-"`
	ShipToAddress address `json:"shipto"`
	PurchaseOrder order `json:"order"`
}

func main() {
	jsonData := []byte(` 
  { 
    "username" :"blackhat", 
    "shipto": 
    { 
      "street": "Sulphur Springs Rd", 
      "city": "Park City", 
      "state": "VA", 
      "zipcode": 12345 
    }, 
    "order": 
    { 
      "paid":false, 
      "orderdetail" : 
      [{ 
        "itemname":"A Guide to the World of zeros and ones", 
        "desc": "book", 
        "qty": 3, 
        "price": 50 
      }] 
    } 
  } 
  `) 

  if !json.Valid(jsonData){
	  fmt.Printf("Not a valid JSON data %s", jsonData)
	  os.Exit(1)
  }

  var c customer

  err := json.Unmarshal(jsonData, &c)
  if err != nil {
	  fmt.Println(err)
	  os.Exit(1)
  }

	game := item{}
	game.Name = "The Rise of the Last Jedai"
	game.Description = "A PS4 Game"
	game.Price = 50
	game.Quantity = 8

	glass := item{} 
	glass.Name = "Crystal Drinking Glass" 
	glass.Quantity = 11 
	glass.Price = 25

	  // add the newly created items to orderdetails
	c.PurchaseOrder.OrderDetail = append(c.PurchaseOrder.OrderDetail, game)
	c.PurchaseOrder.OrderDetail = append(c.PurchaseOrder.OrderDetail, glass)

	c.Total()

	c.PurchaseOrder.IsPaid = true
	c.PurchaseOrder.Fragile = true
	 
  customerOrder, err := json.MarshalIndent(c, "", " ")
  if err != nil {
	  fmt.Println(err)
		os.Exit(1)
  }

  fmt.Println(string(customerOrder))
}

// a method to calculate the total price
func (c *customer) Total() { 
	price := 0 
	for _, item := range c.PurchaseOrder.OrderDetail { 
	  price += item.Quantity * item.Price 
	} 
   c.PurchaseOrder.TotalPrice = price 
  } 

