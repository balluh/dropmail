package dropmail

type SessionResponse struct {
	Data struct {
		IntroduceSession struct {
			ID        string `json:"id"`
			Addresses []struct {
				Address string `json:"address"`
			} `json:"addresses"`
		} `json:"introduceSession"`
	} `json:"data"`
}

type InboxResponse struct {
	Data struct {
		Sessions []struct {
			Mails []struct {
				Text    string `json:"text"`
				Subject string `json:"headerSubject"`
				From    string `json:"fromAddr"`
			} `json:"mails"`
		} `json:"sessions"`
	} `json:"data"`
}

type DomainResponse struct {
	Data struct {
		Domains []struct {
			Name string `json:"name"`
			ID   string `json:"id"`
		} `json:"domains"`
	} `json:"data"`
}
