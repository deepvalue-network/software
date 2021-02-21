package authenticated

import (
	"math/rand"
	"time"

	"github.com/deepvalue-network/software/libs/cryptography/pk/signature"
)

func newRing(pkFactory signature.PrivateKeyFactory, pk signature.PrivateKey, amount int) ([]signature.PublicKey, error) {
	out := []signature.PublicKey{}
	for i := 0; i < amount; i++ {
		genPubKey := pkFactory.Create().PublicKey()
		out = append(out, genPubKey)
	}

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	index := r1.Intn(amount - 1)
	first := append(out[:index], pk.PublicKey())
	out = append(first, out[index+1:]...)
	return out, nil
}
