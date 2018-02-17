package bot

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/asaskevich/govalidator"
)

func TestGettingWebhookURL(t *testing.T) {
	cli := &http.Client{Timeout: 3 * time.Second}
	r, err := cli.Get("http://localhost:4040/api/tunnels/command_line")
	if err != nil {
		t.Errorf("Error in getting config json: %v", err)
	}
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		t.Errorf("Error in reading response: %v", err)
	}

	data := &ngrok{}
	json.Unmarshal(body, &data)

	if !govalidator.IsRequestURL(data.PublicURL) {
		t.Errorf("Wrong webhook")
	}
}

func TestInitConfig(t *testing.T) {
	config := InitConfig()

	if config.TelegramToken == "" {
		t.Error("Empty token")
	}

	if config.WebhookURL == "" {
		t.Error("Empty webhook")
	}
}
