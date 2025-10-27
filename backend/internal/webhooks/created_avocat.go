package webhooks

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"test/internal/models"
	"time"

	"github.com/joho/godotenv"
)
func CreatedAvocatN8nWebhook(avocat models.Avocat){

	godotenv.Load("internal/webhooks/n8n_webhook.env")
	webhookUrl:=os.Getenv("WEBHOOK_URL")
	webhookSecret:=os.Getenv("WEBHOOK_SECRET")
	if webhookUrl=="" || webhookSecret==""{
		log.Println("N8N webhook URL or secret not set")
		return
	}

	playload,err:= json.Marshal(&avocat)
	if err!=nil{
		log.Println("Error marshalling avocat data:", err)
		return
	}

	client:=&http.Client{Timeout: 5*time.Second}

	req,err:=http.NewRequest("POST",webhookUrl,bytes.NewBuffer(playload))
	if err != nil {
    	log.Printf("Failed to trigger n8n webhook: %v", err)
    	return
	}
	req.Header.Set("content-type","application/json")
	req.Header.Set("Authorisation",fmt.Sprintf("Bearer %s",webhookSecret))

	resp,err:=client.Do(req)
	if err!=nil{
		log.Printf("Error triggering n8n webhook: %v", err)
		return
	}

	defer resp.Body.Close()
	log.Printf("Triggered n8n webhook: %s (status %d)\n", webhookUrl, resp.StatusCode)

}