package service

import (
	"encoding/json"

	"github.com/alanwade2001/go-sepa-engine-data/model"
	q "github.com/alanwade2001/go-sepa-q"
)

var DEST_ENGINE_PAYMENT_GROUP_INGESTED string = "queue:portal.payment_group.ingested"
var DEST_ENGINE_PAYMENT_PREPARED string = "queue:portal.payment.prepared"

type Delivery struct {
	Stomp *q.Stomp
}

func NewDelivery(Stomp *q.Stomp) *Delivery {
	i := &Delivery{
		Stomp: Stomp,
	}

	return i
}

func (d *Delivery) PaymentPrepared(p *model.Payment) error {
	if bytes, err := json.Marshal(p); err != nil {
		return err
	} else {
		if err = d.Stomp.SendMessage(DEST_ENGINE_PAYMENT_PREPARED, bytes); err != nil {
			return err
		}
	}

	return nil
}
