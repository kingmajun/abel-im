package main

import (
	"fmt"
	"net/http"
	_"abel-im/routers"
)

func main()  {
	fmt.Println("service start port: 8080")
	http.ListenAndServe(":8080",nil)
}
