package server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	service2 "service/internal/item/service"
	"service/internal/pkg/core"
	"service/internal/pkg/form"
	"service/internal/pkg/grpc/inventory"
	"service/internal/pkg/model"

	"google.golang.org/grpc"
)

type SettingItemBrandServer struct {
	inventory.UnimplementedSettingItemBrandServiceServer

	rollbackData map[string]interface{}
}

func (srv *SettingItemBrandServer) Register(serverRPC *grpc.Server) {
	inventory.RegisterSettingItemBrandServiceServer(serverRPC, srv)
}

func (srv *SettingItemBrandServer) Store(ctx context.Context, in *inventory.SettingItemBrandRequest) (*inventory.Response, error) {
	res, err := core.GRPCErrorHandler(func() (*inventory.Response, error) {
		var brand model.ItemComponentBrand

		service := service2.NewItemComponentBrandService()

		brand = service.Create(form.SettingForm{
			Name: in.GetName(),
		})

		return srv.success(brand)
	})

	return res, err
}

func (srv *SettingItemBrandServer) RollbackStore(ctx context.Context, in *inventory.RollbackRequest) (*inventory.Response, error) {
	res, err := core.GRPCErrorHandler(func() (*inventory.Response, error) {
		err := json.Unmarshal(in.GetData(), &srv.rollbackData)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("Unable to unmarshal data. Err: %s", err))

		}

		service := service2.NewItemComponentBrandService()

		service.Delete(srv.rollbackData["id"].(string))

		return srv.success()
	})

	return res, err
}

/** --- UNEXPORTED FUNCTIONS --- */

func (srv *SettingItemBrandServer) success(result ...any) (*inventory.Response, error) {
	var data []byte
	if len(result) > 0 {
		data, _ = json.Marshal(result[0])
	}

	return &inventory.Response{Message: "Success", Result: data}, nil
}
