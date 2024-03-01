package types

import (
	"strconv"
	"sync"
)

type BitTag struct {
	Tag uint64 `json:"tag,string"`
	mtx sync.RWMutex
}

func NewBitTag() *BitTag {
	return &BitTag{}
}

func NewBitTagWithTag(tag uint64) *BitTag {
	return &BitTag{Tag: tag}
}

func (btag *BitTag) On(bitIdx int) {
	btag.mtx.Lock()
	defer btag.mtx.Unlock()

	btag.Tag |= (1 << bitIdx)
}

func (btag *BitTag) Off(bitIdx int) {
	btag.mtx.Lock()
	defer btag.mtx.Unlock()

	if btag.Tag&(1<<bitIdx) != 0 {
		btag.Tag ^= (1 << bitIdx)
	}
}

func (btag *BitTag) IsOn(bitIdx int) bool {
	btag.mtx.RLock()
	defer btag.mtx.RUnlock()

	return btag.Tag&(1<<bitIdx) != 0
}

func (btag *BitTag) Copy() *BitTag {
	btag.mtx.Lock()
	defer btag.mtx.Unlock()

	return &BitTag{
		Tag: btag.Tag,
	}
}

func (btag *BitTag) Count(onoff bool) int {
	btag.mtx.RLock()
	defer btag.mtx.RUnlock()

	tag := btag.Tag
	cnt := 0
	for i := 0; i < strconv.IntSize; i++ {
		if (tag>>i)&(0x1) != 0 {
			cnt++
		}
	}

	if onoff == false {
		cnt = strconv.IntSize - cnt
	}
	return cnt
}
