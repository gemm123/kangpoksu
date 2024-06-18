package service

import (
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	mock_repository "kopoksu/mock/repository"
	"testing"
)

func setupRecapService(t *testing.T) (*gomock.Controller, *mock_repository.MockPickupOnlineOrderRepository, *mock_repository.MockOnlineOrderRepository, RecapService) {
	ctrl := gomock.NewController(t)
	pickupOnlineOrderRepoMock := mock_repository.NewMockPickupOnlineOrderRepository(ctrl)
	onlineOrderRepoMock := mock_repository.NewMockOnlineOrderRepository(ctrl)
	recapService := NewRecapService(pickupOnlineOrderRepoMock, onlineOrderRepoMock)
	return ctrl, pickupOnlineOrderRepoMock, onlineOrderRepoMock, recapService
}

func TestRecapService_ProfitRecapFormulaMilkPickupOnlineOrder(t *testing.T) {
	ctrl, pickupOnlineOrderRepoMock, _, recapService := setupRecapService(t)
	defer ctrl.Finish()

	t.Run("should success profit recap formula milk pickupOnline order", func(t *testing.T) {
		pickupOnlineOrderRepoMock.EXPECT().RecapProfitFormulaMilkPickupOnlineOrder().Return(0, nil)

		result, err := recapService.ProfitRecapFormulaMilkPickupOnlineOrder()
		assert.NoError(t, err)
		assert.Equal(t, 0, result)
	})
}

func TestRecapService_ProfitRecapBabyDiaperPickupOnlineOrder(t *testing.T) {
	ctrl, pickupOnlineOrderRepoMock, _, recapService := setupRecapService(t)
	defer ctrl.Finish()

	t.Run("should success profit recap baby diaper pickupOnline order", func(t *testing.T) {
		pickupOnlineOrderRepoMock.EXPECT().RecapProfitBabyDiaperPickupOnlineOrder().Return(0, nil)

		result, err := recapService.ProfitRecapBabyDiaperPickupOnlineOrder()
		assert.NoError(t, err)
		assert.Equal(t, 0, result)
	})
}

func TestRecapService_ProfitRecapAdultDiaperPickupOnlineOrder(t *testing.T) {
	ctrl, pickupOnlineOrderRepoMock, _, recapService := setupRecapService(t)
	defer ctrl.Finish()

	t.Run("should success profit recap adult diaper pickupOnline order", func(t *testing.T) {
		pickupOnlineOrderRepoMock.EXPECT().RecapProfitAdultDiaperPickupOnlineOrder().Return(0, nil)

		result, err := recapService.ProfitRecapAdultDiaperPickupOnlineOrder()
		assert.NoError(t, err)
		assert.Equal(t, 0, result)
	})
}

func TestRecapService_ProfitRecapFormulaMilkOnlineOrder(t *testing.T) {
	ctrl, _, onlineOrderRepoMock, recapService := setupRecapService(t)
	defer ctrl.Finish()

	t.Run("should success profit recap formula milk online order", func(t *testing.T) {
		onlineOrderRepoMock.EXPECT().RecapProfitFormulaMilkOnlineOrder().Return(0, nil)

		result, err := recapService.ProfitRecapFormulaMilkOnlineOrder()
		assert.NoError(t, err)
		assert.Equal(t, 0, result)
	})
}

func TestRecapService_ProfitRecapBabyDiaperOnlineOrder(t *testing.T) {
	ctrl, _, onlineOrderRepoMock, recapService := setupRecapService(t)
	defer ctrl.Finish()

	t.Run("should success profit recap baby diaper online order", func(t *testing.T) {
		onlineOrderRepoMock.EXPECT().RecapProfitBabyDiaperOnlineOrder().Return(0, nil)

		result, err := recapService.ProfitRecapBabyDiaperOnlineOrder()
		assert.NoError(t, err)
		assert.Equal(t, 0, result)
	})
}

func TestRecapService_ProfitRecapAdultDiaperOnlineOrder(t *testing.T) {
	ctrl, _, onlineOrderRepoMock, recapService := setupRecapService(t)
	defer ctrl.Finish()

	t.Run("should success profit recap adult diaper online order", func(t *testing.T) {
		onlineOrderRepoMock.EXPECT().RecapProfitAdultDiaperOnlineOrder().Return(0, nil)

		result, err := recapService.ProfitRecapAdultDiaperOnlineOrder()
		assert.NoError(t, err)
		assert.Equal(t, 0, result)
	})
}
