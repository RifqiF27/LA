package helper

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
	"time"

	"github.com/mailersend/mailersend-go"
)



func RandInt(max int) int {
    n, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
    if err != nil {
        return 0 // Default fallback
    }
    return int(n.Int64())
}

func GenerateOTP(length int) string {
	const charset = "0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[RandInt(len(charset))]
	}
	return string(result)
}


var APIKey = "mlsn.d70c1d532bfb5ddb42cb9014964884fd2503447bcae2be9374b003ec02acf519"

func SendDummyEmail(to, subject, body string) error {
	ms := mailersend.NewMailersend(APIKey)

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	Subject := subject
	text := "This is the text content"
	html := body
	

	from := mailersend.From{
		Name:  "sender",
		Email: "MS_HWvbWM@trial-o65qngkw62wgwr12.mlsender.net",
	}

	recipients := []mailersend.Recipient{
		{
			Name:  "sender",
			Email: "rifqifadluloh27@gmail.com",
		},
	}

	// Send in 5 minute
	// sendAt := time.Now().Add(time.Minute * 5).Unix()

	// tags := []string{"foo", "bar"}

	message := ms.Email.NewMessage()

	message.SetFrom(from)
	message.SetRecipients(recipients)
	message.SetSubject(Subject)
	message.SetHTML(html)
	message.SetText(text)
	// message.SetTags(tags)
	// message.SetSendAt(sendAt)
	// message.SetInReplyTo("client-id")

	res, _ := ms.Email.Send(ctx, message)

	fmt.Printf(res.Header.Get("x-message-id"))

	return nil
}

