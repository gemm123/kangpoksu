package service

import (
	"github.com/go-playground/assert/v2"
	"github.com/google/uuid"
	"go.uber.org/mock/gomock"
	"kopoksu/internal/model"
	mock_repository "kopoksu/mock/repository"
	"testing"
)

func TestOfflineOrderService_GetAllOfflineOrder(t *testing.T) {
	ctrl := gomock.NewController(t)
	offlineOrderRepoMock := mock_repository.NewMockOfflineOrderRepository(ctrl)
	productRepoMock := mock_repository.NewMockProductRepository(ctrl)
	offlineOrderService := NewPickupOnlineOrderService(offlineOrderRepoMock, productRepoMock)

	t.Run("should success get all offline order", func(t *testing.T) {
		uuidOfflineOrder := uuid.New()
		offlineOrderRepoMock.EXPECT().GetAllOfflineOrder().Return([]model.PickupOnlineOrder{
			{
				Id: uuidOfflineOrder,
			},
		}, nil)

		result, err := offlineOrderService.GetAllPickupOnlineOrder()

		assert.Equal(t, nil, err)
		assert.Equal(t, []model.PickupOnlineOrder{
			{
				Id: uuidOfflineOrder,
			},
		}, result)
	})
}
