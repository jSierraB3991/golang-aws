package service

import (
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/jsierrab3991/order-service/pkg/dto"
	"github.com/jsierrab3991/order-service/pkg/repository"
)

var s OrderService

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func TestCreateOrderOk(t *testing.T) {
	repo := repository.OrderRepositoryMockOkFirst{}
	s = OrderService{repo: repo}

	testCases := []struct {
		Request     dto.CreateOrderRequest
		ErrorExpect error
	}{
		{
			ErrorExpect: nil,
			Request:     dto.CreateOrderRequest{UserID: uuid.NewString(), Item: "shampoo", Quantity: 1, TotalPrice: 500},
		},
	}

	for _, value := range testCases {
		response, err := s.CreateOrUpdateOrder(value.Request)

		if err != value.ErrorExpect {
			t.Errorf("Error expect error, \nexpect %v, \ngot %v", value.ErrorExpect, err)
		}

		if response.TotalPrice != value.Request.TotalPrice && response.UserID != value.Request.UserID &&
			response.List[0].Item != value.Request.Item && response.List[0].Quantity != value.Request.Quantity {
			t.Errorf("Error item difference, \nexpect %d, \ngot %d", value.Request.TotalPrice, response.TotalPrice)
		}
	}

}

func TestUpdateOrderOk(t *testing.T) {
	preValue := int64(2000)
	repo := repository.OrderRepositoryMockOkOther{PreValue: preValue}
	s = OrderService{repo: repo}

	testCases := []struct {
		Request     dto.CreateOrderRequest
		ErrorExpect error
	}{
		{
			ErrorExpect: nil,
			Request:     dto.CreateOrderRequest{UserID: uuid.NewString(), Item: "shampoo", Quantity: 1, TotalPrice: 500},
		},
	}

	for _, value := range testCases {
		response, err := s.CreateOrUpdateOrder(value.Request)

		if err != value.ErrorExpect {
			t.Errorf("Error expect error, \nexpect %v, \ngot %v", value.ErrorExpect, err)
		}

		if response.TotalPrice != value.Request.TotalPrice+preValue ||
			response.UserID != value.Request.UserID {
			t.Errorf("Error order difference, \nexpect %v, \ngot %v", value.Request, response)
		}

		ifFind := false
		for _, detail := range response.List {
			if detail.Item == value.Request.Item {
				ifFind = true
			}
		}
		if !ifFind {
			t.Errorf("Error item difrerence, \nexpect %v, \ngot %v", value.Request, response)
		}
	}
}
