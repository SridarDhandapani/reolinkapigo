package api

import (
	"encoding/json"
	"fmt"
	"github.com/ReolinkCameraAPI/reolinkapigo/internal/pkg/models"
	"github.com/ReolinkCameraAPI/reolinkapigo/pkg/network/rest"
)

type PtzMixin struct{}

type ptzOperationOptions struct {
	Operation string
	Speed     *int
	Index     *int
}

type ptzPresetOptions struct {
	Index int
	Name  string
}

type OptionPtzOperation func(*ptzOperationOptions)

type OptionPtzPreset func(*ptzPresetOptions)

// helper function for ptz presets
func ptzPreset(enable bool, preset int, name string) interface{} {
	return map[string]interface{}{
		"cmd":    "SetPtzPreset",
		"action": 0,
		"param": map[string]interface{}{
			"channel": 0,
			"enable":  enable,
			"id":      preset,
			"name":    name,
		},
	}
}

// helper function for ptz operations
func ptzOperation(ptzOperation *ptzOperationOptions) interface{} {

	param := map[string]interface{}{
		"channel": 0,
		"op":      ptzOperation.Operation,
	}

	if ptzOperation.Index != nil {
		param["index"] = ptzOperation.Index
	}

	if ptzOperation.Speed != nil {
		param["speed"] = ptzOperation.Speed
	}

	return map[string]interface{}{
		"cmd":    "PtzCtrl",
		"action": 0,
		"param":  param,
	}
}

func (pm *PtzMixin) GetPreset() func(handler *rest.RestHandler) (map[string]int, error) {
	return func(handler *rest.RestHandler) (map[string]int, error) {
		payload := map[string]interface{}{
			"cmd":    "GetPtzPreset",
			"action": 1,
			"param": map[string]interface{}{
				"channel": 0,
			},
		}
		result, err := handler.Request("POST", payload, "GetPtzPreset")

		if err != nil {
			return map[string]int{}, err
		}

		var presets []*models.PtzPreset

		err = json.Unmarshal(result.Value["PtzPreset"], &presets)

		presetMap := make(map[string]int, 2)

		if err != nil {
			return presetMap, err
		}

		for _, p := range presets {
			if p.Enable == 1 {
				presetMap[p.Name] = p.Index
			}
		}

		return presetMap, nil
	}
}

// Moves the camera to the specified preset
// The preset index and speed is optional and will fallback to defaults
// One can also force the preset to have no index by passing api.PtzOptionOpsIndex(nil)
// Defaults:
// index: 1
// speed: 60
func (pm *PtzMixin) GoToPreset(ptzOptions ...OptionPtzOperation) func(handler *rest.RestHandler) (
	bool, error) {
	speed := 60
	index := 1

	ptzPreset := &ptzOperationOptions{
		Operation: "ToPos",
		Speed:     &speed,
		Index:     &index,
	}

	for _, op := range ptzOptions {
		op(ptzPreset)
	}

	return func(handler *rest.RestHandler) (bool, error) {
		payload := ptzOperation(ptzPreset)
		result, err := handler.Request("POST", payload, "PtzCtrl")

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

		return false, fmt.Errorf("camera could not go to ptz preset. camera responded with %v", result.Value)
	}
}

// Create a new preset at the current camera position
// The preset index and name is optional and will fallback to defaults
// Defaults:
// index: 1
// name: pos1
func (pm *PtzMixin) AddPreset(ptzOptions ...OptionPtzPreset) func(handler *rest.RestHandler) (bool, error) {
	presetOptions := &ptzPresetOptions{
		Index: 1,
		Name:  "pos1",
	}

	for _, op := range ptzOptions {
		op(presetOptions)
	}

	return func(handler *rest.RestHandler) (bool, error) {
		payload := ptzPreset(true, presetOptions.Index, presetOptions.Name)

		result, err := handler.Request("POST", payload, "PtzCtrl")

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

		return false, fmt.Errorf("camera could not add preset. camera responded with %v", result.Value)
	}
}

// Remove the specified preset
// The preset index and name is optional and will fallback to defaults
// Defaults:
// index: 1
// name: pos1
func (pm *PtzMixin) RemovePreset(ptzOptions ...OptionPtzPreset) func(handler *rest.RestHandler) (bool, error) {

	presetOptions := &ptzPresetOptions{
		Index: 1,
		Name:  "pos1",
	}

	for _, op := range ptzOptions {
		op(presetOptions)
	}

	return func(handler *rest.RestHandler) (bool, error) {
		payload := ptzPreset(false, presetOptions.Index, presetOptions.Name)

		result, err := handler.Request("POST", payload, "PtzPreset")

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

		return false, fmt.Errorf("camera could not remove preset. camera responded with %v", result.Value)
	}
}

// Move the camera to the right
// The operation speed is optional and will fallback to defaults. Other operations will be ignored.
// Defaults:
// speed: 25
func (pm *PtzMixin) MoveRight(ptzOptions ...OptionPtzOperation) func(handler *rest.RestHandler) (bool,
	error) {

	speed := 25

	ptzOperations := &ptzOperationOptions{
		Operation: "Right",
		Speed:     &speed,
		Index:     nil,
	}

	for _, op := range ptzOptions {
		op(ptzOperations)
	}

	return func(handler *rest.RestHandler) (bool, error) {
		// set the index to nil in case the user passes an option for it
		ptzOperations.Index = nil
		payload := ptzOperation(ptzOperations)

		result, err := handler.Request("POST", payload, "PtzCtrl")

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

		return false, fmt.Errorf("camera could not move right. camera responded with %v", result.Value)
	}
}

// Move the camera to the right Up
// The operation speed is optional and will fallback to defaults. Other operations will be ignored.
// Defaults:
// speed: 25
func (pm *PtzMixin) MoveRightUp(ptzOptions ...OptionPtzOperation) func(handler *rest.RestHandler) (bool,
	error) {

	speed := 25

	ptzOperations := &ptzOperationOptions{
		Operation: "RightUp",
		Speed:     &speed,
		Index:     nil,
	}

	for _, op := range ptzOptions {
		op(ptzOperations)
	}

	return func(handler *rest.RestHandler) (bool, error) {
		// set the index to nil in case the user passes an option for it
		ptzOperations.Index = nil
		payload := ptzOperation(ptzOperations)

		result, err := handler.Request("POST", payload, "PtzCtrl")

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

		return false, fmt.Errorf("camera could not move right up. camera responded with %v", result.Value)
	}
}

// Move the camera to the right Down
// The operation speed is optional and will fallback to defaults. Other operations will be ignored.
// Defaults:
// speed: 25
func (pm *PtzMixin) MoveRightDown(ptzOptions ...OptionPtzOperation) func(handler *rest.RestHandler) (bool,
	error) {

	speed := 25

	ptzOperations := &ptzOperationOptions{
		Operation: "RightDown",
		Speed:     &speed,
		Index:     nil,
	}

	for _, op := range ptzOptions {
		op(ptzOperations)
	}

	return func(handler *rest.RestHandler) (bool, error) {
		// set the index to nil in case the user passes an option for it
		ptzOperations.Index = nil
		payload := ptzOperation(ptzOperations)

		result, err := handler.Request("POST", payload, "PtzCtrl")

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

		return false, fmt.Errorf("camera could not move right down. camera responded with %v", result.Value)
	}
}

// Move the camera to the right Left
// The operation speed is optional and will fallback to defaults. Other operations will be ignored.
// Defaults:
// speed: 25
func (pm *PtzMixin) MoveLeft(ptzOptions ...OptionPtzOperation) func(handler *rest.RestHandler) (bool,
	error) {

	speed := 25

	ptzOperations := &ptzOperationOptions{
		Operation: "Left",
		Speed:     &speed,
		Index:     nil,
	}

	for _, op := range ptzOptions {
		op(ptzOperations)
	}

	return func(handler *rest.RestHandler) (bool, error) {
		// set the index to nil in case the user passes an option for it
		ptzOperations.Index = nil
		payload := ptzOperation(ptzOperations)

		result, err := handler.Request("POST", payload, "PtzCtrl")

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

		return false, fmt.Errorf("camera could not move left. camera responded with %v", result.Value)
	}
}

// Move the camera to the left Up
// The operation speed is optional and will fallback to defaults. Other operations will be ignored.
// Defaults:
// speed: 25
func (pm *PtzMixin) MoveLeftUp(ptzOptions ...OptionPtzOperation) func(handler *rest.RestHandler) (bool,
	error) {

	speed := 25

	ptzOperations := &ptzOperationOptions{
		Operation: "LeftUp",
		Speed:     &speed,
		Index:     nil,
	}

	for _, op := range ptzOptions {
		op(ptzOperations)
	}

	return func(handler *rest.RestHandler) (bool, error) {
		// set the index to nil in case the user passes an option for it
		ptzOperations.Index = nil
		payload := ptzOperation(ptzOperations)

		result, err := handler.Request("POST", payload, "PtzCtrl")

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

		return false, fmt.Errorf("camera could not move left up. camera responded with %v", result.Value)
	}
}

// Move the camera to the left Down
// The operation speed is optional and will fallback to defaults. Other operations will be ignored.
// Defaults:
// speed: 25
func (pm *PtzMixin) MoveLeftDown(ptzOptions ...OptionPtzOperation) func(handler *rest.RestHandler) (bool,
	error) {

	speed := 25

	ptzOperations := &ptzOperationOptions{
		Operation: "LeftDown",
		Speed:     &speed,
		Index:     nil,
	}

	for _, op := range ptzOptions {
		op(ptzOperations)
	}

	return func(handler *rest.RestHandler) (bool, error) {
		// set the index to nil in case the user passes an option for it
		ptzOperations.Index = nil
		payload := ptzOperation(ptzOperations)

		result, err := handler.Request("POST", payload, "PtzCtrl")

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

		return false, fmt.Errorf("camera could not move left down. camera responded with %v", result.Value)
	}
}

// Move the camera to the up
// The operation speed is optional and will fallback to defaults. Other operations will be ignored.
// Defaults:
// speed: 25
func (pm *PtzMixin) MoveUp(ptzOptions ...OptionPtzOperation) func(handler *rest.RestHandler) (bool,
	error) {

	speed := 25

	ptzOperations := &ptzOperationOptions{
		Operation: "Up",
		Speed:     &speed,
		Index:     nil,
	}

	for _, op := range ptzOptions {
		op(ptzOperations)
	}

	return func(handler *rest.RestHandler) (bool, error) {
		// set the index to nil in case the user passes an option for it
		ptzOperations.Index = nil
		payload := ptzOperation(ptzOperations)

		result, err := handler.Request("POST", payload, "PtzCtrl")

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

		return false, fmt.Errorf("camera could not move up. camera responded with %v", result.Value)
	}
}

// Move the camera to the down
// The operation speed is optional and will fallback to defaults. Other operations will be ignored.
// Defaults:
// speed: 25
func (pm *PtzMixin) MoveDown(ptzOptions ...OptionPtzOperation) func(handler *rest.RestHandler) (bool,
	error) {

	speed := 25

	ptzOperations := &ptzOperationOptions{
		Operation: "Down",
		Speed:     &speed,
		Index:     nil,
	}

	for _, op := range ptzOptions {
		op(ptzOperations)
	}

	return func(handler *rest.RestHandler) (bool, error) {
		// set the index to nil in case the user passes an option for it
		ptzOperations.Index = nil
		payload := ptzOperation(ptzOperations)

		result, err := handler.Request("POST", payload, "PtzCtrl")

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

		return false, fmt.Errorf("camera could not move down. camera responded with %v", result.Value)
	}
}

// Stops the cameras current action
func (pm *PtzMixin) StopPtz() func(handler *rest.RestHandler) (bool,
	error) {
	return func(handler *rest.RestHandler) (bool, error) {
		ptzOperations := &ptzOperationOptions{
			Operation: "Stop",
			Speed:     nil,
			Index:     nil,
		}

		payload := ptzOperation(ptzOperations)

		result, err := handler.Request("POST", payload, "PtzCtrl")

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

		return false, fmt.Errorf("camera could not stop ptz operation. camera responded with %v", result.Value)
	}
}

// Move the camera in a clockwise rotation
func (pm *PtzMixin) AutoMovement() func(handler *rest.RestHandler) (bool, error) {
	return func(handler *rest.RestHandler) (bool, error) {
		ptzOperations := &ptzOperationOptions{
			Operation: "Auto",
			Speed:     nil,
			Index:     nil,
		}

		payload := ptzOperation(ptzOperations)

		result, err := handler.Request("POST", payload, "PtzCtrl")

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

		return false, fmt.Errorf("camera could not auto move. camera responded with %v", result.Value)
	}
}

// Set the Ptz Operation Speed
func PtzOptionOpsSpeed(speed int) OptionPtzOperation {
	return func(p *ptzOperationOptions) {
		p.Speed = &speed
	}
}

// Set the Ptz Operation Index
func PtzOptionOpsIndex(index *int) OptionPtzOperation {
	return func(p *ptzOperationOptions) {
		p.Index = index
	}
}

// Set the Ptz Preset Index
func PtzOptionPresetIndex(index int) OptionPtzPreset {
	return func(p *ptzPresetOptions) {
		p.Index = index
	}
}

// Set the Ptz Preset Value
func PtzOptionsPresetName(name string) OptionPtzPreset {
	return func(p *ptzPresetOptions) {
		p.Name = name
	}
}
