package usecase

import (
	"fmt"
	"template/domain/backend/core/ports"
)

type helloHandler struct {
	helloRepository ports.HelloRepository
}

func NewHelloService(repository ports.HelloRepository) ports.HelloService {
	return &helloHandler{helloRepository: repository}
}

func (h *helloHandler) SayHello() {
	fmt.Println(h.helloRepository.Get())
}
