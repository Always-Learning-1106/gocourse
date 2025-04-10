package basics

import (
	"fmt"
	"maps"
)

func main() {
	// var mapVariable map[keyType]valueType
	// mapVariable = make(map[key]valueType)
	// using a Map literal
	// mapVariable = map[keyType]valueType{
	// 	key1:value1,
	// 	key2:value2,
	// }
	myMap := make(map[string]int)
	fmt.Println(myMap)
	myMap["key1"] = 9
	fmt.Println(myMap)
	myMap["code"] = 18
	fmt.Println(myMap)
	fmt.Println(myMap["key1"])
	fmt.Println(myMap["key"])
	myMap["code"] = 21
	fmt.Println(myMap)
	delete(myMap, "key1")
	fmt.Println(myMap)
	myMap["key2"] = 10
	myMap["key3"] = 11
	myMap["key4"] = 12
	fmt.Println(myMap)
	// clear(myMap)
	// fmt.Println(myMap)
	value, unknownValue := myMap["key12"]
	fmt.Println(value)
	fmt.Println(unknownValue)
	myMap2 := map[string]int{"a": 1, "b": 2}
	fmt.Println(myMap2)
	myMap3 := map[string]int{"a": 1, "b": 2}
	if maps.Equal(myMap, myMap2) {
		fmt.Println("myMap and myMap2 are equal")
	} else {
		fmt.Println("maps are not equal")
	}
	if maps.Equal(myMap2, myMap3) {
		fmt.Println("myMap2 and myMap3 are equal")
	} else {
		fmt.Println("maps are not equal")
	}

	for k, v := range myMap3 {
		fmt.Println("Key:", k, "Value", v)
	}
	myMap4 := map[string]int{"c": 3, "d": 4}
	if _, ok := myMap4["c"]; ok {
		fmt.Println("Key c exists in the map")
	}
	var myMap5 map[string]string
	val := myMap5["key"]
	fmt.Println(val)
	myMap5["key"] = "10"
	fmt.Println(myMap5)
}
