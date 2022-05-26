package main

import (
	"fmt"
	"github.com/monkeyskid/dropmail"
	"time"
)

// waiting for email example.
func main() {
	d, err := dropmail.New(dropmail.DomainSpymailOne)
	if err != nil {
		panic(err)
	}

	fmt.Println("created an email address:", d.Address)
	fmt.Println("waiting for emails...")
	fmt.Println()

	t := time.NewTicker(1 * time.Second)
	defer t.Stop()

	for range t.C {
		emails, err := d.GetInbox()
		if err != nil {
			panic(err)
		}

		if len(emails) > 0 {
			fmt.Println("received email")
			fmt.Println("from:", emails[0].From)
			fmt.Println("content:", emails[0].Text)

			break
		}
	}
}

/*
output:

created an email address: 665vkg3z@spymail.one
waiting for emails...

received email
from: drip@mailfence.com
content: this is an email.

-- Sent with https://mailfence.com  Secure and private email
*/
