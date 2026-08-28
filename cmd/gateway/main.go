package main

import (
	"log"
	"net/http"
	"traininggateway/admin"
	"traininggateway/cards"
	"traininggateway/domain"
	"traininggateway/storage"
	"traininggateway/video"
	"traininggateway/web"
	"traininggateway/workflow"
)

func main() {
	s, e := storage.Open("training.db")
	if e != nil {
		log.Fatal(e)
	}
	defer s.Close()
	cs := cards.New(s, domain.DefaultPolicy())
	cat := video.New(s)
	_ = admin.New(s)
	w := workflow.NewCheckout(cs, s)
	log.Fatal(http.ListenAndServe(":8080", web.New(cs, cat, w).Handler()))
}
