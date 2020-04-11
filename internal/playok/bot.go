package playok

import (
	"fmt"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"

	"github.com/pkg/errors"
)

const (
	playokURL        = "https://www.playok.com/"
	playokLoginURL   = "https://www.playok.com/en/login.phtml"
	playokReversiURL = "https://www.playok.com/en/reversi/"
	userAgent        = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36"
)

var (
	playokParsedURL *url.URL
)

func init() {
	var err error
	playokParsedURL, err = url.Parse(playokURL)
	if err != nil {
		panic(err)
	}
}

// Bot contains the state of an automated player on playok.com
type Bot struct {
	userName string
	password string
	browser  *http.Client
}

// NewBot initializes a new bot
func NewBot(userName, password string) *Bot {

	// do not follow redirects
	redirectHandler := func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}

	// skip checking error, this can't go wrong
	cookieJar, _ := cookiejar.New(nil)

	return &Bot{
		userName: userName,
		password: password,
		browser: &http.Client{
			CheckRedirect: redirectHandler,
			Jar:           cookieJar,
		},
	}
}

func checkCookies(jar http.CookieJar, expectedCookies []string) error {
	cookies := jar.Cookies(playokParsedURL)

	for _, cookieName := range expectedCookies {
		var found bool
		for _, cookie := range cookies {
			if cookie.Name == cookieName {
				found = true
				break
			}
		}
		if !found {
			return fmt.Errorf("missing cookie with name %s", cookieName)
		}
	}

	if len(cookies) > len(expectedCookies) {
		log.Printf("warning: received %d cookies while expecing %d cookies", len(cookies), len(expectedCookies))
	}

	return nil
}

func (bot *Bot) login() error {

	if bot.userName == "" || bot.password == "" {
		return errors.New("username and/or password are not set")
	}

	formValues := url.Values{
		"username": {bot.userName},
		"pw":       {bot.password},
		"cc":       {"0"},
	}
	request, err := http.NewRequest("POST", playokLoginURL,
		strings.NewReader(formValues.Encode()))

	if err != nil {
		return errors.Wrap(err, "building request failed")
	}
	request.Header.Add("User-Agent", userAgent)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	response, err := bot.browser.Do(request)
	if err != nil {
		return errors.Wrap(err, "sending request failed")
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusFound {
		return errors.New("unexpected status code")
	}

	expectedCookies := []string{"ku", "ksession"}

	if err := checkCookies(bot.browser.Jar, expectedCookies); err != nil {
		return err
	}

	return nil
}

func (bot *Bot) visitReversiPage() error {
	request, err := http.NewRequest("GET", playokReversiURL, nil)
	if err != nil {
		return errors.Wrap(err, "building request failed")
	}
	request.Header.Add("User-Agent", userAgent)

	// add cookie with constant value
	newCookies := []*http.Cookie{
		&http.Cookie{
			Name:  "kbeta",
			Value: "rv",
		},
	}

	bot.browser.Jar.SetCookies(playokParsedURL, newCookies)

	response, err := bot.browser.Do(request)
	if err != nil {
		return errors.Wrap(err, "sending request failed")
	}

	defer response.Body.Close()

	expectedCookies := []string{"ku", "ksession", "kbeta", "kbexp", "kt"}

	if err := checkCookies(bot.browser.Jar, expectedCookies); err != nil {
		return err
	}

	return nil
}

// Run is the entrypoint of the Bot
func (bot *Bot) Run() error {

	if err := bot.login(); err != nil {
		return errors.Wrap(err, "login failed")
	}

	if err := bot.visitReversiPage(); err != nil {
		return errors.Wrap(err, "visit reversi page failed")
	}

	// connect to ws

	// play games

	return errors.New("not implemented")
}
