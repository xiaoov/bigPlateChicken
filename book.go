package main

import (
	"io/ioutil"
	"encoding/json"
	"fmt"
	"math/rand"
	"math"
	"sort"
	"time"
)
/*
    带权随机
    key(i) = r[i] ** 1/w[i]
    要随机选取一个元素，就去key最大的那个
*/
const MENU_FILE_NAME = "menu"
const AVERAGE = 13.0
type Food struct{
	FoodInfo struct{
		Name string   `json:"name"`
		Price float64 `json:"price"`
		Weight float64 `json:"weight"`
		IsMeat float64 `json:"isMeat"`
	}`json:"food"`
}
type FoodWeight struct{
	food Food
	key  float64
}
type FoodWeightSlice []FoodWeight
func (a FoodWeightSlice) Len() int {    // 重写 Len() 方法
	return len(a)
}
func (a FoodWeightSlice) Swap(i, j int){     // 重写 Swap() 方法
	a[i], a[j] = a[j], a[i]
}
func (a FoodWeightSlice) Less(i, j int) bool {    // 重写 Less() 方法， 从大到小排序
	return a[j].key < a[i].key
}
var (
	cookLists = []Food{}
	foodWeights = []FoodWeight{}
)

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
	//fmt.Println(cookLists)

	for i,v := range cookLists{
		s2 := rand.NewSource(time.Now().UTC().UnixNano() + int64(i))  //刷新随机种子
		r2 := rand.New(s2)
		k := math.Pow(r2.Float64(),(1 / v.FoodInfo.Weight))
		foodWeight := FoodWeight{food:v,key:k}
		foodWeights = append(foodWeights,foodWeight)
		//cookMap[v] = key
	}
	fmt.Println(foodWeights)
	sort.Sort(FoodWeightSlice(foodWeights))
	fmt.Println(foodWeights)
}

func generateCookMenu(total float64) {
	var bookList = make(map[string]float64)
	var bookTotal = 0.0

	for _, v := range foodWeights {
		if bookTotal+v.food.FoodInfo.Price >= total +5 {
			continue
		}

		bookList[v.food.FoodInfo.Name] = v.food.FoodInfo.Price
		bookTotal = bookTotal + v.food.FoodInfo.Price
	}

	fmt.Println(bookList)
	fmt.Println(bookTotal)
}
