# CryptoCLI

CryptoCLI is a command-line interface (CLI) application for obtaining cryptocurrency price information, visualizing data, and performing various tasks related to cryptocurrencies.

## Table of Contents

- [Features](#features)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
- [Usage](#usage)


## Features

- Fetch real-time cryptocurrency prices.
- Visualize cryptocurrency price data with charts.
- Upload and analyze cryptocurrency data files.
- Perform various commands like 'about', 'help', and more.

## Getting Started

These instructions will help you set up and run CryptoCLI on your local machine.

### Prerequisites

Before you begin, make sure you have the following installed:

- Go (Golang): [Installation Guide](https://golang.org/doc/install)
- Node.js and npm (for the frontend): [Installation Guide](https://nodejs.org/en/download/)

### Installation

1. Clone the CryptoCLI repository:

   ```shell
   git clone https://github.com/sheriumair/CryptoCLI.git


### Directions
1. Navigate to the CryptoCLI main repo:
  
    ```shell
    docker-compose up

## Usuage
CryptoCLI offers a set of commands to interact with cryptocurrency data. Here are some common commands and how to use them:

### About
To get information about CryptoCLI, you can use the `about` command:

    ```shell
    about
     

### Help
If you need assistance with using CryptoCLI or want to see a list of available commands, you can use the help command:

     ```shell
     help


### Fetching Crypto Prices
You can fetch real-time cryptocurrency prices using the `fetch-price` command. Replace `PAIR` with the cryptocurrency pair you're interested in (e.g., `BTCUSDT`):
   
    ```shell
    fetch-price PAIR


### Uploading File
To upload and analyze cryptocurrency data files, use the `upload` command. This command allows you to submit a file containing cryptocurrency data for further processing.

     ```shell
     $ upload

### Drawing File Columns
To create visual representations of cryptocurrency data, you can use the `draw` command. This command allows you to specify a file and columns to be plotted.

     ```shell
     $ draw [file] [columns]

### Delete File

     ```shell
     $ delete [file] 





    


