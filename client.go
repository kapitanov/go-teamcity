package teamcity

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type client struct {
	url      string
	userInfo *url.Userinfo
	c        *http.Client
}

// GuestAuth creates an Authorizer for non-authorized access
func GuestAuth() Authorizer {
	return guestAuthorizer{}
}

// BasicAuth creates an Authorizer for an authorized access (using HTTP Basic Auth)
func BasicAuth(username string, password string) Authorizer {
	return basicAuthorizer{username, password}
}

// NewClient creates new TeamCity Client
func NewClient(url string, auth Authorizer) Client {
	url = strings.TrimRight(url, "//")

	if auth == nil {
		auth = GuestAuth()
	}

	url = auth.ResolveBaseURL(url)

	c := client{}
	c.url = url
	c.userInfo = auth.GetUserInfo()
	c.c = new(http.Client)

	client := Client(c)
	return client
}

func (c client) httpGet(uri string, query *url.Values, result interface{}) error {
	uri = c.url + uri
	if query != nil {
		uri = uri + "?" + query.Encode()
	}

	request, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		errorf("GET %s failed with %s (unable to create HTTP request)", uri, err)
		return err
	}

	request.Header.Set("Accept", "application/json")

	if c.userInfo != nil {
		username := c.userInfo.Username()
		password, _ := c.userInfo.Password()
		request.SetBasicAuth(username, password)
	}

	response, err := c.c.Do(request)
	if err != nil {
		errorf("GET %s failed with %s", uri, err)
		return err
	}

	defer response.Body.Close()
	debugf("GET %s -> %s", uri, response.Status)

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		errorf("GET %s failed with %s (unable to read response)", uri, err)
		return err
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		errorf("GET %s failed with %s (malformed response)", uri, err)
		return err
	}

	/*
		session := &napping.Session{}
		session.Header = &http.Header{}
		session.Header.Set("Accept", "application/json")
		session.Userinfo = c.userInfo

		response, err := session.Get(c.url+uri, query, &result, nil)
		if err != nil {
			errorf("GET %s failed with %s", uri, err)
			return err
		}

		debugf("GET %s -> %d", response.Url, response.Status())*/
	return nil
}

type guestAuthorizer struct{}

func (a guestAuthorizer) ResolveBaseURL(baseURL string) string {
	return baseURL + "/guestAuth/app/rest"
}

func (a guestAuthorizer) GetUserInfo() *url.Userinfo {
	return nil
}

type basicAuthorizer struct {
	username string
	password string
}

func (a basicAuthorizer) ResolveBaseURL(baseURL string) string {
	return baseURL + "/httpAuth/app/rest"
}

func (a basicAuthorizer) GetUserInfo() *url.Userinfo {
	return url.UserPassword(a.username, a.password)
}

func errorf(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	log.Panicf("[teamcity] %s", msg)
}

func debugf(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	log.Printf("[teamcity] %s", msg)
}
