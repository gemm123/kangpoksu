package service

import (
	"kopoksu/internal/model"
	"kopoksu/internal/repository"
	"log"
	"time"
)

type recapService struct {
	pickupOnlineOrderRepository repository.PickupOnlineOrderRepository
	onlineOrderRepository       repository.OnlineOrderRepository
}

type RecapService interface {
	ProfitRecapFormulaMilkPickupOnlineOrder() (int, error)
	ProfitRecapBabyDiaperPickupOnlineOrder() (int, error)
	ProfitRecapAdultDiaperPickupOnlineOrder() (int, error)
	ProfitRecapFormulaMilkOnlineOrder() (int, error)
	ProfitRecapBabyDiaperOnlineOrder() (int, error)
	ProfitRecapAdultDiaperOnlineOrder() (int, error)
	RecapSalesFormulaMilkByMonthOnlineOrder() ([]model.RecapSalesByMonth, error)
	RecapSalesBabyDiaperByMonthOnlineOrder() ([]model.RecapSalesByMonth, error)
	RecapSalesAdultDiaperByMonthOnlineOrder() ([]model.RecapSalesByMonth, error)
	RecapSalesFormulaMilkByMonthPickupOnlineOrder() ([]model.RecapSalesByMonth, error)
	RecapSalesBabyDiaperByMonthPickupOnlineOrder() ([]model.RecapSalesByMonth, error)
	RecapSalesAdultDiaperByMonthPickupOnlineOrder() ([]model.RecapSalesByMonth, error)
}

func NewRecapService(pickupOnlineOrderRepository repository.PickupOnlineOrderRepository, onlineOrderRepository repository.OnlineOrderRepository) *recapService {
	return &recapService{
		pickupOnlineOrderRepository: pickupOnlineOrderRepository,
		onlineOrderRepository:       onlineOrderRepository,
	}
}

func (s *recapService) ProfitRecapFormulaMilkPickupOnlineOrder() (int, error) {
	ProfitFormulaMilkPickupOnlineOrder, err := s.pickupOnlineOrderRepository.RecapProfitFormulaMilkPickupOnlineOrder()
	if err != nil {
		log.Println("error: " + err.Error())
		return ProfitFormulaMilkPickupOnlineOrder, err
	}

	return ProfitFormulaMilkPickupOnlineOrder, nil
}

func (s *recapService) ProfitRecapBabyDiaperPickupOnlineOrder() (int, error) {
	ProfitBabyDiaperPickupOnlineOrder, err := s.pickupOnlineOrderRepository.RecapProfitBabyDiaperPickupOnlineOrder()
	if err != nil {
		log.Println("error: " + err.Error())
		return ProfitBabyDiaperPickupOnlineOrder, err
	}

	return ProfitBabyDiaperPickupOnlineOrder, nil
}

func (s *recapService) ProfitRecapAdultDiaperPickupOnlineOrder() (int, error) {
	ProfitAdultDiaperPickupOnlineOrder, err := s.pickupOnlineOrderRepository.RecapProfitAdultDiaperPickupOnlineOrder()
	if err != nil {
		log.Println("error: " + err.Error())
		return ProfitAdultDiaperPickupOnlineOrder, err
	}

	return ProfitAdultDiaperPickupOnlineOrder, nil
}

func (s *recapService) ProfitRecapFormulaMilkOnlineOrder() (int, error) {
	ProfitFormulaMilkOnlineOrder, err := s.onlineOrderRepository.RecapProfitFormulaMilkOnlineOrder()
	if err != nil {
		log.Println("error: " + err.Error())
		return ProfitFormulaMilkOnlineOrder, err
	}

	return ProfitFormulaMilkOnlineOrder, nil
}

func (s *recapService) ProfitRecapBabyDiaperOnlineOrder() (int, error) {
	ProfitBabyDiaperOnlineOrder, err := s.onlineOrderRepository.RecapProfitBabyDiaperOnlineOrder()
	if err != nil {
		log.Println("error: " + err.Error())
		return ProfitBabyDiaperOnlineOrder, err
	}

	return ProfitBabyDiaperOnlineOrder, nil
}

func (s *recapService) ProfitRecapAdultDiaperOnlineOrder() (int, error) {
	ProfitAdultDiaperOnlineOrder, err := s.onlineOrderRepository.RecapProfitAdultDiaperOnlineOrder()
	if err != nil {
		log.Println("error: " + err.Error())
		return ProfitAdultDiaperOnlineOrder, err
	}

	return ProfitAdultDiaperOnlineOrder, nil
}

func (s *recapService) RecapSalesFormulaMilkByMonthOnlineOrder() ([]model.RecapSalesByMonth, error) {
	recapSales, err := s.onlineOrderRepository.RecapSalesFormulaMilkByMonthOnlineOrder()
	if err != nil {
		log.Println("error: " + err.Error())
		return recapSales, err
	}

	now := time.Now()
	var sixMonthsBack []int

	for i := 0; i < 6; i++ {
		month := now.AddDate(0, -i, 0).Month()
		sixMonthsBack = append(sixMonthsBack, int(month))
	}

	for i := 0; i < len(sixMonthsBack)/2; i++ {
		j := len(sixMonthsBack) - i - 1
		sixMonthsBack[i], sixMonthsBack[j] = sixMonthsBack[j], sixMonthsBack[i]
	}

	var salesByMonths []model.RecapSalesByMonth
	for _, smb := range sixMonthsBack {
		var salesByMonth model.RecapSalesByMonth
		salesByMonth.Month = smb
		for _, recapSale := range recapSales {
			if smb == recapSale.Month {
				salesByMonth.Sold = recapSale.Sold
			}
		}
		salesByMonths = append(salesByMonths, salesByMonth)
	}

	return salesByMonths, nil
}

func (s *recapService) RecapSalesBabyDiaperByMonthOnlineOrder() ([]model.RecapSalesByMonth, error) {
	recapSales, err := s.onlineOrderRepository.RecapSalesBabyDiaperByMonthOnlineOrder()
	if err != nil {
		log.Println("error: " + err.Error())
		return recapSales, err
	}

	now := time.Now()
	var sixMonthsBack []int

	for i := 0; i < 6; i++ {
		month := now.AddDate(0, -i, 0).Month()
		sixMonthsBack = append(sixMonthsBack, int(month))
	}

	for i := 0; i < len(sixMonthsBack)/2; i++ {
		j := len(sixMonthsBack) - i - 1
		sixMonthsBack[i], sixMonthsBack[j] = sixMonthsBack[j], sixMonthsBack[i]
	}

	var salesByMonths []model.RecapSalesByMonth
	for _, smb := range sixMonthsBack {
		var salesByMonth model.RecapSalesByMonth
		salesByMonth.Month = smb
		for _, recapSale := range recapSales {
			if smb == recapSale.Month {
				salesByMonth.Sold = recapSale.Sold
			}
		}
		salesByMonths = append(salesByMonths, salesByMonth)
	}

	return salesByMonths, nil
}

func (s *recapService) RecapSalesAdultDiaperByMonthOnlineOrder() ([]model.RecapSalesByMonth, error) {
	recapSales, err := s.onlineOrderRepository.RecapSalesAdultDiaperByMonthOnlineOrder()
	if err != nil {
		log.Println("error: " + err.Error())
		return recapSales, err
	}

	now := time.Now()
	var sixMonthsBack []int

	for i := 0; i < 6; i++ {
		month := now.AddDate(0, -i, 0).Month()
		sixMonthsBack = append(sixMonthsBack, int(month))
	}

	for i := 0; i < len(sixMonthsBack)/2; i++ {
		j := len(sixMonthsBack) - i - 1
		sixMonthsBack[i], sixMonthsBack[j] = sixMonthsBack[j], sixMonthsBack[i]
	}

	var salesByMonths []model.RecapSalesByMonth
	for _, smb := range sixMonthsBack {
		var salesByMonth model.RecapSalesByMonth
		salesByMonth.Month = smb
		for _, recapSale := range recapSales {
			if smb == recapSale.Month {
				salesByMonth.Sold = recapSale.Sold
			}
		}
		salesByMonths = append(salesByMonths, salesByMonth)
	}

	return salesByMonths, nil
}

func (s *recapService) RecapSalesFormulaMilkByMonthPickupOnlineOrder() ([]model.RecapSalesByMonth, error) {
	recapSales, err := s.pickupOnlineOrderRepository.RecapSalesFormulaMilkByMonthPickupOnlineOrder()
	if err != nil {
		log.Println("error: " + err.Error())
		return recapSales, err
	}

	now := time.Now()
	var sixMonthsBack []int

	for i := 0; i < 6; i++ {
		month := now.AddDate(0, -i, 0).Month()
		sixMonthsBack = append(sixMonthsBack, int(month))
	}

	for i := 0; i < len(sixMonthsBack)/2; i++ {
		j := len(sixMonthsBack) - i - 1
		sixMonthsBack[i], sixMonthsBack[j] = sixMonthsBack[j], sixMonthsBack[i]
	}

	var salesByMonths []model.RecapSalesByMonth
	for _, smb := range sixMonthsBack {
		var salesByMonth model.RecapSalesByMonth
		salesByMonth.Month = smb
		for _, recapSale := range recapSales {
			if smb == recapSale.Month {
				salesByMonth.Sold = recapSale.Sold
			}
		}
		salesByMonths = append(salesByMonths, salesByMonth)
	}

	return salesByMonths, nil
}

func (s *recapService) RecapSalesBabyDiaperByMonthPickupOnlineOrder() ([]model.RecapSalesByMonth, error) {
	recapSales, err := s.pickupOnlineOrderRepository.RecapSalesBabyDiaperByMonthPickupOnlineOrder()
	if err != nil {
		log.Println("error: " + err.Error())
		return recapSales, err
	}

	now := time.Now()
	var sixMonthsBack []int

	for i := 0; i < 6; i++ {
		month := now.AddDate(0, -i, 0).Month()
		sixMonthsBack = append(sixMonthsBack, int(month))
	}

	for i := 0; i < len(sixMonthsBack)/2; i++ {
		j := len(sixMonthsBack) - i - 1
		sixMonthsBack[i], sixMonthsBack[j] = sixMonthsBack[j], sixMonthsBack[i]
	}

	var salesByMonths []model.RecapSalesByMonth
	for _, smb := range sixMonthsBack {
		var salesByMonth model.RecapSalesByMonth
		salesByMonth.Month = smb
		for _, recapSale := range recapSales {
			if smb == recapSale.Month {
				salesByMonth.Sold = recapSale.Sold
			}
		}
		salesByMonths = append(salesByMonths, salesByMonth)
	}

	return salesByMonths, nil
}

func (s *recapService) RecapSalesAdultDiaperByMonthPickupOnlineOrder() ([]model.RecapSalesByMonth, error) {
	recapSales, err := s.pickupOnlineOrderRepository.RecapSalesAdultDiaperByMonthPickupOnlineOrder()
	if err != nil {
		log.Println("error: " + err.Error())
		return recapSales, err
	}

	now := time.Now()
	var sixMonthsBack []int

	for i := 0; i < 6; i++ {
		month := now.AddDate(0, -i, 0).Month()
		sixMonthsBack = append(sixMonthsBack, int(month))
	}

	for i := 0; i < len(sixMonthsBack)/2; i++ {
		j := len(sixMonthsBack) - i - 1
		sixMonthsBack[i], sixMonthsBack[j] = sixMonthsBack[j], sixMonthsBack[i]
	}

	var salesByMonths []model.RecapSalesByMonth
	for _, smb := range sixMonthsBack {
		var salesByMonth model.RecapSalesByMonth
		salesByMonth.Month = smb
		for _, recapSale := range recapSales {
			if smb == recapSale.Month {
				salesByMonth.Sold = recapSale.Sold
			}
		}
		salesByMonths = append(salesByMonths, salesByMonth)
	}

	return salesByMonths, nil
}
