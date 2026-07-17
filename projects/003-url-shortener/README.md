## Usage

- Requires Go >= 1.22 + installed

- Build and run the binary or simply run

```bash
  go run shorten.go
```

## Example

You should see a message like

```bash
  2026/07/17 16:28:12 
  Listening on http://localhost:8080
```

- Open your rest client (Postman, EchoAPI, Thunderclient etc...) and send a **POST** request like this

```bash
  POST:
  http://localhost:8080/shorten?url=<url-to-shorten>

  // will return short url as response eg http://localhost:8080/PuyhfS
```

- Paste short url on your browser to trigger a **GET** request and you'll be redirected to the long url
