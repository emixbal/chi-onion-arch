package menu

import (
	"chi-onion-arch/app/database"
	"chi-onion-arch/app/helper"
	"chi-onion-arch/app/helper/util"
	"net/http"
	"time"

	"chi-onion-arch/app/model"
	repoMenu "chi-onion-arch/app/repository/menu"
	"chi-onion-arch/app/request"
	usecaseMenu "chi-onion-arch/app/usecase/menu"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/spf13/viper"
)

type MenuHandler struct {
	MenuUsecase usecaseMenu.MenuUsecase
}

func NewMenuHandler() MenuHandler {
	timeout := time.Duration(viper.GetInt(`context.timeout`)) * time.Second
	db := database.Load()

	repoMenu := repoMenu.NewMenuRepository(db)
	ucaseMenu := usecaseMenu.NewMenuUsecase(repoMenu, timeout)

	return MenuHandler{
		MenuUsecase: ucaseMenu,
	}
}

// List menampilkan daftar menu
func (h *MenuHandler) List(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	apiResponse := h.MenuUsecase.List(ctx)
	util.ResponseJSON(ctx, w, apiResponse.Code, &apiResponse)
}

// Detail menampilkan detail menu berdasarkan ID
func (h *MenuHandler) Detail(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		apiResponse := model.ApiResponse{
			Message: "Invalid ID format",
			Code:    http.StatusBadRequest,
		}
		util.ResponseJSON(ctx, w, apiResponse.Code, &apiResponse)
		return
	}

	apiResponse := h.MenuUsecase.Detail(ctx, id)
	util.ResponseJSON(ctx, w, apiResponse.Code, &apiResponse)
}

// Insert menambahkan menu baru
func (h *MenuHandler) Insert(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var menu model.Menu

	form := new(request.FormInsertMenu)
	err := helper.ParseAndValidateForm(r, form)
	if err != nil {
		apiResponse := model.ApiResponse{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		}
		util.ResponseJSON(ctx, w, apiResponse.Code, &apiResponse)
		return
	}

	menu.Name = form.Name
	menu.Path = form.Path
	menu.CreatedAt = time.Now()
	menu.CreatedBy = r.Header.Get("X-ID")

	apiResponse := h.MenuUsecase.Insert(ctx, menu)
	util.ResponseJSON(ctx, w, apiResponse.Code, &apiResponse)
}

// Update memperbarui data menu berdasarkan ID
func (h *MenuHandler) Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var datas = make(map[string]interface{})

	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		apiResponse := model.ApiResponse{
			Message: "Invalid ID format",
			Code:    http.StatusBadRequest,
		}
		util.ResponseJSON(ctx, w, apiResponse.Code, &apiResponse)
		return
	}

	form := new(request.FormUpdateMenu)
	err = helper.ParseAndValidateForm(r, form)
	if err != nil {
		apiResponse := model.ApiResponse{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		}
		util.ResponseJSON(ctx, w, apiResponse.Code, &apiResponse)
		return
	}

	// Menyusun data dari form untuk dikirim ke usecase
	datas["name"] = form.Name
	datas["path"] = form.Path
	datas["updated_at"] = time.Now()
	datas["updated_by"] = r.Header.Get("X-ID")

	apiResponse := h.MenuUsecase.Update(ctx, datas, id)
	util.ResponseJSON(ctx, w, apiResponse.Code, &apiResponse)
}
