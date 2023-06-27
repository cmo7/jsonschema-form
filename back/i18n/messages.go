package i18n

import (
	"fmt"
)

type Messages = map[MessageKey]string

type DictionaryType = map[Locale]Messages

var Dictionary = DictionaryType{
	ES: {
		REQUIRED:             "%s es requerido",
		NOT_FOUND:            "%s no encontrado",
		INVALID:              "%s es inválido",
		ALREADY_EXISTS:       "%s ya existe",
		NOT_FOUND_OR_INVALID: "%s no encontrado o inválido",
		FOUND:                "%s encontrado",
		CREATED:              "%s creado",
		UPDATED:              "%s actualizado",
		NOT_DELETED:          "%s no eliminado",
		DELETED:              "%s eliminado",
		GET_ALL:              "obtener todos los %s",
		GET:                  "obtener %s",
		CREATE:               "crear %s",
		UPDATE:               "actualizar %s",
		DELETE:               "eliminar %s",
	},
	EN: {
		REQUIRED:             "%s is required",
		NOT_FOUND:            "%s not found",
		INVALID:              "%s is invalid",
		ALREADY_EXISTS:       "%s already exists",
		NOT_FOUND_OR_INVALID: "%s not found or invalid",
		FOUND:                "%s found",
		CREATED:              "%s created",
		UPDATED:              "%s updated",
		NOT_DELETED:          "%s not deleted",
		DELETED:              "%s deleted",
		GET_ALL:              "get all %s",
		GET:                  "get %s",
		CREATE:               "create %s",
		UPDATE:               "update %s",
		DELETE:               "delete %s",
	},
}

func Get(locale Locale, key MessageKey) string {
	return Dictionary[locale][key]
}

func GetWithValue(locale Locale, key MessageKey, value string) string {
	return fmt.Sprintf(Dictionary[locale][key], value)
}
