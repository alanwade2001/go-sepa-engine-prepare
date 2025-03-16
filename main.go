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
	ReposMgr *repository.Manager
	Iban     *service.Iban
	Listener *q.Listener
}

func NewApp() *App {

	infra := inf.NewInfra()

	Stomp := q.NewStomp()
	Postgres := db.NewPersist()
	ReposMgr := repository.NewManager(Postgres)

	Delivery := service.NewDelivery(Stomp)
	Service := service.NewPreparer(ReposMgr, Delivery)
	Receiver := receiver.NewPaymentGroup(Service)

	Listener := q.Newlistener(Stomp, Receiver)

	app := &App{
		infra:    infra,
		Postgres: Postgres,
		ReposMgr: ReposMgr,

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
