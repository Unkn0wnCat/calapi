package ghost

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"
)

type GhostAPI struct {
	BaseURL string
	Jar     *cookiejar.Jar
}

func (api *GhostAPI) get(path string, query string) (response *http.Response, err error) {
	requestUrl, err := url.JoinPath(api.BaseURL, path)
	if err != nil {
		return nil, err
	}

	if api.Jar == nil {
		api.Jar, err = cookiejar.New(nil)
		if err != nil {
			return nil, err
		}
	}

	client := http.Client{Jar: api.Jar}

	response, err = client.Get(requestUrl + query)
	if err != nil {
		return response, err
	}

	if response.StatusCode > 299 {
		return response, fmt.Errorf("server returned status %d", response.StatusCode)
	}

	return response, nil
}

func (api *GhostAPI) post(path string, data interface{}) (response *http.Response, err error) {
	requestUrl, err := url.JoinPath(api.BaseURL, path)
	if err != nil {
		return nil, err
	}

	if api.Jar == nil {
		api.Jar, err = cookiejar.New(nil)
		if err != nil {
			return nil, err
		}
	}

	client := http.Client{Jar: api.Jar}

	marshaledData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	dataReader := bytes.NewReader(marshaledData)

	response, err = client.Post(requestUrl, "application/json", dataReader)
	if err != nil {
		return response, err
	}

	if response.StatusCode > 299 {
		return response, fmt.Errorf("server returned status %d", response.StatusCode)
	}

	return response, nil
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (api *GhostAPI) Login(username string, password string) (token string, err error) {
	data := LoginRequest{
		Username: username,
		Password: password,
	}

	response, err := api.post("/api/admin/session/", data)
	if err != nil {
		return "", err
	}

	cookies := response.Cookies()

	for _, cookie := range cookies {
		if cookie.Name != "ghost-admin-api-session" {
			continue
		}

		token = cookie.Value
		break
	}

	return token, nil
}

type UserData struct {
	Users []User `json:"users"`
}

type User struct {
	ID                                   string    `json:"id"`
	Name                                 string    `json:"name"`
	Slug                                 string    `json:"slug"`
	Email                                string    `json:"email"`
	ProfileImage                         string    `json:"profile_image"`
	CoverImage                           string    `json:"cover_image"`
	Bio                                  string    `json:"bio"`
	Website                              string    `json:"website"`
	Location                             string    `json:"location"`
	Facebook                             string    `json:"facebook"`
	Twitter                              string    `json:"twitter"`
	Accessibility                        string    `json:"accessibility"`
	Status                               string    `json:"status"`
	MetaTitle                            string    `json:"meta_title"`
	MetaDescription                      string    `json:"meta_description"`
	Tour                                 string    `json:"tour"`
	LastSeen                             time.Time `json:"last_seen"`
	CommentNotifications                 bool      `json:"comment_notifications"`
	FreeMemberSignupNotification         bool      `json:"free_member_signup_notification"`
	PaidSubscriptionStartedNotification  bool      `json:"paid_subscription_started_notification"`
	PaidSubscriptionCanceledNotification bool      `json:"paid_subscription_canceled_notification"`
	CreatedAt                            time.Time `json:"created_at"`
	UpdatedAt                            time.Time `json:"updated_at"`
	Roles                                []Role    `json:"roles"`
	Url                                  string    `json:"url"`
}

type Role struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (api *GhostAPI) UserSelf() (user *UserData, err error) {
	response, err := api.get("/api/admin/users/me", "?include=roles")
	if err != nil {
		return nil, err
	}

	user = &UserData{}

	decoder := json.NewDecoder(response.Body)

	err = decoder.Decode(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
