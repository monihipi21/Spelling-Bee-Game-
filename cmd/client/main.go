package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	spellingbeepb "spellingbee_game/api/spellingbee/v1"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()

	client := spellingbeepb.NewSpellingBeeServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	lettersResp, err := client.GetLetters(ctx, &spellingbeepb.GetLettersRequest{})
	if err != nil {
		log.Fatalf("Error getting letters: %v", err)
	}

	fmt.Print("\nToday's letters: ")
	for _, l := range lettersResp.Letters {
		if l == lettersResp.CenterLetter {
			fmt.Printf("(%s) ", l)
		} else {
			fmt.Printf("%s ", l)
		}
	}
	fmt.Println("\n")

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Spelling Bee!")

	for {
		fmt.Print("Enter word > ")
		word, _ := reader.ReadString('\n')
		word = strings.TrimSpace(word)

		if word == "exit" {
			break
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		resp, err := client.SubmitWord(ctx, &spellingbeepb.SubmitWordRequest{Word: word})
		if err != nil {
			log.Printf("Error: %v", err)
			continue
		}

		fmt.Printf("%s (Score: %d, Total: %d)\n", resp.Message, resp.Score, resp.TotalScore)
	}
}
