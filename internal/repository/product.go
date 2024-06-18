package repository

import (
	"fmt"
	"kopoksu/internal/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type productRepository struct {
	DB *gorm.DB
}

type ProductRepository interface {
	GetAllProductsFormulaMilk() ([]model.Product, error)
	SaveProduct(product model.Product) error
	DeleteProduct(product model.Product) error
	GetProductById(id uuid.UUID) (model.Product, error)
	UpdateProduct(product model.Product) error
	GetAllProductsBabyDiaper() ([]model.Product, error)
	GetAllProductsAdultDiaper() ([]model.Product, error)
	GetAllProductsFormulaMilkLimit(limit int) ([]model.Product, error)
	GetAllProductsBabyDiaperLimit(limit int) ([]model.Product, error)
	GetAllProductsAdultDiaperLimit(limit int) ([]model.Product, error)
	SearchProductsByName(name string) ([]model.SearchResult, error)
	ReportSalesFormulaMilkOnlineOrderByMonthYear(month, year int) ([]model.ProductSales, error)
	ReportSalesFormulaMilkOfflineOrderByMonthYear(month, year int) ([]model.ProductSales, error)
	ReportSalesBabyDiaperOnlineOrderByMonthYear(month, year int) ([]model.ProductSales, error)
	ReportSalesBabyDiaperOfflineOrderByMonthYear(month, year int) ([]model.ProductSales, error)
	ReportSalesAdultDiaperOnlineOrderByMonthYear(month, year int) ([]model.ProductSales, error)
	ReportSalesAdultDiaperOfflineOrderByMonthYear(month, year int) ([]model.ProductSales, error)
	ReportSalesFormulaMilkOnlineOrderByDate(date string) ([]model.ProductSales, error)
	ReportSalesFormulaMilkOfflineOrderByDate(date string) ([]model.ProductSales, error)
	ReportSalesBabyDiaperOnlineOrderByDate(date string) ([]model.ProductSales, error)
	ReportSalesBabyDiaperOfflineOrderByDate(date string) ([]model.ProductSales, error)
	ReportSalesAdultDiaperOnlineOrderByDate(date string) ([]model.ProductSales, error)
	ReportSalesAdultDiaperOfflineOrderByDate(date string) ([]model.ProductSales, error)
}

func NewProductRepository(DB *gorm.DB) *productRepository {
	return &productRepository{
		DB: DB,
	}
}

func (r *productRepository) GetAllProductsFormulaMilk() ([]model.Product, error) {
	var products []model.Product
	err := r.DB.Where("category_id = 'ea600c63-283a-415e-8ed1-b10d12c544a0' AND deleted_at is NULL").Find(&products).Error
	return products, err
}

func (r *productRepository) GetAllProductsFormulaMilkLimit(limit int) ([]model.Product, error) {
	var products []model.Product
	err := r.DB.Limit(limit).Where("category_id = 'ea600c63-283a-415e-8ed1-b10d12c544a0' AND deleted_at is NULL").Find(&products).Error
	return products, err
}

func (r *productRepository) GetAllProductsBabyDiaper() ([]model.Product, error) {
	var products []model.Product
	err := r.DB.Where("category_id = '981464fb-3241-4a33-97ae-33b110e2d4aa' AND deleted_at is NULL").Find(&products).Error
	return products, err
}

func (r *productRepository) GetAllProductsBabyDiaperLimit(limit int) ([]model.Product, error) {
	var products []model.Product
	err := r.DB.Limit(limit).Where("category_id = '981464fb-3241-4a33-97ae-33b110e2d4aa' AND deleted_at is NULL").Find(&products).Error
	return products, err
}

func (r *productRepository) GetAllProductsAdultDiaper() ([]model.Product, error) {
	var products []model.Product
	err := r.DB.Where("category_id = 'f5976ce9-7496-4fd2-8322-3beaef36e4d8' AND deleted_at is NULL").Find(&products).Error
	return products, err
}

func (r *productRepository) GetAllProductsAdultDiaperLimit(limit int) ([]model.Product, error) {
	var products []model.Product
	err := r.DB.Limit(limit).Where("category_id = 'f5976ce9-7496-4fd2-8322-3beaef36e4d8' AND deleted_at is NULL").Find(&products).Error
	return products, err
}

func (r *productRepository) SaveProduct(product model.Product) error {
	err := r.DB.Create(&product).Error
	return err
}

func (r *productRepository) DeleteProduct(product model.Product) error {
	return r.DB.Exec("UPDATE products SET deleted_at = NOW() WHERE id = ?", product.Id).Error
}

func (r *productRepository) GetProductById(id uuid.UUID) (model.Product, error) {
	var product model.Product
	err := r.DB.First(&product, "id = ?", id).Error
	return product, err
}

func (r *productRepository) UpdateProduct(product model.Product) error {
	return r.DB.Save(&product).Error
}

func (r *productRepository) SearchProductsByName(name string) ([]model.SearchResult, error) {
	var searchResult []model.SearchResult
	err := r.DB.Table("products").
		Joins("left join categories on products.category_id = categories.id").
		Select("products.id, products.\"name\", categories.\"name\" as category").
		Where("LOWER(products.\"name\") LIKE ?", fmt.Sprintf("%%%s%%", name)).
		Limit(3).Find(&searchResult).Error
	return searchResult, err
}

func (r *productRepository) ReportSalesFormulaMilkOnlineOrderByMonthYear(month, year int) ([]model.ProductSales, error) {
	var proudctSales []model.ProductSales

	query := fmt.Sprintf(`select p.id, p."name", sum(doo.amount) as sold
		from detail_online_orders doo 
		left join products p ON p.id = doo.product_id 
		left join categories c on c.id = p.category_id 
		where EXTRACT(YEAR FROM CAST(doo.created_at AS date)) = %d
			and EXTRACT(MONTH FROM CAST(doo.created_at AS date)) = %d
			and c.id = 'ea600c63-283a-415e-8ed1-b10d12c544a0'
		group by p.id`, year, month)

	err := r.DB.Raw(query).Find(&proudctSales).Error

	return proudctSales, err
}

func (r *productRepository) ReportSalesFormulaMilkOfflineOrderByMonthYear(month, year int) ([]model.ProductSales, error) {
	var proudctSales []model.ProductSales

	query := fmt.Sprintf(`select p.id, p."name", sum(doo.amount) as sold
		from detail_pickup_online_orders doo 
		left join products p ON p.id = doo.product_id 
		left join categories c on c.id = p.category_id 
		where EXTRACT(YEAR FROM CAST(doo.created_at AS date)) = %d
			and EXTRACT(MONTH FROM CAST(doo.created_at AS date)) = %d
			and c.id = 'ea600c63-283a-415e-8ed1-b10d12c544a0'
		group by p.id`, year, month)

	err := r.DB.Raw(query).Find(&proudctSales).Error

	return proudctSales, err
}

func (r *productRepository) ReportSalesBabyDiaperOnlineOrderByMonthYear(month, year int) ([]model.ProductSales, error) {
	var proudctSales []model.ProductSales

	query := fmt.Sprintf(`select p.id, p."name", sum(doo.amount) as sold
		from detail_online_orders doo 
		left join products p ON p.id = doo.product_id 
		left join categories c on c.id = p.category_id 
		where EXTRACT(YEAR FROM CAST(doo.created_at AS date)) = %d
			and EXTRACT(MONTH FROM CAST(doo.created_at AS date)) = %d
			and c.id = '981464fb-3241-4a33-97ae-33b110e2d4aa'
		group by p.id`, year, month)

	err := r.DB.Raw(query).Find(&proudctSales).Error

	return proudctSales, err
}

func (r *productRepository) ReportSalesBabyDiaperOfflineOrderByMonthYear(month, year int) ([]model.ProductSales, error) {
	var proudctSales []model.ProductSales

	query := fmt.Sprintf(`select p.id, p."name", sum(doo.amount) as sold
		from detail_pickup_online_orders doo 
		left join products p ON p.id = doo.product_id 
		left join categories c on c.id = p.category_id 
		where EXTRACT(YEAR FROM CAST(doo.created_at AS date)) = %d
			and EXTRACT(MONTH FROM CAST(doo.created_at AS date)) = %d
			and c.id = '981464fb-3241-4a33-97ae-33b110e2d4aa'
		group by p.id`, year, month)

	err := r.DB.Raw(query).Find(&proudctSales).Error

	return proudctSales, err
}

func (r *productRepository) ReportSalesAdultDiaperOnlineOrderByMonthYear(month, year int) ([]model.ProductSales, error) {
	var proudctSales []model.ProductSales

	query := fmt.Sprintf(`select p.id, p."name", sum(doo.amount) as sold
		from detail_online_orders doo 
		left join products p ON p.id = doo.product_id 
		left join categories c on c.id = p.category_id 
		where EXTRACT(YEAR FROM CAST(doo.created_at AS date)) = %d
			and EXTRACT(MONTH FROM CAST(doo.created_at AS date)) = %d
			and c.id = 'f5976ce9-7496-4fd2-8322-3beaef36e4d8'
		group by p.id`, year, month)

	err := r.DB.Raw(query).Find(&proudctSales).Error

	return proudctSales, err
}

func (r *productRepository) ReportSalesAdultDiaperOfflineOrderByMonthYear(month, year int) ([]model.ProductSales, error) {
	var proudctSales []model.ProductSales

	query := fmt.Sprintf(`select p.id, p."name", sum(doo.amount) as sold
		from detail_pickup_online_orders doo 
		left join products p ON p.id = doo.product_id 
		left join categories c on c.id = p.category_id 
		where EXTRACT(YEAR FROM CAST(doo.created_at AS date)) = %d
			and EXTRACT(MONTH FROM CAST(doo.created_at AS date)) = %d
			and c.id = 'f5976ce9-7496-4fd2-8322-3beaef36e4d8'
		group by p.id`, year, month)

	err := r.DB.Raw(query).Find(&proudctSales).Error

	return proudctSales, err
}

func (r *productRepository) ReportSalesFormulaMilkOnlineOrderByDate(date string) ([]model.ProductSales, error) {
	var proudctSales []model.ProductSales

	query := fmt.Sprintf(`select p.id, p."name", sum(doo.amount) as sold
		from detail_online_orders doo 
			left join products p ON p.id = doo.product_id 
			left join categories c on c.id = p.category_id 
		where date(doo.created_at) = '%s'
			and c.id = 'ea600c63-283a-415e-8ed1-b10d12c544a0'
		group by p.id`, date)

	err := r.DB.Raw(query).Find(&proudctSales).Error

	return proudctSales, err
}

func (r *productRepository) ReportSalesFormulaMilkOfflineOrderByDate(date string) ([]model.ProductSales, error) {
	var proudctSales []model.ProductSales

	query := fmt.Sprintf(`select p.id, p."name", sum(doo.amount) as sold
		from detail_pickup_online_orders doo 
			left join products p ON p.id = doo.product_id 
			left join categories c on c.id = p.category_id 
		where date(doo.created_at) = '%s'
			and c.id = 'ea600c63-283a-415e-8ed1-b10d12c544a0'
		group by p.id`, date)

	err := r.DB.Raw(query).Find(&proudctSales).Error

	return proudctSales, err
}

func (r *productRepository) ReportSalesBabyDiaperOnlineOrderByDate(date string) ([]model.ProductSales, error) {
	var proudctSales []model.ProductSales

	query := fmt.Sprintf(`select p.id, p."name", sum(doo.amount) as sold
		from detail_online_orders doo 
			left join products p ON p.id = doo.product_id 
			left join categories c on c.id = p.category_id 
		where date(doo.created_at) = '%s'
			and c.id = '981464fb-3241-4a33-97ae-33b110e2d4aa'
		group by p.id`, date)

	err := r.DB.Raw(query).Find(&proudctSales).Error

	return proudctSales, err
}

func (r *productRepository) ReportSalesBabyDiaperOfflineOrderByDate(date string) ([]model.ProductSales, error) {
	var proudctSales []model.ProductSales

	query := fmt.Sprintf(`select p.id, p."name", sum(doo.amount) as sold
		from detail_pickup_online_orders doo 
			left join products p ON p.id = doo.product_id 
			left join categories c on c.id = p.category_id 
		where date(doo.created_at) = '%s'
			and c.id = '981464fb-3241-4a33-97ae-33b110e2d4aa'
		group by p.id`, date)

	err := r.DB.Raw(query).Find(&proudctSales).Error

	return proudctSales, err
}

func (r *productRepository) ReportSalesAdultDiaperOnlineOrderByDate(date string) ([]model.ProductSales, error) {
	var proudctSales []model.ProductSales

	query := fmt.Sprintf(`select p.id, p."name", sum(doo.amount) as sold
		from detail_online_orders doo 
			left join products p ON p.id = doo.product_id 
			left join categories c on c.id = p.category_id 
		where date(doo.created_at) = '%s'
			and c.id = 'f5976ce9-7496-4fd2-8322-3beaef36e4d8'
		group by p.id`, date)

	err := r.DB.Raw(query).Find(&proudctSales).Error

	return proudctSales, err
}

func (r *productRepository) ReportSalesAdultDiaperOfflineOrderByDate(date string) ([]model.ProductSales, error) {
	var proudctSales []model.ProductSales

	query := fmt.Sprintf(`select p.id, p."name", sum(doo.amount) as sold
		from detail_pickup_online_orders doo 
			left join products p ON p.id = doo.product_id 
			left join categories c on c.id = p.category_id 
		where date(doo.created_at) = '%s'
			and c.id = 'f5976ce9-7496-4fd2-8322-3beaef36e4d8'
		group by p.id`, date)

	err := r.DB.Raw(query).Find(&proudctSales).Error

	return proudctSales, err
}
