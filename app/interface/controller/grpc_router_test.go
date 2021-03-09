package controller

import (
	"context"
	"log"
	"net"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pepese/golang-CleanArchitecture/app"
	"github.com/pepese/golang-CleanArchitecture/app/domain/model"
	upb "github.com/pepese/golang-CleanArchitecture/app/interface/controller/grpc_gen/user/v1"
	"github.com/pepese/golang-CleanArchitecture/app/usecase/mock_usecase"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

var (
	lis *bufconn.Listener
)

func TestUsersGrpcAPI(t *testing.T) {
	// mock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	app.Config()

	t.Run("Get empty users", func(t *testing.T) {
		// setting mock
		m := mock_usecase.NewMockUserUsecaser(ctrl)
		m.EXPECT().List(gomock.Any(), gomock.Any()).Return(model.Users{}, nil)

		// setting gRPC Server
		lis = bufconn.Listen(1024 * 1024)
		s := grpc.NewServer(
			grpc.UnaryInterceptor(newTestUnaryServierInterceptor()),
		)
		router := NewGrpcRouter(m)
		upb.RegisterUserServiceServer(s, &router.UsrServ)
		go func() {
			if err := s.Serve(lis); err != nil {
				log.Fatal((err))
			}
		}()

		// setting gRPC Client
		ctx := context.Background()
		conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()
		client := upb.NewUserServiceClient(conn)
		_, err = client.ListUsers(ctx, &upb.UserRequest{})
		if err != nil {
			t.Fatal(err)
		}
	})
}

func bufDialer(ctx context.Context, address string) (net.Conn, error) {
	return lis.Dial()
}

func newTestUnaryServierInterceptor() grpc.UnaryServerInterceptor {
	return func(c context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (reply interface{}, err error) {
		ctx := context.Background()
		c = context.WithValue(c, "ctx", ctx)
		reply, err = handler(c, req)
		return
	}
}
