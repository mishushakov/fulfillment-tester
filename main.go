package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"io/ioutil"
	"net/http"
)

// Variables
var (
	Gateway     string // Gateway URL
	Fulfillment string // Fulfillment URL
	Port        string // Port to run the fulfillment-tester at
)

func proxy(res http.ResponseWriter, req *http.Request) {
	// Set response Headers
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Headers", "Content-Type, Cache-Control")
	res.Header().Set("Access-Control-Allow-Methods", "*")

	if req.Method == "GET" {
		// Retrieve Agent
		resp, gateway_conn_err := http.Get(Gateway)
		if gateway_conn_err != nil {
			panic(gateway_conn_err)
		}
		body, gateway_res_err := ioutil.ReadAll(resp.Body)
		if gateway_res_err != nil {
			panic(gateway_res_err)
		}

		res.Write(body)
	} else if req.Method == "POST" {
		// Proxy request to Dialogflow Gateway (Session ID is always fulfillment-tester)
		resp, gateway_conn_err := http.Post(Gateway, "application/json", req.Body)
		if gateway_conn_err != nil {
			panic(gateway_conn_err)
		}

		// Read the response from Dialogflow Gateway
		resp_body, gateway_res_err := ioutil.ReadAll(resp.Body)
		if gateway_res_err != nil {
			panic(gateway_res_err)
		}

		if resp.StatusCode != 200 {
			res.Write(resp_body)
		}

		var result map[string]interface{}
		json.Unmarshal(resp_body, &result)

		// Add (fake) session identifier to response, so our Webhook can see the session
		// Generate new json
		result["session"] = "projects/x/agent/sessions/fulfillment-tester"
		data, _ := json.Marshal(result)

		// Send it to the Fulfillment
		fulfillment_resp, fulfillment_conn_err := http.Post(Fulfillment, "application/json", bytes.NewReader(data))
		if fulfillment_conn_err != nil {
			panic(fulfillment_conn_err)
		}

		// Parse response from fulfillment
		// Set queryResult of Dialogflow Gateway response to the fulfillment's response
		// That's how Dialogflow actually responds with fulfillment option enabled (you don't have to be a Googler to tell that)
		fulfillment_messages, fulfillment_res_err := ioutil.ReadAll(fulfillment_resp.Body)
		if fulfillment_res_err != nil {
			panic(fulfillment_res_err)
		}

		var queryResult = result["queryResult"]
		json.Unmarshal(fulfillment_messages, &queryResult)

		// Convert result back to JSON and send it to the client
		output, _ := json.Marshal(result)
		res.Write(output)
	} else if req.Method == "OPTIONS" {
		// Pre-flight checks
		res.WriteHeader(200)
	} else {
		// Invalid request method
		res.WriteHeader(404)
	}
}

func main() {
	// Parse flags, setup default flags and descriptions
	flag.StringVar(&Gateway, "gateway", "https://dialogflow-web-v2.gateway.dialogflow.cloud.ushakov.co", "Dialogflow Gateway URL")
	flag.StringVar(&Fulfillment, "fulfillment", "https://us-central1-dialogflow-web-v2.cloudfunctions.net/dialogflowFirebaseFulfillment", "URL to fullfillment (remote or local)")
	flag.StringVar(&Port, "port", "8899", "Port to run the fulfillment-tester at")
	flag.Parse()

	// Log some useful information to console
	println("Dialogflow Fulfillment Tester is running 🚀")
	println("Listening on: http://localhost:" + Port)
	println("\nConnection 🔌")
	println("Gateway: " + Gateway)
	println("Fulfillment: " + Fulfillment)
	println("\nHappy Testing!")

	// Setup HTTP Server and Proxy (Handler)
	http.HandleFunc("/", proxy)
	panic(http.ListenAndServe(":"+Port, nil))
}