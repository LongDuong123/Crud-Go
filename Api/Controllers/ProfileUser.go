package controllers

import (
	middleware "crud/Api/Middleware"
	domain "crud/Domain"
	"encoding/json"
	"net/http"
)

type ProfileControllers struct {
	profileInteractor domain.ProfileInteractor
}

func NewProfileControllers(pri domain.ProfileInteractor) *ProfileControllers {
	return &ProfileControllers{
		profileInteractor: pri,
	}
}

func (prc *ProfileControllers) ProfileUser(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIdKey).(int)
	if !ok {
		http.Error(w, "User ID not found in context", http.StatusInternalServerError)
		return
	}
	var profile *domain.Profile
	profile, err := prc.profileInteractor.GetProfileByID(userID)
	if err != nil {
		http.Error(w, "User ID not found", http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(profile)
}

func (prc *ProfileControllers) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIdKey).(int)
	if !ok {
		http.Error(w, "User ID not found in context", http.StatusInternalServerError)
		return
	}
	var profile domain.Profile
	err := json.NewDecoder(r.Body).Decode(&profile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = prc.profileInteractor.UpdateProfile(userID, &profile)
	if err != nil {
		http.Error(w, "User ID not found", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}
