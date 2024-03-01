package chaincode

import (
	"github.com/stretchr/testify/require"
	"github.com/the-medium-tech/util/ecdhm/types"

	"sync"
	"testing"
)

var (
	peerExNum = 12
	peerExs   []*KeyBoxEx
)

func init() {
	peerExs = make([]*KeyBoxEx, peerExNum)
	for i := 0; i < peerExNum; i++ {
		peerExs[i] = NewKeyBoxEx()
	}
}

func TestKeyBoxEx(t *testing.T) {
	err := rpc_callPeer_InitEx()
	require.NoError(t, err)

	for _, peer := range peerExs {
		pv := peer.Update(nil)
		toIdx := (peer.idx + 1) % peerExNum
		for {
			t.Logf("toIdx:%v, pvTag: %v", toIdx, pv.PartsTag())
			//t.Logf("toIdx: %v", toIdx)
			pv = peerExs[toIdx].Update(pv)
			if pv == nil {
				t.Logf("pv is nil")
				break
			}
			toIdx = (toIdx + 1) % peerExNum
		}
		t.Logf("=====================================")
	}

	t.Logf("Peer count: %v, Total Round: %v\n", len(peerExs), 0)
	for i, p := range peerExs {
		p.debugPrint()
		if i > 0 {
			require.Equal(t, peerExs[(i-1)%len(peerExs)].Key(), peerExs[i%len(peerExs)].Key())
			require.NotEqual(t, peerExs[(i-1)%len(peerExs)].preLastKey, peerExs[i%len(peerExs)].preLastKey)
		}
	}

}

func TestAsyncKeyBoxEx(t *testing.T) {
	err := rpc_callPeer_InitEx()
	require.NoError(t, err)
	wg := &sync.WaitGroup{}
	for _, peer := range peerExs {
		wg.Add(1)
		pv := peer.Update(nil)
		toIdx := (peer.idx + 1) % peerExNum
		go async_update(pv, toIdx, wg, t)
		t.Logf("=====================================")
	}
	wg.Wait()

	t.Logf("Peer count: %v, Total Round: %v\n", len(peerExs), 0)
	for i, p := range peerExs {
		p.debugPrint()
		if i > 0 {
			require.Equal(t, peerExs[(i-1)%len(peerExs)].Key(), peerExs[i%len(peerExs)].Key())
			require.NotEqual(t, peerExs[(i-1)%len(peerExs)].preLastKey, peerExs[i%len(peerExs)].preLastKey)
		}
	}

}

func rpc_callPeer_InitEx() error {
	for i, p := range peerExs {
		if err := p.Init(i, peerExNum); err != nil {
			return err
		}
	}
	return nil
}

func async_update(pv *types.PointValue, toIdx int, wg *sync.WaitGroup, t *testing.T) {
	defer wg.Done()

	for {
		t.Logf("toIdx:%v, pvTag: %v", toIdx, pv.PartsTag())
		pv = peerExs[toIdx].Update(pv)
		if pv == nil {
			t.Logf("pv is nil")
			break
		}
		toIdx = (toIdx + 1) % peerExNum
	}
}
