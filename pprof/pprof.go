package pprof

import (
	"log"
	"net/http"
	_ "net/http/pprof"
)

// 可视化 http://42.192.155.29:9990/debug/pprof/
//go tool pprof -http=:8080 "http://42.192.155.29:9990/debug/pprof/heap"
//go tool pprof -http=:8080 "http://42.192.155.29:9990/debug/pprof/heap"

func InitPprofMonitor() {
	go func() {
		log.Println(http.ListenAndServe(":9990", nil))
	}()
}
