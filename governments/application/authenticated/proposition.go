package authenticated

import (
	"github.com/deepvalue-network/software/governments/domain/propositions"
	"github.com/deepvalue-network/software/governments/domain/propositions/votes"
	"github.com/deepvalue-network/software/libs/cryptography/pk/signature"
	"github.com/deepvalue-network/software/libs/hash"
)

type proposition struct {
	identityApp         Identity
	repository          propositions.Repository
	service             propositions.Service
	builder             propositions.Builder
	voteService         votes.Service
	voteBuilder         votes.Builder
	voteContentBuilder  votes.ContentBuilder
	pkFactory           signature.PrivateKeyFactory
	amountPubKeysInRing uint
}

func createProposition(
	identityApp Identity,
	repository propositions.Repository,
	service propositions.Service,
	builder propositions.Builder,
	voteService votes.Service,
	voteBuilder votes.Builder,
	voteContentBuilder votes.ContentBuilder,
	pkFactory signature.PrivateKeyFactory,
	amountPubKeysInRing uint,
) Proposition {
	out := proposition{
		identityApp:         identityApp,
		repository:          repository,
		service:             service,
		builder:             builder,
		voteService:         voteService,
		voteBuilder:         voteBuilder,
		voteContentBuilder:  voteContentBuilder,
		pkFactory:           pkFactory,
		amountPubKeysInRing: amountPubKeysInRing,
	}

	return &out
}

// New creates a new proposition
func (app *proposition) New(content propositions.Content, sigs []signature.RingSignature) error {
	proposition, err := app.builder.Create().WithContent(content).WithSignatures(sigs).Now()
	if err != nil {
		return err
	}

	return app.service.Insert(proposition)
}

// Approve approves a proposition
func (app *proposition) Approve(propositionHash hash.Hash) error {
	return app.vote(propositionHash, true, false, false)
}

// Cancel cancels a proposition
func (app *proposition) Cancel(propositionHash hash.Hash) error {
	return app.vote(propositionHash, false, true, false)
}

// Disapprove disapprove a proposition
func (app *proposition) Disapprove(propositionHash hash.Hash) error {
	return app.vote(propositionHash, false, false, true)
}

func (app *proposition) vote(propositionHash hash.Hash, isApprove bool, isCancel bool, isDisapprove bool) error {
	identity, err := app.identityApp.Retrieve()
	if err != nil {
		return err
	}

	proposition, err := app.repository.Retrieve(propositionHash)
	if err != nil {
		return err
	}

	contentBuilder := app.voteContentBuilder.Create().WithProposition(proposition)
	if isApprove {
		contentBuilder.IsApproved()
	}

	if isCancel {
		contentBuilder.IsCancel()
	}

	if isDisapprove {
		contentBuilder.IsDisapproved()
	}

	content, err := contentBuilder.Now()
	if err != nil {
		return err
	}

	gov := proposition.Content().Government()
	shareHolder, err := identity.ShareHolders().Fetch(gov)
	if err != nil {
		return err
	}

	pk := shareHolder.SigPK()
	ring, err := newRing(app.pkFactory, pk, int(app.amountPubKeysInRing))
	if err != nil {
		return err
	}

	msg := content.Hash().String()
	sig, err := pk.RingSign(msg, ring)
	if err != nil {
		return err
	}

	vote, err := app.voteBuilder.Create().WithContent(content).WithSignature(sig).Now()
	if err != nil {
		return err
	}

	return app.voteService.Insert(vote)
}
