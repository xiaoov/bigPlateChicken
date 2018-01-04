package main

import (
	"io/ioutil"
	"encoding/json"
	"fmt"
)

const MENU_FILE_NAME = "menu"
const AVERAGE = 13.0

var cookLists = make(map[string]float64)

func main() {
	var num = 0.0
	var total = 0.0
	InitBookMenu()

	fmt.Print("please input number of person: ")
	fmt.Scanf("%f", &num)
	fmt.Print("\n")

	total = num * AVERAGE

	fmt.Println("total is :", total)

	generateCookMenu(total)
}

func InitBookMenu() {
	file, err := ioutil.ReadFile(MENU_FILE_NAME)
	if err != nil {
		fmt.Println("can not open menu file")
		return
	}

	err = json.Unmarshal(file, &cookLists)
	if err != nil {
		fmt.Println("can not unmarshal menu file")
		return
	}
}

func generateCookMenu(total float64) {
	var bookList = make(map[string]float64)
	var bookTotal = 0.0

	for k, v := range cookLists {
		if bookTotal+v >= total +5 {
			continue
		}

		bookList[k] = v
		bookTotal = bookTotal + v
	}

	fmt.Println(bookList)
	fmt.Println(bookTotal)
}


