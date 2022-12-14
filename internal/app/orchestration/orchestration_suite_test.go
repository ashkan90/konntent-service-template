package orchestration_test

//
//import (
//	"context"
//	msm "konntent-service-template/internal/app/mobilisim/mocks"
//	"konntent-service-template/internal/app/orchestration"
//	"konntent-service-template/pkg/claimer"
//	cm "konntent-service-template/pkg/claimer/mocks"
//	"konntent-service-template/pkg/event"
//	"konntent-service-template/pkg/utils"
//	"testing"
//
//	"github.com/brianvoe/gofakeit/v6"
//	"github.com/golang/mock/gomock"
//
//	. "github.com/onsi/ginkgo/v2"
//	. "github.com/onsi/gomega"
//)
//
//func TestMobilisimOrchestration(t *testing.T) {
//	RegisterFailHandler(Fail)
//	RunSpecs(t, "Mobilisim Orchestration Suite")
//}
//
//var (
//	ctx                       context.Context
//	mockCtrl                  *gomock.Controller
//	mobilisimEventServiceMock *msm.MockService
//	claimerMock               *cm.MockClaimer
//	orchestrator              orchestration.DummyOrchestrator
//
//	owner     *claimer.Model
//	headerMap event.HeaderMap
//)
//
//var _ = BeforeEach(func() {
//	ctx = context.Background()
//	mockCtrl = gomock.NewController(GinkgoT())
//	mobilisimEventServiceMock = msm.NewMockService(mockCtrl)
//	claimerMock = cm.NewMockClaimer(mockCtrl)
//	orchestrator = orchestration.NewDummyOrchestrator()
//
//	ctx = context.WithValue(ctx, utils.Claimer, claimerMock)
//	ctx = context.WithValue(ctx, utils.HeaderMapCtx, map[string]string{
//		utils.HeaderAuthorizationMobilisim: "api-key",
//	})
//
//	owner = &claimer.Model{
//		UserID: gofakeit.IntRange(1000, 9999),
//		Credit: gofakeit.IntRange(100, 500),
//	}
//	headerMap = event.HeaderMap{
//		utils.HeaderAuthorizationMobilisim: "api-key",
//	}
//})
//
//var _ = AfterEach(func() {
//	mockCtrl.Finish()
//})
