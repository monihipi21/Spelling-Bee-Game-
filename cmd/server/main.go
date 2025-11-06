package main

import (
	"context"
	"fmt"
	"log"
	"net"

	spellingbeepb "spellingbee_game/api/spellingbee/v1"
	"spellingbee_game/internal/game"

	"google.golang.org/grpc"
)

type SpellingBeeServer struct {
	spellingbeepb.UnimplementedSpellingBeeServiceServer
}

func (s *SpellingBeeServer) GetLetters(ctx context.Context, req *spellingbeepb.GetLettersRequest) (*spellingbeepb.GetLettersResponse, error) {
	newGame := game.NewGame()
	currentGame = game.NewLoggingDecorator(newGame)

	resp := &spellingbeepb.GetLettersResponse{
		Letters:      newGame.Letters,
		CenterLetter: newGame.CenterLetter,
	}
	return resp, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	spellingbeepb.RegisterSpellingBeeServiceServer(grpcServer, &SpellingBeeServer{})

	fmt.Println("Server running on port 50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

var currentGame game.GameInterface = game.NewLoggingDecorator(game.NewGame())

func (s *SpellingBeeServer) SubmitWord(ctx context.Context, req *spellingbeepb.SubmitWordRequest) (*spellingbeepb.SubmitWordResponse, error) {
	word := req.GetWord()
	valid, msg := currentGame.ValidateWord(word)
	points := 0

	if valid {
		points = currentGame.AddScore(word)
	}

	resp := &spellingbeepb.SubmitWordResponse{
		Valid:      valid,
		Score:      int32(points),
		Message:    msg,
		TotalScore: int32(currentGame.GetTotalScore()),
	}

	return resp, nil
}
