<p align="center">
  <img src="/brand/rollpay_logo_thumb.png" />
</p>

# Rollpay
An experiment with the Plaid EU payment products

# Local Development
This project is made up of multiple sub projects which power the backend web server as well as various user interfaces.

## Backend Web Service
The backend web service is written in Go and uses PostgreSQL as it's primary datastore. It exposes a REST API for clients to interface with.

Local development is coordinated with `make`. Run `make help` to see available commands for working with the project.

A quickstart will look like the following:
```
git clone https://github.com/syllabix/rollpay

cd backend

make dev.start // go get a coffee while docker images are pulled

make envfile
// add your plaid client id and secret to the generated .env file

make run
// navigate to http://localhost:3000 (unless you have changed the port)
```
