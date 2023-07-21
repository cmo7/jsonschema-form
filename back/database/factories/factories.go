package factories

import "github.com/jaswdr/faker"

// Faker is a global variable that contains a faker instance.
// It is used to generate fake data for the database.
// We use the same faker instance for all the factories, so that
// every invocation of the factory will generate different data.
var Faker = faker.Faker{}

func init() {
	Faker = faker.New()
}

// Factory is an interface for all factories.
// It contains the methods that all factories must implement.
// The methods are used to generate fake data for the database.
type Factory interface {
	// CreateOne creates one entity
	createOne() interface{}
	// CreateMany creates many entities
	createMany(int) []interface{}
	// CreateOneWithData creates one entity with the given data, using fake data for the rest
	createOneWithData(interface{}) interface{}
	// CreateManyWithData creates many entities with the given data, using fake data for the rest
	createManyWithData(int, interface{}) []interface{}
}
