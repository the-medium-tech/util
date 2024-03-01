package types

import (
	"crypto/elliptic"
	"crypto/sha256"
	"math/big"
	"sync"
)

type PointValue struct {
	X     *big.Int
	Y     *big.Int
	Parts *BitTag
	curve elliptic.Curve

	mtx sync.RWMutex
}

func NewPointValue(x, y *big.Int, c elliptic.Curve) *PointValue {
	return &PointValue{
		X:     x,
		Y:     y,
		curve: c,
		Parts: NewBitTag(),
	}
}

func NewPointValueWithTag(x, y *big.Int, c elliptic.Curve, tag uint64) *PointValue {
	return &PointValue{
		X:     x,
		Y:     y,
		curve: c,
		Parts: NewBitTagWithTag(tag),
	}
}

func (p *PointValue) Copy() *PointValue {
	p.mtx.RLock()
	defer p.mtx.RUnlock()

	newPv := NewPointValue(
		new(big.Int).Set(p.X),
		new(big.Int).Set(p.Y),
		p.curve,
	)
	newPv.Parts = p.Parts.Copy()

	return newPv
}

func (p *PointValue) Mul(scalar []byte) {
	p.mtx.Lock()
	defer p.mtx.Unlock()

	p.X, p.Y = p.curve.ScalarMult(p.X, p.Y, scalar)
}

func (p *PointValue) Key() []byte {
	p.mtx.RLock()
	defer p.mtx.RUnlock()

	h := sha256.Sum256(append(p.X.Bytes(), p.Y.Bytes()...))
	return h[:]
}

func (p *PointValue) AddPart(idx int) {
	p.mtx.Lock()
	defer p.mtx.Unlock()

	p.Parts.On(idx)
}

func (p *PointValue) IsPart(id int) bool {
	p.mtx.RLock()
	defer p.mtx.RUnlock()

	return p.Parts.IsOn(id)
}

func (p *PointValue) PartsTag() uint64 {
	p.mtx.RLock()
	defer p.mtx.RUnlock()

	return p.Parts.Tag
}

func (p *PointValue) PartsCount() int {
	p.mtx.RLock()
	defer p.mtx.RUnlock()

	return p.Parts.Count(true)
}

func (p *PointValue) Serialize() []byte {
	p.mtx.RLock()
	defer p.mtx.RUnlock()

	return append(p.X.Bytes(), p.Y.Bytes()...)
}
