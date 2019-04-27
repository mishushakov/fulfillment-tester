package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"io/ioutil"
	"net/http"
)

// Declaring Variables
var (
	Project     string // Project ID
	Fulfillment string // Fulfillment Host
	Port        string // Port to run the fulfillment-tester at
)

func proxy(res http.ResponseWriter, req *http.Request) {
	// Set response Headers
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Headers", "Content-Type, Cache-Control")
	res.Header().Set("Access-Control-Allow-Methods", "*")

	if req.Method == "POST" {
		// Proxy request to Dialogflow Gateway (Session ID is always fulfillment-tester)
		resp, _ := http.Post("https://"+Project+".gateway.dialogflow.cloud.ushakov.co/fulfillment-tester", "application/json", req.Body)
		if resp.StatusCode != 200 {
			res.WriteHeader(403)
		}

		// Read the response from Dialogflow Gateway
		// Add session identifier to response
		// Generate new JSON
		resp_body, _ := ioutil.ReadAll(resp.Body)
		var result map[string]interface{}
		json.Unmarshal(resp_body, &result)
		result["session"] = "projects/" + Project + "/agent/sessions/fulfillment-tester"
		data, _ := json.Marshal(result)

		// Send the response to Fulfillment
		fulfillment_resp, _ := http.Post(Fulfillment, "application/json", bytes.NewReader(data))

		// Parse the fulfillment_messages from the Fulfillment response
		// Overwrite queryResult on response (addr is the address of the queryResult field, that can be pointed to later)
		// That's how Dialogflow actually responds with fulfillment option enabled (you don't have to be a Googler to tell that)
		fulfillment_messages, _ := ioutil.ReadAll(fulfillment_resp.Body)
		var addr = result["queryResult"]
		json.Unmarshal(fulfillment_messages, &addr)

		// Convert output back to JSON and send it to the client
		output, _ := json.Marshal(result)
		res.Write(output)
	} else if req.Method == "OPTIONS" {
		res.WriteHeader(200)
	} else {
		res.WriteHeader(404)
	}
}

func main() {
	// Parse flags, setup default flags and descriptions
	flag.StringVar(&Project, "project", "dialogflow-web-v2", "Dialogflow Gateway Project ID")
	flag.StringVar(&Fulfillment, "fulfillment", "https://us-central1-dialogflow-web-v2.cloudfunctions.net/dialogflowFirebaseFulfillment", "URL to fullfillment (remote or local)")
	flag.StringVar(&Port, "port", "8899", "Port to run the fulfillment-tester at")
	flag.Parse()

	// Log some useful information to console
	println("Dialogflow Fulfillment Tester is running 🚀")
	println("Listening on: http://localhost:" + Port)
	println("Web UI at: http://localhost:" + Port + "/ui" + "\n")
	println("Connection 🔌")
	println("Project ID: " + Project)
	println("Fulfillment URL: " + Fulfillment)
	println("\nHappy Testing!")

	// Setup HTTP Server and Proxy (Handler)
	// Setup the Web Server for the Web Client
	http.HandleFunc("/", proxy)
	http.Handle("/ui/", http.StripPrefix("/ui", http.FileServer(http.Dir("ui"))))
	panic(http.ListenAndServe(":"+Port, nil))
}