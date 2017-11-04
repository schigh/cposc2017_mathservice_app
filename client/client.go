package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	ms "github.com/schigh/cposc2017_mathservice"
	log "github.com/sirupsen/logrus"
	"goji.io"
	"goji.io/pat"
	"google.golang.org/grpc"
)

var client ms.MathServiceClient

func init() {
	log.SetLevel(log.DebugLevel)
}

func main() {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	conn, err := grpc.Dial("math-server:80", opts...)
	if err != nil {
		log.Errorf("gRPC connect error: %+v", err)
		os.Exit(1)
	}
	defer conn.Close()

	client = ms.NewMathServiceClient(conn)

	mux := goji.NewMux()
	mux.HandleFunc(pat.Get("/add/:addend1/:addend2"), handleAddRequest)
	mux.HandleFunc(pat.Get("/average/:comma_separated"), handleAvgerageRequest)

	http.ListenAndServe("0.0.0.0:80", mux)
}

func fail(w http.ResponseWriter, err error) {
	w.WriteHeader(400)
	fmt.Fprintf(w, "Bad request: %s", err.Error())
}

func handleAddRequest(w http.ResponseWriter, r *http.Request) {
	addend1 := pat.Param(r, "addend1")
	addend2 := pat.Param(r, "addend2")

	a1, err := strconv.ParseInt(addend1, 10, 64)
	if err != nil {
		fail(w, err)
		return
	}

	a2, err := strconv.ParseInt(addend2, 10, 64)
	if err != nil {
		fail(w, err)
		return
	}

	ar := &ms.AddRequest{
		Addend1: a1,
		Addend2: a2,
	}
	resp, err := client.Add(r.Context(), ar)
	if err != nil {
		log.Error("Add failed")
		fail(w, err)
		return
	}

	log.Debugf("Call to add %+v and %+v yielded response: %+v", a1, a2, resp.Sum)
	fmt.Fprintf(w, "%+v", resp.Sum)
}

func handleAvgerageRequest(w http.ResponseWriter, r *http.Request) {
	commaSeparated := pat.Param(r, "comma_separated")
	numberStrings := strings.Split(commaSeparated, ",")
	var ints []int64

	for _, s := range numberStrings {
		n, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			fail(w, err)
			return
		}
		ints = append(ints, n)
	}

	ar := &ms.AverageRequest{
		Numbers: ints,
	}
	resp, err := client.Average(r.Context(), ar)
	if err != nil {
		log.Error("Average failed")
		fail(w, err)
		return
	}

	log.Debugf("Call to average %+v yielded response: %+v", ints, resp.Average)
	fmt.Fprintf(w, "%+v", resp.Average)
}
