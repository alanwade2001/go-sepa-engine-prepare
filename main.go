package main

import (
	db "github.com/alanwade2001/go-sepa-db"
	"github.com/alanwade2001/go-sepa-engine-data/repository"
	"github.com/alanwade2001/go-sepa-engine-prepare/internal/receiver"
	"github.com/alanwade2001/go-sepa-engine-prepare/internal/service"

	inf "github.com/alanwade2001/go-sepa-infra"
	q "github.com/alanwade2001/go-sepa-q"
)

type App struct {
	infra    *inf.Infra
	receiver *receiver.PaymentGroup
	Postgres *db.Persist
	PGRepos  *repository.PaymentGroup
	PmtRepos *repository.Payment
	TxRepos  *repository.CreditTransfer
	Iban     *service.Iban
	Listener *q.Listener
}

func NewApp() *App {

	infra := inf.NewInfra()

	Stomp := q.NewStomp()
	Postgres := db.NewPersist()
	PGRepos := repository.NewPaymentGroup(Postgres)
	PmtRepos := repository.NewPayment(Postgres)
	TxRepos := repository.NewCreditTransfer(Postgres)

	Delivery := service.NewDelivery(Stomp)
	Service := service.NewPreparer(PGRepos, PmtRepos, TxRepos, Delivery)
	Receiver := receiver.NewPaymentGroup(Service)

	Listener := q.Newlistener(Stomp, Receiver)

	app := &App{
		infra:    infra,
		Postgres: Postgres,
		PGRepos:  PGRepos,
		PmtRepos: PmtRepos,
		TxRepos:  TxRepos,

		Listener: Listener,
	}

	return app
}

func (a *App) Run() {
	a.Listener.Listen(service.DEST_ENGINE_PAYMENT_GROUP_INGESTED)
}

func main() {
	app := NewApp()
	app.Run()
}
