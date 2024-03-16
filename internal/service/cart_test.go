package service

import (
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"kopoksu/internal/model"
	mock_repository "kopoksu/mock/repository"
	"testing"
)

func TestCartService_GetAccumulationTotalCart(t *testing.T) {
	ctrl := gomock.NewController(t)
	productRepoMock := mock_repository.NewMockProductRepository(ctrl)
	cartService := NewCartService(productRepoMock)

	uuidProduct := uuid.New()
	test := []struct {
		name         string
		cart         []model.Cart
		product      model.Product
		expectResult int
		error        error
	}{
		{
			name: "should failed error get product by id",
			cart: []model.Cart{
				{
					Id: uuidProduct,
				},
			},
			product:      model.Product{},
			expectResult: 0,
			error:        errors.New("got error"),
		},
		{
			name: "should success get accumulation total cart",
			cart: []model.Cart{
				{
					Id:     uuidProduct,
					Amount: 2,
					Name:   "popok bayi",
					Total:  20000,
				},
			},
			product: model.Product{
				Id:    uuidProduct,
				Name:  "popok bayi",
				Price: 10000,
			},
			expectResult: 20000,
			error:        nil,
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			productRepoMock.EXPECT().GetProductById(uuidProduct).Return(tt.product, tt.error)

			result, err := cartService.GetAccumulationTotalCart(tt.cart)

			assert.Equal(t, tt.error, err)
			assert.Equal(t, tt.expectResult, result)
		})
	}
}
