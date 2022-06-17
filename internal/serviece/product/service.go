package product

type Serviece struct {}

func NewServiece() *Serviece{
  return &Serviece{}
}

func (s *Serviece) List() []Product{
  return allProducts
}
