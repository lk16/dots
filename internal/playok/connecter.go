package playok

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"regexp"
	"strings"

	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
)

const (
	playokURL          = "https://www.playok.com/"
	playokLoginURL     = "https://www.playok.com/en/login.phtml"
	playokReversiURL   = "https://www.playok.com/en/reversi/"
	playokWebsocketURL = "wss://x.playok.com:17003/ws/"
	userAgent          = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36"
)

var (
	playokParsedURL *url.URL
	windowApRegex   = regexp.MustCompile("window.ap = (.*);")
	windowGeRegex   = regexp.MustCompile("window.ge = (.*);")
)

// connector sets up the connection to the websocket
type connector struct {
	userName string
	password string
	windowAp string
	windowGe string
	browser  *http.Client
}

func newConnector(userName, password string) *connector {

	// do not follow redirects
	redirectHandler := func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}

	// skip checking error, this can't go wrong
	cookieJar, _ := cookiejar.New(nil)

	return &connector{
		userName: userName,
		password: password,
		browser: &http.Client{
			CheckRedirect: redirectHandler,
			Jar:           cookieJar,
		},
	}
}

func (connector *connector) connect() (*websocket.Conn, error) {

	if err := connector.login(); err != nil {
		return nil, errors.Wrap(err, "login failed")
	}

	if err := connector.visitReversiPage(); err != nil {
		return nil, errors.Wrap(err, "visiting reversi page failed")
	}

	headers := make(http.Header)
	headers.Add("Origin", "https://www.playok.com")
	headers.Add("User-Agent", userAgent)

	dialer := *websocket.DefaultDialer
	dialer.Jar = connector.browser.Jar
	websocket, _, err := dialer.Dial(playokWebsocketURL, headers)

	if err != nil {
		return nil, errors.Wrap(err, "connecting failed")
	}

	return websocket, nil
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
			return errorf("missing cookie with name %s", cookieName)
		}
	}

	if len(cookies) > len(expectedCookies) {
		log.Printf("warning: received %d cookies while expecing %d cookies", len(cookies), len(expectedCookies))
	}

	return nil
}

func (connector *connector) login() error {

	if connector.userName == "" || connector.password == "" {
		return errors.New("username and/or password are not set")
	}

	formValues := url.Values{
		"username": {connector.userName},
		"pw":       {connector.password},
		"cc":       {"0"},
	}
	request, err := http.NewRequest("POST", playokLoginURL,
		strings.NewReader(formValues.Encode()))

	if err != nil {
		return errors.Wrap(err, "building request failed")
	}
	request.Header.Add("User-Agent", userAgent)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	response, err := connector.browser.Do(request)
	if err != nil {
		return errors.Wrap(err, "sending request failed")
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusFound {
		return errors.New("unexpected status code")
	}

	expectedCookies := []string{"ku", "ksession"}

	if err := checkCookies(connector.browser.Jar, expectedCookies); err != nil {
		return err
	}

	return nil
}

func (connector *connector) visitReversiPage() error {
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

	connector.browser.Jar.SetCookies(playokParsedURL, newCookies)

	response, err := connector.browser.Do(request)
	if err != nil {
		return errors.Wrap(err, "sending request failed")
	}

	defer response.Body.Close()

	expectedCookies := []string{"ku", "ksession", "kbeta", "kbexp", "kt"}

	if err := checkCookies(connector.browser.Jar, expectedCookies); err != nil {
		return err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return errors.Wrap(err, "failed to read body")
	}

	windowAp := windowApRegex.FindSubmatch(body)
	windowGe := windowGeRegex.FindSubmatch(body)

	if windowAp == nil || windowGe == nil {
		return errors.New("failed to regex match js vars")
	}

	connector.windowAp = string(windowAp[1])
	connector.windowGe = string(windowGe[1])

	return nil
}
