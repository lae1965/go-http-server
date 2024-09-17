package users

import (
	"encoding/json"
	"http-server/utils"
	"io"
	"net/http"
	"strconv"
)

const userEndpoint = "user"

func UsersController(res http.ResponseWriter, req *http.Request) {
	sendResponse := func(data any, status int) {
		if jsonData, err := json.Marshal(data); err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
		} else {
			res.WriteHeader(status)
			res.Header().Set("Content-Type", "application/json")
			res.Write(jsonData)
		}
	}

	getBody := func() User {
		body, err := io.ReadAll(req.Body)
		if err != nil {
			http.Error(res, err.Error(), http.StatusBadRequest)
		}
		var user User
		err = json.Unmarshal(body, &user)
		if err != nil {
			http.Error(res, err.Error(), http.StatusBadRequest)
		}
		return user
	}

	getId := func() (bool, int64) {
		id, err := utils.ParseUrlParams(req.URL.Path, userEndpoint)
		if err != nil {
			http.Error(res, err.Error(), http.StatusNotFound)
		}
		if id == nil {
			return false, 0
		}
		intId, err := strconv.Atoi(*id)
		if err != nil {
			http.Error(res, err.Error(), http.StatusBadRequest)
		}
		return true, int64(intId)
	}

	switch req.Method {
	case http.MethodGet:
		isIdInUrl, id := getId()
		var data any
		if isIdInUrl {
			user := serviceGetUserById(id)
			if user == nil {
				http.Error(res, "Пользователь не найден", http.StatusNotFound)
				return
			}
			data = user
		} else {
			data = serviceGetAllUsers()
		}
		sendResponse(data, http.StatusOK)
	case http.MethodPost:
		user := getBody()
		defer req.Body.Close()
		data := serviceAddUser(user)
		sendResponse(data, http.StatusCreated)
	case http.MethodPatch:
		isIdInUrl, id := getId()
		if isIdInUrl {
			user := getBody()
			defer req.Body.Close()
			data := serviceUpdateUser(id, user)
			if data == nil {
				http.Error(res, "Пользователь не найден", http.StatusNotFound)
			} else {
				sendResponse(data, http.StatusOK)
			}
		} else {
			http.Error(res, "Не указан идентификатор изменяемого пользователя", http.StatusBadRequest)
		}
	case http.MethodDelete:
		isIdInUrl, id := getId()
		if isIdInUrl {
			isRemoved := serviceRemoveUser(id)
			if isRemoved {
				res.WriteHeader(http.StatusOK)
				res.Write([]byte("Пользователь успешно удален"))
			} else {
				http.Error(res, "Пользователь не найден", http.StatusNotFound)
			}
		} else {
			http.Error(res, "Не указан идентификатор удаляемого пользователя", http.StatusBadRequest)
		}
	}

}
