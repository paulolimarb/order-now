package repository

import (
	"log"
	"order-now/order-consumer/config"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

func CreateOrder(document []byte) (s bool) {

	session, err := mgo.Dial(config.Get().MongoUrl)
	if err != nil {
		log.Fatal(err)
		return false

	}
	defer session.Close()
	var bdoc interface{}
	err = bson.UnmarshalJSON(document, &bdoc)
	if err != nil {
		log.Fatal(err)
		return false
	}
	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("Order").C("orders")
	err = c.Insert(&bdoc)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}
