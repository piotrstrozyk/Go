// package main

// import (
//     "fmt"
//     "net/http"
//     "os"
// )

// func main() {
//     http.HandleFunc("/", handler)
//     err := http.ListenAndServe(":3000", nil)
//     if err != nil {
//         fmt.Printf("Error starting server: %v", err)
//         os.Exit(1)
//     }
// }

// func handler(w http.ResponseWriter, r *http.Request) {
//     html, err := os.ReadFile("index.html")
//     if err != nil {
//         fmt.Fprintf(w, "Error reading file: %v", err)
//         return
//     }
//     w.Header().Set("Content-Type", "text/html")
//     fmt.Fprintf(w, "%s", html)
// }

package main

import (
    "encoding/json"
    "fmt"
    "log"
    "math/rand"
    "net/http"
    "os"
    "sync"
    "time"
)

type Post struct {
    ID                   string `json:"id"`
    Date                 string `json:"date"`
    Year                 string `json:"year"`
    Type                 string `json:"type"`
    Country              string `json:"country"`
    Area                 string `json:"area"`
    Location             string `json:"location"`
}

var (
    posts = make(map[string]Post)
    mu    sync.Mutex
)

func loadPosts() {
    file, err := os.ReadFile("posts.json")
    if err != nil {
        log.Fatal(err)
    }

    var allPosts []Post
    err = json.Unmarshal(file, &allPosts)
    if err != nil {
        log.Fatal(err)
    }

    rand.Seed(time.Now().UnixNano())
    for i := 0; i < 10; i++ {
        index := rand.Intn(len(allPosts))
        p := allPosts[index]
        p.ID = fmt.Sprintf("%s-%d", p.ID, i) // Add index to ID
        posts[p.ID] = p
    }
}

func main() {
    http.HandleFunc("/posts", postsHandler)
    http.HandleFunc("/posts/", postHandler)

    loadPosts()

    fmt.Println("Starting server on port 3000")
    err := http.ListenAndServe(":3000", nil)
    if err != nil {
        log.Fatalf("Error starting server: %v", err)
    }
}

func postsHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "GET":
        getPosts(w, r)
    case "POST":
        createPost(w, r)
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

func postHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "GET":
        getPost(w, r)
    case "DELETE":
        deletePost(w, r)
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

func getPosts(w http.ResponseWriter, r *http.Request) {
    mu.Lock()
    defer mu.Unlock()

    json.NewEncoder(w).Encode(posts)
}

func createPost(w http.ResponseWriter, r *http.Request) {
    var p Post
    err := json.NewDecoder(r.Body).Decode(&p)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    mu.Lock()
    posts[p.ID] = p
    mu.Unlock()
}

func getPost(w http.ResponseWriter, r *http.Request) {
    id := r.URL.Path[len("/posts/"):]

    mu.Lock()
    p, ok := posts[id]
    mu.Unlock()

    if !ok {
        http.Error(w, "Post not found", http.StatusNotFound)
        return
    }

    json.NewEncoder(w).Encode(p)
}

func deletePost(w http.ResponseWriter, r *http.Request) {
    id := r.URL.Path[len("/posts/"):]

    mu.Lock()
    delete(posts, id)
    mu.Unlock()
}