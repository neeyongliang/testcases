package main

import (
	"encoding/json"
	"fmt"
)

type Screen struct {
	Size       float32
	ResX, ResY int
}

type Battery struct {
	Capacity int
}

func getJsonData() []byte {
	raw := &struct {
		Screen
		Battery
		HasTouchID bool
	}{
		Screen: Screen{
			Size: 5.5,
			ResX: 1920,
			ResY: 1080,
		},
		Battery: Battery{
			2910,
		},
		HasTouchID: true,
	}
	// JSON 序列化
	jsonData, _ := json.Marshal(raw)
	return jsonData
}

func main() {
	jsonData := getJsonData()
	fmt.Println(string(jsonData))

	ScreenAndTouch := struct {
		Screen
		HasTouchID bool
	}{}
	// 反序列化
	json.Unmarshal(jsonData, &ScreenAndTouch)
	fmt.Printf("%+v\n", ScreenAndTouch)

	batteryAndTouch := struct {
		Battery
		HasTouchID bool
	}{}
	json.Unmarshal(jsonData, &batteryAndTouch)
	fmt.Printf("%+v\n", batteryAndTouch)
}
