package handlers

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/FkLalita/hano/models"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"

	"net/http"
	"sync"
)

type Message struct {
	Content string `json:"content"`
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var (
	connections = make(map[*websocket.Conn]int)
	clientsLock sync.Mutex
	userCounter int
)

func HandleWebSocket(e echo.Context, db *sql.DB) error {
	conn, err := upgrader.Upgrade(e.Response(), e.Request(), nil)
	if err != nil {
		e.Logger().Error(err)
	}
	defer conn.Close()
	post_id, _ := strconv.Atoi(e.Param("id"))
	fmt.Println(post_id)

	user_id := 1
	clientsLock.Lock()
	connections[conn] = user_id
	clientsLock.Unlock()

	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			e.Logger().Error(err)
			break
		}
		strP := string(p)

		err = models.CreateMessage(db, 4, user_id, strP)
		if err != nil {
			fmt.Println(err)
			e.Logger().Error(err)
			break
		}
		msg := Message{
			Content: strP,
		}

		broadcast(msg, e)
	}
	clientsLock.Lock()
	delete(connections, conn)
	clientsLock.Unlock()

	return nil

}

func broadcast(msg Message, e echo.Context) {
	clientsLock.Lock()
	defer clientsLock.Unlock()

	for client := range connections {

		// Send the content directly
		err := client.WriteJSON(msg)
		if err != nil {
			e.Logger().Error(err)
			client.Close()
			delete(connections, client)

		}
	}
}
