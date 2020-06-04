package controller

import (
	"crypto/sha512"
	"encoding/hex"
	"net/http"

	"github.com/jinzhu/gorm"

	"github.com/mopeneko/novel-gamest/api/model"

	"github.com/labstack/echo/v4"
	"github.com/mopeneko/novel-gamest/api/domain"
	"github.com/mopeneko/novel-gamest/api/interface/database"
)

// UserController is an interface for User Use Case
type UserController struct {
	UserRepository    database.UserRepository
	UserTokenProvider UserTokenProvider
}

// SignUp new user
func (controller *UserController) SignUp(c echo.Context) error {
	// リクエストをバインド
	req := model.UserCreateRequest{}
	resp := model.UserCreateResponse{}
	if err := c.Bind(&req); err != nil {
		resp.Message = "不正なデータです。"
		return c.JSON(http.StatusBadRequest, &resp)
	}

	user, _ := controller.UserRepository.FindByID(req.ID)
	if len(user.UserID) > 0 {
		resp.Message = "そのIDは既に使われています。"
		return c.JSON(http.StatusConflict, &resp)
	}

	// パスワードをハッシュ化
	hashedPassword := sha512.Sum512([]byte(req.Password))
	hashedHexPassword := hex.EncodeToString(hashedPassword[:])

	// ユーザーを作成
	user = domain.User{
		UserID:   req.ID,
		Password: hashedHexPassword,
		Name:     req.ID,
	}
	err := controller.UserRepository.Save(user)
	if err != nil {
		resp.Message = "エラーが発生しました。"
		return c.JSON(http.StatusInternalServerError, &resp)
	}

	// トークンを生成
	token, err := controller.UserTokenProvider.Generate(req.ID)
	if err != nil {
		resp.Message = "エラーが発生しました。"
		return c.JSON(http.StatusInternalServerError, &resp)
	}

	resp.Message = token
	return c.JSON(http.StatusOK, &resp)
}

// GetByID returns an user which has same ID
func (controller *UserController) GetByID(c echo.Context) error {
	resp := model.UserGetByIDResponse{}

	// ユーザーを検索
	user, err := controller.UserRepository.FindByID(c.Param("id"))
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return c.JSON(http.StatusNotFound, &resp)
		}
		return c.JSON(http.StatusInternalServerError, &resp)
	}
	resp.Name = user.Name

	return c.JSON(http.StatusOK, &resp)
}

// SignIn uesr
func (controller *UserController) SignIn(c echo.Context) error {
	// リクエストをバインド
	req := model.UserGetWithAuthenticationRequest{}
	resp := model.UserGetWithAuthenticationResponse{}
	if err := c.Bind(&req); err != nil {
		resp.Message = "不正なデータです。"
		return c.JSON(http.StatusBadRequest, &resp)
	}

	// ユーザーを検索
	user, err := controller.UserRepository.FindByID(req.ID)
	if err != nil || len(user.Name) <= 0 {
		if gorm.IsRecordNotFoundError(err) {
			resp.Message = "そのIDは存在しません。"
			return c.JSON(http.StatusBadRequest, &resp)
		}
		resp.Message = "エラーが発生しました。"
		return c.JSON(http.StatusInternalServerError, &resp)
	}

	// トークンを生成
	token, err := controller.UserTokenProvider.Generate(req.ID)
	if err != nil {
		resp.Message = "エラーが発生しました。"
		return c.JSON(http.StatusInternalServerError, &resp)
	}

	resp.Message = token
	return c.JSON(http.StatusOK, &resp)
}
