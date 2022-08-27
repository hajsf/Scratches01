//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"log"
	pb "walistner/proto"

	_ "github.com/jhump/protoreflect/desc"
	_ "github.com/jhump/protoreflect/dynamic"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"google.golang.org/protobuf/types/known/emptypb"
)

// type server pb.StreamServiceServer
type server struct {
	pb.UnimplementedEventsServer
}

// ctx context.Context, in *pb.Request
// func (s *server) WaMessageStreame(in *emptypb.Empty, srv pb.StreamService_WaMessageStreameServer) error {
func (s server) StreamEvents(_ *emptypb.Empty, srv pb.Events_StreamEventsServer) error {
	if p, ok := peer.FromContext(srv.Context()); ok {
		fmt.Println("Client ip is:", p.Addr.String())
	}
	md := metadata.New(map[string]string{"Content-Type": "text/event-stream", "Connection": "keep-alive"})
	srv.SetHeader(md)
	//use wait group to allow process to be concurrent
	//var wg sync.WaitGroup
	//for i := 0; i < 5; i++ {
	for {
		waMessage := <-passer.data
		fmt.Println("************************")
		fmt.Println("received event:", waMessage.event)
		fmt.Println("received data :", waMessage.message)
		fmt.Println("************************")

		go func() {

			/*	x := NewKnownTypeRegistryWithoutWellKnownTypes()
				dataMessage := x.CreateIfKnown(waMessage.message)
				//dataMessage := NewMessageFactoryWithDefaults().NewMessage() // desc.fa dynnamic.NewKnownTypeRegistryWithDefaults() desc.Descriptor() // NewMessageFactoryWithDefaults() // make(protoreflect.ProtoMessage)
				dataByte := []byte(waMessage.message)
				fmt.Println("************************")
				fmt.Printf("%T %v %T %v\n", dataByte, dataByte, dataMessage, dataMessage)
				fmt.Println("************************")
				err := prototext.Unmarshal(dataByte, dataMessage)
				if err != nil {
					fmt.Println("Failed to prepare the prptotext: ", err)
				}
				data, err := anypb.New(dataMessage)
				if err != nil {
					fmt.Println("Failed to prepare the final datat: ", err)
				}
				fmt.Println(data) */

			resp := pb.EventSource{
				Event: waMessage.event,
				Data:  waMessage.message, //  data,
			}

			if err := srv.Send(&resp); err != nil {
				log.Printf("send error %v", err)
			}

		}()
	}

	//wg.Wait()
	return nil
}
