// CORS (Cross-Origin Resource Sharing) adalah mekanisme keamanan yang digunakan oleh browser untuk mengontrol
// bagaimana sumber daya (resources) di server dapat diakses oleh aplikasi web dari domain lain.
package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/rizkycahyono97/online-shop-api/helpers"
)

func SetupCors() cors.Config {
	return cors.Config{
		// CORS Config
		AllowOrigins:     helpers.ParseEnvList("CORS_ALLOWED_ORIGINS"),
		AllowMethods:     helpers.ParseEnvList("CORS_ALLOWED_METHODS"),
		AllowHeaders:     helpers.ParseEnvList("CORS_ALLOWED_HEADERS"),
		AllowCredentials: helpers.GetEnvBool("CORS_ALLOW_CREDENTIALS", false),
		MaxAge:           12 * 60 * 60, // 12 hours
	}
}
