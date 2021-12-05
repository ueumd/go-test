package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	var provinceList []interface{}
	var districtList []interface{}
	var province map[string]interface{}
	var m map[string]interface{}

	m = make(map[string]interface{})
	m["name"] = "安庆"
	m["items"] = [...]string{"大观区","怀宁县"}
	districtList = append(districtList, m)

	m = make(map[string]interface{})
	m["name"] = "蚌埠"
	m["items"] = [...]string{"龙子湖区","固镇县"}
	districtList = append(districtList, m)

	province = make(map[string]interface{})
	province["name"] = "安徽"
	province["items"] = districtList

	provinceList = append(provinceList, province)

	jsonStr, err := json.Marshal(provinceList)
	if err!=nil{
		log.Fatal(err)
		return
	}

	fmt.Printf("%s\n",jsonStr)
}