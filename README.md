# benfords-law
A Go-based web application that ingests a CSV file and outputs a JSON payload describing whether Benford's law is observed on a specific column. 

A demo is hosted at [https://benford.lph.pw](https://benford.lph.pw).

## Prerequisites
* Go 1.13

## Installation
_git clone https://github.com/rewardian/benfords-law.git_ 
_docker build -t benfords-law ._

main.go is configured to start the HTTP server on :9021. For my demo, I've configured nginx to proxy into the container's exposed port.
