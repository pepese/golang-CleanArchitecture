package controller

// func TestUsersGrpcAPI(t *testing.T) {
// 	// mock controller
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	app.Config()

// 	t.Run("Get empty users", func(t *testing.T) {
// 		// setting mock
// 		m := mock_usecase.NewMockUserUsecaser(ctrl)
// 		m.EXPECT().List(gomock.Any(), gomock.Any()).Return(model.Users{}, nil)

// 		// setting gRPC Server
// 		lis := bufconn.Listen(1024 * 1024)
// 		s := grpc.NewServer()
// 		upb.RegisterUserServiceServer(s, nil)

// 		assert.Equal(t, 200, res.Code)
// 		assert.Equal(t, "[]\n", getBodyString(res))
// 	})
// }
