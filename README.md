# Dialogflow Fulfillment Tester

Test your Fulfillments like never before!

Start it:

![](https://i.imgur.com/Y5Ufr32.png)

And test with style using the User Interface:

![](https://i.imgur.com/YSulY94.png)

Or send the requests manually:

![](https://i.imgur.com/ywbIjrE.jpg)

## Features

- Works on Windows, Linux and Mac (64 bit)
- Small binary with no external dependencies, written in Go (<10 MB)
- Test fulfillments in any programming language and Actions on Google fulfillments
- No More waiting. Never. Just run your fulfillment (locally or remotely) and start testing
- Run automated tests with tools like Jest, Ava, Mocha (or any other of your choice)
- Test with convinience, with a built-in User Interface or on CI/CD
- Works exactly like Dialogflow, 100% accurate testing results guaranteed

Excited? Let's get started!

## Installation

Connect your Dialogflow Agent to [Dialogflow Gateway](https://dialogflow.cloud.ushakov.co), read the guide [here](https://github.com/mishushakov/dialogflow-gateway-docs/blob/master/guide.md)

Install the latest executable for your operating system from the [Releases Page](https://github.com/mishushakov/dialogflow-fulfillment-tester/releases)

Run

```sh
dialogflow-fulfillment-tester --project <YOUR GOOGLE CLOUD PROJECT ID> --fulfillment <URL>
```

Get help:

```sh
dialogflow-fulfillment-tester --help
```

Tip: if you are on node and firebase functions, run your function locally using the firebase functions emulator

## Installing the UI

If you want the UI, clone this repo and put the `ui` folder near the executable.

Notice: when running Dialogflow Fulfillment Tester on a diffrent host/port, make sure to change it in the UI as well (`index.html`):

```js
let url = "http://localhost:8899" // <- Change the url, when running on a different host/port
```

Notice: the UI doesn't display Actions on Google components at the moment. 

Tip: When inspecting using the UI, open the console to see the request/response body

## Making Requests

The request/response format of the Dialogflow Fulfillment Tester is equal to the Dialogflow Gateway request/response format, which is equal to the Dialogflow request/response format. Read the docs [here](https://github.com/mishushakov/dialogflow-gateway-docs/blob/master/api.md)

## Building from source

- Get [Go](https://golang.org/dl/)
- Build using `build.sh` script
- Ready!