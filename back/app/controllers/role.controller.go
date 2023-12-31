package controllers

import (
	"nartex/ngr-stack/app/models"
)

// We create a struct to hold the handlers
// and we initialize it in the init function
// This way we can use the handlers in the routes
// without contaminating the package namespace with handlers

var RoleController = NewStandardCrudController[models.Role](
	CrudControllerOptions{
		ResourceName:       "Role",
		ResourceSlug:       "role",
		ResourcePluralName: "Roles",
		ResourcePluralSlug: "roles",
	},
)
