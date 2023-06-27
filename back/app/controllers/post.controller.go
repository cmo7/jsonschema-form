package controllers

import (
	"nartex/ngr-stack/app/models"
	"nartex/ngr-stack/i18n"
)

var PostController = NewStandardCrudController[models.Post](
	CrudControllerOptions{
		Locale:             i18n.ES,
		ResourceName:       "Post",
		ResourceSlug:       "post",
		ResourcePluralName: "Posts",
		ResourcePluralSlug: "posts",
	},
)
