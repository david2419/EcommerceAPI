package user

import (
	"e-commerce/configs"
	"e-commerce/service/auth"
	"e-commerce/types"
	"e-commerce/utils"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")

}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	var userPayload types.LoginUserPayload
	if err := utils.ParseJSON(r, &userPayload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("error en parsejson"))
		return
	}

	userDB, err := h.store.GetUserByEmail(userPayload.Email)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("error, usuario o password incorrectos"))
	}

	if !auth.ComparePasswords(userDB.Password, []byte(userPayload.Password)) {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("error, usuario o password incorrectos"))
	}

	secret := []byte(configs.Envs.JWTSecret)
	token, err := auth.CreateJWT(secret, userDB.ID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{"token": token})

}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	//get JSON payload
	var user types.RegisterUserPayload
	if err := utils.ParseJSON(r, &user); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("error en parsejson"))
		return
	}
	// check if user exists
	_, err := h.store.GetUserByEmail(user.Email)
	if err == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with email %s already exists", user.Email))
		return
	}

	//create user
	hashedPassword, err := auth.HashPassword(user.Password)
	println(hashedPassword)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("error creando hash del password"))
		return
	}

	err = h.store.CreateUser(types.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  hashedPassword,
	})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("error creando usuario"))
		return
	}

}
