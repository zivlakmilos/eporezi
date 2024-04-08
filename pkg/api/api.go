package api

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/zivlakmilos/eporezi/pkg/scard"
)

type API struct {
	params url.Values
	card   *scard.SCard
}

func NewAPI(card *scard.SCard) *API {
	return &API{
		card: card,
	}
}

func (a *API) ParseUrl(u string) error {
	query := strings.ReplaceAll(u, "eporezi://", "")
	if len(query) > 0 && query[len(query)-1] == '/' {
		query = query[:len(query)-1]
	}

	q, err := url.ParseQuery(query)
	if err != nil {
		return err
	}

	a.params = q
	a.fixUrl()

	return nil
}

func (a *API) IsLogin() bool {
	_, ok := a.params["loginKey"]
	return ok
}

func (a *API) IsSign() bool {
	_, ok := a.params["xmlUrl"]
	return ok
}

func (a *API) Login() error {
	resp, err := a.getLoginXml()
	if err != nil {
		return err
	}
	_ = resp
	fmt.Printf("%s\n", resp)

	return nil
}

func (a *API) Sign() error {
	return nil
}

func (a *API) fixUrl() {
	env, ok := a.params["env"]
	if !ok || len(env) == 0 {
		a.params["env"] = []string{"prod"}
		return
	}

	if !isValidEnv(env[0]) {
		a.params["env"] = []string{"prod"}
	}
}

func (a *API) getLoginXml() (string, error) {
	res, err := http.Get(fmt.Sprintf("%s/checkXML.jsp", baseUrl[a.params["env"][0]]))
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return "", fmt.Errorf("error: response with %s", res.Status)
	}

	resBuf, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(resBuf), nil
}
