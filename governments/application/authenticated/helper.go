package authenticated

import (
	"math/rand"
	"strings"
	"time"

	"github.com/deepvalue-network/software/libs/cryptography/pk/signature"
)

func newRing(pkFactory signature.PrivateKeyFactory, pk signature.PrivateKey, amount int) ([]signature.PublicKey, error) {
	out := []signature.PublicKey{}
	for i := 0; i < amount; i++ {
		genPubKey := pkFactory.Create().PublicKey()
		out = append(out, genPubKey)
	}

	index := genRand(amount - 1)
	first := append(out[:index], pk.PublicKey())
	out = append(first, out[index+1:]...)
	return out, nil
}

func newSeed(length int) string {
	out := []string{}
	characters := "abcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*()_+{}:<>?-=[];',./"
	amount := len(characters)
	for i := 0; i < length; i++ {
		index := genRand(amount - 1)
		out = append(out, string(characters[index]))
	}

	return strings.Join(out, "")
}

func genRand(max int) int {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	return r1.Intn(max)
}
