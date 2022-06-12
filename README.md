# Dropmail wrapper
wrapper for [dropmail.me](https://dropmail.me)

## Installation

```bash
go get github.com/balluh/dropmail
```

## Example

```go
d, err := dropmail.New(dropmail.DomainRandom)
if err != nil {
    panic(err)
}

fmt.Println(d.Address)

mails, err := d.GetInbox()
if err != nil {
    panic(err)
}

for _, mail := range mails {
    fmt.Println(mail.Text, mail.Subject, mail.From)
}
```

[waiting for email example](example/example.go)
