# url-shortener-api
URL Shortening API in GoLang

## Features

- Generate short URLs using a random URL shortening algorithm
- Redirect to orignal URL when the short URL is accessed
- Easy to use REST API

## Getting Started
 Follow these instructions to download and run the project on your local machine

 ### Prerequisites
  
- Go (version 1.18 or higher)
- Git

### Installation

1. Clone the repository:
    
    Navigate to your project directory

    Run the following command in your terminal:
    ```bash
    git clone https://github.com/zy-131/url-shortener-api.git
    cd url-shortener-api
    ```

2. Initialize the Go module:

    ```bash
    go mod init <module-name>
    ```

3. Install dependencies:

    ```bash
    go mod tidy
    ```

### Running the API (Locally)

1. Navigate to your local project's root directory

2. Run the following command to start your local server:
    ```bash
    go run ./src/
    ```

3. With the server running, open a new terminal or Postman and send a `POST` request with the URL that you want to shorten

    Note: A `GET` request without an initial succesful `POST` request will result in a `404 StatusNotFound` error

    Example (Powershell cURL):
    ```bash
    curl http://localhost:8080/shortenURL -METHOD POST -BODY '{"long_url": "<original-url>"}'
    ```

4. Open the shortened URL in a browser

    Ex: ```localhost:8080/abc123```

## API Endpoints

### Create Short URL Endpoint

- #### Endpoint: ```POST /shortenURL```

- #### Request Headers: 
  - ```Content-Type: application/json```

- #### Request Body:
  ```json
  {
    "long_url": "https://www.example.com"
  }
  ```
    The ```long_url``` field is required and will return a ```400 StatusBadRequest``` error if missing

- #### Response
  - Status Code: 
    - ```201 Created``` if successful

  - Response Body:
    - The shortened URL string will be returned as plain text
    ```json
    "abc123"
    ```  

 - #### Error Response:
   - Status Code: ```404 StatusBadRequest``` 
   - Response Body:
    ```json
    {
        "message": "No URL passed in"
    }
    ```

### Redirect to Original URL

- #### Endpoint: ```GET /:shortenURL```

- #### Parameters:
  - ```shortURL```: The shortened URL identifier (ex: ```abc123``` from ```http://localhost:8080/abc123```)

- #### Response
  - Status Code: 
    - ```302 Found``` if shortened URL exists

    - ```Location``` header will contain original URL and cliented will be redirected

 - #### Error Response:
   - Status Code: ```404 StatusNotFound``` if the shortened URL does not exist 
   - Response Body:
    ```json
    {
        "message": "Short URL not found"
    }
    ```

## Database

In progress ...