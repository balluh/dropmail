package dropmail

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

type Mail struct {
	Text    string `json:"text"`
	Subject string `json:"headerSubject"`
	From    string `json:"fromAddr"`
}

func (c *Client) GetInbox() ([]Mail, error) {
	data := `{"query":"query {sessions {id, expiresAt, mails {fromAddr, toAddr, text, headerSubject}}}"}`

	req, err := http.NewRequest("POST", "https://dropmail.me/api/graphql/"+c.Token, strings.NewReader(data))
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	b, _ := io.ReadAll(resp.Body)

	var r InboxResponse

	if err := json.Unmarshal(b, &r); err != nil {
		return nil, err
	}

	var mails []Mail

	for _, mail := range r.Data.Sessions[0].Mails {
		mails = append(mails, Mail{
			Text:    mail.Text,
			Subject: mail.Subject,
			From:    mail.From,
		})
	}

	return mails, nil
}
