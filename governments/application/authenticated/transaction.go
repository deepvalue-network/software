package authenticated

import (
	"strconv"

	"github.com/deepvalue-network/software/governments/domain/governments"
	"github.com/deepvalue-network/software/governments/domain/governments/shareholders/payments"
	"github.com/deepvalue-network/software/governments/domain/governments/shareholders/transfers"
	"github.com/deepvalue-network/software/governments/domain/governments/shareholders/transfers/views"
	identity_payments "github.com/deepvalue-network/software/governments/domain/identities/shareholders/transactions/payments"
	"github.com/deepvalue-network/software/governments/domain/identities/shareholders/transactions/transfers/incomings"
	"github.com/deepvalue-network/software/governments/domain/identities/shareholders/transactions/transfers/outgoings"
	"github.com/deepvalue-network/software/libs/cryptography/pk/signature"
	"github.com/deepvalue-network/software/libs/hash"
	uuid "github.com/satori/go.uuid"
)

type transaction struct {
	identityApp                Identity
	identityPaymentService     identity_payments.Service
	identityPaymentBuilder     identity_payments.PaymentBuilder
	identityOutgoingService    outgoings.Service
	identityOutgoingBuilder    outgoings.OutgoingBuilder
	identityIncomingService    incomings.Service
	identityIncomingBuilder    incomings.IncomingBuilder
	paymentBuilder             payments.Builder
	paymentContentBuilder      payments.ContentBuilder
	transferContentBuilder     transfers.ContentBuilder
	transferBuilder            transfers.Builder
	viewTransferSectionBuilder views.SectionBuilder
	governmentRepository       governments.Repository
	pkFactory                  signature.PrivateKeyFactory
	hashAdapter                hash.Adapter
	amountPubKeysInRing        uint
}

func createTransaction(
	identityApp Identity,
	identityPaymentService identity_payments.Service,
	identityPaymentBuilder identity_payments.PaymentBuilder,
	identityOutgoingService outgoings.Service,
	identityOutgoingBuilder outgoings.OutgoingBuilder,
	identityIncomingService incomings.Service,
	identityIncomingBuilder incomings.IncomingBuilder,
	paymentBuilder payments.Builder,
	paymentContentBuilder payments.ContentBuilder,
	transferContentBuilder transfers.ContentBuilder,
	transferBuilder transfers.Builder,
	viewTransferSectionBuilder views.SectionBuilder,
	governmentRepository governments.Repository,
	pkFactory signature.PrivateKeyFactory,
	hashAdapter hash.Adapter,
	amountPubKeysInRing uint,
) Transaction {
	out := transaction{
		identityApp:                identityApp,
		identityPaymentService:     identityPaymentService,
		identityPaymentBuilder:     identityPaymentBuilder,
		identityOutgoingService:    identityOutgoingService,
		identityOutgoingBuilder:    identityOutgoingBuilder,
		identityIncomingService:    identityIncomingService,
		identityIncomingBuilder:    identityIncomingBuilder,
		paymentBuilder:             paymentBuilder,
		paymentContentBuilder:      paymentContentBuilder,
		transferContentBuilder:     transferContentBuilder,
		transferBuilder:            transferBuilder,
		viewTransferSectionBuilder: viewTransferSectionBuilder,
		governmentRepository:       governmentRepository,
		pkFactory:                  pkFactory,
		hashAdapter:                hashAdapter,
		amountPubKeysInRing:        amountPubKeysInRing,
	}

	return &out
}

// Payment creates a payment
func (app *transaction) Payment(govID *uuid.UUID, amount uint, note string) error {
	gov, err := app.governmentRepository.Retrieve(govID)
	if err != nil {
		return err
	}

	identity, err := app.identityApp.Retrieve()
	if err != nil {
		return err
	}

	shareHolder, err := identity.ShareHolders().Fetch(gov)
	if err != nil {
		return err
	}

	paymentContent, err := app.paymentContentBuilder.Create().WithShareHolder(shareHolder.Public()).WithAmount(amount).Now()
	if err != nil {
		return err
	}

	msg := paymentContent.Hash().String()
	sig, err := shareHolder.SigPK().Sign(msg)
	if err != nil {
		return err
	}

	payment, err := app.paymentBuilder.Create().WithContent(paymentContent).WithSignature(sig).Now()
	if err != nil {
		return err
	}

	identityPayment, err := app.identityPaymentBuilder.Create().WithPayment(payment).WithNote(note).Now()
	if err != nil {
		return err
	}

	return app.identityPaymentService.Insert(identityPayment)
}

// Transfer creates a transfer
func (app *transaction) Transfer(govID *uuid.UUID, amount uint, seed string, to []hash.Hash, note string) error {
	viewTransfer, err := app.View(govID, amount, seed, to)
	if err != nil {
		return err
	}

	outgoing, err := app.identityOutgoingBuilder.Create().WithTransfer(viewTransfer).WithNote(note).Now()
	if err != nil {
		return err
	}

	return app.identityOutgoingService.Insert(outgoing)
}

// View creates a view transfer
func (app *transaction) View(govID *uuid.UUID, amount uint, seed string, to []hash.Hash) (views.Section, error) {
	gov, err := app.governmentRepository.Retrieve(govID)
	if err != nil {
		return nil, err
	}

	identity, err := app.identityApp.Retrieve()
	if err != nil {
		return nil, err
	}

	shareHolder, err := identity.ShareHolders().Fetch(gov)
	if err != nil {
		return nil, err
	}

	origin := shareHolder.Hash()
	seedHash, err := app.hashAdapter.FromBytes([]byte(seed))
	if err != nil {
		return nil, err
	}

	amountHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(seed),
		[]byte(strconv.Itoa(int(amount))),
	})

	if err != nil {
		return nil, err
	}

	pk := shareHolder.SigPK()
	ring, err := newRing(app.pkFactory, pk, int(app.amountPubKeysInRing))
	if err != nil {
		return nil, err
	}

	owners := []hash.Hash{}
	for _, onePubKey := range ring {
		hsh, err := app.hashAdapter.FromBytes([]byte(onePubKey.String()))
		if err != nil {
			return nil, err
		}

		owners = append(owners, *hsh)
	}

	transferContent, err := app.transferContentBuilder.Create().WithOrigin(origin).WithAmount(*amountHash).WithSeed(*seedHash).WithOwner(owners).Now()
	if err != nil {
		return nil, err
	}

	msg := transferContent.Hash().String()
	sig, err := pk.RingSign(msg, ring)
	if err != nil {
		return nil, err
	}

	transfer, err := app.transferBuilder.Create().WithContent(transferContent).WithSignature(sig).Now()
	if err != nil {
		return nil, err
	}

	return app.viewTransferSectionBuilder.Create().WithTransfer(transfer).WithOrigin(shareHolder.Public()).WithSeed(seed).WithAmount(amount).Now()
}

// Receive receives a transfer
func (app *transaction) Receive(view views.Section, pk signature.PrivateKey, note string) error {
	incoming, err := app.identityIncomingBuilder.Create().WithPK(pk).WithTransfer(view).WithNote(note).Now()
	if err != nil {
		return err
	}

	return app.identityIncomingService.Insert(incoming)
}
