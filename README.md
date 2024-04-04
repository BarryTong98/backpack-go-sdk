Introduction
------------

The Backpack Trade Bot is a specialized software utility designed to automate trading tasks on the Backpack exchange
platform. Leveraging the official Backpack
API ([https://docs.backpack.exchange/#section/Introduction](https://docs.backpack.exchange/#section/Introduction)), this
bot enables users to effortlessly manage and execute trades for specified token pairs, streamlining the trading process
and enhancing efficiency.

Features
--------

* **Automated Trading**: Set up your trading bot once and let it automate the trading process for you.
* **Custom Token Pairs**: Easily specify which token pairs you wish to trade, offering you full control over your
  trading strategy.
* **Direct Integration with Backpack API**: Directly integrated with the official Backpack exchange API, ensuring
  reliable and up-to-date interaction with your Backpack account.

Prerequisites
-------------

Before you begin, ensure you have the following:

* A Backpack exchange account
* API keys generated from your Backpack account (refer to Backpack's API documentation for guidance)
* Go installed on your system (visit [https://golang.org/dl/](https://golang.org/dl/) for installation instructions)

Setup
-----

1. **Clone the Repository**

``` bash
git clone https://github.com/yourusername/backpack-trade-bot.git 
cd backpack-trade-bot
```

2. **Configure API Keys**

   Open `config.toml` in a text editor and insert your Backpack API key and secret:

``` toml
[api] 
apisecret = "your_api_secret" 
apikey = "your_api_key"
```

3. **Install Dependencies**

   Run the following command in the project directory:

``` bash
go mod tidy
```

Usage
-----

To start the bot, run the following command from the root of the project directory:

``` bash
go run main.go
```

### Customizing Token Pairs

Edit the `main.go` file or create a configuration section in `config.toml` to specify the token pairs you wish to trade.
Refer to the Backpack API documentation for the exact naming conventions of token pairs.

Contributing
------------

Contributions to the Backpack Trade Bot are welcome! Please refer to the contributing guidelines for more information on
how you can contribute to this project.

License
-------

This project is licensed under the MIT License - see the LICENSE file for details.

