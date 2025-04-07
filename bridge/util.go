package greeting

import (
	"fmt"

	"github.com/google/generative-ai-go/genai"
)

// PrintResponse processes and prints a streaming response from the Gemini API.
// It extracts text parts from the first candidate in the response and prints them
// to standard output. If the response contains no candidates, the function
// returns without printing anything.
//
// Parameters:
//   - resp: A pointer to a GenerateContentResponse from the Gemini API
func PrintResponse(resp *genai.GenerateContentResponse) {
	// Return if there are no candidates in the response.

	if len(resp.Candidates) == 0 {
		return
	}

	// Print each text part from the response.

	for _, part := range resp.Candidates[0].Content.Parts {
		fmt.Print(part)
	}
}

func PrintResponseParts(resp *genai.GenerateContentResponse) {
	// Return if there are no candidates in the response.

	candidateCount := len(resp.Candidates)

	if candidateCount == 0 {
		return
	}

	// Print the number of candidates.
	fmt.Printf("Number of candidates: %d\n", candidateCount)

	for i, candidate := range resp.Candidates {
		fmt.Printf("(( CANDIDATE %d ))\n", i)
		fmt.Printf("%#v\n", candidate.Content)
	}
}
