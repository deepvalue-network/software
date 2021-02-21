package authenticated

import (
	"time"

	"github.com/deepvalue-network/software/governments/domain/governments"
	"github.com/deepvalue-network/software/governments/domain/governments/shareholders/swaps/closes"
	"github.com/deepvalue-network/software/governments/domain/governments/shareholders/swaps/requests"
	"github.com/deepvalue-network/software/governments/domain/governments/shareholders/swaps/trades"
	"github.com/deepvalue-network/software/libs/cryptography/pk/signature"
	"github.com/deepvalue-network/software/libs/hash"
	uuid "github.com/satori/go.uuid"
)

type swap struct {
	identityApp           Identity
	transactionApp        Transaction
	governmentRepository  governments.Repository
	requestRepository     requests.Repository
	requestService        requests.Service
	requestContentBuilder requests.ContentBuilder
	requestBuilder        requests.Builder
	tradeRepository       trades.Repository
	tradeService          trades.Service
	tradeContentBuilder   trades.ContentBuilder
	tradeBuilder          trades.Builder
	closeService          closes.Service
	closeContentBuilder   closes.ContentBuilder
	closeBuilder          closes.Builder
	pkFactory             signature.PrivateKeyFactory
	amountPubKeysInRing   uint
	seedLength            uint
}

func createSwap(
	identityApp Identity,
	transactionApp Transaction,
	governmentRepository governments.Repository,
	requestRepository requests.Repository,
	requestService requests.Service,
	requestContentBuilder requests.ContentBuilder,
	requestBuilder requests.Builder,
	tradeRepository trades.Repository,
	tradeService trades.Service,
	tradeContentBuilder trades.ContentBuilder,
	tradeBuilder trades.Builder,
	closeService closes.Service,
	closeContentBuilder closes.ContentBuilder,
	closeBuilder closes.Builder,
	pkFactory signature.PrivateKeyFactory,
	amountPubKeysInRing uint,
	seedLength uint,
) Swap {
	out := swap{
		identityApp:           identityApp,
		transactionApp:        transactionApp,
		governmentRepository:  governmentRepository,
		requestRepository:     requestRepository,
		requestService:        requestService,
		requestContentBuilder: requestContentBuilder,
		requestBuilder:        requestBuilder,
		tradeRepository:       tradeRepository,
		tradeService:          tradeService,
		tradeContentBuilder:   tradeContentBuilder,
		tradeBuilder:          tradeBuilder,
		closeService:          closeService,
		closeContentBuilder:   closeContentBuilder,
		closeBuilder:          closeBuilder,
		pkFactory:             pkFactory,
		amountPubKeysInRing:   amountPubKeysInRing,
		seedLength:            seedLength,
	}

	return &out
}

// Request creates a swap request
func (app *swap) Request(fromGovID *uuid.UUID, amount uint, seed string, to []hash.Hash, forGovID *uuid.UUID, expireOn time.Time) error {
	fromGov, err := app.governmentRepository.Retrieve(fromGovID)
	if err != nil {
		return err
	}

	identity, err := app.identityApp.Retrieve()
	if err != nil {
		return err
	}

	shareHolder, err := identity.ShareHolders().Fetch(fromGov)
	if err != nil {
		return err
	}

	stake, err := app.transactionApp.View(fromGovID, amount, seed)
	if err != nil {
		return err
	}

	forGov, err := app.governmentRepository.Retrieve(forGovID)
	if err != nil {
		return err
	}

	content, err := app.requestContentBuilder.Create().From(fromGov).WithStake(stake).For(forGov).To(to).WithAmount(amount).ExpiresOn(expireOn).Now()
	if err != nil {
		return err
	}

	msg := content.Hash().String()
	pk := shareHolder.SigPK()
	ring, err := newRing(app.pkFactory, pk, int(app.amountPubKeysInRing))
	if err != nil {
		return err
	}

	sig, err := shareHolder.SigPK().RingSign(msg, ring)
	if err != nil {
		return err
	}

	request, err := app.requestBuilder.Create().WithContent(content).WithSignature(sig).Now()
	if err != nil {
		return err
	}

	return app.requestService.Insert(request)
}

// Trade executes a trade on a request
func (app *swap) Trade(requestHash hash.Hash, expireOn time.Time) error {
	request, err := app.requestRepository.Retrieve(requestHash)
	if err != nil {
		return err
	}

	reqContent := request.Content()
	govID := reqContent.For().ID()
	amount := reqContent.Amount()
	seed := newSeed(int(app.seedLength))
	viewSection, err := app.transactionApp.View(govID, amount, seed)
	if err != nil {
		return err
	}

	newOwner := reqContent.To()
	viewTransfer, err := app.transactionApp.ViewTransfer(viewSection, govID, newOwner)
	if err != nil {
		return err
	}

	forGov, err := app.governmentRepository.Retrieve(govID)
	if err != nil {
		return err
	}

	identity, err := app.identityApp.Retrieve()
	if err != nil {
		return err
	}

	shareHolder, err := identity.ShareHolders().Fetch(forGov)
	if err != nil {
		return err
	}

	keys := shareHolder.Public().Keys()
	content, err := app.tradeContentBuilder.Create().WithRequest(request).WithTransfer(viewTransfer).To(keys).ExpiresOn(expireOn).Now()
	if err != nil {
		return err
	}

	msg := content.Hash().String()
	pk := shareHolder.SigPK()
	ring, err := newRing(app.pkFactory, pk, int(app.amountPubKeysInRing))
	if err != nil {
		return err
	}

	sig, err := shareHolder.SigPK().RingSign(msg, ring)
	if err != nil {
		return err
	}

	trade, err := app.tradeBuilder.Create().WithContent(content).WithSignature(sig).Now()
	if err != nil {
		return err
	}

	return app.tradeService.Insert(trade)
}

// Close closes a trade
func (app *swap) Close(tradeHash hash.Hash) error {
	trade, err := app.tradeRepository.Retrieve(tradeHash)
	if err != nil {
		return err
	}

	content, err := app.closeContentBuilder.Create().WithTrade(trade).Now()
	if err != nil {
		return err
	}

	identity, err := app.identityApp.Retrieve()
	if err != nil {
		return err
	}

	gov := trade.Content().Request().Content().For()
	shareHolder, err := identity.ShareHolders().Fetch(gov)
	if err != nil {
		return err
	}

	msg := content.Hash().String()
	pk := shareHolder.SigPK()
	ring, err := newRing(app.pkFactory, pk, int(app.amountPubKeysInRing))
	if err != nil {
		return err
	}

	sig, err := shareHolder.SigPK().RingSign(msg, ring)
	if err != nil {
		return err
	}

	close, err := app.closeBuilder.Create().WithContent(content).WithSignature(sig).Now()
	if err != nil {
		return err
	}

	return app.closeService.Insert(close)
}
