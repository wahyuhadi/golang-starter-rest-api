/**
	Rahmat wahyu hadi
**/

package middlewares

import "github.com/rs/cors"

// contoh middleware sederhana
func CORSMiddleware() *cors.Cors {
	return cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
	})
}
