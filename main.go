package main

import (
	"fmt"
	"os"

	"github.com/sakiib/go-pattern/db"
	"github.com/sakiib/go-pattern/user"
)

func main() {
	fmt.Println("Accept interfaces, return structs")

	mp := make(map[string]string)
	store, err := db.NewStore(mp)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	service := user.NewUserService(store)

	if err := service.CreateUser("sakib", "appscode"); err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	if err := service.CreateUser("aminul", "google"); err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	u, err := service.RetrieveUser("sakib")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	fmt.Println("user --> ", u)

	u, err = service.RetrieveUser("aminul")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	fmt.Println("user --> ", u)

	u, err = service.RetrieveUser("akash")
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
		os.Exit(0)
	}
	fmt.Println("user --> ", u)
}
