package server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"service/internal/pkg/activity"
	"service/internal/pkg/config"
	"service/internal/pkg/core"
	"service/internal/pkg/form"
	"service/internal/pkg/grpc/inventory"
	"service/internal/pkg/model"
	"service/internal/pkg/parser"
	"service/internal/setting/repository"

	"google.golang.org/grpc"
	"gorm.io/gorm"
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

		err := config.PgSQL.Transaction(func(tx *gorm.DB) error {
			repo := repository.NewSettingItemBrandRepository(tx)

			brand = repo.Create(form.SettingForm{Name: in.GetName()})

			parser := parser.SettingItemBrandParser{Object: brand}

			activity.UseActivity{}.SetReference(brand).SetParser(&parser).
				Save(fmt.Sprintf("gRPC: Create new brand: %s", brand.Name))

			srv.rollbackData = map[string]interface{}{
				"id": brand.ID,
			}

			return nil
		})

		if err != nil {
			return nil, err
		}

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

		err = config.PgSQL.Transaction(func(tx *gorm.DB) error {
			repo := repository.NewSettingItemBrandRepository(tx)

			idInt, _ := srv.rollbackData["id"].(int)
			id := uint(idInt)
			brand := repo.FirstById(id)

			repo.Delete(brand)

			return nil
		})
		if err != nil {
			return nil, err
		}

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
