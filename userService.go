package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func createUserHandler(repo UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var item User
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&item); err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}

		if err := repo.CreateUser(item); err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		respondWithJSON(w, http.StatusCreated, map[string]string{"message": "Data created"})
	}
}

func readUserHandler(repo UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := r.URL.Query()
		id, convErr := strconv.Atoi(params.Get("id"))
		if convErr != nil {
			respondWithError(w, http.StatusNotFound, "Invalid Format For Parameter")
			return
		}

		item, err := repo.GetUser(id)
		if err != nil {
			respondWithError(w, http.StatusNotFound, "Data not found")
			return
		}

		respondWithJSON(w, http.StatusOK, item)
	}
}

func updateUserHandler(repo UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := r.URL.Query()
		id, convErr := strconv.Atoi(params.Get("id"))
		if convErr != nil {
			respondWithError(w, http.StatusNotFound, "Invalid Format For Parameter")
			return
		}

		var item User
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&item); err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}

		if err := repo.UpdateUser(id, item); err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		respondWithJSON(w, http.StatusOK, map[string]string{"message": "Data updated"})
	}
}

func deleteUserHandler(repo UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := r.URL.Query()
		id, convErr := strconv.Atoi(params.Get("id"))
		if convErr != nil {
			respondWithError(w, http.StatusNotFound, "Invalid Format For Parameter")
			return
		}

		if err := repo.DeleteUser(id); err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		respondWithJSON(w, http.StatusOK, map[string]string{"message": "Data deleted"})
	}
}
