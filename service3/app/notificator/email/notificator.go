package email

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/DarkSoul94/services/models"
	"github.com/DarkSoul94/services/service3/app"
)

type emailNotifycator struct{}

func NewEmailNotifycator() app.Notificator {
	return &emailNotifycator{}
}

func (n *emailNotifycator) Notify(msg models.Notification) error {
	time.Sleep(1 * time.Second)

	num, err := rand.Int(rand.Reader, big.NewInt(10))
	if err != nil {
		return err
	}

	if num.Int64() > 1 {
		return nil
	} else {
		return errors.New(fmt.Sprintf("Failed send notification: '%s' to %s", msg.Text, msg.Email))
	}

}
