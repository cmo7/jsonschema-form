package data

import (
	"fmt"
	"nartex/ngr-stack/config"
	"nartex/ngr-stack/database"
)

func GetPage[M any](model M, page int, size int) []M {

	// In case the page or size are invalid, set them to the default values
	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 1
	}
	if size > config.Pagination.MaxPageSize {
		size = config.Pagination.DefaultPageSize
	}

	offset := (page - 1) * size
	limit := size

	fmt.Println("limit: ", limit, "offset: ", offset)

	result := []M{}
	database.DB.Limit(limit).Offset(offset).Find(&result)
	return result
}

func Count[R any](model R) int {
	var total int64
	database.DB.Model(&model).Count(&total)
	total32 := int(total)
	return total32
}
