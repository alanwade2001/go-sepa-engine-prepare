package receiver

import (
	"encoding/json"
	"log"

	"github.com/alanwade2001/go-sepa-engine-data/model"
	"github.com/alanwade2001/go-sepa-engine-prepare/internal/service"
)

type PaymentGroup struct {
	Service *service.Preparer
}

func NewPaymentGroup(Service *service.Preparer) *PaymentGroup {
	i := &PaymentGroup{
		Service: Service,
	}

	return i
}

func (i *PaymentGroup) Process(body []byte) error {
	text := string(body)
	mdl := &model.PaymentGroup{}
	if err := json.Unmarshal([]byte(text), mdl); err != nil {
		log.Println(err)
	} else {

		log.Printf("model:[%v]", mdl.String())
		if err = i.Service.Prepare(mdl); err != nil {
			log.Println(err)
		}
	}

	return nil
}
