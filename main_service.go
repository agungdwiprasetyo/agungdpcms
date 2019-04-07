package main

import (
	resumeDelivery "github.com/agungdwiprasetyo/agungdpcms/src/resume/delivery"
)

type service struct {
	handler *handler
}

type handler struct {
	*resumeDelivery.ResumeHandler
}

func NewService() *service {
	srv := new(service)
	srv.handler = &handler{resumeDelivery.New()}

	return srv
}
