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
	Sockets    map[string]chan []byte
	IncomingCh chan []byte
	mu         sync.Mutex
}

func (s *SocketCollection) AddSocket(id string, ch chan []byte) {
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

func (s *SocketCollection) Fanout(msg []byte) {
	for _, socket := range s.Sockets {
		socket <- msg
	}
}

func NewSocketCollection(ch chan []byte) *SocketCollection {
	s := SocketCollection{IncomingCh: ch}
	s.Sockets = make(map[string]chan []byte)
	go func() {
		for msg := range ch {
			for _, socket := range s.Sockets {
				socket <- msg
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

type RPM struct {
	MaxRPM     float32
	IdleRPM    float32
	CurrentRPM float32
}

func websocketHandler(s *SocketCollection) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ch := make(chan []byte)
		var rpm RPM
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
			reader := bytes.NewReader(msg[8:20])
			err = binary.Read(reader, binary.LittleEndian, &rpm)
			if err != nil {
				_ = conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("Failed marshaling struct: %v", err)))
				return
			}

			err = conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("RPM: %f", rpm.CurrentRPM)))

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
