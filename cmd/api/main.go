package main

import (
	"fmt"
	"log"
	"net/http"
)

const PORT string = ":8080"

func main() {
	mux := http.NewServeMux()

	logMiddleware := LoggingMiddleware(mux)

	fmt.Printf("🚀 서버가 %s 에서 시작되었습니다.\n", PORT)

	if err := http.ListenAndServe(PORT, logMiddleware); err != nil {
		log.Fatal(err)
	}
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%s] %s %s", r.Method, r.URL.Path, r.RemoteAddr)
		next.ServeHTTP(w, r) // 실제 핸들러 실행
	})
}
