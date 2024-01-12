package main

//menambah middleware agar setiap hanya user tertensu saja yang diijinkan mengakses server

import "net/http"

func (app *application) enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Access-Control-Allow-Origin", "*") //ubah "*" dengan domain server anda . saya mencontohkan agar semuanya bisa mengakses
		next.ServeHTTP(rw, r)
	})
}

//end middleware
