package service

import (
	"kopoksu/internal/repository"
	"log"
)

type recapService struct {
	offlineOrderRepository repository.OfflineOrderRepository
	onlineOrderRepository  repository.OnlineOrderRepository
}

type RecapService interface {
	GrossProfitRecapFormulaMilkOfflineOrder() (int, error)
	GrossProfitRecapBabyDiaperOfflineOrder() (int, error)
	GrossProfitRecapAdultDiaperOfflineOrder() (int, error)
	GrossProfitRecapFormulaMilkOnlineOrder() (int, error)
	GrossProfitRecapBabyDiaperOnlineOrder() (int, error)
	GrossProfitRecapAdultDiaperOnlineOrder() (int, error)
}

func NewRecapService(offlineOrderRepository repository.OfflineOrderRepository, onlineOrderRepository repository.OnlineOrderRepository) *recapService {
	return &recapService{
		offlineOrderRepository: offlineOrderRepository,
		onlineOrderRepository:  onlineOrderRepository,
	}
}

func (s *recapService) GrossProfitRecapFormulaMilkOfflineOrder() (int, error) {
	grossProfitFormulaMilkOfflineOrder, err := s.offlineOrderRepository.RecapGrossProfitFormulaMilkOfflineOrder()
	if err != nil {
		log.Println("error: " + err.Error())
		return grossProfitFormulaMilkOfflineOrder, err
	}

	return grossProfitFormulaMilkOfflineOrder, nil
}

func (s *recapService) GrossProfitRecapBabyDiaperOfflineOrder() (int, error) {
	grossProfitBabyDiaperOfflineOrder, err := s.offlineOrderRepository.RecapGrossProfitBabyDiaperOfflineOrder()
	if err != nil {
		log.Println("error: " + err.Error())
		return grossProfitBabyDiaperOfflineOrder, err
	}

	return grossProfitBabyDiaperOfflineOrder, nil
}

func (s *recapService) GrossProfitRecapAdultDiaperOfflineOrder() (int, error) {
	grossProfitAdultDiaperOfflineOrder, err := s.offlineOrderRepository.RecapGrossProfitAdultDiaperOfflineOrder()
	if err != nil {
		log.Println("error: " + err.Error())
		return grossProfitAdultDiaperOfflineOrder, err
	}

	return grossProfitAdultDiaperOfflineOrder, nil
}

func (s *recapService) GrossProfitRecapFormulaMilkOnlineOrder() (int, error) {
	grossProfitFormulaMilkOnlineOrder, err := s.onlineOrderRepository.RecapGrossProfitFormulaMilkOnlineOrder()
	if err != nil {
		log.Println("error: " + err.Error())
		return grossProfitFormulaMilkOnlineOrder, err
	}

	return grossProfitFormulaMilkOnlineOrder, nil
}

func (s *recapService) GrossProfitRecapBabyDiaperOnlineOrder() (int, error) {
	grossProfitBabyDiaperOnlineOrder, err := s.onlineOrderRepository.RecapGrossProfitBabyDiaperOnlineOrder()
	if err != nil {
		log.Println("error: " + err.Error())
		return grossProfitBabyDiaperOnlineOrder, err
	}

	return grossProfitBabyDiaperOnlineOrder, nil
}

func (s *recapService) GrossProfitRecapAdultDiaperOnlineOrder() (int, error) {
	grossProfitAdultDiaperOnlineOrder, err := s.onlineOrderRepository.RecapGrossProfitAdultDiaperOnlineOrder()
	if err != nil {
		log.Println("error: " + err.Error())
		return grossProfitAdultDiaperOnlineOrder, err
	}

	return grossProfitAdultDiaperOnlineOrder, nil
}
