## Quiz Application

This application allows users to test their knowledge by taking quizzes on various topics. It's built with the Go programming language and consists of a server and a client component.

## Overview

The server provides quiz questions and manages user interaction. The client allows users to see the questions, submit answers, and receive feedback.

## Requirements

* Go version >= 1.20
* (Optional) Additional dependencies (list any specific dependencies)

## Installation

1. Clone this repository:

`git clone github.com/oliveryanh/quiz`

2. Navigate to the project directory:

`cd quiz`

3. Build the server and client binaries:

`make all`

## Usage

### Starting the Server

`./bin/server`

This will build and run the server on the default port.

### Running the Client

`./bin/client`

### Cleaning Up

To clean up the build artifacts, run:

`make clean`
