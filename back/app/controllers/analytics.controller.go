package controllers

import (
	"nartex/ngr-stack/app/models"
)

var AnalyticsController = NewStandardCrudController[models.Analytic](
	CrudControllerOptions{
		ResourceName:       "Analytic",
		ResourceSlug:       "analytic",
		ResourcePluralName: "Analytics",
		ResourcePluralSlug: "analytics",
	},
)
