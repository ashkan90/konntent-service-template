//go:build integration
// +build integration

package listener_test

//
//import (
//	"context"
//	mcm "konntent-service-template/pkg/dummyclient/mocks"
//	rm "konntent-service-template/pkg/rabbit/mocks"
//	"testing"
//
//	"github.com/golang/mock/gomock"
//	. "github.com/onsi/ginkgo/v2"
//	. "github.com/onsi/gomega"
//)
//
//func TestCustomHandlerGroupIntegration(t *testing.T) {
//	RegisterFailHandler(Fail)
//	RunSpecs(t, "Consumer Integration Suite")
//}
//
//var (
//	ctx                 = context.Background()
//	mockController      = gomock.NewController(GinkgoT())
//	mobilisimClientMock = mcm.NewMockClient(mockController)
//	rabbitClientMock    = rm.NewMockClient(mockController)
//)
