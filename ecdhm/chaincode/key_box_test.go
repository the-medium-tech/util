package chaincode

import (
	"github.com/stretchr/testify/require"
	"github.com/the-medium-tech/util/ecdhm/types"

	"testing"
)

var (
	peerNum = 12
	peers   []*KeyBox
)

func init() {
	peers = make([]*KeyBox, peerNum)
	for i := 0; i < peerNum; i++ {
		peers[i] = NewKeyBox() // i 를 식별정보(bit index) 로 사용.
	}
}

func TestKeyBox(t *testing.T) {
	var rets []*types.PointValue
	var err error

	round := 0
	for {

		if round == 0 {
			rets, err = rpc_callPeer_Init()
			require.NoError(t, err)
		} else {
			reqCnt := len(rets)
			rets, err = rpc_callPeer_Update(rets)
			t.Log("Round", round, "Request", reqCnt, "Response", len(rets))

			require.NoError(t, err)
		}

		if len(rets) == 0 {
			break
		}

		//for _, r := range rets {
		//	t.Logf("\t%x\n", r.Key())
		//}
		//t.Log("---")
		//t.Log(" ")

		round++
	}

	t.Logf("Peer count: %v, Total Round: %v\n", len(peers), round)
	for i, p := range peers {
		p.debugPrint()
		if i > 0 {
			require.Equal(t, peers[(i-1)%len(peers)].Key(), peers[i%len(peers)].key)
			require.NotEqual(t, peers[(i-1)%len(peers)].preLastKey, peers[i%len(peers)].preLastKey)
		}
	}

}

func rpc_callPeer_Init() ([]*types.PointValue, error) {
	var responses []*types.PointValue

	for i, p := range peers {
		if pt, err := p.Init(i); err != nil {
			return nil, err
		} else {
			responses = append(responses, pt)
		}
	}
	return responses, nil
}

func rpc_callPeer_Update(preResults []*types.PointValue) ([]*types.PointValue, error) {
	var results []*types.PointValue
	for peer_idx, peer := range peers {
		var reqs []*types.PointValue
		for _, preRet := range preResults {
			if !preRet.IsPart(peer_idx) {
				reqs = append(reqs, preRet.Copy())
			}
		}

		rets := peer.Update(reqs)
		if rets == nil {
			// this peer finishes the last calculation.
			continue
		}
		results = merge(rets, results)
	}
	return results, nil
}

func merge(from, to []*types.PointValue) []*types.PointValue {
	merged := make([]*types.PointValue, len(to))
	copy(merged, to)

	for _, r1 := range from {

		found := false
		for _, r2 := range merged {
			if r1.PartsTag() == r2.PartsTag() {
				found = true
				break
			}
		}

		if !found {
			merged = append(merged, r1)
		}
	}

	return merged
}
