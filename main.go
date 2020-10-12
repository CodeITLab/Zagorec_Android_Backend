package main

import (
	"fmt"
	"net/http"
)

func main()  {
	fmt.Println("Dobro mi došel prijatelj, vu staru zagorsku klet.")
	http.ListenAndServe(":8891", nil)
}