package main

import (
	"context"
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"log"
	"net/http"
)

var (
	oauthConfig *oauth2.Config
	oauthState  string
)

func init() {
	// Set up OAuth2 configuration
	oauthConfig = &oauth2.Config{
		ClientID:     "YOUR_CLIENT_ID",
		ClientSecret: "YOUR_CLIENT_SECRET",
		RedirectURL:  "YOUR_REDIRECT_URI",
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.profile", "https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
	// Generate random state string
	oauthState = "random"
}
func main() {
	http.HandleFunc("/", handleMain)
	http.HandleFunc("/login", handleGoogleLogin)
	http.HandleFunc("/callback", handleGoogleCallback)
	fmt.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
func handleMain(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Go Google Login Example!")
}
func handleGoogleLogin(w http.ResponseWriter, r *http.Request) {
	url := oauthConfig.AuthCodeURL(oauthState)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}
func handleGoogleCallback(w http.ResponseWriter, r *http.Request) {
	state := r.FormValue("state")
	if state != oauthState {
		fmt.Println("Invalid oauth state")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	code := r.FormValue("code")
	token, err := oauthConfig.Exchange(context.Background(), code)
	if err != nil {
		fmt.Printf("Error exchanging code: %s\n", err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	client := oauthConfig.Client(context.Background(), token)
	service, err := oauth2.New(oauth2.NoContext, client)
	if err != nil {
		fmt.Printf("Error creating OAuth2 service: %s\n", err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	userInfo, err := oauth2.New(oauth2.NoContext, client)
	if err != nil {
		fmt.Printf("Error getting user info: %s\n", err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	profile, err := userInfo.Userinfo.Get().Do()
	if err != nil {
		fmt.Printf("Error getting user profile: %s\n", err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	fmt.Fprintf(w, "Welcome, %s! Email: %s", profile.Name, profile.Email)
}
