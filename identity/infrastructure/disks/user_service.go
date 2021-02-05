package disks

import "github.com/deepvalue-network/software/identity/domain/users"

type userService struct {
	basePath string
}

func createUserService(basePath string) users.Service {
	out := userService{
		basePath: basePath,
	}

	return &out
}

// Insert inserts a new user
func (app *userService) Insert(user users.User, password string) error {
	//filePath := filepath.Join(app.basePath, name)

	return nil
}

// Update updates an existing user
func (app *userService) Update(original users.User, updated users.User, originalPass string) error {
	return nil
}

// UpdateWithPassword updates an existing user and changes its password
func (app *userService) UpdateWithPassword(original users.User, updated users.User, originalPass string, updatedPassword string) error {
	return nil
}

// Delete deletes a user
func (app *userService) Delete(user users.User, password string) error {
	return nil
}
