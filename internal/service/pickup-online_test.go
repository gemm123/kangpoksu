package service

import (
	"github.com/go-playground/assert/v2"
	"github.com/google/uuid"
	"go.uber.org/mock/gomock"
	"kopoksu/internal/model"
	mock_repository "kopoksu/mock/repository"
	"testing"
)

func TestPikupOnlineOrderService_GetAllPickupOnlineOrder(t *testing.T) {
	ctrl := gomock.NewController(t)
	pickupOnlineOrderRepoMock := mock_repository.NewMockPickupOnlineOrderRepository(ctrl)
	productRepoMock := mock_repository.NewMockProductRepository(ctrl)
	pickupOnlineOrderService := NewPickupOnlineOrderService(pickupOnlineOrderRepoMock, productRepoMock)

	t.Run("should success get all pickupOnline order", func(t *testing.T) {
		uuidOfflineOrder := uuid.New()
		pickupOnlineOrderRepoMock.EXPECT().GetAllPickupOnlineOrder().Return([]model.PickupOnlineOrder{
			{
				Id: uuidOfflineOrder,
			},
		}, nil)

		result, err := pickupOnlineOrderService.GetAllPickupOnlineOrder()

		assert.Equal(t, nil, err)
		assert.Equal(t, []model.PickupOnlineOrder{
			{
				Id: uuidOfflineOrder,
			},
		}, result)
	})
}
