package infobip

import (
	"flag"
	"github.com/fairyhunter13/dotenv"
	"log"
	"os"
	"sync"
	"testing"
	"time"
)

var (
	once sync.Once
	c    *APIClient
)

func setupClient() {
	once.Do(func() {
		err := dotenv.Load2(
			dotenv.WithPaths(".env", "../.env"),
		)
		if err != nil {
			log.Fatalln(err)
		}

		// c, err = New(
		// 	WithBasicAuth(os.Getenv(""), os.Getenv("")),
		// 	WithCustomIPs(os.Getenv("")),
		// )
		// if err != nil {
		// 	log.Fatalln(err)
		// }
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
