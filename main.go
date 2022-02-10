package main

import (
	// 模組名稱/套件名稱
	"fmt"
	"math"
	"net/http"
	"practice/geo"
	"sort"
)

type myList []int

func (list myList) Len() int {
	return len(list)
}

func (list myList) Less(i, j int) bool {
	return list[i] < list[j]
}

func (list myList) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

func handler(w http.ResponseWriter, r *http.Request) {
	// ResponseWriter 是一個 Interface{}
	// 不需要實作，因為 HandlerFunc 已經負責實作了
	// 只要呼叫 Write() 方法就可以了
	w.Write([]byte("Hello World\n:))")) // Write() 需要一個 []byte 參數，不過可以將 String 轉型再放入
}

func main() {
	// practice/geo package
	c1 := &geo.Circle{
		Radius: 4.2,
	}

	geo.PrintInfo(c1)

	// math package
	// 有很多實用的函式、常數
	r := 10.0
	fmt.Println(r * r * math.Pi)

	// sort package
	// 必需按照要求實作由套件定義的 interface
	list := []int{1, 4, 8, 3, 5, 9, 6}
	newList := myList(list) // 轉型
	fmt.Println(newList)
	sort.Sort(newList)
	fmt.Println(newList)

	// net/http package
	// 有已經定義的型態可以直接使用
	// func NewServeMux() *ServeMux
	mux := http.NewServeMux() // 利用 http 包提供的函式新增一個 *ServeMux 型態

	// func (mux *ServeMux) HandleFunc(Pattern string, handler func(ResponseWriter, *Request))
	mux.HandleFunc("/", handler) // 第一個參數是字串，第二個參數是一個函式

	// 可以呼叫 http 包的 New 開頭函式新增一個變數
	// 也可用套件本身提供的 struct 型態
	serve := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	// 呼叫 http 包裡對 *Server 定義的方法 ListenAndServe()
	serve.ListenAndServe()
}
