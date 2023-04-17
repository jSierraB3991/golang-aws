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
			t.Errorf("Error %v, got %v", value.ErrorExpect, err)
		}

		if response.TotalPrice != value.Request.TotalPrice {
			t.Errorf("Error %d, got %d", value.Request.TotalPrice, response.TotalPrice)
		}
	}

}

func TestUpdateOrder(t *testing.T) {
	repo := repository.OrderRepositoryMockOkOther{}
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
			t.Errorf("Error %v, got %v", value.ErrorExpect, err)
		}

		if response.TotalPrice > value.Request.TotalPrice {
			t.Errorf("Error %d, got %d", value.Request.TotalPrice, response.TotalPrice)
		}
	}
}
