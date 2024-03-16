package service

import (
	"github.com/go-playground/assert/v2"
	"github.com/google/uuid"
	"go.uber.org/mock/gomock"
	"kopoksu/internal/model"
	mock_repository "kopoksu/mock/repository"
	"testing"
)

func TestOnlineOrderService_GetAllOnlineOrder(t *testing.T) {
	ctrl := gomock.NewController(t)
	onlineOrderRepoMock := mock_repository.NewMockOnlineOrderRepository(ctrl)
	productRepoMock := mock_repository.NewMockProductRepository(ctrl)
	onlineOrderService := NewOnlineOrderService(onlineOrderRepoMock, productRepoMock)

	t.Run("should success get all online order", func(t *testing.T) {
		uuidOnlineOrder := uuid.New()
		onlineOrderRepoMock.EXPECT().GetAllOnlineOrder().Return([]model.OnlineOrder{
			{
				Id: uuidOnlineOrder,
			},
		}, nil)

		result, err := onlineOrderService.GetAllOnlineOrder()

		assert.Equal(t, nil, err)
		assert.Equal(t, []model.OnlineOrder{
			{
				Id: uuidOnlineOrder,
			},
		}, result)
	})
}
