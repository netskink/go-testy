
package main

import "github.com/google/generative-ai-go/genai"
import "google.golang.org/api/option"

import (
    "fmt"
    "log"
    "context"
    "os"
)

func main() {
    fmt.Println("hello from go")
    ctx := context.Background()
    // Access your API key as an environment variable (see "Set up your API key" above)
    client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GOOGLE_API_KEY")))
    if err != nil {
      log.Fatal(err)
    }
    defer client.Close()

    // For text-only input, use the gemini-pro model
    model := client.GenerativeModel("gemini-pro")
    resp, err := model.GenerateContent(ctx, genai.Text("Write a story about a magic backpack."))
    if err != nil {
      log.Fatal(err)
    }
    fmt.Println("resp: ", resp)
    fmt.Println("*resp: ", *resp)
    //fmt.Println("resp[0]: ", *resp[0])
    //fmt.Println("resp[1]: ", *resp[1])
    fmt.Println("resp.feedback: ", resp{1})
}
