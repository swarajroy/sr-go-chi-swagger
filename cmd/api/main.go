package main

import (
	"log"

	"github.com/swarajroy/sr-go-chi-swagger/internal/env"
	"go.uber.org/zap"
)

var (
	version = "0.0.1"
)

//	@title			Sample API
//	@description	API for Samaple.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath	    /v1
//
// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
// @description
func main() {
	cfg := config{
		addr:   env.GetString("ADDR", ":8080"),
		apiURL: env.GetString("API_URL", "http://localhost:8080"),
		env:    env.GetString("ENV", "developement"),
	}

	//Logger
	logger := zap.Must(zap.NewProduction()).Sugar()
	defer logger.Sync()

	app := application{
		config: cfg,
		logger: logger,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))

}
