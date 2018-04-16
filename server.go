package main

import (
	"log"
	"os"
	"encoding/json"
	"github.com/armon/go-socks5"
)

func main() {
	var creds socks5.StaticCredentials
	json.Unmarshal([]byte(os.Getenv("PROXY_AUTH")), &creds)

	auth := socks5.UserPassAuthenticator{Credentials: creds}

	conf := &socks5.Config{
		AuthMethods: []socks5.Authenticator{auth},
		Logger:      log.New(os.Stdout, "", log.LstdFlags),
	}

	server, err := socks5.New(conf)
	if err != nil {
		panic(err)
	}

	if err := server.ListenAndServe("tcp", "0.0.0.0:1080"); err != nil {
		panic(err)
	}
}
