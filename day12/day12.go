package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"reflect"
)

func main() {

	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var obj interface{}
	err = json.Unmarshal(b, &obj)
	if err != nil {
		panic(err)
	}

	fmt.Println("Solution 1:", part1(obj, 0.0))
	fmt.Println("Solution 2:", part2(obj, 0.0))
}

func part1(obj interface{}, total float64) float64 {
	switch v := obj.(type) {
	case float64:
		total += v
	case []interface{}:
		for _, v2 := range obj.([]interface{}) {
			total = part1(v2, total)
		}
	case map[string]interface{}:
		for _, v2 := range obj.(map[string]interface{}) {
			total = part1(v2, total)
		}
	}
	return total
}

func part2(obj interface{}, total float64) float64 {
	switch v := obj.(type) {
	case float64:
		total += v
	case []interface{}:
		for _, v2 := range obj.([]interface{}) {
			total = part2(v2, total)
		}
	case map[string]interface{}:
		for _, v2 := range obj.(map[string]interface{}) {
			if reflect.TypeOf(v2).Kind() == reflect.String {
				if reflect.ValueOf(v2).String() == "red" {
					return total
				}
			}
		}
		for _, v2 := range obj.(map[string]interface{}) {
			total = part2(v2, total)
		}
	}
	return total
}
