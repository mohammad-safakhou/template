package main

import (
	"bizpooly/logger"
	"bizpooly/transport/rest"
)
func main() {
	l := logger.ZapLogger()
	s := l.Sugar()
	rest.StartServer(s)
}
