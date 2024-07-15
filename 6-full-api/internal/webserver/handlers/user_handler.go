package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/jwtauth"
	"github.com/stefanoMat/boost/6-full-api/internal/dto"
	"github.com/stefanoMat/boost/6-full-api/internal/entity"
	"github.com/stefanoMat/boost/6-full-api/internal/infra/database"
)

type UserHandler struct {
	UserDB       database.UserInterface
	TokenAuth    *jwtauth.JWTAuth
	JWTExpiresIn int
}

func NewUserHandler(userDB database.UserInterface, tokenAuth *jwtauth.JWTAuth, jwtExpiresIn int) *UserHandler {
	return &UserHandler{UserDB: userDB, TokenAuth: tokenAuth, JWTExpiresIn: jwtExpiresIn}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user dto.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userEntity, err := entity.NewUser(user.Name, user.Email, user.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.UserDB.Create(userEntity)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(userEntity)
}

func (h *UserHandler) GetJWT(w http.ResponseWriter, r *http.Request) {

	var user dto.GetJWTInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userEntity, err := h.UserDB.FindByEmail(user.Email)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if !userEntity.ValidatePassword(user.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	_, tokenString, err := h.TokenAuth.Encode(map[string]interface{}{
		"sub": userEntity.ID.String(),
		"exp": time.Now().Add(time.Second * time.Duration(h.JWTExpiresIn)).Unix(),
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	jwtOutput := dto.GetJWTOutput{AccessToken: tokenString}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(jwtOutput)
}
