package chaincode

import (
	"github.com/the-medium-tech/util/ecdhm/types"

	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"sync"
)

type KeyBox struct {
	idx int
	prv *ecdsa.PrivateKey
	key []byte

	mtx sync.RWMutex

	// only for debug
	preLastKey   []byte
	callCnt      int
	calculateCnt int
}

func NewKeyBox() *KeyBox {
	return &KeyBox{}
}

func (k *KeyBox) Init(idx int) (*types.PointValue, error) {
	_prv, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, err
	}

	k.mtx.Lock()
	defer k.mtx.Unlock()

	k.callCnt++

	k.idx = idx
	k.prv = _prv
	k.key = nil

	pt := types.NewPointValue(k.prv.PublicKey.X, k.prv.PublicKey.Y, k.prv.PublicKey.Curve)
	pt.AddPart(k.idx)
	k.calculateCnt++

	return pt, nil
}

func (k *KeyBox) Update(ptVals []*types.PointValue) []*types.PointValue {
	k.mtx.Lock()
	defer k.mtx.Unlock()

	k.callCnt++

	if len(ptVals) == 1 {
		k.preLastKey = ptVals[0].Key()
	}

	for _, p := range ptVals {
		// it's for debug
		if p.IsPart(k.idx) {
			panic("Receive point value that is already calculated by me")
		}
		p.Mul(k.prv.D.Bytes())
		p.AddPart(k.idx)
		k.calculateCnt++
	}

	if len(ptVals) == 1 {
		k.key = ptVals[0].Key()
		return nil // finished. The last result MUST NOT BE revealed
	}
	return ptVals
}

func (k *KeyBox) Key() []byte {
	k.mtx.RLock()
	defer k.mtx.RUnlock()

	return k.key
}

func (k *KeyBox) debugPrint() {
	fmt.Printf("KeyBox[%2d] : shared key: %x <- %x, call:%d, calculate:%d, prv: %x, pub: (%x,%x)\n", k.idx, k.key, k.preLastKey, k.callCnt, k.calculateCnt, k.prv.D.Bytes(), k.prv.PublicKey.X.Bytes(), k.prv.PublicKey.Y.Bytes())
}
