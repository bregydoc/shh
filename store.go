package shh

type PairEvent string

const EventPairExpired PairEvent = "expired"

type Store interface {
	RegisterNewPair(token string, pair *Pair) error
	ObservePublic(token, publicKey string, callback func (e *PairEvent)) error
	GetPair(token, publicKey string) (*Pair, error)
}
