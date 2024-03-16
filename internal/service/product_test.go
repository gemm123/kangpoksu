package service

import (
	"github.com/go-playground/assert/v2"
	"github.com/google/uuid"
	"go.uber.org/mock/gomock"
	"kopoksu/internal/model"
	mock_repository "kopoksu/mock/repository"
	"testing"
)

func TestProductService_GetAllProductsFormulaMilk(t *testing.T) {
	ctrl := gomock.NewController(t)
	productRepoMock := mock_repository.NewMockProductRepository(ctrl)
	productService := NewProductService(productRepoMock)

	t.Run("should success get all products formula milk", func(t *testing.T) {
		uuidProduct := uuid.New()
		productRepoMock.EXPECT().GetAllProductsFormulaMilk().Return([]model.Product{
			{
				Id: uuidProduct,
			},
		}, nil)

		result, err := productService.GetAllProductsFormulaMilk()

		assert.Equal(t, nil, err)
		assert.Equal(t, []model.Product{
			{
				Id:                uuidProduct,
				PriceFormatted:    "Rp. 0",
				BuyPriceFormatted: "Rp. 0",
			},
		}, result)
	})
}
