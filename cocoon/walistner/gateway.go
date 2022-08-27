package main

import (
	"fmt"
	"log"
	pb "walistner/proto"

	"google.golang.org/grpc/peer"
	"google.golang.org/protobuf/types/known/emptypb"
)

// type server pb.StreamServiceServer
type server struct {
	pb.UnimplementedStreamServiceServer
	// pb.StreamServiceServer
}

// ctx context.Context, in *pb.Request
// func (s *server) WaMessageStreame(in *emptypb.Empty, srv pb.StreamService_WaMessageStreameServer) error {
func (s server) WaMessageStreame(_ *emptypb.Empty, srv pb.StreamService_WaMessageStreameServer) error {
	if p, ok := peer.FromContext(srv.Context()); ok {
		fmt.Println("Client ip is:", p.Addr.String())
	}
	//md := metadata.New(map[string]string{"Content-Type": "text/event-stream", "Connection": "keep-alive"})
	//srv.SetHeader(md)
	//use wait group to allow process to be concurrent
	//var wg sync.WaitGroup
	//for i := 0; i < 5; i++ {
	for {
		waMessage := <-passer.data
		fmt.Println(waMessage)
		//	wg.Add(1)
		//go func(count int64) {
		go func() {
			//		defer wg.Done()
			//w.Header().Set("Content-Type", "text/event-stream")
			//time sleep to simulate server process time
			//	time.Sleep(time.Duration(count) * time.Second)
			//
			resp := pb.Response{
				SenderType:     "Hasan test",
				MessageGroup:   waMessage.event,
				MessageSender:  waMessage.message,
				SenderName:     "",
				MessageTime:    "",
				MessageID:      "",
				MessageType:    "",
				MessageText:    "",
				MessageCaption: "",
				MessageUri:     "",
			}
			if err := srv.Send(&resp); err != nil {
				log.Printf("send error %v", err)
			}
			//log.Printf("finishing request number : %d", count)
			//}(int64(i))
		}()
	}

	//	wg.Wait()
	return nil
}
