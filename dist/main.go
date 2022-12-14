package main

import (
	"STACK-GO/dataresource"
	routes "STACK-GO/modules"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// .Env
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	// Database
	DSN := os.Getenv("ADM_USER") + ":" + os.Getenv("ADM_PASS") + "@tcp(" + os.Getenv("IP_WHITELIST") + ":" + os.Getenv("PORT") + ")/" + os.Getenv("DATABASE") + "?parseTime=true"
	fmt.Println(DSN)
	database, err := dataresource.New(os.Getenv("DIALECT"), DSN, 10, 10)

	if err != nil {
		fmt.Printf("%s", err)
	}

	// Modules
	r := gin.New()
	routes.InitializeRoutes(r, database)

	// Running
	if err := r.Run("localhost:8019"); err != nil {
		panic(err.Error())
	}

}
