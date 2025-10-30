package web

import (
	"service/internal/app/api/web/handler"
	ItemHandler "service/internal/app/api/web/handler/Item"
	SettingHandler "service/internal/app/api/web/handler/Setting"

	"github.com/gorilla/mux"
)

func Register(router *mux.Router) {
	activityRouter(router)
	itemRouter(router)
	settingRouter(router)
}

func activityRouter(router *mux.Router) {
	router.HandleFunc("/activities", handler.ActivityHandler{}.Get).Methods("GET")
}

// TODO: Hanya contoh. nanti langsung hapus saja
func testingRouter(router *mux.Router) {
	var testingHandler handler.TestingHandler
	router.HandleFunc("/testings", testingHandler.Get).Methods("GET")
	router.HandleFunc("/testings", testingHandler.Create).Methods("POST")
	router.HandleFunc("/testings/upload/file", testingHandler.UploadByFile).Methods("POST")
	router.HandleFunc("/testings/upload/content", testingHandler.UploadByContent).Methods("POST")
}

func itemRouter(router *mux.Router) {
	var itemHandler ItemHandler.ItemHandler
	router.HandleFunc("/items", itemHandler.Get).Methods("GET")
}

func settingRouter(router *mux.Router) {
	router = router.PathPrefix("/settings").Subrouter()

	var itemTypeHandler SettingHandler.ItemTypeHandler
	router.HandleFunc("/item-types", itemTypeHandler.Get).Methods("GET")
	router.HandleFunc("/item-types", itemTypeHandler.Create).Methods("POST")
	router.HandleFunc("/item-types/{id}", itemTypeHandler.Update).Methods("PUT")
	router.HandleFunc("/item-types/{id}", itemTypeHandler.Delete).Methods("DELETE")

	var itemBrandHandler SettingHandler.ItemBrandHandler
	router.HandleFunc("/item-brands", itemBrandHandler.Get).Methods("GET")
	router.HandleFunc("/item-brands", itemBrandHandler.Create).Methods("POST")
	router.HandleFunc("/item-brands/{id}", itemBrandHandler.Update).Methods("PUT")
	router.HandleFunc("/item-brands/{id}", itemBrandHandler.Delete).Methods("DELETE")

	var itemUnitHandler SettingHandler.ItemUnitHandler
	router.HandleFunc("/item-units", itemUnitHandler.Get).Methods("GET")
	router.HandleFunc("/item-units", itemUnitHandler.Create).Methods("POST")
	router.HandleFunc("/item-units/{id}", itemUnitHandler.Update).Methods("PUT")
	router.HandleFunc("/item-units/{id}", itemUnitHandler.Delete).Methods("DELETE")

	var itemCategoryHandler SettingHandler.ItemCategoryHandler
	router.HandleFunc("/item-categories", itemCategoryHandler.Get).Methods("GET")
	router.HandleFunc("/item-categories", itemCategoryHandler.Create).Methods("POST")
	router.HandleFunc("/item-categories/{id}", itemCategoryHandler.Update).Methods("PUT")
	router.HandleFunc("/item-categories/{id}", itemCategoryHandler.Delete).Methods("DELETE")
}
