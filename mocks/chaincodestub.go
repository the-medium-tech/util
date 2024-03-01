package mocks

import (
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-protos-go/peer"

	"sync"
)

type ChaincodeStub struct {
	CreateCompositeKeyStub        func(string, []string) (string, error)
	createCompositeKeyMutex       sync.RWMutex
	createCompositeKeyArgsForCall []struct {
		arg1 string
		arg2 []string
	}
	createCompositeKeyReturns struct {
		result1 string
		result2 error
	}
	createCompositeKeyReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	DelPrivateDataStub        func(string, string) error
	delPrivateDataMutex       sync.RWMutex
	delPrivateDataArgsForCall []struct {
		arg1 string
		arg2 string
	}
	delPrivateDataReturns struct {
		result1 error
	}
	delPrivateDataReturnsOnCall map[int]struct {
		result1 error
	}
	DelStateStub        func(string) error
	delStateMutex       sync.RWMutex
	delStateArgsForCall []struct {
		arg1 string
	}
	delStateReturns struct {
		result1 error
	}
	delStateReturnsOnCall map[int]struct {
		result1 error
	}
	GetArgsStub        func() [][]byte
	getArgsMutex       sync.RWMutex
	getArgsArgsForCall []struct {
	}
	getArgsReturns struct {
		result1 [][]byte
	}
	getArgsReturnsOnCall map[int]struct {
		result1 [][]byte
	}
	GetArgsSliceStub        func() ([]byte, error)
	getArgsSliceMutex       sync.RWMutex
	getArgsSliceArgsForCall []struct {
	}
	getArgsSliceReturns struct {
		result1 []byte
		result2 error
	}
	getArgsSliceReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	GetBindingStub        func() ([]byte, error)
	getBindingMutex       sync.RWMutex
	getBindingArgsForCall []struct {
	}
	getBindingReturns struct {
		result1 []byte
		result2 error
	}
	getBindingReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	GetChannelIDStub        func() string
	getChannelIDMutex       sync.RWMutex
	getChannelIDArgsForCall []struct {
	}
	getChannelIDReturns struct {
		result1 string
	}
	getChannelIDReturnsOnCall map[int]struct {
		result1 string
	}
	GetCreatorStub        func() ([]byte, error)
	getCreatorMutex       sync.RWMutex
	getCreatorArgsForCall []struct {
	}
	getCreatorReturns struct {
		result1 []byte
		result2 error
	}
	getCreatorReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	GetDecorationsStub        func() map[string][]byte
	getDecorationsMutex       sync.RWMutex
	getDecorationsArgsForCall []struct {
	}
	getDecorationsReturns struct {
		result1 map[string][]byte
	}
	getDecorationsReturnsOnCall map[int]struct {
		result1 map[string][]byte
	}
	GetFunctionAndParametersStub        func() (string, []string)
	getFunctionAndParametersMutex       sync.RWMutex
	getFunctionAndParametersArgsForCall []struct {
	}
	getFunctionAndParametersReturns struct {
		result1 string
		result2 []string
	}
	getFunctionAndParametersReturnsOnCall map[int]struct {
		result1 string
		result2 []string
	}
	GetHistoryForKeyStub        func(string) (shim.HistoryQueryIteratorInterface, error)
	getHistoryForKeyMutex       sync.RWMutex
	getHistoryForKeyArgsForCall []struct {
		arg1 string
	}
	getHistoryForKeyReturns struct {
		result1 shim.HistoryQueryIteratorInterface
		result2 error
	}
	getHistoryForKeyReturnsOnCall map[int]struct {
		result1 shim.HistoryQueryIteratorInterface
		result2 error
	}
	GetPrivateDataStub        func(string, string) ([]byte, error)
	getPrivateDataMutex       sync.RWMutex
	getPrivateDataArgsForCall []struct {
		arg1 string
		arg2 string
	}
	getPrivateDataReturns struct {
		result1 []byte
		result2 error
	}
	getPrivateDataReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	GetPrivateDataByPartialCompositeKeyStub        func(string, string, []string) (shim.StateQueryIteratorInterface, error)
	getPrivateDataByPartialCompositeKeyMutex       sync.RWMutex
	getPrivateDataByPartialCompositeKeyArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 []string
	}
	getPrivateDataByPartialCompositeKeyReturns struct {
		result1 shim.StateQueryIteratorInterface
		result2 error
	}
	getPrivateDataByPartialCompositeKeyReturnsOnCall map[int]struct {
		result1 shim.StateQueryIteratorInterface
		result2 error
	}
	GetPrivateDataByRangeStub        func(string, string, string) (shim.StateQueryIteratorInterface, error)
	getPrivateDataByRangeMutex       sync.RWMutex
	getPrivateDataByRangeArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
	}
	getPrivateDataByRangeReturns struct {
		result1 shim.StateQueryIteratorInterface
		result2 error
	}
	getPrivateDataByRangeReturnsOnCall map[int]struct {
		result1 shim.StateQueryIteratorInterface
		result2 error
	}
	GetPrivateDataHashStub        func(string, string) ([]byte, error)
	getPrivateDataHashMutex       sync.RWMutex
	getPrivateDataHashArgsForCall []struct {
		arg1 string
		arg2 string
	}
	getPrivateDataHashReturns struct {
		result1 []byte
		result2 error
	}
	getPrivateDataHashReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	GetPrivateDataQueryResultStub        func(string, string) (shim.StateQueryIteratorInterface, error)
	getPrivateDataQueryResultMutex       sync.RWMutex
	getPrivateDataQueryResultArgsForCall []struct {
		arg1 string
		arg2 string
	}
	getPrivateDataQueryResultReturns struct {
		result1 shim.StateQueryIteratorInterface
		result2 error
	}
	getPrivateDataQueryResultReturnsOnCall map[int]struct {
		result1 shim.StateQueryIteratorInterface
		result2 error
	}
	GetPrivateDataValidationParameterStub        func(string, string) ([]byte, error)
	getPrivateDataValidationParameterMutex       sync.RWMutex
	getPrivateDataValidationParameterArgsForCall []struct {
		arg1 string
		arg2 string
	}
	getPrivateDataValidationParameterReturns struct {
		result1 []byte
		result2 error
	}
	getPrivateDataValidationParameterReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	GetQueryResultStub        func(string) (shim.StateQueryIteratorInterface, error)
	getQueryResultMutex       sync.RWMutex
	getQueryResultArgsForCall []struct {
		arg1 string
	}
	getQueryResultReturns struct {
		result1 shim.StateQueryIteratorInterface
		result2 error
	}
	getQueryResultReturnsOnCall map[int]struct {
		result1 shim.StateQueryIteratorInterface
		result2 error
	}
	GetQueryResultWithPaginationStub        func(string, int32, string) (shim.StateQueryIteratorInterface, *peer.QueryResponseMetadata, error)
	getQueryResultWithPaginationMutex       sync.RWMutex
	getQueryResultWithPaginationArgsForCall []struct {
		arg1 string
		arg2 int32
		arg3 string
	}
	getQueryResultWithPaginationReturns struct {
		result1 shim.StateQueryIteratorInterface
		result2 *peer.QueryResponseMetadata
		result3 error
	}
	getQueryResultWithPaginationReturnsOnCall map[int]struct {
		result1 shim.StateQueryIteratorInterface
		result2 *peer.QueryResponseMetadata
		result3 error
	}
	GetSignedProposalStub        func() (*peer.SignedProposal, error)
	getSignedProposalMutex       sync.RWMutex
	getSignedProposalArgsForCall []struct {
	}
	getSignedProposalReturns struct {
		result1 *peer.SignedProposal
		result2 error
	}
	getSignedProposalReturnsOnCall map[int]struct {
		result1 *peer.SignedProposal
		result2 error
	}
	GetStateStub        func(string) ([]byte, error)
	getStateMutex       sync.RWMutex
	getStateArgsForCall []struct {
		arg1 string
	}
	getStateReturns struct {
		result1 []byte
		result2 error
	}
	getStateReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	GetStateByPartialCompositeKeyStub        func(string, []string) (shim.StateQueryIteratorInterface, error)
	getStateByPartialCompositeKeyMutex       sync.RWMutex
	getStateByPartialCompositeKeyArgsForCall []struct {
		arg1 string
		arg2 []string
	}
	getStateByPartialCompositeKeyReturns struct {
		result1 shim.StateQueryIteratorInterface
		result2 error
	}
	getStateByPartialCompositeKeyReturnsOnCall map[int]struct {
		result1 shim.StateQueryIteratorInterface
		result2 error
	}
	GetStateByPartialCompositeKeyWithPaginationStub        func(string, []string, int32, string) (shim.StateQueryIteratorInterface, *peer.QueryResponseMetadata, error)
	getStateByPartialCompositeKeyWithPaginationMutex       sync.RWMutex
	getStateByPartialCompositeKeyWithPaginationArgsForCall []struct {
		arg1 string
		arg2 []string
		arg3 int32
		arg4 string
	}
	getStateByPartialCompositeKeyWithPaginationReturns struct {
		result1 shim.StateQueryIteratorInterface
		result2 *peer.QueryResponseMetadata
		result3 error
	}
	getStateByPartialCompositeKeyWithPaginationReturnsOnCall map[int]struct {
		result1 shim.StateQueryIteratorInterface
		result2 *peer.QueryResponseMetadata
		result3 error
	}
	GetStateByRangeStub        func(string, string) (shim.StateQueryIteratorInterface, error)
	getStateByRangeMutex       sync.RWMutex
	getStateByRangeArgsForCall []struct {
		arg1 string
		arg2 string
	}
	getStateByRangeReturns struct {
		result1 shim.StateQueryIteratorInterface
		result2 error
	}
	getStateByRangeReturnsOnCall map[int]struct {
		result1 shim.StateQueryIteratorInterface
		result2 error
	}
	GetStateByRangeWithPaginationStub        func(string, string, int32, string) (shim.StateQueryIteratorInterface, *peer.QueryResponseMetadata, error)
	getStateByRangeWithPaginationMutex       sync.RWMutex
	getStateByRangeWithPaginationArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 int32
		arg4 string
	}
	getStateByRangeWithPaginationReturns struct {
		result1 shim.StateQueryIteratorInterface
		result2 *peer.QueryResponseMetadata
		result3 error
	}
	getStateByRangeWithPaginationReturnsOnCall map[int]struct {
		result1 shim.StateQueryIteratorInterface
		result2 *peer.QueryResponseMetadata
		result3 error
	}
	GetStateValidationParameterStub        func(string) ([]byte, error)
	getStateValidationParameterMutex       sync.RWMutex
	getStateValidationParameterArgsForCall []struct {
		arg1 string
	}
	getStateValidationParameterReturns struct {
		result1 []byte
		result2 error
	}
	getStateValidationParameterReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	GetStringArgsStub        func() []string
	getStringArgsMutex       sync.RWMutex
	getStringArgsArgsForCall []struct {
	}
	getStringArgsReturns struct {
		result1 []string
	}
	getStringArgsReturnsOnCall map[int]struct {
		result1 []string
	}
	GetTransientStub        func() (map[string][]byte, error)
	getTransientMutex       sync.RWMutex
	getTransientArgsForCall []struct {
	}
	getTransientReturns struct {
		result1 map[string][]byte
		result2 error
	}
	getTransientReturnsOnCall map[int]struct {
		result1 map[string][]byte
		result2 error
	}
	GetTxIDStub        func() string
	getTxIDMutex       sync.RWMutex
	getTxIDArgsForCall []struct {
	}
	getTxIDReturns struct {
		result1 string
	}
	getTxIDReturnsOnCall map[int]struct {
		result1 string
	}
	GetTxTimestampStub        func() (*timestamp.Timestamp, error)
	getTxTimestampMutex       sync.RWMutex
	getTxTimestampArgsForCall []struct {
	}
	getTxTimestampReturns struct {
		result1 *timestamp.Timestamp
		result2 error
	}
	getTxTimestampReturnsOnCall map[int]struct {
		result1 *timestamp.Timestamp
		result2 error
	}
	InvokeChaincodeStub        func(string, [][]byte, string) peer.Response
	invokeChaincodeMutex       sync.RWMutex
	invokeChaincodeArgsForCall []struct {
		arg1 string
		arg2 [][]byte
		arg3 string
	}
	invokeChaincodeReturns struct {
		result1 peer.Response
	}
	invokeChaincodeReturnsOnCall map[int]struct {
		result1 peer.Response
	}
	PurgePrivateDataStub        func(string, string) error
	purgePrivateDataMutex       sync.RWMutex
	purgePrivateDataArgsForCall []struct {
		arg1 string
		arg2 string
	}
	purgePrivateDataReturns struct {
		result1 error
	}
	purgePrivateDataReturnsOnCall map[int]struct {
		result1 error
	}
	PutPrivateDataStub        func(string, string, []byte) error
	putPrivateDataMutex       sync.RWMutex
	putPrivateDataArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 []byte
	}
	putPrivateDataReturns struct {
		result1 error
	}
	putPrivateDataReturnsOnCall map[int]struct {
		result1 error
	}
	PutStateStub        func(string, []byte) error
	putStateMutex       sync.RWMutex
	putStateArgsForCall []struct {
		arg1 string
		arg2 []byte
	}
	putStateReturns struct {
		result1 error
	}
	putStateReturnsOnCall map[int]struct {
		result1 error
	}
	SetEventStub        func(string, []byte) error
	setEventMutex       sync.RWMutex
	setEventArgsForCall []struct {
		arg1 string
		arg2 []byte
	}
	setEventReturns struct {
		result1 error
	}
	setEventReturnsOnCall map[int]struct {
		result1 error
	}
	SetPrivateDataValidationParameterStub        func(string, string, []byte) error
	setPrivateDataValidationParameterMutex       sync.RWMutex
	setPrivateDataValidationParameterArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 []byte
	}
	setPrivateDataValidationParameterReturns struct {
		result1 error
	}
	setPrivateDataValidationParameterReturnsOnCall map[int]struct {
		result1 error
	}
	SetStateValidationParameterStub        func(string, []byte) error
	setStateValidationParameterMutex       sync.RWMutex
	setStateValidationParameterArgsForCall []struct {
		arg1 string
		arg2 []byte
	}
	setStateValidationParameterReturns struct {
		result1 error
	}
	setStateValidationParameterReturnsOnCall map[int]struct {
		result1 error
	}
	SplitCompositeKeyStub        func(string) (string, []string, error)
	splitCompositeKeyMutex       sync.RWMutex
	splitCompositeKeyArgsForCall []struct {
		arg1 string
	}
	splitCompositeKeyReturns struct {
		result1 string
		result2 []string
		result3 error
	}
	splitCompositeKeyReturnsOnCall map[int]struct {
		result1 string
		result2 []string
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *ChaincodeStub) CreateCompositeKey(arg1 string, arg2 []string) (string, error) {
	var arg2Copy []string
	if arg2 != nil {
		arg2Copy = make([]string, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.createCompositeKeyMutex.Lock()
	ret, specificReturn := fake.createCompositeKeyReturnsOnCall[len(fake.createCompositeKeyArgsForCall)]
	fake.createCompositeKeyArgsForCall = append(fake.createCompositeKeyArgsForCall, struct {
		arg1 string
		arg2 []string
	}{arg1, arg2Copy})
	fake.recordInvocation("CreateCompositeKey", []interface{}{arg1, arg2Copy})
	fake.createCompositeKeyMutex.Unlock()
	if fake.CreateCompositeKeyStub != nil {
		return fake.CreateCompositeKeyStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.createCompositeKeyReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *ChaincodeStub) CreateCompositeKeyCallCount() int {
	fake.createCompositeKeyMutex.RLock()
	defer fake.createCompositeKeyMutex.RUnlock()
	return len(fake.createCompositeKeyArgsForCall)
}

func (fake *ChaincodeStub) CreateCompositeKeyCalls(stub func(string, []string) (string, error)) {
	fake.createCompositeKeyMutex.Lock()
	defer fake.createCompositeKeyMutex.Unlock()
	fake.CreateCompositeKeyStub = stub
}

func (fake *ChaincodeStub) CreateCompositeKeyArgsForCall(i int) (string, []string) {
	fake.createCompositeKeyMutex.RLock()
	defer fake.createCompositeKeyMutex.RUnlock()
	argsForCall := fake.createCompositeKeyArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *ChaincodeStub) CreateCompositeKeyReturns(result1 string, result2 error) {
	fake.createCompositeKeyMutex.Lock()
	defer fake.createCompositeKeyMutex.Unlock()
	fake.CreateCompositeKeyStub = nil
	fake.createCompositeKeyReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *ChaincodeStub) CreateCompositeKeyReturnsOnCall(i int, result1 string, result2 error) {
	fake.createCompositeKeyMutex.Lock()
	defer fake.createCompositeKeyMutex.Unlock()
	fake.CreateCompositeKeyStub = nil
	if fake.createCompositeKeyReturnsOnCall == nil {
		fake.createCompositeKeyReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.createCompositeKeyReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *ChaincodeStub) DelPrivateData(arg1 string, arg2 string) error {
	fake.delPrivateDataMutex.Lock()
	ret, specificReturn := fake.delPrivateDataReturnsOnCall[len(fake.delPrivateDataArgsForCall)]
	fake.delPrivateDataArgsForCall = append(fake.delPrivateDataArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("DelPrivateData", []interface{}{arg1, arg2})
	fake.delPrivateDataMutex.Unlock()
	if fake.DelPrivateDataStub != nil {
		return fake.DelPrivateDataStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.delPrivateDataReturns
	return fakeReturns.result1
}

func (fake *ChaincodeStub) DelPrivateDataCallCount() int {
	fake.delPrivateDataMutex.RLock()
	defer fake.delPrivateDataMutex.RUnlock()
	return len(fake.delPrivateDataArgsForCall)
}

func (fake *ChaincodeStub) DelPrivateDataCalls(stub func(string, string) error) {
	fake.delPrivateDataMutex.Lock()
	defer fake.delPrivateDataMutex.Unlock()
	fake.DelPrivateDataStub = stub
}

func (fake *ChaincodeStub) DelPrivateDataArgsForCall(i int) (string, string) {
	fake.delPrivateDataMutex.RLock()
	defer fake.delPrivateDataMutex.RUnlock()
	argsForCall := fake.delPrivateDataArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *ChaincodeStub) DelPrivateDataReturns(result1 error) {
	fake.delPrivateDataMutex.Lock()
	defer fake.delPrivateDataMutex.Unlock()
	fake.DelPrivateDataStub = nil
	fake.delPrivateDataReturns = struct {
		result1 error
	}{result1}
}

func (fake *ChaincodeStub) DelPrivateDataReturnsOnCall(i int, result1 error) {
	fake.delPrivateDataMutex.Lock()
	defer fake.delPrivateDataMutex.Unlock()
	fake.DelPrivateDataStub = nil
	if fake.delPrivateDataReturnsOnCall == nil {
		fake.delPrivateDataReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.delPrivateDataReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *ChaincodeStub) DelState(arg1 string) error {
	fake.delStateMutex.Lock()
	ret, specificReturn := fake.delStateReturnsOnCall[len(fake.delStateArgsForCall)]
	fake.delStateArgsForCall = append(fake.delStateArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("DelState", []interface{}{arg1})
	fake.delStateMutex.Unlock()
	if fake.DelStateStub != nil {
		return fake.DelStateStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.delStateReturns
	return fakeReturns.result1
}

func (fake *ChaincodeStub) DelStateCallCount() int {
	fake.delStateMutex.RLock()
	defer fake.delStateMutex.RUnlock()
	return len(fake.delStateArgsForCall)
}

func (fake *ChaincodeStub) DelStateCalls(stub func(string) error) {
	fake.delStateMutex.Lock()
	defer fake.delStateMutex.Unlock()
	fake.DelStateStub = stub
}

func (fake *ChaincodeStub) DelStateArgsForCall(i int) string {
	fake.delStateMutex.RLock()
	defer fake.delStateMutex.RUnlock()
	argsForCall := fake.delStateArgsForCall[i]
	return argsForCall.arg1
}

func (fake *ChaincodeStub) DelStateReturns(result1 error) {
	fake.delStateMutex.Lock()
	defer fake.delStateMutex.Unlock()
	fake.DelStateStub = nil
	fake.delStateReturns = struct {
		result1 error
	}{result1}
}

func (fake *ChaincodeStub) DelStateReturnsOnCall(i int, result1 error) {
	fake.delStateMutex.Lock()
	defer fake.delStateMutex.Unlock()
	fake.DelStateStub = nil
	if fake.delStateReturnsOnCall == nil {
		fake.delStateReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.delStateReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *ChaincodeStub) GetArgs() [][]byte {
	fake.getArgsMutex.Lock()
	ret, specificReturn := fake.getArgsReturnsOnCall[len(fake.getArgsArgsForCall)]
	fake.getArgsArgsForCall = append(fake.getArgsArgsForCall, struct {
	}{})
	fake.recordInvocation("GetArgs", []interface{}{})
	fake.getArgsMutex.Unlock()
	if fake.GetArgsStub != nil {
		return fake.GetArgsStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.getArgsReturns
	return fakeReturns.result1
}

func (fake *ChaincodeStub) GetArgsCallCount() int {
	fake.getArgsMutex.RLock()
	defer fake.getArgsMutex.RUnlock()
	return len(fake.getArgsArgsForCall)
}

func (fake *ChaincodeStub) GetArgsCalls(stub func() [][]byte) {
	fake.getArgsMutex.Lock()
	defer fake.getArgsMutex.Unlock()
	fake.GetArgsStub = stub
}

func (fake *ChaincodeStub) GetArgsReturns(result1 [][]byte) {
	fake.getArgsMutex.Lock()
	defer fake.getArgsMutex.Unlock()
	fake.GetArgsStub = nil
	fake.getArgsReturns = struct {
		result1 [][]byte
	}{result1}
}

func (fake *ChaincodeStub) GetArgsReturnsOnCall(i int, result1 [][]byte) {
	fake.getArgsMutex.Lock()
	defer fake.getArgsMutex.Unlock()
	fake.GetArgsStub = nil
	if fake.getArgsReturnsOnCall == nil {
		fake.getArgsReturnsOnCall = make(map[int]struct {
			result1 [][]byte
		})
	}
	fake.getArgsReturnsOnCall[i] = struct {
		result1 [][]byte
	}{result1}
}

func (fake *ChaincodeStub) GetArgsSlice() ([]byte, error) {
	fake.getArgsSliceMutex.Lock()
	ret, specificReturn := fake.getArgsSliceReturnsOnCall[len(fake.getArgsSliceArgsForCall)]
	fake.getArgsSliceArgsForCall = append(fake.getArgsSliceArgsForCall, struct {
	}{})
	fake.recordInvocation("GetArgsSlice", []interface{}{})
	fake.getArgsSliceMutex.Unlock()
	if fake.GetArgsSliceStub != nil {
		return fake.GetArgsSliceStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getArgsSliceReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *ChaincodeStub) GetArgsSliceCallCount() int {
	fake.getArgsSliceMutex.RLock()
	defer fake.getArgsSliceMutex.RUnlock()
	return len(fake.getArgsSliceArgsForCall)
}

func (fake *ChaincodeStub) GetArgsSliceCalls(stub func() ([]byte, error)) {
	fake.getArgsSliceMutex.Lock()
	defer fake.getArgsSliceMutex.Unlock()
	fake.GetArgsSliceStub = stub
}

func (fake *ChaincodeStub) GetArgsSliceReturns(result1 []byte, result2 error) {
	fake.getArgsSliceMutex.Lock()
	defer fake.getArgsSliceMutex.Unlock()
	fake.GetArgsSliceStub = nil
	fake.getArgsSliceReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *ChaincodeStub) GetArgsSliceReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.getArgsSliceMutex.Lock()
	defer fake.getArgsSliceMutex.Unlock()
	fake.GetArgsSliceStub = nil
	if fake.getArgsSliceReturnsOnCall == nil {
		fake.getArgsSliceReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.getArgsSliceReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *ChaincodeStub) GetBinding() ([]byte, error) {
	fake.getBindingMutex.Lock()
	ret, specificReturn := fake.getBindingReturnsOnCall[len(fake.getBindingArgsForCall)]
	fake.getBindingArgsForCall = append(fake.getBindingArgsForCall, struct {
	}{})
	fake.recordInvocation("GetBinding", []interface{}{})
	fake.getBindingMutex.Unlock()
	if fake.GetBindingStub != nil {
		return fake.GetBindingStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getBindingReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *ChaincodeStub) GetBindingCallCount() int {
	fake.getBindingMutex.RLock()
	defer fake.getBindingMutex.RUnlock()
	return len(fake.getBindingArgsForCall)
}

func (fake *ChaincodeStub) GetBindingCalls(stub func() ([]byte, error)) {
	fake.getBindingMutex.Lock()
	defer fake.getBindingMutex.Unlock()
	fake.GetBindingStub = stub
}

func (fake *ChaincodeStub) GetBindingReturns(result1 []byte, result2 error) {
	fake.getBindingMutex.Lock()
	defer fake.getBindingMutex.Unlock()
	fake.GetBindingStub = nil
	fake.getBindingReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *ChaincodeStub) GetBindingReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.getBindingMutex.Lock()
	defer fake.getBindingMutex.Unlock()
	fake.GetBindingStub = nil
	if fake.getBindingReturnsOnCall == nil {
		fake.getBindingReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.getBindingReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *ChaincodeStub) GetChannelID() string {
	fake.getChannelIDMutex.Lock()
	ret, specificReturn := fake.getChannelIDReturnsOnCall[len(fake.getChannelIDArgsForCall)]
	fake.getChannelIDArgsForCall = append(fake.getChannelIDArgsForCall, struct {
	}{})
	fake.recordInvocation("GetChannelID", []interface{}{})
	fake.getChannelIDMutex.Unlock()
	if fake.GetChannelIDStub != nil {
		return fake.GetChannelIDStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.getChannelIDReturns
	return fakeReturns.result1
}

func (fake *ChaincodeStub) GetChannelIDCallCount() int {
	fake.getChannelIDMutex.RLock()
	defer fake.getChannelIDMutex.RUnlock()
	return len(fake.getChannelIDArgsForCall)
}

func (fake *ChaincodeStub) GetChannelIDCalls(stub func() string) {
	fake.getChannelIDMutex.Lock()
	defer fake.getChannelIDMutex.Unlock()
	fake.GetChannelIDStub = stub
}

func (fake *ChaincodeStub) GetChannelIDReturns(result1 string) {
	fake.getChannelIDMutex.Lock()
	defer fake.getChannelIDMutex.Unlock()
	fake.GetChannelIDStub = nil
	fake.getChannelIDReturns = struct {
		result1 string
	}{result1}
}

func (fake *ChaincodeStub) GetChannelIDReturnsOnCall(i int, result1 string) {
	fake.getChannelIDMutex.Lock()
	defer fake.getChannelIDMutex.Unlock()
	fake.GetChannelIDStub = nil
	if fake.getChannelIDReturnsOnCall == nil {
		fake.getChannelIDReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.getChannelIDReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *ChaincodeStub) GetCreator() ([]byte, error) {
	fake.getCreatorMutex.Lock()
	ret, specificReturn := fake.getCreatorReturnsOnCall[len(fake.getCreatorArgsForCall)]
	fake.getCreatorArgsForCall = append(fake.getCreatorArgsForCall, struct {
	}{})
	fake.recordInvocation("GetCreator", []interface{}{})
	fake.getCreatorMutex.Unlock()
	if fake.GetCreatorStub != nil {
		return fake.GetCreatorStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getCreatorReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *ChaincodeStub) GetCreatorCallCount() int {
	fake.getCreatorMutex.RLock()
	defer fake.getCreatorMutex.RUnlock()
	return len(fake.getCreatorArgsForCall)
}

func (fake *ChaincodeStub) GetCreatorCalls(stub func() ([]byte, error)) {
	fake.getCreatorMutex.Lock()
	defer fake.getCreatorMutex.Unlock()
	fake.GetCreatorStub = stub
}

func (fake *ChaincodeStub) GetCreatorReturns(result1 []byte, result2 error) {
	fake.getCreatorMutex.Lock()
	defer fake.getCreatorMutex.Unlock()
	fake.GetCreatorStub = nil
	fake.getCreatorReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *ChaincodeStub) GetCreatorReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.getCreatorMutex.Lock()
	defer fake.getCreatorMutex.Unlock()
	fake.GetCreatorStub = nil
	if fake.getCreatorReturnsOnCall == nil {
		fake.getCreatorReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.getCreatorReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *ChaincodeStub) GetDecorations() map[string][]byte {
	fake.getDecorationsMutex.Lock()
	ret, specificReturn := fake.getDecorationsReturnsOnCall[len(fake.getDecorationsArgsForCall)]
	fake.getDecorationsArgsForCall = append(fake.getDecorationsArgsForCall, struct {
	}{})
	fake.recordInvocation("GetDecorations", []interface{}{})
	fake.getDecorationsMutex.Unlock()
	if fake.GetDecorationsStub != nil {
		return fake.GetDecorationsStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.getDecorationsReturns
	return fakeReturns.result1
}

func (fake *ChaincodeStub) GetDecorationsCallCount() int {
	fake.getDecorationsMutex.RLock()
	defer fake.getDecorationsMutex.RUnlock()
	return len(fake.getDecorationsArgsForCall)
}

func (fake *ChaincodeStub) GetDecorationsCalls(stub func() map[string][]byte) {
	fake.getDecorationsMutex.Lock()
	defer fake.getDecorationsMutex.Unlock()
	fake.GetDecorationsStub = stub
}

func (fake *ChaincodeStub) GetDecorationsReturns(result1 map[string][]byte) {
	fake.getDecorationsMutex.Lock()
	defer fake.getDecorationsMutex.Unlock()
	fake.GetDecorationsStub = nil
	fake.getDecorationsReturns = struct {
		result1 map[string][]byte
	}{result1}
}

func (fake *ChaincodeStub) GetDecorationsReturnsOnCall(i int, result1 map[string][]byte) {
	fake.getDecorationsMutex.Lock()
	defer fake.getDecorationsMutex.Unlock()
	fake.GetDecorationsStub = nil
	if fake.getDecorationsReturnsOnCall == nil {
		fake.getDecorationsReturnsOnCall = make(map[int]struct {
			result1 map[string][]byte
		})
	}
	fake.getDecorationsReturnsOnCall[i] = struct {
		result1 map[string][]byte
	}{result1}
}

func (fake *ChaincodeStub) GetFunctionAndParameters() (string, []string) {
	fake.getFunctionAndParametersMutex.Lock()
	ret, specificReturn := fake.getFunctionAndParametersReturnsOnCall[len(fake.getFunctionAndParametersArgsForCall)]
	fake.getFunctionAndParametersArgsForCall = append(fake.getFunctionAndParametersArgsForCall, struct {
	}{})
	fake.recordInvocation("GetFunctionAndParameters", []interface{}{})
	fake.getFunctionAndParametersMutex.Unlock()
	if fake.GetFunctionAndParametersStub != nil {
		return fake.GetFunctionAndParametersStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getFunctionAndParametersReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *ChaincodeStub) GetFunctionAndParametersCallCount() int {
	fake.getFunctionAndParametersMutex.RLock()
	defer fake.getFunctionAndParametersMutex.RUnlock()
	return len(fake.getFunctionAndParametersArgsForCall)
}

func (fake *ChaincodeStub) GetFunctionAndParametersCalls(stub func() (string, []string)) {
	fake.getFunctionAndParametersMutex.Lock()
	defer fake.getFunctionAndParametersMutex.Unlock()
	fake.GetFunctionAndParametersStub = stub
}

func (fake *ChaincodeStub) GetFunctionAndParametersReturns(result1 string, result2 []string) {
	fake.getFunctionAndParametersMutex.Lock()
	defer fake.getFunctionAndParametersMutex.Unlock()
	fake.GetFunctionAndParametersStub = nil
	fake.getFunctionAndParametersReturns = struct {
		result1 string
		result2 []string
	}{result1, result2}
}

func (fake *ChaincodeStub) GetFunctionAndParametersReturnsOnCall(i int, result1 string, result2 []string) {
	fake.getFunctionAndParametersMutex.Lock()
	defer fake.getFunctionAndParametersMutex.Unlock()
	fake.GetFunctionAndParametersStub = nil
	if fake.getFunctionAndParametersReturnsOnCall == nil {
		fake.getFunctionAndParametersReturnsOnCall = make(map[int]struct {
			result1 string
			result2 []string
		})
	}
	fake.getFunctionAndParametersReturnsOnCall[i] = struct {
		result1 string
		result2 []string
	}{result1, result2}
}

func (fake *ChaincodeStub) GetHistoryForKey(arg1 string) (shim.HistoryQueryIteratorInterface, error) {
	fake.getHistoryForKeyMutex.Lock()
	ret, specificReturn := fake.getHistoryForKeyReturnsOnCall[len(fake.getHistoryForKeyArgsForCall)]
	fake.getHistoryForKeyArgsForCall = append(fake.getHistoryForKeyArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetHistoryForKey", []interface{}{arg1})
	fake.getHistoryForKeyMutex.Unlock()
	if fake.GetHistoryForKeyStub != nil {
		return fake.GetHistoryForKeyStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getHistoryForKeyReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *ChaincodeStub) GetHistoryForKeyCallCount() int {
	fake.getHistoryForKeyMutex.RLock()
	defer fake.getHistoryForKeyMutex.RUnlock()
	return len(fake.getHistoryForKeyArgsForCall)
}

func (fake *ChaincodeStub) GetHistoryForKeyCalls(stub func(string) (shim.HistoryQueryIteratorInterface, error)) {
	fake.getHistoryForKeyMutex.Lock()
	defer fake.getHistoryForKeyMutex.Unlock()
	fake.GetHistoryForKeyStub = stub
}

func (fake *ChaincodeStub) GetHistoryForKeyArgsForCall(i int) string {
	fake.getHistoryForKeyMutex.RLock()
	defer fake.getHistoryForKeyMutex.RUnlock()
	argsForCall := fake.getHistoryForKeyArgsForCall[i]
	return argsForCall.arg1
}

func (fake *ChaincodeStub) GetHistoryForKeyReturns(result1 shim.HistoryQueryIteratorInterface, result2 error) {
	fake.getHistoryForKeyMutex.Lock()
	defer fake.getHistoryForKeyMutex.Unlock()
	fake.GetHistoryForKeyStub = nil
	fake.getHistoryForKeyReturns = struct {
		result1 shim.HistoryQueryIteratorInterface
		result2 error
	}{result1, result2}
}

func (fake *ChaincodeStub) GetHistoryForKeyReturnsOnCall(i int, result1 shim.HistoryQueryIteratorInterface, result2 error) {
	fake.getHistoryForKeyMutex.Lock()
	defer fake.getHistoryForKeyMutex.Unlock()
	fake.GetHistoryForKeyStub = nil
	if fake.getHistoryForKeyReturnsOnCall == nil {
		fake.getHistoryForKeyReturnsOnCall = make(map[int]struct {
			result1 shim.HistoryQueryIteratorInterface
			result2 error
		})
	}
	fake.getHistoryForKeyReturnsOnCall[i] = struct {
		result1 shim.HistoryQueryIteratorInterface
		result2 error
	}{result1, result2}
}

func (fake *ChaincodeStub) GetPrivateData(arg1 string, arg2 string) ([]byte, error) {
	fake.getPrivateDataMutex.Lock()
	ret, specificReturn := fake.getPrivateDataReturnsOnCall[len(fake.getPrivateDataArgsForCall)]
	fake.getPrivateDataArgsForCall = append(fake.getPrivateDataArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("GetPrivateData", []interface{}{arg1, arg2})
	fake.getPrivateDataMutex.Unlock()
	if fake.GetPrivateDataStub != nil {
		return fake.GetPrivateDataStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getPrivateDataReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *ChaincodeStub) GetPrivateDataCallCount() int {
	fake.getPrivateDataMutex.RLock()
	defer fake.getPrivateDataMutex.RUnlock()
	return len(fake.getPrivateDataArgsForCall)
}

func (fake *ChaincodeStub) GetPrivateDataCalls(stub func(string, string) ([]byte, error)) {
	fake.getPrivateDataMutex.Lock()
	defer fake.getPrivateDataMutex.Unlock()
	fake.GetPrivateDataStub = stub
}

func (fake *ChaincodeStub) GetPrivateDataArgsForCall(i int) (string, string) {
	fake.getPrivateDataMutex.RLock()
	defer fake.getPrivateDataMutex.RUnlock()
	argsForCall := fake.getPrivateDataArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *ChaincodeStub) GetPrivateDataReturns(result1 []byte, result2 error) {
	fake.getPrivateDataMutex.Lock()
	defer fake.getPrivateDataMutex.Unlock()
	fake.GetPrivateDataStub = nil
	fake.getPrivateDataReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *ChaincodeStub) GetPrivateDataReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.getPrivateDataMutex.Lock()
	defer fake.getPrivateDataMutex.Unlock()
	fake.GetPrivateDataStub = nil
	if fake.getPrivateDataReturnsOnCall == nil {
		fake.getPrivateDataReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.getPrivateDataReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *ChaincodeStub) GetPrivateDataByPartialCompositeKey(arg1 string, arg2 string, arg3 []string) (shim.StateQueryIteratorInterface, error) {
	var arg3Copy []string
	if arg3 != nil {
		arg3Copy = make([]string, len(arg3))
		copy(arg3Copy, arg3)
	}
	fake.getPrivateDataByPartialCompositeKeyMutex.Lock()
	ret, specificReturn := fake.getPrivateDataByPartialCompositeKeyReturnsOnCall[len(fake.getPrivateDataByPartialCompositeKeyArgsForCall)]
	fake.getPrivateDataByPartialCompositeKeyArgsForCall = append(fake.getPrivateDataByPartialCompositeKeyArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 []string
	}{arg1, arg2, arg3Copy})
	fake.recordInvocation("GetPrivateDataByPartialCompositeKey", []interface{}{arg1, arg2, arg3Copy})
	fake.getPrivateDataByPartialCompositeKeyMutex.Unlock()
	if fake.GetPrivateDataByPartialCompositeKeyStub != nil {
		return fake.GetPrivateDataByPartialCompositeKeyStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getPrivateDataByPartialCompositeKeyReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *ChaincodeStub) GetPrivateDataByPartialCompositeKeyCallCount() int {
	fake.getPrivateDataByPartialCompositeKeyMutex.RLock()
	defer fake.getPrivateDataByPartialCompositeKeyMutex.RUnlock()
	return len(fake.getPrivateDataByPartialCompositeKeyArgsForCall)
}

func (fake *ChaincodeStub) GetPrivateDataByPartialCompositeKeyCalls(stub func(string, string, []string) (shim.StateQueryIteratorInterface, error)) {
	fake.getPrivateDataByPartialCompositeKeyMutex.Lock()
	defer fake.getPrivateDataByPartialCompositeKeyMutex.Unlock()
	fake.GetPrivateDataByPartialCompositeKeyStub = stub
}

func (fake *ChaincodeStub) GetPrivateDataByPartialCompositeKeyArgsForCall(i int) (string, string, []string) {
	fake.getPrivateDataByPartialCompositeKeyMutex.RLock()
	defer fake.getPrivateDataByPartialCompositeKeyMutex.RUnlock()
	argsForCall := fake.getPrivateDataByPartialCompositeKeyArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *ChaincodeStub) GetPrivateDataByPartialCompositeKeyReturns(result1 shim.StateQueryIteratorInterface, result2 error) {
	fake.getPrivateDataByPartialCompositeKeyMutex.Lock()
	defer fake.getPrivateDataByPartialCompositeKeyMutex.Unlock()
	fake.GetPrivateDataByPartialCompositeKeyStub = nil
	fake.getPrivateDataByPartialCompositeKeyReturns = struct {
		result1 shim.StateQueryIteratorInterface
		result2 error
	}{result1, result2}
}

func (fake *ChaincodeStub) GetPrivateDataByPartialCompositeKeyReturnsOnCall(i int, result1 shim.StateQueryIteratorInterface, result2 error) {
	fake.getPrivateDataByPartialCompositeKeyMutex.Lock()
	defer fake.getPrivateDataByPartialCompositeKeyMutex.Unlock()
	fake.GetPrivateDataByPartialCompositeKeyStub = nil
	if fake.getPrivateDataByPartialCompositeKeyReturnsOnCall == nil {
		fake.getPrivateDataByPartialCompositeKeyReturnsOnCall = make(map[int]struct {
			result1 shim.StateQueryIteratorInterface
			result2 error
		})
	}
	fake.getPrivateDataByPartialCompositeKeyReturnsOnCall[i] = struct {
		result1 shim.StateQueryIteratorInterface
		result2 error
	}{result1, result2}
}

func (fake *ChaincodeStub) GetPrivateDataByRange(arg1 string, arg2 string, arg3 string) (shim.StateQueryIteratorInterface, error) {
	fake.getPrivateDataByRangeMutex.Lock()
	ret, specificReturn := fake.getPrivateDataByRangeReturnsOnCall[len(fake.getPrivateDataByRangeArgsForCall)]
	fake.getPrivateDataByRangeArgsForCall = append(fake.getPrivateDataByRangeArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	fake.recordInvocation("GetPrivateDataByRange", []interface{}{arg1, arg2, arg3})
	fake.getPrivateDataByRangeMutex.Unlock()
	if fake.GetPrivateDataByRangeStub != nil {
		return fake.GetPrivateDataByRangeStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getPrivateDataByRangeReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *ChaincodeStub) GetPrivateDataByRangeCallCount() int {
	fake.getPrivateDataByRangeMutex.RLock()
	defer fake.getPrivateDataByRangeMutex.RUnlock()
	return len(fake.getPrivateDataByRangeArgsForCall)
}

func (fake *ChaincodeStub) GetPrivateDataByRangeCalls(stub func(string, string, string) (shim.StateQueryIteratorInterface, error)) {
	fake.getPrivateDataByRangeMutex.Lock()
	defer fake.getPrivateDataByRangeMutex.Unlock()
	fake.GetPrivateDataByRangeStub = stub
}

func (fake *ChaincodeStub) GetPrivateDataByRangeArgsForCall(i int) (string, string, string) {
	fake.getPrivateDataByRangeMutex.RLock()
	defer fake.getPrivateDataByRangeMutex.RUnlock()
	argsForCall := fake.getPrivateDataByRangeArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *ChaincodeStub) GetPrivateDataByRangeReturns(result1 shim.StateQueryIteratorInterface, result2 error) {
	fake.getPrivateDataByRangeMutex.Lock()
	defer fake.getPrivateDataByRangeMutex.Unlock()
	fake.GetPrivateDataByRangeStub = nil
	fake.getPrivateDataByRangeReturns = struct {
		result1 shim.StateQueryIteratorInterface
		result2 error
	}{result1, result2}
}

func (fake *ChaincodeStub) GetPrivateDataByRangeReturnsOnCall(i int, result1 shim.StateQueryIteratorInterface, result2 error) {
	fake.getPrivateDataByRangeMutex.Lock()
	defer fake.getPrivateDataByRangeMutex.Unlock()
	fake.GetPrivateDataByRangeStub = nil
	if fake.getPrivateDataByRangeReturnsOnCall == nil {
		fake.getPrivateDataByRangeReturnsOnCall = make(map[int]struct {
			result1 shim.StateQueryIteratorInterface
			result2 error
		})
	}
	fake.getPrivateDataByRangeReturnsOnCall[i] = struct {
		result1 shim.StateQueryIteratorInterface
		result2 error
	}{result1, result2}
}

func (fake *ChaincodeStub) GetPrivateDataHash(arg1 string, arg2 string) ([]byte, error) {
	fake.getPrivateDataHashMutex.Lock()
	ret, specificReturn := fake.getPrivateDataHashReturnsOnCall[len(fake.getPrivateDataHashArgsForCall)]
	fake.getPrivateDataHashArgsForCall = append(fake.getPrivateDataHashArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("GetPrivateDataHash", []interface{}{arg1, arg2})
	fake.getPrivateDataHashMutex.Unlock()
	if fake.GetPrivateDataHashStub != nil {
		return fake.GetPrivateDataHashStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getPrivateDataHashReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *ChaincodeStub) GetPrivateDataHashCallCount() int {
	fake.getPrivateDataHashMutex.RLock()
	defer fake.getPrivateDataHashMutex.RUnlock()
	return len(fake.getPrivateDataHashArgsForCall)
}

func (fake *ChaincodeStub) GetPrivateDataHashCalls(stub func(string, string) ([]byte, error)) {
	fake.getPrivateDataHashMutex.Lock()
	defer fake.getPrivateDataHashMutex.Unlock()
	fake.GetPrivateDataHashStub = stub
}

func (fake *ChaincodeStub) GetPrivateDataHashArgsForCall(i int) (string, string) {
	fake.getPrivateDataHashMutex.RLock()
	defer fake.getPrivateDataHashMutex.RUnlock()
	argsForCall := fake.getPrivateDataHashArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *ChaincodeStub) GetPrivateDataHashReturns(result1 []byte, result2 error) {
	fake.getPrivateDataHashMutex.Lock()
	defer fake.getPrivateDataHashMutex.Unlock()
	fake.GetPrivateDataHashStub = nil
	fake.getPrivateDataHashReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *ChaincodeStub) GetPrivateDataHashReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.getPrivateDataHashMutex.Lock()
	defer fake.getPrivateDataHashMutex.Unlock()
	fake.GetPrivateDataHashStub = nil
	if fake.getPrivateDataHashReturnsOnCall == nil {
		fake.getPrivateDataHashReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.getPrivateDataHashReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *ChaincodeStub) GetPrivateDataQueryResult(arg1 string, arg2 string) (shim.StateQueryIteratorInterface, error) {
	fake.getPrivateDataQueryResultMutex.Lock()
	ret, specificReturn := fake.getPrivateDataQueryResultReturnsOnCall[len(fake.getPrivateDataQueryResultArgsForCall)]
	fake.getPrivateDataQueryResultArgsForCall = append(fake.getPrivateDataQueryResultArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("GetPrivateDataQueryResult", []interface{}{arg1, arg2})
	fake.getPrivateDataQueryResultMutex.Unlock()
	if fake.GetPrivateDataQueryResultStub != nil {
		return fake.GetPrivateDataQueryResultStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getPrivateDataQueryResultReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *ChaincodeStub) GetPrivateDataQueryResultCallCount() int {
	fake.getPrivateDataQueryResultMutex.RLock()
	defer fake.getPrivateDataQueryResultMutex.RUnlock()
	return len(fake.getPrivateDataQueryResultArgsForCall)
}

func (fake *ChaincodeStub) GetPrivateDataQueryResultCalls(stub func(string, string) (shim.StateQueryIteratorInterface, error)) {
	fake.getPrivateDataQueryResultMutex.Lock()
	defer fake.getPrivateDataQueryResultMutex.Unlock()
	fake.GetPrivateDataQueryResultStub = stub
}

func (fake *ChaincodeStub) GetPrivateDataQueryResultArgsForCall(i int) (string, string) {
	fake.getPrivateDataQueryResultMutex.RLock()
	defer fake.getPrivateDataQueryResultMutex.RUnlock()
	argsForCall := fake.getPrivateDataQueryResultArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *ChaincodeStub) GetPrivateDataQueryResultReturns(result1 shim.StateQueryIteratorInterface, result2 error) {
	fake.getPrivateDataQueryResultMutex.Lock()
	defer fake.getPrivateDataQueryResultMutex.Unlock()
	fake.GetPrivateDataQueryResultStub = nil
	fake.getPrivateDataQueryResultReturns = struct {
		result1 shim.StateQueryIteratorInterface
		result2 error
	}{result1, result2}
}

func (fake *ChaincodeStub) GetPrivateDataQueryResultReturnsOnCall(i int, result1 shim.StateQueryIteratorInterface, result2 error) {
	fake.getPrivateDataQueryResultMutex.Lock()
	defer fake.getPrivateDataQueryResultMutex.Unlock()
	fake.GetPrivateDataQueryResultStub = nil
	if fake.getPrivateDataQueryResultReturnsOnCall == nil {
		fake.getPrivateDataQueryResultReturnsOnCall = make(map[int]struct {
			result1 shim.StateQueryIteratorInterface
			result2 error
		})
	}
	fake.getPrivateDataQueryResultReturnsOnCall[i] = struct {
		result1 shim.StateQueryIteratorInterface
		result2 error
	}{result1, result2}
}

func (fake *ChaincodeStub) GetPrivateDataValidationParameter(arg1 string, arg2 string) ([]byte, error) {
	fake.getPrivateDataValidationParameterMutex.Lock()
	ret, specificReturn := fake.getPrivateDataValidationParameterReturnsOnCall[len(fake.getPrivateDataValidationParameterArgsForCall)]
	fake.getPrivateDataValidationParameterArgsForCall = append(fake.getPrivateDataValidationParameterArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("GetPrivateDataValidationParameter", []interface{}{arg1, arg2})
	fake.getPrivateDataValidationParameterMutex.Unlock()
	if fake.GetPrivateDataValidationParameterStub != nil {
		return fake.GetPrivateDataValidationParameterStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getPrivateDataValidationParameterReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *ChaincodeStub) GetPrivateDataValidationParameterCallCount() int {
	fake.getPrivateDataValidationParameterMutex.RLock()
	defer fake.getPrivateDataValidationParameterMutex.RUnlock()
	return len(fake.getPrivateDataValidationParameterArgsForCall)
}

func (fake *ChaincodeStub) GetPrivateDataValidationParameterCalls(stub func(string, string) ([]byte, error)) {
	fake.getPrivateDataValidationParameterMutex.Lock()
	defer fake.getPrivateDataValidationParameterMutex.Unlock()
	fake.GetPrivateDataValidationParameterStub = stub
}

func (fake *ChaincodeStub) GetPrivateDataValidationParameterArgsForCall(i int) (string, string) {
	fake.getPrivateDataValidationParameterMutex.RLock()
	defer fake.getPrivateDataValidationParameterMutex.RUnlock()
	argsForCall := fake.getPrivateDataValidationParameterArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *ChaincodeStub) GetPrivateDataValidationParameterReturns(result1 []byte, result2 error) {
	fake.getPrivateDataValidationParameterMutex.Lock()
	defer fake.getPrivateDataValidationParameterMutex.Unlock()
	fake.GetPrivateDataValidationParameterStub = nil
	fake.getPrivateDataValidationParameterReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *ChaincodeStub) GetPrivateDataValidationParameterReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.getPrivateDataValidationParameterMutex.Lock()
	defer fake.getPrivateDataValidationParameterMutex.Unlock()
	fake.GetPrivateDataValidationParameterStub = nil
	if fake.getPrivateDataValidationParameterReturnsOnCall == nil {
		fake.getPrivateDataValidationParameterReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.getPrivateDataValidationParameterReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *ChaincodeStub) GetQueryResult(arg1 string) (shim.StateQueryIteratorInterface, error) {
	fake.getQueryResultMutex.Lock()
	ret, specificReturn := fake.getQueryResultReturnsOnCall[len(fake.getQueryResultArgsForCall)]
	fake.getQueryResultArgsForCall = append(fake.getQueryResultArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetQueryResult", []interface{}{arg1})
	fake.getQueryResultMutex.Unlock()
	if fake.GetQueryResultStub != nil {
		return fake.GetQueryResultStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getQueryResultReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *ChaincodeStub) GetQueryResultCallCount() int {
	fake.getQueryResultMutex.RLock()
	defer fake.getQueryResultMutex.RUnlock()
	return len(fake.getQueryResultArgsForCall)
}

func (fake *ChaincodeStub) GetQueryResultCalls(stub func(string) (shim.StateQueryIteratorInterface, error)) {
	fake.getQueryResultMutex.Lock()
	defer fake.getQueryResultMutex.Unlock()
	fake.GetQueryResultStub = stub
}

func (fake *ChaincodeStub) GetQueryResultArgsForCall(i int) string {
	fake.getQueryResultMutex.RLock()
	defer fake.getQueryResultMutex.RUnlock()
	argsForCall := fake.getQueryResultArgsForCall[i]
	return argsForCall.arg1
}

func (fake *ChaincodeStub) GetQueryResultReturns(result1 shim.StateQueryIteratorInterface, result2 error) {
	fake.getQueryResultMutex.Lock()
	defer fake.getQueryResultMutex.Unlock()
	fake.GetQueryResultStub = nil
	fake.getQueryResultReturns = struct {
		result1 shim.StateQueryIteratorInterface
		result2 error
	}{result1, result2}
}

func (fake *ChaincodeStub) GetQueryResultReturnsOnCall(i int, result1 shim.StateQueryIteratorInterface, result2 error) {
	fake.getQueryResultMutex.Lock()
	defer fake.getQueryResultMutex.Unlock()
	fake.GetQueryResultStub = nil
	if fake.getQueryResultReturnsOnCall == nil {
		fake.getQueryResultReturnsOnCall = make(map[int]struct {
			result1 shim.StateQueryIteratorInterface
			result2 error
		})
	}
	fake.getQueryResultReturnsOnCall[i] = struct {
		result1 shim.StateQueryIteratorInterface
		result2 error
	}{result1, result2}
}

func (fake *ChaincodeStub) GetQueryResultWithPagination(arg1 string, arg2 int32, arg3 string) (shim.StateQueryIteratorInterface, *peer.QueryResponseMetadata, error) {
	fake.getQueryResultWithPaginationMutex.Lock()
	ret, specificReturn := fake.getQueryResultWithPaginationReturnsOnCall[len(fake.getQueryResultWithPaginationArgsForCall)]
	fake.getQueryResultWithPaginationArgsForCall = append(fake.getQueryResultWithPaginationArgsForCall, struct {
		arg1 string
		arg2 int32
		arg3 string
	}{arg1, arg2, arg3})
	fake.recordInvocation("GetQueryResultWithPagination", []interface{}{arg1, arg2, arg3})
	fake.getQueryResultWithPaginationMutex.Unlock()
	if fake.GetQueryResultWithPaginationStub != nil {
		return fake.GetQueryResultWithPaginationStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getQueryResultWithPaginationReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *ChaincodeStub) GetQueryResultWithPaginationCallCount() int {
	fake.getQueryResultWithPaginationMutex.RLock()
	defer fake.getQueryResultWithPaginationMutex.RUnlock()
	return len(fake.getQueryResultWithPaginationArgsForCall)
}

func (fake *ChaincodeStub) GetQueryResultWithPaginationCalls(stub func(string, int32, string) (shim.StateQueryIteratorInterface, *peer.QueryResponseMetadata, error)) {
	fake.getQueryResultWithPaginationMutex.Lock()
	defer fake.getQueryResultWithPaginationMutex.Unlock()
	fake.GetQueryResultWithPaginationStub = stub
}

func (fake *ChaincodeStub) GetQueryResultWithPaginationArgsForCall(i int) (string, int32, string) {
	fake.getQueryResultWithPaginationMutex.RLock()
	defer fake.getQueryResultWithPaginationMutex.RUnlock()
	argsForCall := fake.getQueryResultWithPaginationArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *ChaincodeStub) GetQueryResultWithPaginationReturns(result1 shim.StateQueryIteratorInterface, result2 *peer.QueryResponseMetadata, result3 error) {
	fake.getQueryResultWithPaginationMutex.Lock()
	defer fake.getQueryResultWithPaginationMutex.Unlock()
	fake.GetQueryResultWithPaginationStub = nil
	fake.getQueryResultWithPaginationReturns = struct {
		result1 shim.StateQueryIteratorInterface
		result2 *peer.QueryResponseMetadata
		result3 error
	}{result1, result2, result3}
}

func (fake *ChaincodeStub) GetQueryResultWithPaginationReturnsOnCall(i int, result1 shim.StateQueryIteratorInterface, result2 *peer.QueryResponseMetadata, result3 error) {
	fake.getQueryResultWithPaginationMutex.Lock()
	defer fake.getQueryResultWithPaginationMutex.Unlock()
	fake.GetQueryResultWithPaginationStub = nil
	if fake.getQueryResultWithPaginationReturnsOnCall == nil {
		fake.getQueryResultWithPaginationReturnsOnCall = make(map[int]struct {
			result1 shim.StateQueryIteratorInterface
			result2 *peer.QueryResponseMetadata
			result3 error
		})
	}
	fake.getQueryResultWithPaginationReturnsOnCall[i] = struct {
		result1 shim.StateQueryIteratorInterface
		result2 *peer.QueryResponseMetadata
		result3 error
	}{result1, result2, result3}
}

func (fake *ChaincodeStub) GetSignedProposal() (*peer.SignedProposal, error) {
	fake.getSignedProposalMutex.Lock()
	ret, specificReturn := fake.getSignedProposalReturnsOnCall[len(fake.getSignedProposalArgsForCall)]
	fake.getSignedProposalArgsForCall = append(fake.getSignedProposalArgsForCall, struct {
	}{})
	fake.recordInvocation("GetSignedProposal", []interface{}{})
	fake.getSignedProposalMutex.Unlock()
	if fake.GetSignedProposalStub != nil {
		return fake.GetSignedProposalStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getSignedProposalReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *ChaincodeStub) GetSignedProposalCallCount() int {
	fake.getSignedProposalMutex.RLock()
	defer fake.getSignedProposalMutex.RUnlock()
	return len(fake.getSignedProposalArgsForCall)
}

func (fake *ChaincodeStub) GetSignedProposalCalls(stub func() (*peer.SignedProposal, error)) {
	fake.getSignedProposalMutex.Lock()
	defer fake.getSignedProposalMutex.Unlock()
	fake.GetSignedProposalStub = stub
}

func (fake *ChaincodeStub) GetSignedProposalReturns(result1 *peer.SignedProposal, result2 error) {
	fake.getSignedProposalMutex.Lock()
	defer fake.getSignedProposalMutex.Unlock()
	fake.GetSignedProposalStub = nil
	fake.getSignedProposalReturns = struct {
		result1 *peer.SignedProposal
		result2 error
	}{result1, result2}
}

func (fake *ChaincodeStub) GetSignedProposalReturnsOnCall(i int, result1 *peer.SignedProposal, result2 error) {
	fake.getSignedProposalMutex.Lock()
	defer fake.getSignedProposalMutex.Unlock()
	fake.GetSignedProposalStub = nil
	if fake.getSignedProposalReturnsOnCall == nil {
		fake.getSignedProposalReturnsOnCall = make(map[int]struct {
			result1 *peer.SignedProposal
			result2 error
		})
	}
	fake.getSignedProposalReturnsOnCall[i] = struct {
		result1 *peer.SignedProposal
		result2 error
	}{result1, result2}
}

func (fake *ChaincodeStub) GetState(arg1 string) ([]byte, error) {
	fake.getStateMutex.Lock()
	ret, specificReturn := fake.getStateReturnsOnCall[len(fake.getStateArgsForCall)]
	fake.getStateArgsForCall = append(fake.getStateArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetState", []interface{}{arg1})
	fake.getStateMutex.Unlock()
	if fake.GetStateStub != nil {
		return fake.GetStateStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getStateReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *ChaincodeStub) GetStateCallCount() int {
	fake.getStateMutex.RLock()
	defer fake.getStateMutex.RUnlock()
	return len(fake.getStateArgsForCall)
}

func (fake *ChaincodeStub) GetStateCalls(stub func(string) ([]byte, error)) {
	fake.getStateMutex.Lock()
	defer fake.getStateMutex.Unlock()
	fake.GetStateStub = stub
}

func (fake *ChaincodeStub) GetStateArgsForCall(i int) string {
	fake.getStateMutex.RLock()
	defer fake.getStateMutex.RUnlock()
	argsForCall := fake.getStateArgsForCall[i]
	return argsForCall.arg1
}

func (fake *ChaincodeStub) GetStateReturns(result1 []byte, result2 error) {
	fake.getStateMutex.Lock()
	defer fake.getStateMutex.Unlock()
	fake.GetStateStub = nil
	fake.getStateReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *ChaincodeStub) GetStateReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.getStateMutex.Lock()
	defer fake.getStateMutex.Unlock()
	fake.GetStateStub = nil
	if fake.getStateReturnsOnCall == nil {
		fake.getStateReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.getStateReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *ChaincodeStub) GetStateByPartialCompositeKey(arg1 string, arg2 []string) (shim.StateQueryIteratorInterface, error) {
	var arg2Copy []string
	if arg2 != nil {
		arg2Copy = make([]string, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.getStateByPartialCompositeKeyMutex.Lock()
	ret, specificReturn := fake.getStateByPartialCompositeKeyReturnsOnCall[len(fake.getStateByPartialCompositeKeyArgsForCall)]
	fake.getStateByPartialCompositeKeyArgsForCall = append(fake.getStateByPartialCompositeKeyArgsForCall, struct {
		arg1 string
		arg2 []string
	}{arg1, arg2Copy})
	fake.recordInvocation("GetStateByPartialCompositeKey", []interface{}{arg1, arg2Copy})
	fake.getStateByPartialCompositeKeyMutex.Unlock()
	if fake.GetStateByPartialCompositeKeyStub != nil {
		return fake.GetStateByPartialCompositeKeyStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getStateByPartialCompositeKeyReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *ChaincodeStub) GetStateByPartialCompositeKeyCallCount() int {
	fake.getStateByPartialCompositeKeyMutex.RLock()
	defer fake.getStateByPartialCompositeKeyMutex.RUnlock()
	return len(fake.getStateByPartialCompositeKeyArgsForCall)
}

func (fake *ChaincodeStub) GetStateByPartialCompositeKeyCalls(stub func(string, []string) (shim.StateQueryIteratorInterface, error)) {
	fake.getStateByPartialCompositeKeyMutex.Lock()
	defer fake.getStateByPartialCompositeKeyMutex.Unlock()
	fake.GetStateByPartialCompositeKeyStub = stub
}

func (fake *ChaincodeStub) GetStateByPartialCompositeKeyArgsForCall(i int) (string, []string) {
	fake.getStateByPartialCompositeKeyMutex.RLock()
	defer fake.getStateByPartialCompositeKeyMutex.RUnlock()
	argsForCall := fake.getStateByPartialCompositeKeyArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *ChaincodeStub) GetStateByPartialCompositeKeyReturns(result1 shim.StateQueryIteratorInterface, result2 error) {
	fake.getStateByPartialCompositeKeyMutex.Lock()
	defer fake.getStateByPartialCompositeKeyMutex.Unlock()
	fake.GetStateByPartialCompositeKeyStub = nil
	fake.getStateByPartialCompositeKeyReturns = struct {
		result1 shim.StateQueryIteratorInterface
		result2 error
	}{result1, result2}
}

func (fake *ChaincodeStub) GetStateByPartialCompositeKeyReturnsOnCall(i int, result1 shim.StateQueryIteratorInterface, result2 error) {
	fake.getStateByPartialCompositeKeyMutex.Lock()
	defer fake.getStateByPartialCompositeKeyMutex.Unlock()
	fake.GetStateByPartialCompositeKeyStub = nil
	if fake.getStateByPartialCompositeKeyReturnsOnCall == nil {
		fake.getStateByPartialCompositeKeyReturnsOnCall = make(map[int]struct {
			result1 shim.StateQueryIteratorInterface
			result2 error
		})
	}
	fake.getStateByPartialCompositeKeyReturnsOnCall[i] = struct {
		result1 shim.StateQueryIteratorInterface
		result2 error
	}{result1, result2}
}

func (fake *ChaincodeStub) GetStateByPartialCompositeKeyWithPagination(arg1 string, arg2 []string, arg3 int32, arg4 string) (shim.StateQueryIteratorInterface, *peer.QueryResponseMetadata, error) {
	var arg2Copy []string
	if arg2 != nil {
		arg2Copy = make([]string, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.getStateByPartialCompositeKeyWithPaginationMutex.Lock()
	ret, specificReturn := fake.getStateByPartialCompositeKeyWithPaginationReturnsOnCall[len(fake.getStateByPartialCompositeKeyWithPaginationArgsForCall)]
	fake.getStateByPartialCompositeKeyWithPaginationArgsForCall = append(fake.getStateByPartialCompositeKeyWithPaginationArgsForCall, struct {
		arg1 string
		arg2 []string
		arg3 int32
		arg4 string
	}{arg1, arg2Copy, arg3, arg4})
	fake.recordInvocation("GetStateByPartialCompositeKeyWithPagination", []interface{}{arg1, arg2Copy, arg3, arg4})
	fake.getStateByPartialCompositeKeyWithPaginationMutex.Unlock()
	if fake.GetStateByPartialCompositeKeyWithPaginationStub != nil {
		return fake.GetStateByPartialCompositeKeyWithPaginationStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getStateByPartialCompositeKeyWithPaginationReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *ChaincodeStub) GetStateByPartialCompositeKeyWithPaginationCallCount() int {
	fake.getStateByPartialCompositeKeyWithPaginationMutex.RLock()
	defer fake.getStateByPartialCompositeKeyWithPaginationMutex.RUnlock()
	return len(fake.getStateByPartialCompositeKeyWithPaginationArgsForCall)
}

func (fake *ChaincodeStub) GetStateByPartialCompositeKeyWithPaginationCalls(stub func(string, []string, int32, string) (shim.StateQueryIteratorInterface, *peer.QueryResponseMetadata, error)) {
	fake.getStateByPartialCompositeKeyWithPaginationMutex.Lock()
	defer fake.getStateByPartialCompositeKeyWithPaginationMutex.Unlock()
	fake.GetStateByPartialCompositeKeyWithPaginationStub = stub
}

func (fake *ChaincodeStub) GetStateByPartialCompositeKeyWithPaginationArgsForCall(i int) (string, []string, int32, string) {
	fake.getStateByPartialCompositeKeyWithPaginationMutex.RLock()
	defer fake.getStateByPartialCompositeKeyWithPaginationMutex.RUnlock()
	argsForCall := fake.getStateByPartialCompositeKeyWithPaginationArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *ChaincodeStub) GetStateByPartialCompositeKeyWithPaginationReturns(result1 shim.StateQueryIteratorInterface, result2 *peer.QueryResponseMetadata, result3 error) {
	fake.getStateByPartialCompositeKeyWithPaginationMutex.Lock()
	defer fake.getStateByPartialCompositeKeyWithPaginationMutex.Unlock()
	fake.GetStateByPartialCompositeKeyWithPaginationStub = nil
	fake.getStateByPartialCompositeKeyWithPaginationReturns = struct {
		result1 shim.StateQueryIteratorInterface
		result2 *peer.QueryResponseMetadata
		result3 error
	}{result1, result2, result3}
}

func (fake *ChaincodeStub) GetStateByPartialCompositeKeyWithPaginationReturnsOnCall(i int, result1 shim.StateQueryIteratorInterface, result2 *peer.QueryResponseMetadata, result3 error) {
	fake.getStateByPartialCompositeKeyWithPaginationMutex.Lock()
	defer fake.getStateByPartialCompositeKeyWithPaginationMutex.Unlock()
	fake.GetStateByPartialCompositeKeyWithPaginationStub = nil
	if fake.getStateByPartialCompositeKeyWithPaginationReturnsOnCall == nil {
		fake.getStateByPartialCompositeKeyWithPaginationReturnsOnCall = make(map[int]struct {
			result1 shim.StateQueryIteratorInterface
			result2 *peer.QueryResponseMetadata
			result3 error
		})
	}
	fake.getStateByPartialCompositeKeyWithPaginationReturnsOnCall[i] = struct {
		result1 shim.StateQueryIteratorInterface
		result2 *peer.QueryResponseMetadata
		result3 error
	}{result1, result2, result3}
}

func (fake *ChaincodeStub) GetStateByRange(arg1 string, arg2 string) (shim.StateQueryIteratorInterface, error) {
	fake.getStateByRangeMutex.Lock()
	ret, specificReturn := fake.getStateByRangeReturnsOnCall[len(fake.getStateByRangeArgsForCall)]
	fake.getStateByRangeArgsForCall = append(fake.getStateByRangeArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("GetStateByRange", []interface{}{arg1, arg2})
	fake.getStateByRangeMutex.Unlock()
	if fake.GetStateByRangeStub != nil {
		return fake.GetStateByRangeStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getStateByRangeReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *ChaincodeStub) GetStateByRangeCallCount() int {
	fake.getStateByRangeMutex.RLock()
	defer fake.getStateByRangeMutex.RUnlock()
	return len(fake.getStateByRangeArgsForCall)
}

func (fake *ChaincodeStub) GetStateByRangeCalls(stub func(string, string) (shim.StateQueryIteratorInterface, error)) {
	fake.getStateByRangeMutex.Lock()
	defer fake.getStateByRangeMutex.Unlock()
	fake.GetStateByRangeStub = stub
}

func (fake *ChaincodeStub) GetStateByRangeArgsForCall(i int) (string, string) {
	fake.getStateByRangeMutex.RLock()
	defer fake.getStateByRangeMutex.RUnlock()
	argsForCall := fake.getStateByRangeArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *ChaincodeStub) GetStateByRangeReturns(result1 shim.StateQueryIteratorInterface, result2 error) {
	fake.getStateByRangeMutex.Lock()
	defer fake.getStateByRangeMutex.Unlock()
	fake.GetStateByRangeStub = nil
	fake.getStateByRangeReturns = struct {
		result1 shim.StateQueryIteratorInterface
		result2 error
	}{result1, result2}
}

func (fake *ChaincodeStub) GetStateByRangeReturnsOnCall(i int, result1 shim.StateQueryIteratorInterface, result2 error) {
	fake.getStateByRangeMutex.Lock()
	defer fake.getStateByRangeMutex.Unlock()
	fake.GetStateByRangeStub = nil
	if fake.getStateByRangeReturnsOnCall == nil {
		fake.getStateByRangeReturnsOnCall = make(map[int]struct {
			result1 shim.StateQueryIteratorInterface
			result2 error
		})
	}
	fake.getStateByRangeReturnsOnCall[i] = struct {
		result1 shim.StateQueryIteratorInterface
		result2 error
	}{result1, result2}
}

func (fake *ChaincodeStub) GetStateByRangeWithPagination(arg1 string, arg2 string, arg3 int32, arg4 string) (shim.StateQueryIteratorInterface, *peer.QueryResponseMetadata, error) {
	fake.getStateByRangeWithPaginationMutex.Lock()
	ret, specificReturn := fake.getStateByRangeWithPaginationReturnsOnCall[len(fake.getStateByRangeWithPaginationArgsForCall)]
	fake.getStateByRangeWithPaginationArgsForCall = append(fake.getStateByRangeWithPaginationArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 int32
		arg4 string
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("GetStateByRangeWithPagination", []interface{}{arg1, arg2, arg3, arg4})
	fake.getStateByRangeWithPaginationMutex.Unlock()
	if fake.GetStateByRangeWithPaginationStub != nil {
		return fake.GetStateByRangeWithPaginationStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getStateByRangeWithPaginationReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *ChaincodeStub) GetStateByRangeWithPaginationCallCount() int {
	fake.getStateByRangeWithPaginationMutex.RLock()
	defer fake.getStateByRangeWithPaginationMutex.RUnlock()
	return len(fake.getStateByRangeWithPaginationArgsForCall)
}

func (fake *ChaincodeStub) GetStateByRangeWithPaginationCalls(stub func(string, string, int32, string) (shim.StateQueryIteratorInterface, *peer.QueryResponseMetadata, error)) {
	fake.getStateByRangeWithPaginationMutex.Lock()
	defer fake.getStateByRangeWithPaginationMutex.Unlock()
	fake.GetStateByRangeWithPaginationStub = stub
}

func (fake *ChaincodeStub) GetStateByRangeWithPaginationArgsForCall(i int) (string, string, int32, string) {
	fake.getStateByRangeWithPaginationMutex.RLock()
	defer fake.getStateByRangeWithPaginationMutex.RUnlock()
	argsForCall := fake.getStateByRangeWithPaginationArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *ChaincodeStub) GetStateByRangeWithPaginationReturns(result1 shim.StateQueryIteratorInterface, result2 *peer.QueryResponseMetadata, result3 error) {
	fake.getStateByRangeWithPaginationMutex.Lock()
	defer fake.getStateByRangeWithPaginationMutex.Unlock()
	fake.GetStateByRangeWithPaginationStub = nil
	fake.getStateByRangeWithPaginationReturns = struct {
		result1 shim.StateQueryIteratorInterface
		result2 *peer.QueryResponseMetadata
		result3 error
	}{result1, result2, result3}
}

func (fake *ChaincodeStub) GetStateByRangeWithPaginationReturnsOnCall(i int, result1 shim.StateQueryIteratorInterface, result2 *peer.QueryResponseMetadata, result3 error) {
	fake.getStateByRangeWithPaginationMutex.Lock()
	defer fake.getStateByRangeWithPaginationMutex.Unlock()
	fake.GetStateByRangeWithPaginationStub = nil
	if fake.getStateByRangeWithPaginationReturnsOnCall == nil {
		fake.getStateByRangeWithPaginationReturnsOnCall = make(map[int]struct {
			result1 shim.StateQueryIteratorInterface
			result2 *peer.QueryResponseMetadata
			result3 error
		})
	}
	fake.getStateByRangeWithPaginationReturnsOnCall[i] = struct {
		result1 shim.StateQueryIteratorInterface
		result2 *peer.QueryResponseMetadata
		result3 error
	}{result1, result2, result3}
}

func (fake *ChaincodeStub) GetStateValidationParameter(arg1 string) ([]byte, error) {
	fake.getStateValidationParameterMutex.Lock()
	ret, specificReturn := fake.getStateValidationParameterReturnsOnCall[len(fake.getStateValidationParameterArgsForCall)]
	fake.getStateValidationParameterArgsForCall = append(fake.getStateValidationParameterArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetStateValidationParameter", []interface{}{arg1})
	fake.getStateValidationParameterMutex.Unlock()
	if fake.GetStateValidationParameterStub != nil {
		return fake.GetStateValidationParameterStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getStateValidationParameterReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *ChaincodeStub) GetStateValidationParameterCallCount() int {
	fake.getStateValidationParameterMutex.RLock()
	defer fake.getStateValidationParameterMutex.RUnlock()
	return len(fake.getStateValidationParameterArgsForCall)
}

func (fake *ChaincodeStub) GetStateValidationParameterCalls(stub func(string) ([]byte, error)) {
	fake.getStateValidationParameterMutex.Lock()
	defer fake.getStateValidationParameterMutex.Unlock()
	fake.GetStateValidationParameterStub = stub
}

func (fake *ChaincodeStub) GetStateValidationParameterArgsForCall(i int) string {
	fake.getStateValidationParameterMutex.RLock()
	defer fake.getStateValidationParameterMutex.RUnlock()
	argsForCall := fake.getStateValidationParameterArgsForCall[i]
	return argsForCall.arg1
}

func (fake *ChaincodeStub) GetStateValidationParameterReturns(result1 []byte, result2 error) {
	fake.getStateValidationParameterMutex.Lock()
	defer fake.getStateValidationParameterMutex.Unlock()
	fake.GetStateValidationParameterStub = nil
	fake.getStateValidationParameterReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *ChaincodeStub) GetStateValidationParameterReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.getStateValidationParameterMutex.Lock()
	defer fake.getStateValidationParameterMutex.Unlock()
	fake.GetStateValidationParameterStub = nil
	if fake.getStateValidationParameterReturnsOnCall == nil {
		fake.getStateValidationParameterReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.getStateValidationParameterReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *ChaincodeStub) GetStringArgs() []string {
	fake.getStringArgsMutex.Lock()
	ret, specificReturn := fake.getStringArgsReturnsOnCall[len(fake.getStringArgsArgsForCall)]
	fake.getStringArgsArgsForCall = append(fake.getStringArgsArgsForCall, struct {
	}{})
	fake.recordInvocation("GetStringArgs", []interface{}{})
	fake.getStringArgsMutex.Unlock()
	if fake.GetStringArgsStub != nil {
		return fake.GetStringArgsStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.getStringArgsReturns
	return fakeReturns.result1
}

func (fake *ChaincodeStub) GetStringArgsCallCount() int {
	fake.getStringArgsMutex.RLock()
	defer fake.getStringArgsMutex.RUnlock()
	return len(fake.getStringArgsArgsForCall)
}

func (fake *ChaincodeStub) GetStringArgsCalls(stub func() []string) {
	fake.getStringArgsMutex.Lock()
	defer fake.getStringArgsMutex.Unlock()
	fake.GetStringArgsStub = stub
}

func (fake *ChaincodeStub) GetStringArgsReturns(result1 []string) {
	fake.getStringArgsMutex.Lock()
	defer fake.getStringArgsMutex.Unlock()
	fake.GetStringArgsStub = nil
	fake.getStringArgsReturns = struct {
		result1 []string
	}{result1}
}

func (fake *ChaincodeStub) GetStringArgsReturnsOnCall(i int, result1 []string) {
	fake.getStringArgsMutex.Lock()
	defer fake.getStringArgsMutex.Unlock()
	fake.GetStringArgsStub = nil
	if fake.getStringArgsReturnsOnCall == nil {
		fake.getStringArgsReturnsOnCall = make(map[int]struct {
			result1 []string
		})
	}
	fake.getStringArgsReturnsOnCall[i] = struct {
		result1 []string
	}{result1}
}

func (fake *ChaincodeStub) GetTransient() (map[string][]byte, error) {
	fake.getTransientMutex.Lock()
	ret, specificReturn := fake.getTransientReturnsOnCall[len(fake.getTransientArgsForCall)]
	fake.getTransientArgsForCall = append(fake.getTransientArgsForCall, struct {
	}{})
	fake.recordInvocation("GetTransient", []interface{}{})
	fake.getTransientMutex.Unlock()
	if fake.GetTransientStub != nil {
		return fake.GetTransientStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getTransientReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *ChaincodeStub) GetTransientCallCount() int {
	fake.getTransientMutex.RLock()
	defer fake.getTransientMutex.RUnlock()
	return len(fake.getTransientArgsForCall)
}

func (fake *ChaincodeStub) GetTransientCalls(stub func() (map[string][]byte, error)) {
	fake.getTransientMutex.Lock()
	defer fake.getTransientMutex.Unlock()
	fake.GetTransientStub = stub
}

func (fake *ChaincodeStub) GetTransientReturns(result1 map[string][]byte, result2 error) {
	fake.getTransientMutex.Lock()
	defer fake.getTransientMutex.Unlock()
	fake.GetTransientStub = nil
	fake.getTransientReturns = struct {
		result1 map[string][]byte
		result2 error
	}{result1, result2}
}

func (fake *ChaincodeStub) GetTransientReturnsOnCall(i int, result1 map[string][]byte, result2 error) {
	fake.getTransientMutex.Lock()
	defer fake.getTransientMutex.Unlock()
	fake.GetTransientStub = nil
	if fake.getTransientReturnsOnCall == nil {
		fake.getTransientReturnsOnCall = make(map[int]struct {
			result1 map[string][]byte
			result2 error
		})
	}
	fake.getTransientReturnsOnCall[i] = struct {
		result1 map[string][]byte
		result2 error
	}{result1, result2}
}

func (fake *ChaincodeStub) GetTxID() string {
	fake.getTxIDMutex.Lock()
	ret, specificReturn := fake.getTxIDReturnsOnCall[len(fake.getTxIDArgsForCall)]
	fake.getTxIDArgsForCall = append(fake.getTxIDArgsForCall, struct {
	}{})
	fake.recordInvocation("GetTxID", []interface{}{})
	fake.getTxIDMutex.Unlock()
	if fake.GetTxIDStub != nil {
		return fake.GetTxIDStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.getTxIDReturns
	return fakeReturns.result1
}

func (fake *ChaincodeStub) GetTxIDCallCount() int {
	fake.getTxIDMutex.RLock()
	defer fake.getTxIDMutex.RUnlock()
	return len(fake.getTxIDArgsForCall)
}

func (fake *ChaincodeStub) GetTxIDCalls(stub func() string) {
	fake.getTxIDMutex.Lock()
	defer fake.getTxIDMutex.Unlock()
	fake.GetTxIDStub = stub
}

func (fake *ChaincodeStub) GetTxIDReturns(result1 string) {
	fake.getTxIDMutex.Lock()
	defer fake.getTxIDMutex.Unlock()
	fake.GetTxIDStub = nil
	fake.getTxIDReturns = struct {
		result1 string
	}{result1}
}

func (fake *ChaincodeStub) GetTxIDReturnsOnCall(i int, result1 string) {
	fake.getTxIDMutex.Lock()
	defer fake.getTxIDMutex.Unlock()
	fake.GetTxIDStub = nil
	if fake.getTxIDReturnsOnCall == nil {
		fake.getTxIDReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.getTxIDReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *ChaincodeStub) GetTxTimestamp() (*timestamp.Timestamp, error) {
	fake.getTxTimestampMutex.Lock()
	ret, specificReturn := fake.getTxTimestampReturnsOnCall[len(fake.getTxTimestampArgsForCall)]
	fake.getTxTimestampArgsForCall = append(fake.getTxTimestampArgsForCall, struct {
	}{})
	fake.recordInvocation("GetTxTimestamp", []interface{}{})
	fake.getTxTimestampMutex.Unlock()
	if fake.GetTxTimestampStub != nil {
		return fake.GetTxTimestampStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getTxTimestampReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *ChaincodeStub) GetTxTimestampCallCount() int {
	fake.getTxTimestampMutex.RLock()
	defer fake.getTxTimestampMutex.RUnlock()
	return len(fake.getTxTimestampArgsForCall)
}

func (fake *ChaincodeStub) GetTxTimestampCalls(stub func() (*timestamp.Timestamp, error)) {
	fake.getTxTimestampMutex.Lock()
	defer fake.getTxTimestampMutex.Unlock()
	fake.GetTxTimestampStub = stub
}

func (fake *ChaincodeStub) GetTxTimestampReturns(result1 *timestamp.Timestamp, result2 error) {
	fake.getTxTimestampMutex.Lock()
	defer fake.getTxTimestampMutex.Unlock()
	fake.GetTxTimestampStub = nil
	fake.getTxTimestampReturns = struct {
		result1 *timestamp.Timestamp
		result2 error
	}{result1, result2}
}

func (fake *ChaincodeStub) GetTxTimestampReturnsOnCall(i int, result1 *timestamp.Timestamp, result2 error) {
	fake.getTxTimestampMutex.Lock()
	defer fake.getTxTimestampMutex.Unlock()
	fake.GetTxTimestampStub = nil
	if fake.getTxTimestampReturnsOnCall == nil {
		fake.getTxTimestampReturnsOnCall = make(map[int]struct {
			result1 *timestamp.Timestamp
			result2 error
		})
	}
	fake.getTxTimestampReturnsOnCall[i] = struct {
		result1 *timestamp.Timestamp
		result2 error
	}{result1, result2}
}

func (fake *ChaincodeStub) InvokeChaincode(arg1 string, arg2 [][]byte, arg3 string) peer.Response {
	var arg2Copy [][]byte
	if arg2 != nil {
		arg2Copy = make([][]byte, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.invokeChaincodeMutex.Lock()
	ret, specificReturn := fake.invokeChaincodeReturnsOnCall[len(fake.invokeChaincodeArgsForCall)]
	fake.invokeChaincodeArgsForCall = append(fake.invokeChaincodeArgsForCall, struct {
		arg1 string
		arg2 [][]byte
		arg3 string
	}{arg1, arg2Copy, arg3})
	fake.recordInvocation("InvokeChaincode", []interface{}{arg1, arg2Copy, arg3})
	fake.invokeChaincodeMutex.Unlock()
	if fake.InvokeChaincodeStub != nil {
		return fake.InvokeChaincodeStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.invokeChaincodeReturns
	return fakeReturns.result1
}

func (fake *ChaincodeStub) InvokeChaincodeCallCount() int {
	fake.invokeChaincodeMutex.RLock()
	defer fake.invokeChaincodeMutex.RUnlock()
	return len(fake.invokeChaincodeArgsForCall)
}

func (fake *ChaincodeStub) InvokeChaincodeCalls(stub func(string, [][]byte, string) peer.Response) {
	fake.invokeChaincodeMutex.Lock()
	defer fake.invokeChaincodeMutex.Unlock()
	fake.InvokeChaincodeStub = stub
}

func (fake *ChaincodeStub) InvokeChaincodeArgsForCall(i int) (string, [][]byte, string) {
	fake.invokeChaincodeMutex.RLock()
	defer fake.invokeChaincodeMutex.RUnlock()
	argsForCall := fake.invokeChaincodeArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *ChaincodeStub) InvokeChaincodeReturns(result1 peer.Response) {
	fake.invokeChaincodeMutex.Lock()
	defer fake.invokeChaincodeMutex.Unlock()
	fake.InvokeChaincodeStub = nil
	fake.invokeChaincodeReturns = struct {
		result1 peer.Response
	}{result1}
}

func (fake *ChaincodeStub) InvokeChaincodeReturnsOnCall(i int, result1 peer.Response) {
	fake.invokeChaincodeMutex.Lock()
	defer fake.invokeChaincodeMutex.Unlock()
	fake.InvokeChaincodeStub = nil
	if fake.invokeChaincodeReturnsOnCall == nil {
		fake.invokeChaincodeReturnsOnCall = make(map[int]struct {
			result1 peer.Response
		})
	}
	fake.invokeChaincodeReturnsOnCall[i] = struct {
		result1 peer.Response
	}{result1}
}

func (fake *ChaincodeStub) PurgePrivateData(arg1 string, arg2 string) error {
	fake.purgePrivateDataMutex.Lock()
	ret, specificReturn := fake.purgePrivateDataReturnsOnCall[len(fake.purgePrivateDataArgsForCall)]
	fake.purgePrivateDataArgsForCall = append(fake.purgePrivateDataArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("PurgePrivateData", []interface{}{arg1, arg2})
	fake.purgePrivateDataMutex.Unlock()
	if fake.PurgePrivateDataStub != nil {
		return fake.PurgePrivateDataStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.purgePrivateDataReturns
	return fakeReturns.result1
}

func (fake *ChaincodeStub) PurgePrivateDataCallCount() int {
	fake.purgePrivateDataMutex.RLock()
	defer fake.purgePrivateDataMutex.RUnlock()
	return len(fake.purgePrivateDataArgsForCall)
}

func (fake *ChaincodeStub) PurgePrivateDataCalls(stub func(string, string) error) {
	fake.purgePrivateDataMutex.Lock()
	defer fake.purgePrivateDataMutex.Unlock()
	fake.PurgePrivateDataStub = stub
}

func (fake *ChaincodeStub) PurgePrivateDataArgsForCall(i int) (string, string) {
	fake.purgePrivateDataMutex.RLock()
	defer fake.purgePrivateDataMutex.RUnlock()
	argsForCall := fake.purgePrivateDataArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *ChaincodeStub) PurgePrivateDataReturns(result1 error) {
	fake.purgePrivateDataMutex.Lock()
	defer fake.purgePrivateDataMutex.Unlock()
	fake.PurgePrivateDataStub = nil
	fake.purgePrivateDataReturns = struct {
		result1 error
	}{result1}
}

func (fake *ChaincodeStub) PurgePrivateDataReturnsOnCall(i int, result1 error) {
	fake.purgePrivateDataMutex.Lock()
	defer fake.purgePrivateDataMutex.Unlock()
	fake.PurgePrivateDataStub = nil
	if fake.purgePrivateDataReturnsOnCall == nil {
		fake.purgePrivateDataReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.purgePrivateDataReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *ChaincodeStub) PutPrivateData(arg1 string, arg2 string, arg3 []byte) error {
	var arg3Copy []byte
	if arg3 != nil {
		arg3Copy = make([]byte, len(arg3))
		copy(arg3Copy, arg3)
	}
	fake.putPrivateDataMutex.Lock()
	ret, specificReturn := fake.putPrivateDataReturnsOnCall[len(fake.putPrivateDataArgsForCall)]
	fake.putPrivateDataArgsForCall = append(fake.putPrivateDataArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 []byte
	}{arg1, arg2, arg3Copy})
	fake.recordInvocation("PutPrivateData", []interface{}{arg1, arg2, arg3Copy})
	fake.putPrivateDataMutex.Unlock()
	if fake.PutPrivateDataStub != nil {
		return fake.PutPrivateDataStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.putPrivateDataReturns
	return fakeReturns.result1
}

func (fake *ChaincodeStub) PutPrivateDataCallCount() int {
	fake.putPrivateDataMutex.RLock()
	defer fake.putPrivateDataMutex.RUnlock()
	return len(fake.putPrivateDataArgsForCall)
}

func (fake *ChaincodeStub) PutPrivateDataCalls(stub func(string, string, []byte) error) {
	fake.putPrivateDataMutex.Lock()
	defer fake.putPrivateDataMutex.Unlock()
	fake.PutPrivateDataStub = stub
}

func (fake *ChaincodeStub) PutPrivateDataArgsForCall(i int) (string, string, []byte) {
	fake.putPrivateDataMutex.RLock()
	defer fake.putPrivateDataMutex.RUnlock()
	argsForCall := fake.putPrivateDataArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *ChaincodeStub) PutPrivateDataReturns(result1 error) {
	fake.putPrivateDataMutex.Lock()
	defer fake.putPrivateDataMutex.Unlock()
	fake.PutPrivateDataStub = nil
	fake.putPrivateDataReturns = struct {
		result1 error
	}{result1}
}

func (fake *ChaincodeStub) PutPrivateDataReturnsOnCall(i int, result1 error) {
	fake.putPrivateDataMutex.Lock()
	defer fake.putPrivateDataMutex.Unlock()
	fake.PutPrivateDataStub = nil
	if fake.putPrivateDataReturnsOnCall == nil {
		fake.putPrivateDataReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.putPrivateDataReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *ChaincodeStub) PutState(arg1 string, arg2 []byte) error {
	var arg2Copy []byte
	if arg2 != nil {
		arg2Copy = make([]byte, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.putStateMutex.Lock()
	ret, specificReturn := fake.putStateReturnsOnCall[len(fake.putStateArgsForCall)]
	fake.putStateArgsForCall = append(fake.putStateArgsForCall, struct {
		arg1 string
		arg2 []byte
	}{arg1, arg2Copy})
	fake.recordInvocation("PutState", []interface{}{arg1, arg2Copy})
	fake.putStateMutex.Unlock()
	if fake.PutStateStub != nil {
		return fake.PutStateStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.putStateReturns
	return fakeReturns.result1
}

func (fake *ChaincodeStub) PutStateCallCount() int {
	fake.putStateMutex.RLock()
	defer fake.putStateMutex.RUnlock()
	return len(fake.putStateArgsForCall)
}

func (fake *ChaincodeStub) PutStateCalls(stub func(string, []byte) error) {
	fake.putStateMutex.Lock()
	defer fake.putStateMutex.Unlock()
	fake.PutStateStub = stub
}

func (fake *ChaincodeStub) PutStateArgsForCall(i int) (string, []byte) {
	fake.putStateMutex.RLock()
	defer fake.putStateMutex.RUnlock()
	argsForCall := fake.putStateArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *ChaincodeStub) PutStateReturns(result1 error) {
	fake.putStateMutex.Lock()
	defer fake.putStateMutex.Unlock()
	fake.PutStateStub = nil
	fake.putStateReturns = struct {
		result1 error
	}{result1}
}

func (fake *ChaincodeStub) PutStateReturnsOnCall(i int, result1 error) {
	fake.putStateMutex.Lock()
	defer fake.putStateMutex.Unlock()
	fake.PutStateStub = nil
	if fake.putStateReturnsOnCall == nil {
		fake.putStateReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.putStateReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *ChaincodeStub) SetEvent(arg1 string, arg2 []byte) error {
	var arg2Copy []byte
	if arg2 != nil {
		arg2Copy = make([]byte, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.setEventMutex.Lock()
	ret, specificReturn := fake.setEventReturnsOnCall[len(fake.setEventArgsForCall)]
	fake.setEventArgsForCall = append(fake.setEventArgsForCall, struct {
		arg1 string
		arg2 []byte
	}{arg1, arg2Copy})
	fake.recordInvocation("SetEvent", []interface{}{arg1, arg2Copy})
	fake.setEventMutex.Unlock()
	if fake.SetEventStub != nil {
		return fake.SetEventStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.setEventReturns
	return fakeReturns.result1
}

func (fake *ChaincodeStub) SetEventCallCount() int {
	fake.setEventMutex.RLock()
	defer fake.setEventMutex.RUnlock()
	return len(fake.setEventArgsForCall)
}

func (fake *ChaincodeStub) SetEventCalls(stub func(string, []byte) error) {
	fake.setEventMutex.Lock()
	defer fake.setEventMutex.Unlock()
	fake.SetEventStub = stub
}

func (fake *ChaincodeStub) SetEventArgsForCall(i int) (string, []byte) {
	fake.setEventMutex.RLock()
	defer fake.setEventMutex.RUnlock()
	argsForCall := fake.setEventArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *ChaincodeStub) SetEventReturns(result1 error) {
	fake.setEventMutex.Lock()
	defer fake.setEventMutex.Unlock()
	fake.SetEventStub = nil
	fake.setEventReturns = struct {
		result1 error
	}{result1}
}

func (fake *ChaincodeStub) SetEventReturnsOnCall(i int, result1 error) {
	fake.setEventMutex.Lock()
	defer fake.setEventMutex.Unlock()
	fake.SetEventStub = nil
	if fake.setEventReturnsOnCall == nil {
		fake.setEventReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.setEventReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *ChaincodeStub) SetPrivateDataValidationParameter(arg1 string, arg2 string, arg3 []byte) error {
	var arg3Copy []byte
	if arg3 != nil {
		arg3Copy = make([]byte, len(arg3))
		copy(arg3Copy, arg3)
	}
	fake.setPrivateDataValidationParameterMutex.Lock()
	ret, specificReturn := fake.setPrivateDataValidationParameterReturnsOnCall[len(fake.setPrivateDataValidationParameterArgsForCall)]
	fake.setPrivateDataValidationParameterArgsForCall = append(fake.setPrivateDataValidationParameterArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 []byte
	}{arg1, arg2, arg3Copy})
	fake.recordInvocation("SetPrivateDataValidationParameter", []interface{}{arg1, arg2, arg3Copy})
	fake.setPrivateDataValidationParameterMutex.Unlock()
	if fake.SetPrivateDataValidationParameterStub != nil {
		return fake.SetPrivateDataValidationParameterStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.setPrivateDataValidationParameterReturns
	return fakeReturns.result1
}

func (fake *ChaincodeStub) SetPrivateDataValidationParameterCallCount() int {
	fake.setPrivateDataValidationParameterMutex.RLock()
	defer fake.setPrivateDataValidationParameterMutex.RUnlock()
	return len(fake.setPrivateDataValidationParameterArgsForCall)
}

func (fake *ChaincodeStub) SetPrivateDataValidationParameterCalls(stub func(string, string, []byte) error) {
	fake.setPrivateDataValidationParameterMutex.Lock()
	defer fake.setPrivateDataValidationParameterMutex.Unlock()
	fake.SetPrivateDataValidationParameterStub = stub
}

func (fake *ChaincodeStub) SetPrivateDataValidationParameterArgsForCall(i int) (string, string, []byte) {
	fake.setPrivateDataValidationParameterMutex.RLock()
	defer fake.setPrivateDataValidationParameterMutex.RUnlock()
	argsForCall := fake.setPrivateDataValidationParameterArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *ChaincodeStub) SetPrivateDataValidationParameterReturns(result1 error) {
	fake.setPrivateDataValidationParameterMutex.Lock()
	defer fake.setPrivateDataValidationParameterMutex.Unlock()
	fake.SetPrivateDataValidationParameterStub = nil
	fake.setPrivateDataValidationParameterReturns = struct {
		result1 error
	}{result1}
}

func (fake *ChaincodeStub) SetPrivateDataValidationParameterReturnsOnCall(i int, result1 error) {
	fake.setPrivateDataValidationParameterMutex.Lock()
	defer fake.setPrivateDataValidationParameterMutex.Unlock()
	fake.SetPrivateDataValidationParameterStub = nil
	if fake.setPrivateDataValidationParameterReturnsOnCall == nil {
		fake.setPrivateDataValidationParameterReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.setPrivateDataValidationParameterReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *ChaincodeStub) SetStateValidationParameter(arg1 string, arg2 []byte) error {
	var arg2Copy []byte
	if arg2 != nil {
		arg2Copy = make([]byte, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.setStateValidationParameterMutex.Lock()
	ret, specificReturn := fake.setStateValidationParameterReturnsOnCall[len(fake.setStateValidationParameterArgsForCall)]
	fake.setStateValidationParameterArgsForCall = append(fake.setStateValidationParameterArgsForCall, struct {
		arg1 string
		arg2 []byte
	}{arg1, arg2Copy})
	fake.recordInvocation("SetStateValidationParameter", []interface{}{arg1, arg2Copy})
	fake.setStateValidationParameterMutex.Unlock()
	if fake.SetStateValidationParameterStub != nil {
		return fake.SetStateValidationParameterStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.setStateValidationParameterReturns
	return fakeReturns.result1
}

func (fake *ChaincodeStub) SetStateValidationParameterCallCount() int {
	fake.setStateValidationParameterMutex.RLock()
	defer fake.setStateValidationParameterMutex.RUnlock()
	return len(fake.setStateValidationParameterArgsForCall)
}

func (fake *ChaincodeStub) SetStateValidationParameterCalls(stub func(string, []byte) error) {
	fake.setStateValidationParameterMutex.Lock()
	defer fake.setStateValidationParameterMutex.Unlock()
	fake.SetStateValidationParameterStub = stub
}

func (fake *ChaincodeStub) SetStateValidationParameterArgsForCall(i int) (string, []byte) {
	fake.setStateValidationParameterMutex.RLock()
	defer fake.setStateValidationParameterMutex.RUnlock()
	argsForCall := fake.setStateValidationParameterArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *ChaincodeStub) SetStateValidationParameterReturns(result1 error) {
	fake.setStateValidationParameterMutex.Lock()
	defer fake.setStateValidationParameterMutex.Unlock()
	fake.SetStateValidationParameterStub = nil
	fake.setStateValidationParameterReturns = struct {
		result1 error
	}{result1}
}

func (fake *ChaincodeStub) SetStateValidationParameterReturnsOnCall(i int, result1 error) {
	fake.setStateValidationParameterMutex.Lock()
	defer fake.setStateValidationParameterMutex.Unlock()
	fake.SetStateValidationParameterStub = nil
	if fake.setStateValidationParameterReturnsOnCall == nil {
		fake.setStateValidationParameterReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.setStateValidationParameterReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *ChaincodeStub) SplitCompositeKey(arg1 string) (string, []string, error) {
	fake.splitCompositeKeyMutex.Lock()
	ret, specificReturn := fake.splitCompositeKeyReturnsOnCall[len(fake.splitCompositeKeyArgsForCall)]
	fake.splitCompositeKeyArgsForCall = append(fake.splitCompositeKeyArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("SplitCompositeKey", []interface{}{arg1})
	fake.splitCompositeKeyMutex.Unlock()
	if fake.SplitCompositeKeyStub != nil {
		return fake.SplitCompositeKeyStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.splitCompositeKeyReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *ChaincodeStub) SplitCompositeKeyCallCount() int {
	fake.splitCompositeKeyMutex.RLock()
	defer fake.splitCompositeKeyMutex.RUnlock()
	return len(fake.splitCompositeKeyArgsForCall)
}

func (fake *ChaincodeStub) SplitCompositeKeyCalls(stub func(string) (string, []string, error)) {
	fake.splitCompositeKeyMutex.Lock()
	defer fake.splitCompositeKeyMutex.Unlock()
	fake.SplitCompositeKeyStub = stub
}

func (fake *ChaincodeStub) SplitCompositeKeyArgsForCall(i int) string {
	fake.splitCompositeKeyMutex.RLock()
	defer fake.splitCompositeKeyMutex.RUnlock()
	argsForCall := fake.splitCompositeKeyArgsForCall[i]
	return argsForCall.arg1
}

func (fake *ChaincodeStub) SplitCompositeKeyReturns(result1 string, result2 []string, result3 error) {
	fake.splitCompositeKeyMutex.Lock()
	defer fake.splitCompositeKeyMutex.Unlock()
	fake.SplitCompositeKeyStub = nil
	fake.splitCompositeKeyReturns = struct {
		result1 string
		result2 []string
		result3 error
	}{result1, result2, result3}
}

func (fake *ChaincodeStub) SplitCompositeKeyReturnsOnCall(i int, result1 string, result2 []string, result3 error) {
	fake.splitCompositeKeyMutex.Lock()
	defer fake.splitCompositeKeyMutex.Unlock()
	fake.SplitCompositeKeyStub = nil
	if fake.splitCompositeKeyReturnsOnCall == nil {
		fake.splitCompositeKeyReturnsOnCall = make(map[int]struct {
			result1 string
			result2 []string
			result3 error
		})
	}
	fake.splitCompositeKeyReturnsOnCall[i] = struct {
		result1 string
		result2 []string
		result3 error
	}{result1, result2, result3}
}

func (fake *ChaincodeStub) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createCompositeKeyMutex.RLock()
	defer fake.createCompositeKeyMutex.RUnlock()
	fake.delPrivateDataMutex.RLock()
	defer fake.delPrivateDataMutex.RUnlock()
	fake.delStateMutex.RLock()
	defer fake.delStateMutex.RUnlock()
	fake.getArgsMutex.RLock()
	defer fake.getArgsMutex.RUnlock()
	fake.getArgsSliceMutex.RLock()
	defer fake.getArgsSliceMutex.RUnlock()
	fake.getBindingMutex.RLock()
	defer fake.getBindingMutex.RUnlock()
	fake.getChannelIDMutex.RLock()
	defer fake.getChannelIDMutex.RUnlock()
	fake.getCreatorMutex.RLock()
	defer fake.getCreatorMutex.RUnlock()
	fake.getDecorationsMutex.RLock()
	defer fake.getDecorationsMutex.RUnlock()
	fake.getFunctionAndParametersMutex.RLock()
	defer fake.getFunctionAndParametersMutex.RUnlock()
	fake.getHistoryForKeyMutex.RLock()
	defer fake.getHistoryForKeyMutex.RUnlock()
	fake.getPrivateDataMutex.RLock()
	defer fake.getPrivateDataMutex.RUnlock()
	fake.getPrivateDataByPartialCompositeKeyMutex.RLock()
	defer fake.getPrivateDataByPartialCompositeKeyMutex.RUnlock()
	fake.getPrivateDataByRangeMutex.RLock()
	defer fake.getPrivateDataByRangeMutex.RUnlock()
	fake.getPrivateDataHashMutex.RLock()
	defer fake.getPrivateDataHashMutex.RUnlock()
	fake.getPrivateDataQueryResultMutex.RLock()
	defer fake.getPrivateDataQueryResultMutex.RUnlock()
	fake.getPrivateDataValidationParameterMutex.RLock()
	defer fake.getPrivateDataValidationParameterMutex.RUnlock()
	fake.getQueryResultMutex.RLock()
	defer fake.getQueryResultMutex.RUnlock()
	fake.getQueryResultWithPaginationMutex.RLock()
	defer fake.getQueryResultWithPaginationMutex.RUnlock()
	fake.getSignedProposalMutex.RLock()
	defer fake.getSignedProposalMutex.RUnlock()
	fake.getStateMutex.RLock()
	defer fake.getStateMutex.RUnlock()
	fake.getStateByPartialCompositeKeyMutex.RLock()
	defer fake.getStateByPartialCompositeKeyMutex.RUnlock()
	fake.getStateByPartialCompositeKeyWithPaginationMutex.RLock()
	defer fake.getStateByPartialCompositeKeyWithPaginationMutex.RUnlock()
	fake.getStateByRangeMutex.RLock()
	defer fake.getStateByRangeMutex.RUnlock()
	fake.getStateByRangeWithPaginationMutex.RLock()
	defer fake.getStateByRangeWithPaginationMutex.RUnlock()
	fake.getStateValidationParameterMutex.RLock()
	defer fake.getStateValidationParameterMutex.RUnlock()
	fake.getStringArgsMutex.RLock()
	defer fake.getStringArgsMutex.RUnlock()
	fake.getTransientMutex.RLock()
	defer fake.getTransientMutex.RUnlock()
	fake.getTxIDMutex.RLock()
	defer fake.getTxIDMutex.RUnlock()
	fake.getTxTimestampMutex.RLock()
	defer fake.getTxTimestampMutex.RUnlock()
	fake.invokeChaincodeMutex.RLock()
	defer fake.invokeChaincodeMutex.RUnlock()
	fake.purgePrivateDataMutex.RLock()
	defer fake.purgePrivateDataMutex.RUnlock()
	fake.putPrivateDataMutex.RLock()
	defer fake.putPrivateDataMutex.RUnlock()
	fake.putStateMutex.RLock()
	defer fake.putStateMutex.RUnlock()
	fake.setEventMutex.RLock()
	defer fake.setEventMutex.RUnlock()
	fake.setPrivateDataValidationParameterMutex.RLock()
	defer fake.setPrivateDataValidationParameterMutex.RUnlock()
	fake.setStateValidationParameterMutex.RLock()
	defer fake.setStateValidationParameterMutex.RUnlock()
	fake.splitCompositeKeyMutex.RLock()
	defer fake.splitCompositeKeyMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *ChaincodeStub) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}
