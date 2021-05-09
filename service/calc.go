package service

import (
	"calc/model"
)

type CalcService struct{}

func NewCalcService() (*CalcService, error) {
	return &CalcService{}, nil
}

func (CalcService) Add(p model.CalcAddPaylod) (int, error) {
	return *p.X+*p.Y, nil
}

func (CalcService) Div(p model.CalcDivPayload) (int, error) {
	return *p.X / *p.Y, nil
}