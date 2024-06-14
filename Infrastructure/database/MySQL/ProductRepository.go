package mysql

import (
	domain "crud/Domain"
)

type ProductRepository struct {
	db *databaseMySQL
}

func NewProductRepository(_db *databaseMySQL) domain.ProductRepository {
	return &ProductRepository{db: _db}
}

func (ur *ProductRepository) GetByID(id int) (*domain.Product, error) {
	Product, err := ur.db.Mysql.Query("SELECT * FROM product WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	var getProduct domain.Product
	if Product.Next() {
		err := Product.Scan(&getProduct.ID, &getProduct.Name, &getProduct.Image_url, &getProduct.Price, &getProduct.Create_By)
		if err != nil {
			return nil, err
		}
	}
	return &getProduct, nil
}

func (ur *ProductRepository) GetByPage(start int, end int) ([]domain.Product, error) {
	Product, err := ur.db.Mysql.Query("SELECT * FROM product WHERE id >= ? and id <= ?", start, end)
	if err != nil {
		return nil, err
	}
	var getListProduct []domain.Product
	for Product.Next() {
		var getProduct domain.Product
		err := Product.Scan(&getProduct.ID, &getProduct.Name, &getProduct.Image_url, &getProduct.Price, &getProduct.Create_By)
		if err != nil {
			return nil, err
		}
		getListProduct = append(getListProduct, getProduct)
	}
	return getListProduct, nil
}

func (ur *ProductRepository) Create(Product *domain.Product) error {
	_, err := ur.db.Mysql.Exec("INSERT INTO product (name,image_url,price,is_created_by) VALUES (?,?,?,?)", Product.Name, Product.Image_url, Product.Price, Product.Create_By)
	if err != nil {
		return err
	}
	return nil
}

func (ur *ProductRepository) UpdateByID(id int, Product *domain.Product) error {
	_, err := ur.db.Mysql.Exec("UPDATE product SET name = ?, image_url = ?, price = ?, is_created_by = ? WHERE id = ?", Product.Name, Product.Image_url, Product.Price, Product.Create_By, id)
	if err != nil {
		return err
	}
	return nil
}

func (ur *ProductRepository) DeleteByID(id int) error {
	_, err := ur.db.Mysql.Exec("DELETE FROM product WHERE id =?", id)
	if err != nil {
		return err
	}
	return nil
}
