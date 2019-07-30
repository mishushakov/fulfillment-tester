# Fulfillment Tester

Test your Dialogflow/Actions on Google Fulfillments like never before!

- Works on Windows, Linux and Mac (64 bit)
- Small, embeddable binary with no external dependencies, written in Go (<10 MB)
- Supports any dialogflow fulfillment in any programming language, platform-specific responses and Actions on Google
- No More waiting. Never. Run your fulfillment (locally or remotely) and start testing
- Run automated tests with tools like Jest, Ava, Mocha (or any other of your choice)
- Test with a built-in User Interface or on CI/CD
- Acts exactly like Dialogflow, 100% accurate testing results guaranteed

Just start it:

![](https://i.imgur.com/yE2aiid.png)

And test with style using [Dialogflow for Web](https://github.com/mishushakov/fulfillment-tester):

![](https://i.imgur.com/sZR5c63.png)
![](https://i.imgur.com/VmPCd0u.png)
![](https://i.imgur.com/BWbropG.png)

Or send requests manually (programmatically):

![](https://i.imgur.com/ywbIjrE.jpg)

## How it works

![](https://svgur.com/i/EKw.svg)

Fulfillment tester acts like a reverse-proxy between Dialogflow and your fulfillment. It fetches Dialogflow response for a given query, forwards it to the specified fulfillment and then responds with a result.

## Installation

1. Install Dialogflow Gateway first. Fulfillment tester uses Dialogflow Gateway as its backend.
   
   Dialogflow Gateway enables third-party integrations to securely access the Dialogflow V2 API

   - [Documentation](https://github.com/mishushakov/dialogflow-gateway-docs)
   - [Implementations](https://github.com/mishushakov/dialogflow-gateway-docs#implementations)

2. Install the latest executable for your operating system from the [Releases Page](https://github.com/mishushakov/fulfillment-tester/releases)

Run

```sh
fulfillment-tester --gateway <GATEWAY URL> --fulfillment <FULFILLMENT URL>
```

Get help:

```sh
fulfillment-tester --help
```

Tip: if you are on node and firebase functions, run your function locally using the firebase functions emulator

## Accessing the UI

1. Follow the installation instructions in the [Dialogflow for Web repo](https://github.com/mishushakov/dialogflow-web-v2)
2. Change the Gateway URL in the `config.js` to match your fulfillment tester URL

    Example:

    ```js
    export default {
        app: {
            gateway: "http://localhost:8899"
            [...]
        }
    }

    [...]
    ```

## Making Requests

The request/response format of Fulfillment Tester is equal to the [Dialogflow Gateway API](https://github.com/mishushakov/dialogflow-gateway-docs#api) request/response format

## Building from source

- Get [Go](https://golang.org/dl/)
- Build using `build.sh` script
- Ready!