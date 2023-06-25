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

type Factory interface {
	createOne() interface{}
	createMany(int) []interface{}
	createOneWithData(interface{}) interface{}
	createManyWithData(int, interface{}) []interface{}
}
