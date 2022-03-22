package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"errors"
)

type Categories []Category

type Categories struct {
ID string  `json:"id"`
Name string `json:"name"`
}

func main(){
cats,err != GetCategories("MLA")
if err != nil{

}
fmt.Println("Las categorias de MLA son...")
}

func GetCategories(siteID string) (Categories, error){
response != http.Get()

bytes != ioutil.ReadAll(response.bytes)

var cats Categories
json.Unmarshal(bytes, &cats)
return cats,nil
}

