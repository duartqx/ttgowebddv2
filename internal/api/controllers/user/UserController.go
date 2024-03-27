package user

import (
	"encoding/json"
	"net/http"

	"github.com/duartqx/ttgowebddv2/src/api/dto"
	h "github.com/duartqx/ttgowebddv2/src/api/http"
	e "github.com/duartqx/ttgowebddv2/src/common/errors"
	a "github.com/duartqx/ttgowebddv2/src/domains/auth"
	u "github.com/duartqx/ttgowebddv2/src/domains/user"
)

type UserController struct {
	userService    u.IUserService
	sessionService a.ISessionService
}

var userController *UserController

func GetUserController(
	userService u.IUserService, sessionService a.ISessionService,
) *UserController {
	if userController == nil {
		userController = &UserController{
			userService: userService, sessionService: sessionService,
		}
	}
	return userController
}

func (uc UserController) Create(w http.ResponseWriter, r *http.Request) {

	var userDTO dto.UserCreate

	if err := json.NewDecoder(r.Body).Decode(&userDTO); err != nil {
		http.Error(w, e.BadRequestError.Error(), http.StatusBadRequest)
		return
	}

	user := u.GetNewUser().
		SetName(userDTO.Name).
		SetEmail(userDTO.Email).
		SetPassword(userDTO.Password)

	if err := uc.userService.Create(user); err != nil {
		h.ErrorResponse(w, err)
		return
	}

	h.JsonResponse(w, http.StatusCreated, user)
}

func (uc UserController) Get(w http.ResponseWriter, r *http.Request) {

	user := uc.sessionService.GetSessionUser(r.Context())
	if user == nil {
		http.Error(w, e.ForbiddenError.Error(), http.StatusForbidden)
		return
	}

	h.JsonResponse(w, http.StatusOK, user)
}

func (uc UserController) UpdatePassword(w http.ResponseWriter, r *http.Request) {

	user := uc.sessionService.GetSessionUser(r.Context())
	if user == nil {
		http.Error(w, e.ForbiddenError.Error(), http.StatusForbidden)
		return
	}

	var p dto.UserPassword

	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userToUpdate := u.GetNewUser().SetId(user.GetId()).SetPassword(p.Password)

	if err := uc.userService.UpdatePassword(userToUpdate); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
