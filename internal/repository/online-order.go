package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"kopoksu/internal/model"
)

type onlineOrderRepository struct {
	DB *gorm.DB
}

type OnlineOrderRepository interface {
	SaveOnlineOrder(onlineOrder model.OnlineOrder) error
	SaveDetailOnlineOrder(detailOnlineOrder model.DetailOnlineOrder) error
	GetAllOnlineOrder() ([]model.OnlineOrder, error)
	GetOnlineOrderById(id uuid.UUID) (model.OnlineOrder, error)
	GetDetailOnlineOrderByOnlineOrderId(id uuid.UUID) ([]model.DetailOnlineOrderResponse, error)
	GetAllDetailOnlineOrderByOnlineOrderId(id uuid.UUID) ([]model.DetailOnlineOrder, error)
	UpdateOnlineOrder(id uuid.UUID, data map[string]interface{}) error
	DeleteOnlineOrder(onlineOrder model.OnlineOrder) error
	DeleteDetailOnlineOrder(detailOnlineOrder model.DetailOnlineOrder) error
	CountOnlineOrderByStatus(status string) (int, error)
	RecapGrossProfitFormulaMilkOnlineOrder() (int, error)
	RecapGrossProfitBabyDiaperOnlineOrder() (int, error)
	RecapGrossProfitAdultDiaperOnlineOrder() (int, error)
}

func NewOnlineOrderRepository(DB *gorm.DB) *onlineOrderRepository {
	return &onlineOrderRepository{
		DB: DB,
	}
}

func (r *onlineOrderRepository) SaveOnlineOrder(onlineOrder model.OnlineOrder) error {
	err := r.DB.Create(&onlineOrder).Error
	return err
}

func (r *onlineOrderRepository) SaveDetailOnlineOrder(detailOnlineOrder model.DetailOnlineOrder) error {
	err := r.DB.Create(&detailOnlineOrder).Error
	return err
}

func (r *onlineOrderRepository) GetAllOnlineOrder() ([]model.OnlineOrder, error) {
	var onlineOrders []model.OnlineOrder
	err := r.DB.Find(&onlineOrders).Error
	return onlineOrders, err
}

func (r *onlineOrderRepository) GetOnlineOrderById(id uuid.UUID) (model.OnlineOrder, error) {
	var onlineOrder model.OnlineOrder
	err := r.DB.First(&onlineOrder, "id = ?", id).Error
	return onlineOrder, err
}

func (r *onlineOrderRepository) GetDetailOnlineOrderByOnlineOrderId(id uuid.UUID) ([]model.DetailOnlineOrderResponse, error) {
	var detailOnlineOrdersResponse []model.DetailOnlineOrderResponse
	err := r.DB.Table("detail_online_orders").
		Select("products.name, detail_online_orders.amount, detail_online_orders.amount * products.price as price").
		Joins("LEFT JOIN products ON products.id = detail_online_orders.product_id").
		Where("detail_online_orders.online_order_id = ?", id).
		Scan(&detailOnlineOrdersResponse).
		Error
	return detailOnlineOrdersResponse, err
}

func (r *onlineOrderRepository) GetAllDetailOnlineOrderByOnlineOrderId(id uuid.UUID) ([]model.DetailOnlineOrder, error) {
	var detailOnlineOrders []model.DetailOnlineOrder
	err := r.DB.Table("detail_online_orders").Where("online_order_id = ?", id).Find(&detailOnlineOrders).Error
	return detailOnlineOrders, err
}

func (r *onlineOrderRepository) UpdateOnlineOrder(id uuid.UUID, data map[string]interface{}) error {
	err := r.DB.Table("online_orders").Where("id = ?", id).Updates(data).Error
	return err
}

func (r *onlineOrderRepository) DeleteOnlineOrder(onlineOrder model.OnlineOrder) error {
	return r.DB.Delete(&onlineOrder).Error
}

func (r *onlineOrderRepository) DeleteDetailOnlineOrder(detailOnlineOrder model.DetailOnlineOrder) error {
	return r.DB.Delete(&detailOnlineOrder).Error
}

func (r *onlineOrderRepository) CountOnlineOrderByStatus(status string) (int, error) {
	var count int64
	err := r.DB.Table("online_orders").Where("status = ?", status).Count(&count).Error
	return int(count), err
}

func (r *onlineOrderRepository) RecapGrossProfitFormulaMilkOnlineOrder() (int, error) {
	var GrossProfit int

	query := `select coalesce(sum(gross_profit_online_order), 0) as total_gross_profit_online_order  
		from (
			SELECT p.price * doo.amount AS gross_profit_online_order
    		FROM detail_online_orders doo
    		INNER JOIN products p ON doo.product_id = p.id
    		inner join categories c on p.category_id = c.id
    		where c.id = 'ea600c63-283a-415e-8ed1-b10d12c544a0'
		) as gross_profit_online_order;`

	err := r.DB.Raw(query).First(&GrossProfit).Error

	return GrossProfit, err
}

func (r *onlineOrderRepository) RecapGrossProfitBabyDiaperOnlineOrder() (int, error) {
	var GrossProfit int

	query := `select coalesce(sum(gross_profit_online_order), 0) as total_gross_profit_online_order  
		from (
			SELECT p.price * doo.amount AS gross_profit_online_order
    		FROM detail_online_orders doo
    		INNER JOIN products p ON doo.product_id = p.id
    		inner join categories c on p.category_id = c.id
    		where c.id = '981464fb-3241-4a33-97ae-33b110e2d4aa'
		) as gross_profit_online_order;`

	err := r.DB.Raw(query).First(&GrossProfit).Error

	return GrossProfit, err
}

func (r *onlineOrderRepository) RecapGrossProfitAdultDiaperOnlineOrder() (int, error) {
	var GrossProfit int

	query := `select coalesce(sum(gross_profit_online_order), 0) as total_gross_profit_online_order  
		from (
			SELECT p.price * doo.amount AS gross_profit_online_order
    		FROM detail_online_orders doo
    		INNER JOIN products p ON doo.product_id = p.id
    		inner join categories c on p.category_id = c.id
    		where c.id = 'f5976ce9-7496-4fd2-8322-3beaef36e4d8'
		) as gross_profit_online_order;`

	err := r.DB.Raw(query).First(&GrossProfit).Error

	return GrossProfit, err
}
