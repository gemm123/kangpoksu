package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"kopoksu/internal/model"
)

type offlineOrderRepository struct {
	DB *gorm.DB
}

type OfflineOrderRepository interface {
	SaveOfflineOrder(offlineOrder model.OfflineOrder) error
	SaveDetailOfflineOrder(detailOfflineOrder model.DetailOfflineOrder) error
	GetAllOfflineOrder() ([]model.OfflineOrder, error)
	GetOfflineOrderById(id uuid.UUID) (model.OfflineOrder, error)
	GetDetailOfflineOrderByOfflineOrderId(id uuid.UUID) ([]model.DetailOfflineOrderResponse, error)
	UpdateOfflineOrder(id uuid.UUID, data map[string]interface{}) error
	DeleteOfflineOrder(offlineOrder model.OfflineOrder) error
	GetAllDetailOfflineOrderByOfflineOrderId(id uuid.UUID) ([]model.DetailOfflineOrder, error)
	DeleteDetailOfflineOrder(detailOfflineOrder model.DetailOfflineOrder) error
	CountOfflineOrderByStatus(status string) (int, error)
	RecapGrossProfitFormulaMilkOfflineOrder() (int, error)
	RecapGrossProfitBabyDiaperOfflineOrder() (int, error)
	RecapGrossProfitAdultDiaperOfflineOrder() (int, error)
	RecapNetProfitFormulaMilkOfflineOrder() (int, error)
	RecapNetProfitBabyDiaperOfflineOrder() (int, error)
	RecapNetProfitAdultDiaperOfflineOrder() (int, error)
}

func NewOfflineOrderRepository(DB *gorm.DB) *offlineOrderRepository {
	return &offlineOrderRepository{
		DB: DB,
	}
}

func (r *offlineOrderRepository) SaveOfflineOrder(offlineOrder model.OfflineOrder) error {
	err := r.DB.Create(&offlineOrder).Error
	return err
}

func (r *offlineOrderRepository) SaveDetailOfflineOrder(detailOfflineOrder model.DetailOfflineOrder) error {
	err := r.DB.Create(&detailOfflineOrder).Error
	return err
}

func (r *offlineOrderRepository) GetAllOfflineOrder() ([]model.OfflineOrder, error) {
	var offlineOrders []model.OfflineOrder
	err := r.DB.Find(&offlineOrders).Error
	return offlineOrders, err
}

func (r *offlineOrderRepository) GetOfflineOrderById(id uuid.UUID) (model.OfflineOrder, error) {
	var offlineOrder model.OfflineOrder
	err := r.DB.First(&offlineOrder, "id = ?", id).Error
	return offlineOrder, err
}

func (r *offlineOrderRepository) GetDetailOfflineOrderByOfflineOrderId(id uuid.UUID) ([]model.DetailOfflineOrderResponse, error) {
	var detailOfflineOrdersResponse []model.DetailOfflineOrderResponse
	err := r.DB.Table("detail_offline_orders").
		Select("products.name, detail_offline_orders.amount, detail_offline_orders.amount * products.price as price").
		Joins("LEFT JOIN products ON products.id = detail_offline_orders.product_id").
		Where("detail_offline_orders.offline_order_id = ?", id).
		Scan(&detailOfflineOrdersResponse).
		Error
	return detailOfflineOrdersResponse, err
}

func (r *offlineOrderRepository) GetAllDetailOfflineOrderByOfflineOrderId(id uuid.UUID) ([]model.DetailOfflineOrder, error) {
	var detailOfflineOrders []model.DetailOfflineOrder
	err := r.DB.Table("detail_offline_orders").Where("offline_order_id = ?", id).Find(&detailOfflineOrders).Error
	return detailOfflineOrders, err
}

func (r *offlineOrderRepository) UpdateOfflineOrder(id uuid.UUID, data map[string]interface{}) error {
	err := r.DB.Table("offline_orders").Where("id = ?", id).Updates(data).Error
	return err
}

func (r *offlineOrderRepository) DeleteOfflineOrder(offlineOrder model.OfflineOrder) error {
	return r.DB.Delete(&offlineOrder).Error
}

func (r *offlineOrderRepository) DeleteDetailOfflineOrder(detailOfflineOrder model.DetailOfflineOrder) error {
	return r.DB.Delete(&detailOfflineOrder).Error
}

func (r *offlineOrderRepository) CountOfflineOrderByStatus(status string) (int, error) {
	var count int64
	err := r.DB.Table("offline_orders").Where("status = ?", status).Count(&count).Error
	return int(count), err
}

func (r *offlineOrderRepository) RecapGrossProfitFormulaMilkOfflineOrder() (int, error) {
	var GrossProfit int

	query := `select coalesce(sum(gross_profit_offline_order), 0) as total_gross_profit_offline_order  
		from (
			SELECT p.price * doo.amount AS gross_profit_offline_order
    		FROM detail_offline_orders doo
    		INNER JOIN products p ON doo.product_id = p.id
    		inner join categories c on p.category_id = c.id
    		where c.id = 'ea600c63-283a-415e-8ed1-b10d12c544a0'
		) as gross_profit_offline_order;`

	err := r.DB.Raw(query).First(&GrossProfit).Error

	return GrossProfit, err
}

func (r *offlineOrderRepository) RecapGrossProfitBabyDiaperOfflineOrder() (int, error) {
	var GrossProfit int

	query := `select coalesce(sum(gross_profit_offline_order), 0) as total_gross_profit_offline_order  
		from (
			SELECT p.price * doo.amount AS gross_profit_offline_order
    		FROM detail_offline_orders doo
    		INNER JOIN products p ON doo.product_id = p.id
    		inner join categories c on p.category_id = c.id
    		where c.id = '981464fb-3241-4a33-97ae-33b110e2d4aa'
		) as gross_profit_offline_order;`

	err := r.DB.Raw(query).First(&GrossProfit).Error

	return GrossProfit, err
}

func (r *offlineOrderRepository) RecapGrossProfitAdultDiaperOfflineOrder() (int, error) {
	var GrossProfit int

	query := `select coalesce(sum(gross_profit_offline_order), 0) as total_gross_profit_offline_order  
		from (
			SELECT p.price * doo.amount AS gross_profit_offline_order
    		FROM detail_offline_orders doo
    		INNER JOIN products p ON doo.product_id = p.id
    		inner join categories c on p.category_id = c.id
    		where c.id = 'f5976ce9-7496-4fd2-8322-3beaef36e4d8'
		) as gross_profit_offline_order;`

	err := r.DB.Raw(query).First(&GrossProfit).Error

	return GrossProfit, err
}

func (r *offlineOrderRepository) RecapNetProfitFormulaMilkOfflineOrder() (int, error) {
	var NetProfit int

	query := `select coalesce(sum(net_profit_offline_order), 0) as total_net_profit_offline_order  
		from (
			SELECT (p.price - p.buy_price) * doo.amount AS net_profit_offline_order
    		FROM detail_offline_orders doo
    		INNER JOIN products p ON doo.product_id = p.id
    		inner join categories c on p.category_id = c.id
    		where c.id = 'ea600c63-283a-415e-8ed1-b10d12c544a0'
		) as net_profit_offline_order;`

	err := r.DB.Raw(query).First(&NetProfit).Error

	return NetProfit, err
}

func (r *offlineOrderRepository) RecapNetProfitBabyDiaperOfflineOrder() (int, error) {
	var NetProfit int

	query := `select coalesce(sum(net_profit_offline_order), 0) as total_net_profit_offline_order  
		from (
			SELECT (p.price - p.buy_price) * doo.amount AS net_profit_offline_order
    		FROM detail_offline_orders doo
    		INNER JOIN products p ON doo.product_id = p.id
    		inner join categories c on p.category_id = c.id
    		where c.id = '981464fb-3241-4a33-97ae-33b110e2d4aa'
		) as net_profit_offline_order;`

	err := r.DB.Raw(query).First(&NetProfit).Error

	return NetProfit, err
}

func (r *offlineOrderRepository) RecapNetProfitAdultDiaperOfflineOrder() (int, error) {
	var NetProfit int

	query := `select coalesce(sum(net_profit_offline_order), 0) as total_net_profit_offline_order  
		from (
			SELECT (p.price - p.buy_price) * doo.amount AS net_profit_offline_order
    		FROM detail_offline_orders doo
    		INNER JOIN products p ON doo.product_id = p.id
    		inner join categories c on p.category_id = c.id
    		where c.id = 'f5976ce9-7496-4fd2-8322-3beaef36e4d8'
		) as net_profit_offline_order;`

	err := r.DB.Raw(query).First(&NetProfit).Error

	return NetProfit, err
}
