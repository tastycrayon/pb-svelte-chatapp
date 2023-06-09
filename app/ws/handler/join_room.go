package handler

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/models"
	"github.com/tastycrayon/pb-svelte-chatapp/app/ws"
	"nhooyr.io/websocket"
)

const readLimit = 1024

func JoinRoom(h *ws.Hub) echo.HandlerFunc {
	return func(c echo.Context) error {
		r, w := c.Request(), c.Response()

		roomSlug := c.PathParam("roomSlug")
		if _, ok := h.Rooms[roomSlug]; !ok {
			return apis.NewApiError(http.StatusNotFound, "room not found", nil)
		}
		// get user
		authRecord, _ := c.Get(apis.ContextAuthRecordKey).(*models.Record)

		opts := &websocket.AcceptOptions{
			OriginPatterns: []string{"127.0.0.1:5173", "127.0.0.1:8090", "127.0.0.1:8080"},
		}

		conn, err := websocket.Accept(w, r, opts)
		if err != nil {
			return apis.NewApiError(http.StatusInternalServerError, "could not accept websocket connection", nil)
		}
		conn.SetReadLimit(readLimit)
		avatar := authRecord.Get("avatar").(string)
		participantId := authRecord.Id
		username := authRecord.Username()

		participant := ws.NewParticipant(h, conn, participantId, username, roomSlug, avatar)
		h.Register <- participant

		// be ready to write new messages to this user
		go participant.WriteMessage(r.Context(), h)

		// read messages from this user and forward the message to channel
		participant.ReadMessage(r.Context(), h) // subscriber

		return nil
	}
}
