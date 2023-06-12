package resto_test

import (
	"context"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/dzikrurrohmani/golang-echo-rest-api/internal/mocks"
	"github.com/dzikrurrohmani/golang-echo-rest-api/internal/model"
	"github.com/dzikrurrohmani/golang-echo-rest-api/internal/model/constant"
	"github.com/dzikrurrohmani/golang-echo-rest-api/internal/usecase/resto"
)

var _ = Describe("GinkgoResto", func() {
	var usecase resto.Usecase
	var menuRepoMock *mocks.MockMenuRepository
	var orderRepoMock *mocks.MockOrderRepository
	var userRepoMock *mocks.MockUserRepository
	var mockCtrl *gomock.Controller

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		menuRepoMock = mocks.NewMockMenuRepository(mockCtrl)
		orderRepoMock = mocks.NewMockOrderRepository(mockCtrl)
		userRepoMock = mocks.NewMockUserRepository(mockCtrl)

		usecase = resto.GetUsecase(menuRepoMock, orderRepoMock, userRepoMock)
	})
	AfterEach(func() {
		mockCtrl.Finish()
	})

	Describe("requesting order info", func() {
		Context("it gave the correct inputs", func() {
			inputs := model.GetOrderInfoRequest{
				OrderID: "orderID",
				UserID:  "userID",
			}
			When("the requested orderID is not the user's", func() {
				BeforeEach(func() {
					orderRepoMock.EXPECT().GetOrderInfo(gomock.Any(), inputs.OrderID).
						Times(1).
						Return(model.Order{
							ID:            "orderID",
							UserID:        "notUserID",
							Status:        constant.OrderStatusFinished,
							ProductOrders: []model.ProductOrder{},
							ReferenceID:   "randomref",
						}, nil)
				})
				It("returns unauthorized error", func() {
					res, err := usecase.GetOrderInfo(context.Background(), inputs)
					Expect(err).Should(HaveOccurred())
					Expect(err.Error()).To(BeEquivalentTo("unauthorized"))
					Expect(res).To(BeEquivalentTo(model.Order{}))
				})
			})

			When("the requested orderID is the user's", func() {
				BeforeEach(func() {
					orderRepoMock.EXPECT().GetOrderInfo(gomock.Any(), inputs.OrderID).
						Times(1).
						Return(model.Order{
							ID:            "orderID",
							UserID:        "userID",
							Status:        constant.OrderStatusFinished,
							ProductOrders: []model.ProductOrder{},
							ReferenceID:   "randomref",
						}, nil)
				})
				It("returns unauthorized error", func() {
					res, err := usecase.GetOrderInfo(context.Background(), inputs)
					Expect(err).ShouldNot(HaveOccurred())
					Expect(res).To(BeEquivalentTo(model.Order{
						ID:            "orderID",
						UserID:        "userID",
						Status:        constant.OrderStatusFinished,
						ProductOrders: []model.ProductOrder{},
						ReferenceID:   "randomref",
					}))
				})
			})
		})
	})
})
