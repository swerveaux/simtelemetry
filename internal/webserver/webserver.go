package webserver

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"

	uuid "github.com/satori/go.uuid"

	"github.com/gorilla/websocket"
)

type SocketCollection struct {
	Sockets    map[string]chan Telemetry
	IncomingCh chan []byte
	mu         sync.Mutex
}

func (s *SocketCollection) AddSocket(id string, ch chan Telemetry) {
	s.mu.Lock()
	s.Sockets[id] = ch
	s.mu.Unlock()
}

func (s *SocketCollection) RemoveSocket(id string) {
	s.mu.Lock()
	close(s.Sockets[id])
	delete(s.Sockets, id)
	s.mu.Unlock()
}

func NewSocketCollection(ch chan []byte) *SocketCollection {
	s := SocketCollection{IncomingCh: ch}
	s.Sockets = make(map[string]chan Telemetry)
	var rpm Telemetry
	go func() {
		for msg := range ch {
			r := bytes.NewReader(msg)
			err := binary.Read(r, binary.LittleEndian, &rpm)
			if err != nil {
				fmt.Println(err)
				return
			}
			for _, socket := range s.Sockets {
				socket <- rpm
			}
		}
	}()
	return &s
}

func Run(port int, ch chan []byte) error {
	socketCollection := NewSocketCollection(ch)
	http.HandleFunc("/", pageHandler)
	http.HandleFunc("/websocket", websocketHandler(socketCollection))

	http.ListenAndServe("0.0.0.0:"+strconv.Itoa(port), nil)
	return nil
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func websocketHandler(s *SocketCollection) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ch := make(chan Telemetry)
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}
		u, err := uuid.NewV4()
		if err != nil {
			log.Println(err)
			return
		}
		s.AddSocket(u.String(), ch)
		for msg := range ch {
			//irpm := binary.LittleEndian.Uint32(msg[16:20])
			// rpm := math.Float32frombits(irpm)
			err = conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("RPM: %f", msg.CurrentEngineRpm)))

			if err != nil {
				s.RemoveSocket(u.String())
				return
			}
		}
	}
}

func pageHandler(w http.ResponseWriter, r *http.Request) {
	page := `<html>
<head>
<title>This is a test</title>
<script type="text/javascript">
window.addEventListener("DOMContentLoaded", function() {
var exampleSocket = new WebSocket("ws://192.168.86.35:10001/websocket");
var msgfield = document.getElementById("msgfield");
exampleSocket.onmessage = function(event) {
  msgfield.innerText = event.data;
}
}, false);
</script>
</head>
<body>
<h1>Whassup?</h1><br />
<div id="msgfield"></div>
</body>
</html>`
	fmt.Fprintf(w, page)
}
