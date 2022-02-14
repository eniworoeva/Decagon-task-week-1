package main

import (
	"fmt"
	"github.com/elliotchance/orderedmap"
)

func main() {
	ListOfSuperstores := orderedmap.NewOrderedMap()
	ListOfSuperstores.Set("Kada", "Benin")
	ListOfSuperstores.Set("Tivo", "Warri")
	ListOfSuperstores.Set("Spar", "Lagos")
	ListOfSuperstores.Set("Next", "Abuja")
	ListOfSuperstores.Set("Yamen", "Benin")
	ListOfSuperstores.Set("Game", "Lagos")
	ListOfSuperstores.Set("Chuck", "Lokoja")

	for el := ListOfSuperstores.Front(); el != nil; el = el.Next() {
		fmt.Println(el.Key, el.Value)
	}
}
