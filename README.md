# Dropmail wrapper
wrapper for [dropmail.me](https://dropmail.me)

## Installation

```bash
go get github.com/monkeyskid/dropmail
```

## Example

```go
d, err := dropmail.New(dropmail.DomainRandom)
if err != nil {
    panic(err)
}

fmt.Println(d.Address)

mails, err := d.GetMails()
if err != nil {
    panic(err)
}

for _, mail := range mails {
    fmt.Println(mail.Text, mail.Subject, mail.From)
}
```
