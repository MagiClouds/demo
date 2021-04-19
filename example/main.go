package main

import (
	"fmt"
	"log"
	"time"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/store"
)

func main() {

	srv := service.New(service.Name("example"))
	srv.Init()

	records, err := store.Read("mykey")
	if err != nil {
		log.Fatal(err)
	}

	if len(records) == 0 {
		fmt.Println("no records")

	}

	for _, v := range records {

		fmt.Println(v.Key, string(v.Value))

	}

	time.Sleep(time.Second * 5)

}
