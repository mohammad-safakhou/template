package main

import (
	"template/logger"
	"template/transport/rest"
)
func main() {
	l := logger.ZapLogger()
	s := l.Sugar()
	rest.StartRestServer(s)
}
