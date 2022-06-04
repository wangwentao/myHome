package remote

import "fmt"

type IHelloService interface {
	Hello(name string) string
}

func NewHelloSerivce() IHelloService {

	return helloService{}
}

type helloService struct {
}

func (hs helloService) Hello(name string) string {

	return fmt.Sprintf("Hello %s", name)
}
