package controller

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/gorilla/mux"
)

const (
	// Max wait time when writing message to peer
	writeWait = 10 * time.Second

	// Max time till next pong from peer
	pongWait = 60 * time.Second

	// Send ping interval, must be less then pong wait time
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 10000
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  4096,
	WriteBufferSize: 4096,
}

var WsTickets = make(map[string] *WsTicket)

type WsTicket struct {
	ticket_id     string
	Customer_conn *websocket.Conn
	Shop_conn       *websocket.Conn
	broadcast     chan *string
}

func newWsTicket(ticket_id string, Customer_conn *websocket.Conn) *WsTicket {
	return &WsTicket{
		ticket_id:     ticket_id,
		Customer_conn: Customer_conn,
		Shop_conn:       nil,
		broadcast:     make(chan *string),
	}

}

func CustomerWs(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	vars := mux.Vars(r)

	fmt.Print("ticket_id: ")
	fmt.Println(vars["ticketID"]);

	WsTickets[vars["ticketID"]] = newWsTicket(vars["ticketID"], conn)
	// go WsTickets[r.FormValue("ticket_id")].readPump()
}

func (ticket *WsTicket) readShop() {

	ticket.Shop_conn.SetReadLimit(maxMessageSize)
	ticket.Shop_conn.SetReadDeadline(time.Now().Add(pongWait))
	ticket.Shop_conn.SetPongHandler(func(string) error { ticket.Shop_conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, jsonMessage, err := ticket.Shop_conn.ReadMessage()
		fmt.Println(string(jsonMessage))
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("unexpected close error: %v", err)
			}
			break
		}

		ticket.handleNewMessage(jsonMessage)
	}
}

func (ticket *WsTicket) handleNewMessage(jsonMessage []byte) {
	ticket.Customer_conn.WriteMessage(1, jsonMessage)
}


func ShopWs(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	vars := mux.Vars(r)
	
	WsTickets[vars["ticketID"]].Shop_conn = conn

	go WsTickets[vars["ticketID"]].readShop()

}
