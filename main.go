package main

import (
	"github.com/moomman/tour2-blog/internal/routing/router"
	"github.com/moomman/tour2-blog/internal/setting"
	"net/http"
	"time"
)

func initSetting() {
	setting.AllInit()
}

func main() {
	initSetting()

	router := router.NewRouter()
	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
