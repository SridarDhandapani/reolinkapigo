package api

import (
	"encoding/json"
	"fmt"
	"github.com/ReolinkCameraAPI/reolinkapigo/internal/pkg/enum"
	"github.com/ReolinkCameraAPI/reolinkapigo/internal/pkg/models"
	"github.com/ReolinkCameraAPI/reolinkapigo/pkg/network/rest"
)

type UserMixin struct{}

// Retrieves a slice of Online Users
func (um *UserMixin) GetOnlineUsers() func(handler *rest.RestHandler) ([]*models.User, error) {
	return func(handler *rest.RestHandler) ([]*models.User, error) {
		payload := map[string]interface{}{
			"cmd":    "GetOnline",
			"action": 1,
			"param":  map[string]interface{}{},
		}

		result, err := handler.Request("POST", payload, "GetOnline")

		if err != nil {
			return nil, err
		}

		var usersData []*models.User

		err = json.Unmarshal(result.Value["User"], &usersData)

		if err != nil {
			return nil, err
		}

		return usersData, nil
	}
}

// Retrieves a slice of Users
func (um *UserMixin) GetUsers() func(handler *rest.RestHandler) ([]*models.User, error) {
	return func(handler *rest.RestHandler) ([]*models.User, error) {
		payload := map[string]interface{}{
			"cmd":    "GetUser",
			"action": 1,
			"param":  map[string]interface{}{},
		}

		result, err := handler.Request("POST", payload, "GetUser")

		if err != nil {
			return nil, err
		}

		var usersData []*models.User

		err = json.Unmarshal(result.Value["User"], &usersData)

		if err != nil {
			return nil, err
		}

		return usersData, nil
	}
}

// Add a User to the camera
func (um *UserMixin) AddUser(
	username string,
	password string,
	level enum.UserLevel) func(handler *rest.RestHandler) (bool, error) {
	return func(handler *rest.RestHandler) (bool, error) {
		payload := map[string]interface{}{
			"cmd":    "AddUser",
			"action": 0,
			"param": map[string]interface{}{
				"User": map[string]interface{}{
					"userName": username,
					"password": password,
					"level":    level.Value(),
				},
			},
		}

		result, err := handler.Request("POST", payload, "AddUser")

		if err != nil {
			return false, err
		}

		var respCode int

		err = json.Unmarshal(result.Value["rspCode"], &respCode)

		if respCode == 200 {
			return true, nil
		}

		return false, fmt.Errorf("could not add user. camera responded with %v", result.Value)
	}
}

// Update the User's password
func (um *UserMixin) UpdateUserPassword(username string, password string) func(handler *rest.RestHandler) (bool,
	error) {
	return func(handler *rest.RestHandler) (bool, error) {
		payload := map[string]interface{}{
			"cmd":    "ModifyUser",
			"action": 0,
			"param": map[string]interface{}{
				"User": map[string]interface{}{
					"userName": username,
					"password": password,
				},
			},
		}

		result, err := handler.Request("POST", payload, "ModifyUser")

		if err != nil {
			return false, err
		}

		var respCode int

		err = json.Unmarshal(result.Value["rspCode"], &respCode)

		if err != nil {
			return false, err
		}

		if respCode == 200 {
			return true, nil
		}

		return false, fmt.Errorf("camera could not update user's password. camera responded with %v", result.Value)
	}
}

// Delete the User account
func (um *UserMixin) DeleteUser(username string) func(handler *rest.RestHandler) (bool, error) {
	return func(handler *rest.RestHandler) (bool, error) {
		payload := map[string]interface{}{
			"cmd":    "DelUser",
			"action": 0,
			"param": map[string]interface{}{
				"User": map[string]interface{}{
					"userName": username,
				},
			},
		}

		result, err := handler.Request("POST", payload, "DelUser")

		if err != nil {
			return false, err
		}

		var respCode int

		err = json.Unmarshal(result.Value["rspCode"], &respCode)

		if err != nil {
			return false, err
		}

		if respCode == 200 {
			return true, nil
		}

		return false, fmt.Errorf("camera could not delete user. camera responded with %v", result.Value)
	}
}
