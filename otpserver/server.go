package otpserver

import (
	"fmt"
	"github.com/skyjia/go-otpserver/Godeps/_workspace/src/github.com/ant0ine/go-json-rest/rest"
	"log"
	"net/http"
)

// ServerConfig defines a server configuration
type ServerConfig struct {
	Host string
	Port uint
}

var (
	// DefaultConfig is the default server config.
	DefaultConfig = ServerConfig{
		Host: "localhost",
		Port: 8080,
	}
)

// StartServer starts a HTTP server
func StartServer(cfg ServerConfig) {
	api := rest.NewApi()
	api.Use(rest.DefaultProdStack...)

	router, err := rest.MakeRouter(
		rest.Post("/totp/authurl", generateTOTPAuthURLHandler),
		rest.Post("/totp/verify", verifyTOTPHandler),
	)

	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)

	listenAddress := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	log.Println("Server is listening at: ", listenAddress)
	log.Fatal(http.ListenAndServe(listenAddress, api.MakeHandler()))
}
