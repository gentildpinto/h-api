package service

type (
	Dependencies struct{}

	Service struct{}
)

func NewServices(deps Dependencies) *Service {
	return &Service{}
}
