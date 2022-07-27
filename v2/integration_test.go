//go:build integration
// +build integration

package infobip

import (
	"context"
	"flag"
	"fmt"
	"github.com/fairyhunter13/dotenv"
	"github.com/fairyhunter13/reflecthelper/v5"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"sync"
	"testing"
	"time"
)

var (
	once   sync.Once
	c      *APIClient
	ctx    context.Context
	sender string
)

func setupClient() {
	once.Do(func() {
		err := dotenv.Load2(
			dotenv.WithPaths(".env", "../.env"),
		)
		if err != nil {
			log.Fatalln(err)
		}

		cfg := NewConfiguration()
		cfg.Host = os.Getenv("INFOBIP_HOST")
		cfg.Debug = reflecthelper.GetBool(os.Getenv("INFOBIP_DEBUG"))
		ctx = context.WithValue(context.Background(), ContextAPIKey, os.Getenv("INFOBIP_API_KEY"))
		c = NewAPIClient(cfg)
		sender = os.Getenv("INFOBIP_SENDER")
	})
}

// Run integration tests.
// Notes: Run this test only on local, not on CI/CD.
func TestMain(m *testing.M) {
	flag.Parse()
	setupClient()

	os.Exit(m.Run())
}

func formatTime(in time.Time) string {
	return in.Format("2006-01-02 15:04:05")
}

func TestSendMessageSuccess(t *testing.T) {
	msg := fmt.Sprintf("Hello, testing with time %s. [Infobip Integration Testing]", formatTime(time.Now()))
	resp, httpResp, err := c.SendSmsApi.
		SendSmsMessage(ctx).
		SmsAdvancedTextualRequest(SmsAdvancedTextualRequest{
		Messages: &([]SmsTextualMessage{
			{
				Destinations: &([]SmsDestination{
					{
						To: os.Getenv("PHONE_NUMBER"),
					},
				}),
				From: &sender,
				Text: &msg,
			},
		}),
	}).
		Execute()

	assert.Nil(t, err)
	assert.Nil(t, resp.BulkId)
	assert.NotEmpty(t, resp.Messages)
	assert.NotNil(t, httpResp)
}

func TestSendMessageFailed(t *testing.T) {
	t.Parallel()
	t.Run("no phone number", func(t *testing.T) {
		msg := fmt.Sprintf("Hello, testing with time %s. [Infobip Integration Testing]", formatTime(time.Now()))
		resp, httpResp, err := c.SendSmsApi.
			SendSmsMessage(ctx).
			SmsAdvancedTextualRequest(SmsAdvancedTextualRequest{
			Messages: &([]SmsTextualMessage{
				{
					Destinations: &([]SmsDestination{
						{},
					}),
					From: &sender,
					Text: &msg,
				},
			}),
		}).
			Execute()

		assert.NotNil(t, err)
		assert.Nil(t, resp.BulkId)
		assert.Nil(t, resp.Messages)
		assert.NotNil(t, httpResp)
	})
}
