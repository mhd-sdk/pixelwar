package models

import "github.com/gofiber/contrib/websocket"

type Client struct {
	Username string `json:"username"`
	Conn     *websocket.Conn
}

type Message struct {
	Author  string `json:"author"`
	Content string `json:"content"`
	Date    string `json:"date"`
}

type UserList []string

type ChatHandlerWsUpdate struct {
	Messages    []Message   `json:"messages"`
	LoggedUsers UserList    `json:"loggedUsers"`
	PixelMatrix PixelMatrix `json:"pixelMatrix"`
}

type Pixel struct {
	X     int    `json:"x"`
	Y     int    `json:"y"`
	Color string `json:"color"`
}

type PixelMatrix [][]Pixel
