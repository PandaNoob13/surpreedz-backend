package manager

import "surpreedz-backend/repository"

type RepositoryManager interface {
	AccountRepo() repository.AccountRepository
	AccountDetailRepo() repository.AccountDetailRepository
	PhotoProfileRepo() repository.PhotoProfileRepository
	ServiceDetailRepo() repository.ServiceDetailRepository
	ServicePriceRepo() repository.ServicePriceRepository
	VideoProfileRepo() repository.VideoProfileRepository
}

type repositoryManager struct {
	infra Infra
}

func (r *repositoryManager) AccountRepo() repository.AccountRepository {
	return repository.NewAccountRepository(r.infra.SqlDb())
}

func (r *repositoryManager) AccountDetailRepo() repository.AccountDetailRepository {
	return repository.NewAccountDetailRepository(r.infra.SqlDb())
}

func (r *repositoryManager) PhotoProfileRepo() repository.PhotoProfileRepository {
	return repository.NewPhotoProfileRepository(r.infra.SqlDb())
}

func (r *repositoryManager) ServiceDetailRepo() repository.ServiceDetailRepository {
	return repository.NewServiceDetailRepository(r.infra.SqlDb())
}

func (r *repositoryManager) ServicePriceRepo() repository.ServicePriceRepository {
	return repository.NewServicePriceRepository(r.infra.SqlDb())
}

func (r *repositoryManager) VideoProfileRepo() repository.VideoProfileRepository {
	return repository.NewVideoProfileRepository(r.infra.SqlDb())
}

func NewRepositoryManager(infra Infra) RepositoryManager {
	return &repositoryManager{
		infra: infra,
	}
}
