# benfords-law
A Go-based web application that ingests a CSV file and outputs a JSON payload describing whether Benford's law is observed on a specific column. 

A demo is hosted on [Google App Engine](https://cloud.google.com/appengine/) at [https://benfordlaw.appspot.com](https://benfordlaw.appspot.com).

## Prerequisites
* Go 1.13
* Mux ([https://github.com/gorilla/mux](https://github.com/gorilla/mux))
  * _go get -u github.com/gorilla/mux_
  
## Installation
_git clone https://github.com/rewardian/benfords-law.git_ 

main.go is configured to start the HTTP server on :8080, which is totally fine with Google App Engine.
