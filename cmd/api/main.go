package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/rocksun/mvclayout/pkg/application"
	"github.com/rocksun/mvclayout/pkg/exithandler"
)

// func main() {
// 	fmt.Println("something is running-----!")
// 	if err := godotenv.Load(); err != nil {
// 		// logger.Info.Println("failed to load env vars")
// 	}

// 	app, err := application.Get()
// 	if err != nil {
// 		// logger.Error.Fatal(err.Error())
// 	}

// 	srv := server.
// 		Get().
// 		WithAddr(app.Cfg.GetAPIPort()).
// 		WithRouter(router.Get(app))
// 		// WithErrLogger(logger.Error)

// 	go func() {
// 		// logger.Info.Printf("starting server at %s", app.Cfg.GetAPIPort())
// 		if err := srv.Start(); err != nil {
// 			// logger.Error.Fatal(err.Error())
// 		}
// 	}()

// 	exithandler.Init(func() {
// 		if err := srv.Close(); err != nil {
// 			// logger.Error.Println(err.Error())
// 		}

// 		if err := app.DB.Close(); err != nil {
// 			// logger.Error.Println(err.Error())
// 		}
// 	})
// }

func main() {
	fmt.Println("Happy------")
	if err := godotenv.Load(); err != nil {
		log.Println("failed to load env vars")
	}

	app, err := application.Get()
	if err != nil {
		log.Fatal(err.Error())
	}

	exithandler.Init(func() {
		if err := app.DB.Close(); err != nil {
			log.Println(err.Error())
		}
	})
}
