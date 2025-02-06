package avanza_api

import (
	"encoding/json"
	"fmt"
	"log"

	api_models "github.com/Melle101/mom-bot-v3/avanza-api/api-models"
	"github.com/gorilla/websocket"
)

const avanzaWS = "wss://www.avanza.se/_push/cometd"

func (c *ApiClient) ConnectWebSocketDeals(orderbookID string) {

	// Set WebSocket headers for authentication
	headers := api_models.DEF_HEADERS
	//headers.Add("Cookie", "AZAMFA="+c.credInfo.Secret_token)
	// Establish connection
	conn, resp, err := websocket.DefaultDialer.Dial(avanzaWS, headers)
	fmt.Println(resp)
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	// Handshake with Avanza's WebSocket server
	handshakeMsg := fmt.Sprintf(`[
    	{
        	"advice": {"timeout": 60000, "interval": 0},
         	"channel": "/meta/handshake",
          	"ext": {"subscriptionId": "%s"},
          	"minimumVersion": "1.0",
           	"supportedConnectionTypes": [
            	"websocket",
             	"long-polling",
              	"callback-polling"
            ],
            "version": "1.0"
        }
        ]`, c.CredInfo.Push_sub_id)
	err = conn.WriteMessage(websocket.TextMessage, []byte(handshakeMsg))
	if err != nil {
		log.Fatalf("Handshake failed: %v", err)
	}

	_, message, err := conn.ReadMessage()
	var response api_models.HandshakeResponse
	err = json.Unmarshal(message, &response)
	if err != nil {
		log.Printf("Error parsing JSON: %v", err)
	}
	// Subscribe to new deals for the given orderbook ID
	subscribeMsg := fmt.Sprintf(`[{ "channel": "/meta/subscribe", "subscription": "/quotes/%s", "clientId": "%s", "connectionType": "websocket"}]`, orderbookID, response[0].ClientID)
	err = conn.WriteMessage(websocket.TextMessage, []byte(subscribeMsg))
	if err != nil {
		log.Fatalf("Subscription failed: %v", err)
	}

	fmt.Println("Subscribed to market deals for orderbook ID:", orderbookID)

	// Listen for new messages
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Error reading message: %v", err)
			break
		}
		fmt.Println("Received:", string(message))
	}
}
