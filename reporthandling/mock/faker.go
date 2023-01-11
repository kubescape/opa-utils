package mock

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/armosec/armoapi-go/armotypes"
	"github.com/go-faker/faker/v4"
	mockoptions "github.com/go-faker/faker/v4/pkg/options"
)

var mockOpts []mockoptions.OptionFunc

func init() {
	registerFakers()
}

func registerFakers() {
	faker.SetRandomSource(faker.NewSafeSource(rand.NewSource(time.Now().UnixNano()))) //nolint:nosec
	mockOpts = []mockoptions.OptionFunc{
		mockoptions.WithIgnoreInterface(true),
		mockoptions.WithCustomFieldProvider("PortalBase", mockPortalBase),
		mockoptions.WithNilIfLenIsZero(true), // to assert omitempty behavior
	}
}

func MockData[T any]() T {
	var receiver T
	doOrDieTrying(faker.FakeData(&receiver, mockOpts...))

	return receiver
}

func mockPortalBase() (interface{}, error) {
	base := MockData[armotypes.PortalBase]()
	for k := range base.Attributes {
		base.Attributes[k] = faker.Name()
	}

	return base, nil
}

func doOrDieTrying(err error) {
	if err != nil {
		panic(fmt.Sprintf("failed to build faker: %v", err))
	}
}
