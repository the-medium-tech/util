package chaincode

import (
	"github.com/the-medium-tech/util/ecdhm/types"

	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"sync"
)

type KeyBoxEx struct {
	idx     int
	prv     *ecdsa.PrivateKey
	key     []byte
	partCnt int

	mtx sync.RWMutex

	// only for debug
	preLastKey   []byte
	callCnt      int
	calculateCnt int
}

func NewKeyBoxEx() *KeyBoxEx {
	return &KeyBoxEx{}
}

func (k *KeyBoxEx) NewGenerateKey() (*ecdsa.PrivateKey, error) {
	_prv, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, err
	}

	return _prv, nil
}

func (k *KeyBoxEx) Init(idx, partCnt int) error {
	_prv, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return err
	}

	k.mtx.Lock()
	defer k.mtx.Unlock()

	k.callCnt++

	k.idx = idx
	k.prv = _prv
	k.key = nil
	k.partCnt = partCnt

	return nil
}

func (k *KeyBoxEx) InitWithKey(idx, partCnt int, prv *ecdsa.PrivateKey) error {
	k.mtx.Lock()
	defer k.mtx.Unlock()

	k.callCnt++

	k.idx = idx
	k.prv = prv
	k.key = nil
	k.partCnt = partCnt

	return nil
}

func (k *KeyBoxEx) Update(pv *types.PointValue) *types.PointValue {
	k.mtx.Lock()
	defer k.mtx.Unlock()

	k.callCnt++

	if pv == nil {
		// return my public key (d x G)
		pv = types.NewPointValue(k.prv.PublicKey.X, k.prv.PublicKey.Y, k.prv.PublicKey.Curve)
		pv.AddPart(k.idx)
		k.calculateCnt++
		return pv
	}

	// it's for debug
	if pv.PartsCount() == k.partCnt-1 {
		k.preLastKey = pv.Key()
	}

	if pv.IsPart(k.idx) {
		panic("Receive point value that is already calculated by me")
	}

	pv.Mul(k.prv.D.Bytes())
	pv.AddPart(k.idx)
	k.calculateCnt++

	if pv.PartsCount() == k.partCnt {
		// finished.
		k.key = pv.Key()
		return nil
	}

	return pv
}

func (k *KeyBoxEx) Key() []byte {
	k.mtx.RLock()
	defer k.mtx.RUnlock()

	return k.key
}

func (k *KeyBoxEx) Index() int {
	k.mtx.RLock()
	defer k.mtx.RUnlock()

	return k.idx
}

func (k *KeyBoxEx) debugPrint() {
	fmt.Printf("KeyBox[%2d] : shared key: %x <- %x, call:%d, calculate:%d, prv: %x, pub: (%x,%x)\n", k.idx, k.key, k.preLastKey, k.callCnt, k.calculateCnt, k.prv.D.Bytes(), k.prv.PublicKey.X.Bytes(), k.prv.PublicKey.Y.Bytes())
}
