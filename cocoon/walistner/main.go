package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "walistner/proto"
)

//type eventSourceMarshaler struct {
//	runtime.JSONPb
//}

// A custom marshaler implementation, that doesn't implement the delimited interface
type EventSourceMarshaler struct {
	JSONPb runtime.JSONPb
}

func (c *EventSourceMarshaler) Marshal(v interface{}) ([]byte, error) { return c.JSONPb.Marshal(v) }
func (c *EventSourceMarshaler) Unmarshal(data []byte, v interface{}) error {
	return c.JSONPb.Unmarshal(data, v)
}
func (c *EventSourceMarshaler) NewDecoder(r io.Reader) runtime.Decoder { return c.JSONPb.NewDecoder(r) }
func (c *EventSourceMarshaler) NewEncoder(w io.Writer) runtime.Encoder { return c.JSONPb.NewEncoder(w) }

//func (c *EventSourceMarshaler) ContentType(v interface{}) string           { return "Custom-Content-Type" }

func (m *EventSourceMarshaler) ContentType(v interface{}) string {
	return "text/event-stream"
}

func main() {
	gRPcPort := ":50005"
	// Create a gRPC listener on TCP port
	lis, err := net.Listen("tcp", gRPcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create a gRPC server object
	s := grpc.NewServer()
	// Attach the Greeter service to the server
	pb.RegisterStreamServiceServer(s, server{})
	//pb.RegisterEventsServer(s, server{})

	log.Println("Serving gRPC server on 0.0.0.0:50005")
	grpcTerminated := make(chan struct{})
	// Serve gRPC server
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
			close(grpcTerminated) // In case server is terminated without us requesting this
		}
	}()

	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	//	gateWayTarget := fmt.Sprintf("0.0.0.0%s", gRPcPort)
	conn, err := grpc.DialContext(
		context.Background(),
		gRPcPort,
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}
	// gwmux := runtime.NewServeMux()
	gwmux := runtime.NewServeMux(runtime.WithMarshalerOption("text/event-stream", &EventSourceMarshaler{JSONPb: runtime.JSONPb{}})) //

	//http.HandleFunc("/sse", passer.HandleSignal)
	// Register custom route for  GET /hello/{name}
	err = gwmux.HandlePath("GET", "/sse", passer.HandleSignal)
	if err != nil {
		fmt.Println("Error:", err)
	}

	//gwmux.HandlePath("GET", "/", NotFoundHandler) // http.NotFound
	// Register custom route for  GET /hello/{name}
	err = gwmux.HandlePath("GET", "/hello/{name}", func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		w.Write([]byte("hello " + pathParams["name"]))
	})
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Register Event handler
	err = pb.RegisterStreamServiceHandler(context.Background(), gwmux, conn)
	//err = pb.RegisterEventsHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    ":8090",
		Handler: allowCORS(gwmux),
	}

	// log.Fatal(gwServer.ListenAndServe()) // <- This line alone could be enough ang no need for all the lines after,
	go gwServer.ListenAndServe()

	log.Println("Serving gRPC-Gateway on http://localhost:8090")

	fmt.Println("run POST request of: http://localhost:8090/v1/rawdata/stream")
	fmt.Println("or run curl -X GET -k http://localhost:8090/v1/rawdata/stream")
	passer.data <- sseData{
		event:   "start",
		message: "the server",
	}
	// the application is probably doing other things and you will want to be
	// able to shutdown cleanly; passing in a context is a good method..
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // Ensure cancel function is called eventually

	grpcWebTerminated := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := gwServer.ListenAndServe(); err != nil {
			fmt.Printf("Web server (GRPC) shutdown: %s", err)
		}
		close(grpcWebTerminated) // In case server is terminated without us requesting this
	}()

	// Wait for the web server to shutdown OR the context to be cancelled...
	select {
	case <-ctx.Done():
		// Shutdown the servers (there are shutdown commands to request this)
	case <-grpcTerminated:
		// You may want to exit if this happens (will be due to unexpected error)
	case <-grpcWebTerminated:
		// You may want to exit if this happens (will be due to unexpected error)
	}
	// Wait for the goRoutines to complete
	<-grpcTerminated
	<-grpcWebTerminated

	// Listen to Ctrl+C (you can also do something else that prevents the program from exiting)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c
	if client.IsConnected() {
		passer.data <- sseData{
			event:   "notification",
			message: "Server is shut down at the host machine...",
		}
		client.Disconnect()
	}
}

// allowCORS allows Cross Origin Resoruce Sharing from any origin.
// Don't do this without consideration in production systems.
func allowCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
		}
		h.ServeHTTP(w, r)
	})
}
