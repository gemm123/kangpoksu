package service

import (
	"kopoksu/internal/model"
	"kopoksu/internal/repository"
	"log"
	"time"
)

type recapService struct {
	offlineOrderRepository repository.OfflineOrderRepository
	onlineOrderRepository  repository.OnlineOrderRepository
}

type RecapService interface {
	//GrossProfitRecapFormulaMilkOfflineOrder() (int, error)
	//GrossProfitRecapBabyDiaperOfflineOrder() (int, error)
	//GrossProfitRecapAdultDiaperOfflineOrder() (int, error)
	//GrossProfitRecapFormulaMilkOnlineOrder() (int, error)
	//GrossProfitRecapBabyDiaperOnlineOrder() (int, error)
	//GrossProfitRecapAdultDiaperOnlineOrder() (int, error)
	ProfitRecapFormulaMilkOfflineOrder() (int, error)
	ProfitRecapBabyDiaperOfflineOrder() (int, error)
	ProfitRecapAdultDiaperOfflineOrder() (int, error)
	ProfitRecapFormulaMilkOnlineOrder() (int, error)
	ProfitRecapBabyDiaperOnlineOrder() (int, error)
	ProfitRecapAdultDiaperOnlineOrder() (int, error)
	RecapSalesFormulaMilkByMonthOnlineOrder() ([]model.RecapSalesByMonth, error)
	RecapSalesBabyDiaperByMonthOnlineOrder() ([]model.RecapSalesByMonth, error)
	RecapSalesAdultDiaperByMonthOnlineOrder() ([]model.RecapSalesByMonth, error)
	RecapSalesFormulaMilkByMonthOfflineOrder() ([]model.RecapSalesByMonth, error)
	RecapSalesBabyDiaperByMonthOfflineOrder() ([]model.RecapSalesByMonth, error)
	RecapSalesAdultDiaperByMonthOfflineOrder() ([]model.RecapSalesByMonth, error)
}

func NewRecapService(offlineOrderRepository repository.OfflineOrderRepository, onlineOrderRepository repository.OnlineOrderRepository) *recapService {
	return &recapService{
		offlineOrderRepository: offlineOrderRepository,
		onlineOrderRepository:  onlineOrderRepository,
	}
}

//func (s *recapService) GrossProfitRecapFormulaMilkOfflineOrder() (int, error) {
//	grossProfitFormulaMilkOfflineOrder, err := s.offlineOrderRepository.RecapGrossProfitFormulaMilkOfflineOrder()
//	if err != nil {
//		log.Println("error: " + err.Error())
//		return grossProfitFormulaMilkOfflineOrder, err
//	}
//
//	return grossProfitFormulaMilkOfflineOrder, nil
//}
//
//func (s *recapService) GrossProfitRecapBabyDiaperOfflineOrder() (int, error) {
//	grossProfitBabyDiaperOfflineOrder, err := s.offlineOrderRepository.RecapGrossProfitBabyDiaperOfflineOrder()
//	if err != nil {
//		log.Println("error: " + err.Error())
//		return grossProfitBabyDiaperOfflineOrder, err
//	}
//
//	return grossProfitBabyDiaperOfflineOrder, nil
//}
//
//func (s *recapService) GrossProfitRecapAdultDiaperOfflineOrder() (int, error) {
//	grossProfitAdultDiaperOfflineOrder, err := s.offlineOrderRepository.RecapGrossProfitAdultDiaperOfflineOrder()
//	if err != nil {
//		log.Println("error: " + err.Error())
//		return grossProfitAdultDiaperOfflineOrder, err
//	}
//
//	return grossProfitAdultDiaperOfflineOrder, nil
//}
//
//func (s *recapService) GrossProfitRecapFormulaMilkOnlineOrder() (int, error) {
//	grossProfitFormulaMilkOnlineOrder, err := s.onlineOrderRepository.RecapGrossProfitFormulaMilkOnlineOrder()
//	if err != nil {
//		log.Println("error: " + err.Error())
//		return grossProfitFormulaMilkOnlineOrder, err
//	}
//
//	return grossProfitFormulaMilkOnlineOrder, nil
//}
//
//func (s *recapService) GrossProfitRecapBabyDiaperOnlineOrder() (int, error) {
//	grossProfitBabyDiaperOnlineOrder, err := s.onlineOrderRepository.RecapGrossProfitBabyDiaperOnlineOrder()
//	if err != nil {
//		log.Println("error: " + err.Error())
//		return grossProfitBabyDiaperOnlineOrder, err
//	}
//
//	return grossProfitBabyDiaperOnlineOrder, nil
//}
//
//func (s *recapService) GrossProfitRecapAdultDiaperOnlineOrder() (int, error) {
//	grossProfitAdultDiaperOnlineOrder, err := s.onlineOrderRepository.RecapGrossProfitAdultDiaperOnlineOrder()
//	if err != nil {
//		log.Println("error: " + err.Error())
//		return grossProfitAdultDiaperOnlineOrder, err
//	}
//
//	return grossProfitAdultDiaperOnlineOrder, nil
//}

func (s *recapService) ProfitRecapFormulaMilkOfflineOrder() (int, error) {
	ProfitFormulaMilkOfflineOrder, err := s.offlineOrderRepository.RecapProfitFormulaMilkOfflineOrder()
	if err != nil {
		log.Println("error: " + err.Error())
		return ProfitFormulaMilkOfflineOrder, err
	}

	return ProfitFormulaMilkOfflineOrder, nil
}

func (s *recapService) ProfitRecapBabyDiaperOfflineOrder() (int, error) {
	ProfitBabyDiaperOfflineOrder, err := s.offlineOrderRepository.RecapProfitBabyDiaperOfflineOrder()
	if err != nil {
		log.Println("error: " + err.Error())
		return ProfitBabyDiaperOfflineOrder, err
	}

	return ProfitBabyDiaperOfflineOrder, nil
}

func (s *recapService) ProfitRecapAdultDiaperOfflineOrder() (int, error) {
	ProfitAdultDiaperOfflineOrder, err := s.offlineOrderRepository.RecapProfitAdultDiaperOfflineOrder()
	if err != nil {
		log.Println("error: " + err.Error())
		return ProfitAdultDiaperOfflineOrder, err
	}

	return ProfitAdultDiaperOfflineOrder, nil
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

func (s *recapService) RecapSalesFormulaMilkByMonthOfflineOrder() ([]model.RecapSalesByMonth, error) {
	recapSales, err := s.offlineOrderRepository.RecapSalesFormulaMilkByMonthOfflineOrder()
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

func (s *recapService) RecapSalesBabyDiaperByMonthOfflineOrder() ([]model.RecapSalesByMonth, error) {
	recapSales, err := s.offlineOrderRepository.RecapSalesBabyDiaperByMonthOfflineOrder()
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

func (s *recapService) RecapSalesAdultDiaperByMonthOfflineOrder() ([]model.RecapSalesByMonth, error) {
	recapSales, err := s.offlineOrderRepository.RecapSalesAdultDiaperByMonthOfflineOrder()
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
