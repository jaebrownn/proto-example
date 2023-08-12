# Proto Example
This repository contains some basic sample code to illustrate the use of protobufs in setting up gRPC servers in Go, 
running on Google Cloud Run. It also contains some useful workflows 

# Setting up Secrets for Buf related Github actions
1. Create an account at [buf.build](https://buf.build)
2. Navigate to settings by clicking on your profile in the top right corner.
3. Create a new token name "Github Actions" and copy the value.
4. In your repo settings, navigate to "Secrets and variables" > "Actions" and select "New repository secret"
5. Create a secret with name "BUF_API_TOKEN" and paste the value of your token.
6. Add another secret with name "BUF_USER" and provide your buf username.