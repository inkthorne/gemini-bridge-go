package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"

	greeting "github.com/inkthorne/gemini-bridge-go/bridge"
)

func dontStream() {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-2.0-flash")
	resp, err := model.GenerateContent(ctx, genai.Text("How does AI work?"))
	if err != nil {
		log.Fatal(err)
	}

	// Print the response parts (individual text chunks)
	for _, part := range resp.Candidates[0].Content.Parts {
		fmt.Println(part)
	}

	// Print the entire response object
	fmt.Printf("\nComplete response object:\n%+v\n", resp)
}

func stream() {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()
	model := client.GenerativeModel("gemini-1.5-flash")
	iter := model.GenerateContentStream(ctx, genai.Text("Write a story about a magic backpack."))
	for {
		resp, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		greeting.PrintResponseParts(resp)
	}
}

func main() {
	// Uncomment the function you want to run
	// dontStream()
	stream()
}
