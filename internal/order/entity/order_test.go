package entity_test

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"fullcycle-order/internal/order/entity"
	"testing"
)

func TestCreateNewOrderWithSuccessfully(t *testing.T) {
	order, err := entity.NewOrder(10, 7)
	assert.NoError(t, err)
	assert.NotEmpty(t, order.ID)
	assert.Equal(t, 10.0, order.Price)
	assert.Equal(t, 7.0, order.Tax)
}

func TestGivenAnEmptyPriceShouldReceiveAnError(t *testing.T) {
	order := entity.Order{ID: uuid.New().String()}
	assert.Error(t, order.IsValid(), "invalid price")
}

func TestGivenAnEmptyTaxShouldReceiveAndError(t *testing.T) {
	order := entity.Order{ID: uuid.New().String(), Price: 11}
	assert.Error(t, order.IsValid(), "invalid tax")
}

func TestGivenAllValidParamsCalculateFinalPrice(t *testing.T) {
	order, err := entity.NewOrder(10, 7)
	assert.NoError(t, err)
	assert.NotEmpty(t, order.ID)
	assert.Equal(t, 10.0, order.Price)
	assert.Equal(t, 7.0, order.Tax)
	err = order.CalculateFinalPrice()
	assert.NoError(t, err)
	assert.Equal(t, 70.0, order.FinalPrice)
}
