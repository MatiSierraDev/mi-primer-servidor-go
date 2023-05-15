package utils

import (
	"github.com/matiraw/mi-primer-servidor-go/db"
	"github.com/matiraw/mi-primer-servidor-go/models"
)

type Options struct {
	IdUser *string
	IdTask *string
}

func CheckUserId(userId string) bool {

	var user models.User

	db.DB.First(&user, userId)

	// if user.ID == 0 {
	// 	return false
	// }

	// return true
	return user.ID != 0
}

func CheckTaskId(taskId string) bool {

	var task models.Task
	db.DB.First(&task, taskId)
	return task.ID != 0

}

// documentar el struct para options por parametro de forma default como nil por el valor cero del pointer
// prueba de valores por parametro por defecto
func CheckId(o Options) bool {

	if o.IdTask != nil {
		var task models.Task
		db.DB.First(&task, o.IdTask)
		return task.ID != 0
	}

	if o.IdUser != nil {
		var user models.User
		db.DB.First(&user, o.IdUser)
		return user.ID != 0
	}

	return true
}
