package main

import (
	"fmt"
	"net/http"

	"gopkg.in/loremipsum.v1"

	"go.uber.org/zap"
)

var (
	logger              *zap.Logger
	loremIpsumGenerator *loremipsum.LoremIpsum
)

func main() {
	var err error
	logger, err = zap.NewProduction()
	if err != nil {
		panic(err)
	}

	loremIpsumGenerator = loremipsum.New()

	http.HandleFunc("/", hello)
	if err := http.ListenAndServe(":80", nil); err != nil {
		panic(err)
	}
}

func hello(w http.ResponseWriter, req *http.Request) {
	logger.Info("new request", zap.String("RemoteAddr", req.RemoteAddr))
	w.Header().Add("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, loremIpsumGenerator.Paragraph())
}
