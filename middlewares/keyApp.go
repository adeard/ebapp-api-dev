package middlewares

import "net/http"

func KeyAppMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := r.Header.Get("App-Key")
		if key == "" {
			http.Error(w, "Missing App-Key header", http.StatusUnauthorized)
			return
		}
		// Logika untuk memeriksa validitas key bisa ditambahkan di sini
		next.ServeHTTP(w, r)
	})
}
