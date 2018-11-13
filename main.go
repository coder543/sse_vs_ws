package main

import (
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/AndrewBurian/eventsource"
	"github.com/gorilla/websocket"
	"github.com/nytimes/gziphandler"
	"golang.org/x/crypto/acme/autocert"
)

var upgrader = websocket.Upgrader{} // use default options


// note: the gzip implementation being used right now does not flush the buffer until a certain limit is hit,
// therefore the ping never gets sent if it is too short.
//
// TODO: use a flushable gzip implementation
const ping = "pingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingpingping"

func wsEcho(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	sampleTotal := time.Duration(0)
	var i uint64
	for i = 0; i < 10000; i++ {
		start := time.Now()
		err = c.WriteMessage(websocket.TextMessage, []byte(ping))
		if err != nil {
			log.Println("write:", err)
			break
		}
		_, _, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		elapsed := time.Since(start)
		sampleTotal += elapsed
		if i%100 == 0 && i != 0 {
			log.Printf("%d%%: %s\n", i/100, sampleTotal/time.Duration(i))
		}
	}
	log.Printf("ping/pong average: %s\n", sampleTotal/time.Duration(i))
}

var sseSync = make(chan struct{})

func sseEcho(w http.ResponseWriter, r *http.Request) {
	client := eventsource.NewClient(w)
	if client == nil {
		http.Error(w, "bad connection", 400)
		return
	}

	sampleTotal := time.Duration(0)
	var i uint64
	for i = 0; i < 10000; i++ {
		start := time.Now()
		err := client.Send(eventsource.DataEvent(ping))
		if err != nil {
			log.Println("write:", err)
			break
		}
		<-sseSync
		elapsed := time.Since(start)
		sampleTotal += elapsed
		if i%100 == 0 && i != 0 {
			log.Printf("%d%%: %s\n", i/100, sampleTotal/time.Duration(i))
		}
	}
	log.Printf("ping/pong average: %s\n", sampleTotal/time.Duration(i))

	client.Wait()
}

func ssePong(w http.ResponseWriter, r *http.Request) {
	// read the body just to be fair
	_, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("pong failed:", err)
	}
	sseSync <- struct{}{}
	io.WriteString(w, "")
}

func home(w http.ResponseWriter, r *http.Request) {
	homeTemplate.Execute(w, r.Host)
}

func main() {
	http.HandleFunc("/ws_echo", wsEcho)
	http.Handle("/sse_echo", gziphandler.GzipHandler(http.HandlerFunc(sseEcho)))
	http.HandleFunc("/sse_pong", ssePong)
	http.HandleFunc("/", home)
	m := &autocert.Manager{
		Cache:      autocert.DirCache("secret-dir"),
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("ceres0.space", "www.ceres0.space"),
	}
	s := &http.Server{
		Addr:      "0.0.0.0:https",
		TLSConfig: m.TLSConfig(),
	}
	log.Fatal(s.ListenAndServeTLS("", ""))
}

var homeTemplate = template.Must(template.New("").Parse(`
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<script>  
window.addEventListener("load", function(evt) {
    var output = document.getElementById("output");
    var input = document.getElementById("input");
    var ws;
    var sse;
    var print = function(message) {
        var d = document.createElement("div");
        d.innerHTML = message;
        output.appendChild(d);
    };
    document.getElementById("open_ws").onclick = function(evt) {
        if (ws) {
            return false;
        }
        ws = new WebSocket("wss://{{.}}/ws_echo");
        ws.onopen = function(evt) {
            print("OPEN");
        }
        ws.onclose = function(evt) {
            print("CLOSE");
            ws = null;
        }
        ws.onmessage = function(evt) {
			ws.send(evt.data);
        }
        ws.onerror = function(evt) {
            print("ERROR: " + evt.data);
        }
        return false;
    };
    document.getElementById("close_ws").onclick = function(evt) {
        if (!ws) {
            return false;
        }
        ws.close();
        return false;
    };
    document.getElementById("open_sse").onclick = function(evt) {
        if (sse) {
            return false;
        }
        sse = new EventSource("https://{{.}}/sse_echo");
        sse.onopen = function(evt) {
            print("OPEN");
        }
        sse.onmessage = function(evt) {
			fetch('https://{{.}}/sse_pong', {body: evt.data, method: 'post'}).then(data=>{return data});
        }
        sse.onerror = function(evt) {
            print("ERROR: " + evt.data);
        }
        return false;
    };
    document.getElementById("close_sse").onclick = function(evt) {
        if (!sse) {
            return false;
        }
		print("CLOSE");
		sse.close();
		sse = null;
        return false;
    };
});
</script>
</head>
<body>
<table>
<tr><td valign="top" width="50%">
<p>Click "Open" to create a connection to the server, 
"Send" to send a message to the server and "Close" to close the connection. 
You can change the message and send multiple times.
<p>
<form>
<button id="open_ws">Open WS</button>
<button id="close_ws">Close WS</button>
<button id="open_sse">Open SSE</button>
<button id="close_sse">Close SSE</button>
</form>
</td><td valign="top" width="50%">
<div id="output"></div>
</td></tr></table>
</body>
</html>
`))
