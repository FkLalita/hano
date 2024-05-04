package handlers

import (
	"github.com/gorilla/websocket"
  "github.com/labstack/echo/v4"

  "net/http"
  "sync"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var (
    connections = make(map[int]*websocket.Conn)
    clientLock          sync.Mutex
)


func HandleWebSocket(e echo.Context) error {
  conn, err := upgrader.Upgrade(e.Response(), e.Request(), nil)
	if err != nil {
		return err
	}
	defer conn.Close()
  post_id, _ := strconv.Atoi(e.Param("id"))

	user_id := 1
  connections[user_id] = conn
	

		

  for {
    _, p , err := conn.ReadMessage()
    if err != nil {
      e.Logger().Error(err)
    }

    err := models.CreateMessage(db, post_id, user_id, string(p))
    if err != nil {
      return e.Logger().Error(err)
    }
  }   
  
}
