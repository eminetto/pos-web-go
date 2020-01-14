package beer

//define a interface com as funções que serão usadas pelo restante do projeto
type UseCase interface {
	GetAll() ([]*Beer, error)
	Get(ID int64) (*Beer, error)
	Store(b *Beer) error
	Update(b *Beer) error
	Remove(ID int64) error
}

//em Go qualquer coisa que implemente as funções de uma interface
//passa a ser uma implementação válida. Não existe a palavra "implements" como em Java ou PHP
//desta forma, uma struct, uma string, um inteiro, etc, qualquer coisa pode ser válido, desde
//que implemente todas as funções
type Service struct{}

//esta função retorna um ponteiro em memória para uma estrutura
func NewService() *Service {
	return &Service{}
}

//vamos implementar as funções na próxima etapa
func (s *Service) GetAll() ([]*Beer, error) {
	return nil, nil
}

func (s *Service) Get(ID int64) (*Beer, error) {
	return nil, nil
}
func (s *Service) Store(b *Beer) error {
	return nil
}
func (s *Service) Update(b *Beer) error {
	return nil
}
func (s *Service) Remove(ID int64) error {
	return nil
}
