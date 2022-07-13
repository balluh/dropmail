package dropmail

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/rs/xid"
)

type Client struct {
	Token   string
	Address string
}

func New(domain Domain) (*Client, error) {
	token := xid.New().String()
	var data string

	if domain == "" {
		data = `{"query":"mutation {introduceSession {id, addresses {address}}}"}`
	} else {
		data = fmt.Sprintf(`{"query":"mutation {introduceSession(input: {withAddress: true, domainId: \"%v\"}) {id, addresses {address}}}"}`, domain)
	}

	req, err := http.NewRequest("POST", "https://dropmail.me/api/graphql/"+token, strings.NewReader(data))
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	b, _ := io.ReadAll(resp.Body)

	var r SessionResponse

	if err := json.Unmarshal(b, &r); err != nil {
		return nil, err
	}

	return &Client{
		Token:   token,
		Address: r.Data.IntroduceSession.Addresses[0].Address,
	}, nil
}
