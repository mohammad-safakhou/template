package authentication

import "backend-service/transport"

type AuthContext struct {
	*transport.ApplicationContext
}