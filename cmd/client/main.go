package main

import (
	"context"
	"io"
	"log"
	"time"

	"github.com/google/uuid"
	newsv1 "github.com/supLano/go-grpc-proto/api/news/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

func main() {
	conn, err := grpc.NewClient(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := newsv1.NewNewsServiceClient(conn)

	for i := 0; i < 5; i++ {
		response, err := client.CreateNews(context.Background(), &newsv1.CreateNewsRequest{
			Id:       uuid.New().String(),
			Topic:    "Farming",
			Language: "English",
			Country:  "France",
			Author:   "John Man",
			Content:  "xxx-xxx-xxx",
			Keywords: []string{"farming", "agriculture", "food"},
		})
		if err != nil {
			log.Fatalf("could not create news: %v", err)
		}
		log.Printf("News created: %v", response)
	}

	GottenNewsList := make([]*newsv1.ListNewsResponse, 0)
	stream, error := client.ListNews(context.Background(), &emptypb.Empty{})
	if error != nil {
		log.Fatalf("could not get news: %v", error)
	}
	for {
		response, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("could not receive news: %v", err)
		}
		GottenNewsList = append(GottenNewsList, response)
		// log.Printf("News got: %v", response)
	}

	// clientStream , error := client.UpdateNews(context.Background())

	// if error != nil {
	// 	log.Fatalf("could not open stream for update news: %v", error)
	// }
	// for _, news := range GottenNewsList {
	// 	// log.Printf("News got: %v", news)

	// 	if error != nil {
	// 		log.Fatalf("could not open stream for update news: %v", error)
	// 	}
	// 	error = clientStream.Send(&newsv1.UpdateNewsRequest{
	// 		Id:       news.Id,
	// 		Topic:    "Farming updated",
	// 		Language: "English updated",
	// 		Country:  "France updated",
	// 		Author:   "John Man updated",
	// 		Content:  "xxx-xxx-xxx updated",
	// 		Keywords: []string{"farming updated", "agriculture updated", "food updated"},
	// 	})
	// 	time.Sleep(1 * time. Second)
	// }
	// response, error := clientStream.CloseAndRecv()
	// if error != nil {
	// 	log.Fatalf("could not close stream for update news: %v", error)
	// }
	// log.Printf("News updated: %v", response)

	clientStream, error := client.DeleteNews(context.Background())

	waitChannel := make(chan struct{})
 
	go func() {
		for {
			response, err := clientStream.Recv()
			if err == io.EOF {
				clientStream.CloseSend()
				close(waitChannel)
				break
			}
			if err != nil {
				log.Fatalf("could not receive stream for delete news: %v", err)
			}
			log.Printf("News deleted: %v", response)
		}

	}()

	for _, news := range GottenNewsList {
		// log.Printf("News got: %v", news)

		if error != nil {
			log.Fatalf("could not open stream for delete news: %v", error)
		}
		error = clientStream.Send(&newsv1.DeleteNewsRequest{
			Id: news.Id,
		})
		time.Sleep(1 * time.Second)
	}
	clientStream.CloseSend()

	<-waitChannel

	GottenNewsList = make([]*newsv1.ListNewsResponse, 0)
	stream, error = client.ListNews(context.Background(), &emptypb.Empty{})
	if error != nil {
		log.Fatalf("could not get news: %v", error)
	}
	for {
		response, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("could not receive news: %v", err)
		}
		GottenNewsList = append(GottenNewsList, response)
	}
	log.Println("News list after deletion:", GottenNewsList)

}
