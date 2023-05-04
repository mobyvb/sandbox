# Simple OAuth App

    A simple app demonstrating OAuth2 authentication with GitHub.

    ## Running the Program

    ### With Go

    1. Set your environment variables with your GitHub OAuth app credentials:

    
    export OAUTH2_CLIENT_ID=your_client_id
    export OAUTH2_CLIENT_SECRET=your_client_secret
    

    2. Run the server:

    
    go run main.go
    

    ### With Docker

    1. Build the Docker image:

    
    docker build -t simple-oauth-app .
    

    2. Run the Docker container:

    
    docker run -d -p 8080:8080 --env OAUTH2_CLIENT_ID=your_client_id --env OAUTH2_CLIENT_SECRET=your_client_secret simple-oauth-app