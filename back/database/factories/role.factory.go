package factories

import (
	"nartex/ngr-stack/app/models"
)

type RoleFactory struct{}

func (RoleFactory) CreateOne() *models.Role {

	role := models.Role{
		Name:               Faker.Lorem().Word(),
		DefaultForNewUsers: false,
	}
	return &role
}

func (RoleFactory) CreateMany(count int) []*models.Role {
	var Roles []*models.Role
	for i := 0; i < count; i++ {
		role := RoleFactory{}.CreateOne()
		Roles = append(Roles, role)
	}
	return Roles
}

func (RoleFactory) CreateOneWithData(data *models.Role) *models.Role {
	role := models.Role{}
	if data.Name != "" {
		role.Name = data.Name
	} else {
		role.Name = Faker.Lorem().Word()
	}

	if data.DefaultForNewUsers {
		role.DefaultForNewUsers = data.DefaultForNewUsers
	} else {
		role.DefaultForNewUsers = false
	}

	return &role
}

func (RoleFactory) CreateManyWithData(count int, data *models.Role) []*models.Role {
	var roles []*models.Role
	role := &models.Role{}
	for i := 0; i < count; i++ {
		role = RoleFactory{}.CreateOneWithData(data)
		roles = append(roles, role)
	}
	return roles
}
