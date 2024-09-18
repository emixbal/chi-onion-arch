package menu

import (
	"chi-onion-arch/app/helper/error"
	"chi-onion-arch/app/model"
	menu "chi-onion-arch/app/repository/menu"
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/google/uuid"
)

type MenuUsecase interface {
	List(ctx context.Context) model.ApiResponse
	Detail(ctx context.Context, id uuid.UUID) model.ApiResponse
	Insert(ctx context.Context, role model.Menu) model.ApiResponse
	Update(ctx context.Context, datas map[string]interface{}, id uuid.UUID) model.ApiResponse
}

type menuUsecase struct {
	menuRepo       menu.MenuRepository
	contextTimeout time.Duration
}

func NewMenuUsecase(menuRepo menu.MenuRepository, contextTimeout time.Duration) MenuUsecase {
	return &menuUsecase{
		menuRepo:       menuRepo,
		contextTimeout: contextTimeout,
	}
}

func (u *menuUsecase) List(c context.Context) model.ApiResponse {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	menus, err := u.menuRepo.List(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.ApiResponse{
				Code:    200,
				Message: error.ErrMenuNotFound.Error(),
				Data:    nil,
			}
		}

		log.Println(err.Error())
		return model.ApiResponse{
			Code:    500,
			Message: error.InternalServerError.Error(),
			Data:    nil,
		}
	}
	return model.ApiResponse{
		Code:    200,
		Message: "Success",
		Data:    menus,
	}
}

func (u *menuUsecase) Detail(c context.Context, id uuid.UUID) model.ApiResponse {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	menu, err := u.menuRepo.Detail(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.ApiResponse{
				Code:    200,
				Message: error.ErrMenuNotFound.Error(),
				Data:    nil,
			}
		}
		log.Println(err.Error())
		return model.ApiResponse{
			Code:    500,
			Message: error.InternalServerError.Error(),
			Data:    nil,
		}
	}
	return model.ApiResponse{
		Code:    200,
		Message: "Success",
		Data:    menu,
	}
}

func (u *menuUsecase) Insert(c context.Context, menu model.Menu) model.ApiResponse {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	menuResult, err := u.menuRepo.Insert(ctx, menu)
	if err != nil {
		if err == error.ErrMenuOrPathAlreadyExists {
			return model.ApiResponse{
				Code:    409,
				Message: error.ErrMenuOrPathAlreadyExists.Error(),
				Data:    nil,
			}
		}

		log.Println(err.Error())
		return model.ApiResponse{
			Code:    500,
			Message: error.InternalServerError.Error(),
			Data:    nil,
		}
	}
	return model.ApiResponse{
		Code:    200,
		Message: "Menu inserted successfully",
		Data:    menuResult,
	}
}

func (u *menuUsecase) Update(c context.Context, datas map[string]interface{}, id uuid.UUID) model.ApiResponse {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	// Cek apakah Menu dengan ID yang diberikan ada di database
	_, err := u.menuRepo.Detail(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.ApiResponse{
				Code:    200,
				Message: error.ErrMenuNotFound.Error(),
				Data:    nil,
			}
		}
		log.Println(err.Error())
		return model.ApiResponse{
			Code:    500,
			Message: error.InternalServerError.Error(),
			Data:    nil,
		}
	}

	newName := datas["name"].(string)
	newPath := datas["path"].(string)

	duplicate, err := u.menuRepo.CheckDuplicateNameOrPath(ctx, newName, newPath, id)
	if err != nil {
		log.Println(err.Error())
		return model.ApiResponse{
			Code:    500,
			Message: error.InternalServerError.Error(),
			Data:    nil,
		}
	}

	if duplicate != nil {
		return model.ApiResponse{
			Code:    409,
			Message: error.ErrMenuOrPathAlreadyExists.Error(),
			Data:    nil,
		}
	}

	err = u.menuRepo.Update(ctx, datas, id)
	if err != nil {
		log.Println(err.Error())
		return model.ApiResponse{
			Code:    500,
			Message: error.InternalServerError.Error(),
			Data:    nil,
		}
	}

	updatedMenu, err := u.menuRepo.Detail(ctx, id)
	if err != nil {
		log.Println(err.Error())
		return model.ApiResponse{
			Code:    500,
			Message: error.InternalServerError.Error(),
			Data:    nil,
		}
	}

	return model.ApiResponse{
		Code:    200,
		Message: "Menu updated successfully",
		Data:    updatedMenu,
	}
}
