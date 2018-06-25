package main

import (
	"fmt"

	// mgo package
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

func GetDBData() (nItmOne, nItmTwo int) {

	// Open MongoDB session
	session, err := mgo.Dial("mongodb://localhost/mydb")
	if err != nil {
		fmt.Println("Error", err)
	}

	// Connect with the DB
	c := session.DB("TheDB").C("ThCollection")

	// Count the documents in collection
	n, _ := c.Count()
	fmt.Println("The amount of elements in collection is", n)

	// The document model for this example might be
	// doc : { ...
	//				"item" : {
	//					         "subitem_one" : integer
	//					         "subitem_two" : integer
	//								 }
	//         ...
	//			 }
	// Can implement queries as recommended by MongoDB's documentation
	// so, we can count how many documents match our criteria,
	nItmOne, _ = c.Find(bson.M{"item.subitem_one": map[string]int{"$gte": 1}}).Count()
	nItmTwo, _ = c.Find(bson.M{"item.subitem_two": 1}).Count()

	fmt.Println(nItmOne, nItmTwo)

	// Close the session
	session.Close()
	return
}

func main() {
	// Get the counts
	itemOne, itemTwo := GetDBData()
}
