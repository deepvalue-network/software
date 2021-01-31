package application

import "github.com/steve-care-software/products/identity/domain/users"

// NewApplication creates a new application instance
func NewApplication(
	repository users.Repository,
	service users.Service,
	encPkBitrate int,
) Application {
	builder := users.NewBuilder(encPkBitrate)
	return createApplication(repository, service, builder)
}

// NewUpdateBuilder creates a new update builder instance
func NewUpdateBuilder() UpdateBuilder {
	return createUpdateBuilder()
}

// Application represents an identity application
type Application interface {
	List() ([]string, error)
	Retrieve(name string, seed string, password string) (users.User, error)
	Insert(name string, seed string, password string) error
	Update(name string, seed string, password string, update Update) error
	Delete(name string, seed string, password string) error
}

// UpdateBuilder represents an update builder
type UpdateBuilder interface {
	Create() UpdateBuilder
	WithName(name string) UpdateBuilder
	WithPassword(pass string) UpdateBuilder
	Now() (Update, error)
}

// Update represents an update instance
type Update interface {
	HasName() bool
	Name() string
	HasPassword() bool
	Password() string
}
