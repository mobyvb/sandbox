package main

    import (
        "io/ioutil"
        "log"
        "net/http"
        "os"

        "golang.org/x/oauth2"
        "golang.org/x/oauth2/github"
    )

    var (
        githubOauthConfig = &oauth2.Config{
            RedirectURL:  "http://localhost:8080/oauth2/github/callback",
            ClientID:     os.Getenv("OAUTH2_CLIENT_ID"),
            ClientSecret: os.Getenv("OAUTH2_CLIENT_SECRET"),
            Scopes:       []string{},
            Endpoint:     github.Endpoint,
        }
    )

    func main() {
        http.HandleFunc("/", handleIndex)
        http.HandleFunc("/oauth2/github/login", handleGithubLogin)
        http.HandleFunc("/oauth2/github/callback", handleGithubCallback)

        http.ListenAndServe(":8080", nil)
    }

    func handleIndex(w http.ResponseWriter, r *http.Request) {
        file, _ := ioutil.ReadFile("index.html")
        w.Header().Set("Content-Type", "text/html")
        w.Write(file)
    }

    func handleGithubLogin(w http.ResponseWriter, r *http.Request) {
        url := githubOauthConfig.AuthCodeURL("state", oauth2.AccessTypeOnline)
        http.Redirect(w, r, url, http.StatusFound)
    }

    func handleGithubCallback(w http.ResponseWriter, r *http.Request) {
        code := r.URL.Query().Get("code")
        token, err := githubOauthConfig.Exchange(r.Context(), code)

        if err != nil {
            http.Error(w, "Failed to get token", http.StatusInternalServerError)
            return
        }

        log.Printf("User authenticated. Access token: %s", token.AccessToken)
    }
    

  -