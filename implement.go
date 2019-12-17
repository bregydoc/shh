package shh

import (
	"context"

	"github.com/bregydoc/shh/proto"
)

func (w *Wizard) GeneratePublicKey(c context.Context, req *proto.Claims) (*proto.PublicKey, error) {
	pair, err := w.craftNewDefaultPair()
	if err != nil {
		return nil, err
	}

	token := req.Username + ":" + req.Password
	if err = w.verifyToken(token); err != nil {
		return nil, err
	}

	if err = w.store.RegisterNewPair(token, pair); err != nil {
		return nil, err
	}

	return &proto.PublicKey{
		Pem: pair.PublicKeyPem,
	}, nil
}

func (w *Wizard) UnfoldMessage(c context.Context, req *proto.MessageToUnfold) (*proto.Message, error) {
	token := req.Claims.Username + ":" + req.Claims.Password
	if err := w.verifyToken(token); err != nil {
		return nil, err
	}

	pair, err := w.store.GetPair(token, req.PublicKey)
	if err != nil {
		return nil, err
	}

	message, err := w.decryptCipher(pair.PrivateKeyPem, req.EncodedMessage)
	if err != nil {
		return nil, err
	}

	return &proto.Message{
		Message: message,
	}, nil

}

func (w *Wizard) FoldMessage(c context.Context, req *proto.MessageToFold) (*proto.EncodedMessage, error) {
	encryptedMessage, err := w.encrypt(req.PublicKey, req.Message)
	if err != nil {
		return nil, err
	}

	return &proto.EncodedMessage{
		EncodedMessage: encryptedMessage,
	}, nil
}