package service

import (
	"fmt"

	"github.com/Azure/open-service-broker-azure/pkg/crypto/noop"
)

type ArbitraryType struct {
	Foo string `json:"foo"`
}

const fooValue = "bar"

var (
	testArbitraryObject = &ArbitraryType{
		Foo: fooValue,
	}
	testArbitraryObjectJSON = []byte(fmt.Sprintf(`{"foo":"%s"}`, fooValue))
	noopCodec               = noop.NewCodec()
)
