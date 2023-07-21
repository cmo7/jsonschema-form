package controllers

import (
	"nartex/ngr-stack/app/models"
)

var PostController = NewStandardCrudController[models.Post](
	CrudControllerOptions{
		ResourceName:       "Post",
		ResourceSlug:       "post",
		ResourcePluralName: "Posts",
		ResourcePluralSlug: "posts",
	},
)
