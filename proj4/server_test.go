package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"testing"
	"time"
	"github.com/awoodbeck/gnp/ch09/handlers"
)


func TestSimpleHTTPServer(t *testing.T){
	srv := &http.Server{
		Addr: "127.0.0.1:8081",
		Handler: http.TimeoutHandler(handlers.DefaultHandler(), 2*time.Minute, ""),
		ReadHeaderTimeout: time.Minute,
	}
	l, err := net.Listen("tcp", srv.Addr)
	if err != nil {
		t.Fatal(err)
	}
	go func ()  {
		err := srv.Serve(l)
		if err != nil {
			t.Error(err)
		}
	}()

}
