package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"rsc.io/quote"
)

var (
	router = gin.Default()
)

func main() {
	fmt.Println(quote.Go())
	log.Fatal(router.Run(":8080"))
}
