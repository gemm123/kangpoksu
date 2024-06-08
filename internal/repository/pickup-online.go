package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"kopoksu/internal/model"
)

type pickupOnlineOrderRepository struct {
	DB *gorm.DB
}

type PickupOnlineOrderRepository interface {
	SavePickupOnlineOrder(PickupOnlineOrder model.PickupOnlineOrder) error
	SaveDetailPickupOnlineOrder(detailPickupOnlineOrder model.DetailPickupOnlineOrder) error
	GetAllPickupOnlineOrder() ([]model.PickupOnlineOrder, error)
	GetPickupOnlineOrderById(id uuid.UUID) (model.PickupOnlineOrder, error)
	GetDetailPickupOnlineOrderByPickupOnlineOrderId(id uuid.UUID) ([]model.DetailPickupOnlineOrderResponse, error)
	UpdatePickupOnlineOrder(id uuid.UUID, data map[string]interface{}) error
	DeletePickupOnlineOrder(PickupOnlineOrder model.PickupOnlineOrder) error
	GetAllDetailPickupOnlineOrderByPickupOnlineOrderId(id uuid.UUID) ([]model.DetailPickupOnlineOrder, error)
	DeleteDetailPickupOnlineOrder(detailPickupOnlineOrder model.DetailPickupOnlineOrder) error
	CountPickupOnlineOrderByStatus(status string) (int, error)
	RecapProfitFormulaMilkPickupOnlineOrder() (int, error)
	RecapProfitBabyDiaperPickupOnlineOrder() (int, error)
	RecapProfitAdultDiaperPickupOnlineOrder() (int, error)
	RecapSalesFormulaMilkByMonthPickupOnlineOrder() ([]model.RecapSalesByMonth, error)
	RecapSalesBabyDiaperByMonthPickupOnlineOrder() ([]model.RecapSalesByMonth, error)
	RecapSalesAdultDiaperByMonthPickupOnlineOrder() ([]model.RecapSalesByMonth, error)
}

func NewPickupOnlineOrderRepository(DB *gorm.DB) *pickupOnlineOrderRepository {
	return &pickupOnlineOrderRepository{
		DB: DB,
	}
}

func (r *pickupOnlineOrderRepository) SavePickupOnlineOrder(pickupOnlineOrder model.PickupOnlineOrder) error {
	err := r.DB.Table("pickup_online_orders").Create(&pickupOnlineOrder).Error
	return err
}

func (r *pickupOnlineOrderRepository) SaveDetailPickupOnlineOrder(detailPickupOnlineOrder model.DetailPickupOnlineOrder) error {
	err := r.DB.Table("detail_pickup_online_orders").Create(&detailPickupOnlineOrder).Error
	return err
}

func (r *pickupOnlineOrderRepository) GetAllPickupOnlineOrder() ([]model.PickupOnlineOrder, error) {
	var pickupOnlineOrders []model.PickupOnlineOrder
	err := r.DB.Table("pickup_online_orders").Find(&pickupOnlineOrders).Error
	return pickupOnlineOrders, err
}

func (r *pickupOnlineOrderRepository) GetPickupOnlineOrderById(id uuid.UUID) (model.PickupOnlineOrder, error) {
	var pickupOnlineOrder model.PickupOnlineOrder
	err := r.DB.Table("pickup_online_orders").First(&pickupOnlineOrder, "id = ?", id).Error
	return pickupOnlineOrder, err
}

func (r *pickupOnlineOrderRepository) GetDetailPickupOnlineOrderByPickupOnlineOrderId(id uuid.UUID) ([]model.DetailPickupOnlineOrderResponse, error) {
	var detailPickupOnlineOrdersResponse []model.DetailPickupOnlineOrderResponse
	err := r.DB.Table("detail_pickup_online_orders").
		Select("products.name, detail_pickup_online_orders.amount, detail_pickup_online_orders.amount * products.price as price").
		Joins("LEFT JOIN products ON products.id = detail_pickup_online_orders.product_id").
		Where("detail_pickup_online_orders.pickup_online_order_id = ?", id).
		Scan(&detailPickupOnlineOrdersResponse).
		Error
	return detailPickupOnlineOrdersResponse, err
}

func (r *pickupOnlineOrderRepository) GetAllDetailPickupOnlineOrderByPickupOnlineOrderId(id uuid.UUID) ([]model.DetailPickupOnlineOrder, error) {
	var detailPickupOnlineOrders []model.DetailPickupOnlineOrder
	err := r.DB.Table("detail_pickup_online_orders").Where("pickup_online_order_id = ?", id).Find(&detailPickupOnlineOrders).Error
	return detailPickupOnlineOrders, err
}

func (r *pickupOnlineOrderRepository) UpdatePickupOnlineOrder(id uuid.UUID, data map[string]interface{}) error {
	err := r.DB.Table("pickup_online_orders").Where("id = ?", id).Updates(data).Error
	return err
}

func (r *pickupOnlineOrderRepository) DeletePickupOnlineOrder(pickupOnlineOrder model.PickupOnlineOrder) error {
	return r.DB.Table("pickup_online_orders").Delete(&pickupOnlineOrder).Error
}

func (r *pickupOnlineOrderRepository) DeleteDetailPickupOnlineOrder(detailPickupOnlineOrder model.DetailPickupOnlineOrder) error {
	return r.DB.Table("detail_pickup_online_orders").Delete(&detailPickupOnlineOrder).Error
}

func (r *pickupOnlineOrderRepository) CountPickupOnlineOrderByStatus(status string) (int, error) {
	var count int64
	err := r.DB.Table("pickup_online_orders").Where("status = ?", status).Count(&count).Error
	return int(count), err
}

func (r *pickupOnlineOrderRepository) RecapProfitFormulaMilkPickupOnlineOrder() (int, error) {
	var NetProfit int

	query := `select coalesce(sum(net_profit_pickup_online_order), 0) as total_net_profit_pickup_online_order  
		from (
			SELECT (p.price - p.buy_price) * doo.amount AS net_profit_pickup_online_order
    		FROM detail_pickup_online_orders doo
    		INNER JOIN products p ON doo.product_id = p.id
    		inner join categories c on p.category_id = c.id
    		where c.id = 'ea600c63-283a-415e-8ed1-b10d12c544a0'
		) as net_profit_pickup_online_order;`

	err := r.DB.Raw(query).First(&NetProfit).Error

	return NetProfit, err
}

func (r *pickupOnlineOrderRepository) RecapProfitBabyDiaperPickupOnlineOrder() (int, error) {
	var NetProfit int

	query := `select coalesce(sum(net_profit_pickup_online_order), 0) as total_net_profit_pickup_online_order  
		from (
			SELECT (p.price - p.buy_price) * doo.amount AS net_profit_pickup_online_order
    		FROM detail_pickup_online_orders doo
    		INNER JOIN products p ON doo.product_id = p.id
    		inner join categories c on p.category_id = c.id
    		where c.id = '981464fb-3241-4a33-97ae-33b110e2d4aa'
		) as net_profit_pickup_online_order;`

	err := r.DB.Raw(query).First(&NetProfit).Error

	return NetProfit, err
}

func (r *pickupOnlineOrderRepository) RecapProfitAdultDiaperPickupOnlineOrder() (int, error) {
	var NetProfit int

	query := `select coalesce(sum(net_profit_pickup_online_order), 0) as total_net_profit_pickup_online_order  
		from (
			SELECT (p.price - p.buy_price) * doo.amount AS net_profit_pickup_online_order
    		FROM detail_pickup_online_orders doo
    		INNER JOIN products p ON doo.product_id = p.id
    		inner join categories c on p.category_id = c.id
    		where c.id = 'f5976ce9-7496-4fd2-8322-3beaef36e4d8'
		) as net_profit_pickup_online_order;`

	err := r.DB.Raw(query).First(&NetProfit).Error

	return NetProfit, err
}

func (r *pickupOnlineOrderRepository) RecapSalesFormulaMilkByMonthPickupOnlineOrder() ([]model.RecapSalesByMonth, error) {
	var RecapSalesByMonth []model.RecapSalesByMonth

	query := `SELECT to_char(date_trunc('month', doo.created_at), 'MM') AS bulan, SUM(doo.amount) AS terjual
		FROM detail_pickup_online_orders doo  
			left join products p on doo.product_id = p.id 
			left join categories c on p.category_id = c.id 
		WHERE doo.created_at >= current_date - interval '6 months' 
			AND doo.created_at < current_date + interval '1 day' 
		    and c.id = 'ea600c63-283a-415e-8ed1-b10d12c544a0'
		group by date_trunc('month', doo.created_at)
		order by bulan;`

	err := r.DB.Raw(query).Scan(&RecapSalesByMonth).Error

	return RecapSalesByMonth, err
}

func (r *pickupOnlineOrderRepository) RecapSalesBabyDiaperByMonthPickupOnlineOrder() ([]model.RecapSalesByMonth, error) {
	var RecapSalesByMonth []model.RecapSalesByMonth

	query := `SELECT to_char(date_trunc('month', doo.created_at), 'MM') AS bulan, SUM(doo.amount) AS terjual
		FROM detail_pickup_online_orders doo  
			left join products p on doo.product_id = p.id 
			left join categories c on p.category_id = c.id 
		WHERE doo.created_at >= current_date - interval '6 months' 
			AND doo.created_at < current_date + interval '1 day' 
		    and c.id = '981464fb-3241-4a33-97ae-33b110e2d4aa'
		group by date_trunc('month', doo.created_at)
		order by bulan;`

	err := r.DB.Raw(query).Scan(&RecapSalesByMonth).Error

	return RecapSalesByMonth, err
}

func (r *pickupOnlineOrderRepository) RecapSalesAdultDiaperByMonthPickupOnlineOrder() ([]model.RecapSalesByMonth, error) {
	var RecapSalesByMonth []model.RecapSalesByMonth

	query := `SELECT to_char(date_trunc('month', doo.created_at), 'MM') AS bulan, SUM(doo.amount) AS terjual
		FROM detail_pickup_online_orders doo  
			left join products p on doo.product_id = p.id 
			left join categories c on p.category_id = c.id 
		WHERE doo.created_at >= current_date - interval '6 months' 
			AND doo.created_at < current_date + interval '1 day' 
		    and c.id = 'f5976ce9-7496-4fd2-8322-3beaef36e4d8'
		group by date_trunc('month', doo.created_at)
		order by bulan;`

	err := r.DB.Raw(query).Scan(&RecapSalesByMonth).Error

	return RecapSalesByMonth, err
}
