package main

import "net/http"
import "fmt"
import "log"
import "html"
import "os"
import "runtime"
import "time"
import "sort"
import "strings"
import "github.com/vmihailenco/redis"

func sortit(m map[string]string) []string {
  arr := make([]string, len(m))
  for key := range m {
    arr = append(arr,m[key])
  }
  sort.Strings(arr)
  return arr
}

func main() {

  password := ""  // no password set
  db := int64(-1) // use default DB
  client := redis.NewTCPClient("localhost:6379", password, db)

  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    log.Println("incoming request ", r.Header.Get("x-token"))
    hostname, err := os.Hostname()
    when := time.Now().Format("Mon Jan 2 15:04:05 -0700 MST 2006")
    if err != nil {
      fmt.Printf("error: ", err)
      return
    }

    fmt.Fprintf(w, "   _____       _                    \n")
    fmt.Fprintf(w, "  / ____|     | |                   \n")
    fmt.Fprintf(w, " | |  __  ___ | | __ _ _ __   __ _  \n")
    fmt.Fprintf(w, " | | |_ |/ _ \\| |/ _` | '_ \\ / _` | \n")
    fmt.Fprintf(w, " | |__| | (_) | | (_| | | | | (_| | \n")
    fmt.Fprintf(w, "  \\_____|\\___/|_|\\__,_|_| |_|\\__, | \n")
    fmt.Fprintf(w, "                              __/ | \n")
    fmt.Fprintf(w, "                             |___/  \n\n")

    token := r.Header.Get("x-token")
    members := client.SMembers( "route::"+token )
    vals := members.Val()
    fmt.Fprintf(w, strings.Join(vals, ", "))

    fmt.Fprintf(w, "\nHello World\n")
    fmt.Fprintf(w, "Path->%s\n", html.EscapeString(r.URL.Path))
    fmt.Fprintf(w, "Golang %s\n", runtime.Version())
    fmt.Fprintf(w, "%s %s %s %s %d %s\n", runtime.GOOS, hostname, runtime.Version(), when, runtime.NumCPU(), runtime.GOARCH)
    fmt.Fprintf(w, "headers->\n")
    keys := make([]string,0)
    for it := range r.Header {
      keys = append(keys,it)
    }
    sort.Strings(keys)
    for it := range keys {
      fmt.Fprintf(w, "%s: %s\n", keys[it], r.Header.Get(keys[it]))
    }

  })

  log.Fatal(http.ListenAndServe(":31338", nil))
} 
