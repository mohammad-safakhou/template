package main
//go:generate sqlboiler --wipe psql
import (
	"backend-service/logger"
	"backend-service/transport/rest"
)

func main() {
	l := logger.ZapLogger()
	s := l.Sugar()
	rest.StartRestServer(s)
}
