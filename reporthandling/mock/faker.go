package mock

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/armosec/armoapi-go/armotypes"
	"github.com/go-faker/faker/v4"
	ik8s "github.com/kubescape/k8s-interface/workloadinterface"
	mockoptions "github.com/go-faker/faker/v4/pkg/options"
)

var mockOpts []mockoptions.OptionFunc

func init() {
	registerFakers()
}

func registerFakers() {
	faker.SetRandomSource(faker.NewSafeSource(rand.NewSource(time.Now().UnixNano()))) //nolint:nosec
	log.Println("initializing faker")
	mockOpts = []mockoptions.OptionFunc{
		mockoptions.WithIgnoreInterface(true),
		mockoptions.WithCustomFieldProvider("PortalBase", mockPortalBase),
		mockoptions.WithCustomFieldProvider("IMetadata", mockIMetadata),
		mockoptions.WithNilIfLenIsZero(true), // to assert omitempty behavior
	}
	_ = mockoptions.SetRandomMapAndSliceMinSize(2)
	_ = mockoptions.SetRandomMapAndSliceMaxSize(5)
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

func mockIMetadata() (interface{}, error) {
	return ik8s.NewWorkloadMock(nil), nil
}

func doOrDieTrying(err error) {
	if err != nil {
		panic(fmt.Sprintf("failed to build faker: %v", err))
	}
}
