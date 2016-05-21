package util

import (
	"strings"
	//"reflect"
	//"fmt"
	abi "github.com/ethereum/go-ethereum/accounts/abi"
)

func MakeAbi(Abi string) (abi.ABI, error) {
	evmABI, err := abi.JSON(strings.NewReader(Abi))
	if err != nil {
		return abi.ABI{}, err
	}
	return evmABI, nil
}

func Pack(abiInstance abi.ABI, method string, args ...[]byte) ([]byte, error){
	return abiInstance.Pack(method, args)
}

func Unpack(abiInstance abi.ABI, method string, output []byte, result interface{}) error {
	return abiInstance.Unpack(result, method, output)
}

/*func ConvertTypes(types interface{}) ([]byte, error) {
	s := reflect.ValueOf(types)
	if s.Kind() != 
}*/