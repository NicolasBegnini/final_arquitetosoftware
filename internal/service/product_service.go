package service // <-- Mudança aqui

type ProductService struct {
	Repository repositories.ProductRepository
}

func (s ProductService) Create(product *models.Product) error {
	return s.Repository.Create(product)
}

func (s ProductService) GetAll() ([]models.Product, error) {
	return s.Repository.FindAll()
}

func (s ProductService) GetByID(id uint) (models.Product, error) {
	return s.Repository.FindByID(id)
}

func (s ProductService) GetByName(name string) ([]models.Product, error) {
	return s.Repository.FindByName(name)
}

func (s ProductService) Update(product *models.Product) error {
	return s.Repository.Update(product)
}

func (s ProductService) Delete(id uint) error {
	return s.Repository.Delete(id)
}

func (s ProductService) Count() (int64, error) {
	return s.Repository.Count()
}
