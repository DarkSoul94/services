package app

import "github.com/DarkSoul94/services/models"

type Notificator interface {
	Notify(msg models.Notification) error
}
