package main

import (
	"fmt"
	"go-mvc-app/go-mvc-app/model"
	"go-mvc-app/go-mvc-app/routes"
	"go-mvc-app/go-mvc-app/service"
	"go-mvc-app/go-mvc-app/util"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

func main() {
	util.LoadConfig()
	util.ConnectDatabase()
	util.DB.AutoMigrate(&model.User{})

	router := mux.NewRouter()
	routes.RegisterRoutes(router)

	go func() {
		jobchan := make(chan int, 10)
		var wg sync.WaitGroup
		wg.Add(3)
		service.StartWorkPool(jobchan, &wg)

		for i := 0; i < 5; i++ {
			jobchan <- i
		}

		close(jobchan)
		wg.Wait()
	}()

	port := util.AppConfig.Service.Port
	fmt.Println("Server started running on port:" + port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatalf("Server exited: %v", err)
	}

}
