package service

import (
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	mock_repository "kopoksu/mock/repository"
	"testing"
)

func setupRecapService(t *testing.T) (*gomock.Controller, *mock_repository.MockOfflineOrderRepository, *mock_repository.MockOnlineOrderRepository, RecapService) {
	ctrl := gomock.NewController(t)
	offlineOrderRepoMock := mock_repository.NewMockOfflineOrderRepository(ctrl)
	onlineOrderRepoMock := mock_repository.NewMockOnlineOrderRepository(ctrl)
	recapService := NewRecapService(offlineOrderRepoMock, onlineOrderRepoMock)
	return ctrl, offlineOrderRepoMock, onlineOrderRepoMock, recapService
}

func TestRecapService_ProfitRecapFormulaMilkOfflineOrder(t *testing.T) {
	ctrl, offlineOrderRepoMock, _, recapService := setupRecapService(t)
	defer ctrl.Finish()

	t.Run("should success profit recap formula milk offline order", func(t *testing.T) {
		offlineOrderRepoMock.EXPECT().RecapProfitFormulaMilkOfflineOrder().Return(0, nil)

		result, err := recapService.ProfitRecapFormulaMilkOfflineOrder()
		assert.NoError(t, err)
		assert.Equal(t, 0, result)
	})
}

func TestRecapService_ProfitRecapBabyDiaperOfflineOrder(t *testing.T) {
	ctrl, offlineOrderRepoMock, _, recapService := setupRecapService(t)
	defer ctrl.Finish()

	t.Run("should success profit recap baby diaper offline order", func(t *testing.T) {
		offlineOrderRepoMock.EXPECT().RecapProfitBabyDiaperOfflineOrder().Return(0, nil)

		result, err := recapService.ProfitRecapBabyDiaperOfflineOrder()
		assert.NoError(t, err)
		assert.Equal(t, 0, result)
	})
}

func TestRecapService_ProfitRecapAdultDiaperOfflineOrder(t *testing.T) {
	ctrl, offlineOrderRepoMock, _, recapService := setupRecapService(t)
	defer ctrl.Finish()

	t.Run("should success profit recap adult diaper offline order", func(t *testing.T) {
		offlineOrderRepoMock.EXPECT().RecapProfitAdultDiaperOfflineOrder().Return(0, nil)

		result, err := recapService.ProfitRecapAdultDiaperOfflineOrder()
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
