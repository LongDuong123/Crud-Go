package usecase

import domain "crud/Domain"

type productUseCase struct {
	Repository      domain.ProductRepository
	RepositoryRedis domain.ProductRepository
}

func NewProductUseCase(rp domain.ProductRepository, rpredis domain.ProductRepository) *productUseCase {
	return &productUseCase{Repository: rp, RepositoryRedis: rpredis}
}

func (pr *productUseCase) GetByID(id int) (*domain.Product, error) {
	product, err := pr.RepositoryRedis.GetByID(id)
	if err != nil {
		return nil, err
	}
	if product == nil {
		product, err = pr.Repository.GetByID(id)
		if err != nil {
			return nil, err
		}
		err = pr.RepositoryRedis.Create(product)
		if err != nil {
			return nil, err
		}
	}
	return product, nil
}

func (pr *productUseCase) GetByPage(page_size int, page_number int) ([]domain.Product, error) {
	return pr.Repository.GetByPage(page_size, page_number)
}

func (pr *productUseCase) Create(product *domain.Product) error {
	return pr.Repository.Create(product)
}

func (pr *productUseCase) UpdateByID(id int, product *domain.Product) error {
	err := pr.RepositoryRedis.Create(product)
	if err != nil {
		return err
	}
	return pr.Repository.UpdateByID(id, product)
}

func (pr *productUseCase) DeleteByID(id int) error {
	err := pr.RepositoryRedis.DeleteByID(id)
	if err != nil {
		return err
	}
	return pr.Repository.DeleteByID(id)
}
