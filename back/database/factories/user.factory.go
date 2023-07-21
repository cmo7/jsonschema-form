package factories

import (
	"nartex/ngr-stack/app/models"
)

type UserFactory struct{}

func (UserFactory) CreateOne() *models.User {
	user := models.User{
		Name:     Faker.Person().Name(),
		Email:    Faker.Internet().Email(),
		Password: Faker.Internet().Password(),
		Avatar:   "https://i.pravatar.cc/300?u=" + Faker.UUID().V4(),
		Provider: "local",
	}
	return &user
}

func (UserFactory) CreateMany(count int) []*models.User {
	var users []*models.User
	var user *models.User
	for i := 0; i < count; i++ {
		user = UserFactory{}.CreateOne()
		users = append(users, user)
	}
	return users
}

func (UserFactory) CreateOneWithData(data *models.User) *models.User {
	user := models.User{}
	if data.Name != "" {
		user.Name = data.Name
	} else {
		user.Name = Faker.Person().Name()
	}
	if data.Email != "" {
		user.Email = data.Email
	} else {
		user.Email = Faker.Internet().Email()
	}
	if data.Password != "" {
		user.Password = data.Password
	} else {
		user.Password = Faker.Internet().Password()
	}
	if data.Avatar != "" {
		user.Avatar = data.Avatar
	} else {
		user.Avatar = "https://i.pravatar.cc/300"
	}
	if data.Provider != "" {
		user.Provider = data.Provider
	} else {
		user.Provider = "local"
	}
	return &user
}

func (UserFactory) CreateManyWithData(count int, data *models.User) []*models.User {
	var users []*models.User
	user := &models.User{}
	for i := 0; i < count; i++ {
		user = UserFactory{}.CreateOneWithData(data)
		users = append(users, user)
	}
	return users
}
