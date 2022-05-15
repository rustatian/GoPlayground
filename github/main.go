package main

import (
	"context"
	"log"
	"net/http"

	"github.com/google/go-github/v42/github"
	"golang.org/x/oauth2"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

type user struct {
	IssueTitle        string
	UserBio           string
	Twitter           string
	NumberOfFollowers int
	Email             string
	Username          string
	LoginURL          string
	Company           string
	IssueLink         string
}

func main() {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: ""})
	ghClient := github.NewClient(oauth2.NewClient(ctx, ts))

	users := make([]*user, 0, 10)

	for i := 1; i < 1009; i++ {
		u := &user{}

		log.Printf("issue: %d\n", i)
		issue, resp, err := ghClient.Issues.Get(context.Background(), "", "", i)
		if err != nil {
			continue
		}
		if issue.IsPullRequest() {
			continue
		}

		if resp.StatusCode != http.StatusOK {
			log.Println("bad issue")
			continue
		}

		if *issue.User.Login == "" || *issue.User.Login == "" || *issue.User.Login == "" {
			log.Println("skipped")
			continue
		}

		u.Username = *issue.User.Login
		u.IssueTitle = *issue.Title

		usr, resp2, err := ghClient.Users.Get(context.Background(), *issue.User.Login)
		if err != nil {
			continue
		}

		if resp2.StatusCode != http.StatusOK {
			log.Println("bad usr")
			continue
		}

		if usr.Company != nil {
			u.Company = *usr.Company
		} else {
			u.Company = "no-company"
		}

		u.LoginURL = *usr.HTMLURL
		u.IssueLink = *issue.HTMLURL

		if usr.Bio != nil {
			u.UserBio = *usr.Bio
		}

		if usr.Email != nil {
			u.Email = *usr.Email
		}

		if usr.TwitterUsername != nil {
			u.Twitter = *usr.TwitterUsername
		}

		if usr.Followers != nil {
			u.NumberOfFollowers = *usr.Followers
		}

		users = append(users, u)
	}

	source := oauth2.StaticTokenSource(&oauth2.Token{TokenType: "Bearer", RefreshToken: "", AccessToken: ""})
	client := oauth2.NewClient(context.Background(), source)

	srv, err := sheets.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	spreadsheetId := ""
	writeRange := "A2:I10000"

	var vr sheets.ValueRange
	for i := 0; i < len(users); i++ {
		tmpVal := []interface{}{users[i].Username, users[i].Company, users[i].UserBio, users[i].Email, users[i].NumberOfFollowers, users[i].Twitter, users[i].IssueLink, users[i].IssueTitle, users[i].LoginURL}
		vr.Values = append(vr.Values, tmpVal)
	}

	updCall := srv.Spreadsheets.Values.Update(spreadsheetId, writeRange, &vr)
	updCall.ValueInputOption("RAW")

	_, err = updCall.Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}
}
