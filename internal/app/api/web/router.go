package web

import (
	"service/internal/app/api/web/handler"
	ItemHandler "service/internal/app/api/web/handler/Item"

	"github.com/gorilla/mux"
)

func Register(router *mux.Router) {
	activityRouter(router)
	itemRouter(router)
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
