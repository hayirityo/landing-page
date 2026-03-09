package main
import (
   "context"
   "fmt"
   "log"
   "os"
   "github.com/google/go-github/v81/github"
   "golang.org/x/oauth2"
)
func main() {
   ctx := context.Background()
   ts := oauth2.StaticTokenSource(
       &oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
   )
   tc := oauth2.NewClient(ctx, ts)
   client := github.NewClient(tc)
   // Example: List repositories for authenticated user
   repos, _, err := client.Repositories.List(ctx, "", nil)
   if err != nil {
       log.Fatal(err)
   }
   for _, repo := range repos {
       fmt.Println(repo.GetName())
   }
}
