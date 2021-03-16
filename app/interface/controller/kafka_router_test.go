package controller

import (
	"context"
	"encoding/json"
	"log"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pepese/golang-CleanArchitecture/app"
	"github.com/pepese/golang-CleanArchitecture/app/domain/model"
	"github.com/pepese/golang-CleanArchitecture/app/usecase/mock_usecase"
)

func TestKafkaRouter(t *testing.T) {
	// mock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	// init settings
	app.Config()

	////////////////////////////////////////////////////////////////////////
	// TEST kafka_router_input
	////////////////////////////////////////////////////////////////////////
	t.Run("[Users Topic] Parse Test", func(t *testing.T) {
		msg := `{"method":"create","message":{"first_name":"first","last_name":"last"}}`
		var input UsersTopic
		if err := json.Unmarshal([]byte(msg), &input); err != nil {
			log.Println("Kafka Users Topic Parse Error.")
			log.Println(err)
		} else {
			log.Printf("input: %v\n", input)
		}
	})

	////////////////////////////////////////////////////////////////////////
	// TEST Users Topic
	////////////////////////////////////////////////////////////////////////
	t.Run("[Users Topic] Create", func(t *testing.T) {
		// setting mock
		m := mock_usecase.NewMockUserUsecaser(ctrl)
		m.EXPECT().Create(gomock.Any(), gomock.Any()).Return(&model.User{}, nil)

		msg := UsersTopic{
			Method:  "create",
			Message: model.User{},
		}
		kRouter := NewKafkaRouter(m)
		_, err := kRouter.KafkaRouter(getKafkaContext(), msg)
		if err != nil {
			t.Fail()
		}
	})

	t.Run("[Users Topic] Update", func(t *testing.T) {
		// setting mock
		m := mock_usecase.NewMockUserUsecaser(ctrl)
		m.EXPECT().Update(gomock.Any(), gomock.Any()).Return(&model.User{}, nil)

		msg := UsersTopic{
			Method:  "update",
			Message: model.User{},
		}
		kRouter := NewKafkaRouter(m)
		_, err := kRouter.KafkaRouter(getKafkaContext(), msg)
		if err != nil {
			t.Fail()
		}
	})

	t.Run("[Users Topic] Delete", func(t *testing.T) {
		// setting mock
		m := mock_usecase.NewMockUserUsecaser(ctrl)
		m.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(&model.User{}, nil)

		msg := UsersTopic{
			Method:  "delete",
			Message: model.User{},
		}
		kRouter := NewKafkaRouter(m)
		_, err := kRouter.KafkaRouter(getKafkaContext(), msg)
		if err != nil {
			t.Fail()
		}
	})
}

////////////////////////////////////////////////////////////////////////
// TEST Utilities
////////////////////////////////////////////////////////////////////////

func getKafkaContext() context.Context {
	c := context.Background()
	return context.WithValue(c, "ctx", context.Background())
}
