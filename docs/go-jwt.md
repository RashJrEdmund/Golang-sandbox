## JWT

- For password hashing, we're using [argon2id](https://github.com/alexedwards/argon2id)
  Which provides a convenience wrapper around Go's [argon2](https://pkg.go.dev/golang.org/x/crypto/argon2?tab=doc) implementation,
  making it simpler to securely hash and verify passwords using Argon2.

  It enforces use of the Argon2id algorithm variant and cryptographically-secure random salts.

  _Argon2id is a great choice. [SHA-256](https://www.boot.dev/blog/computer-science/how-sha-2-works-step-by-step-sha-256/) and [MD5](https://en.wikipedia.org/wiki/MD5) are not.

- **For session signing**, we're using [golang-JWT](https://github.com/golang-jwt/jwt)

  Create a new token.

  - Use [jwt.NewWithClaims](https://pkg.go.dev/github.com/golang-jwt/jwt/v5#NewWithClaims)
  - Use [jwt.SigningMethodHS256](https://pkg.go.dev/github.com/golang-jwt/jwt/v5#SigningMethodHS256) as the signing method.
  - Use [jwt.RegisteredClaims](https://pkg.go.dev/github.com/golang-jwt/jwt/v5#RegisteredClaims) as the claims.
    - Set the Issuer to "chirpy-access"
    - Set IssuedAt to the current time in UTC
    - Set ExpiresAt to the current time plus the expiration time (expiresIn)
    - Set the Subject to a stringified version of the user's id
  - Use [token.SignedString](https://pkg.go.dev/github.com/golang-jwt/jwt/v5#Token.SignedString) to sign the token with the secret key. Refer to the [docs](https://golang-jwt.github.io/jwt/usage/signing_methods/#signing-methods-and-key-types) for an overview of the
  
  All of this should yeild

  ```go
    func SignJWT(userID uuid.UUID, tokenSecret string, expiresIn time.Duration) (string, error) {
      now := time.Now().UTC()
      claims := jwt.RegisteredClaims{
        Issuer:    "chirpy-access",
        IssuedAt:  jwt.NewNumericDate(now),
        ExpiresAt: jwt.NewNumericDate(now.Add(expiresIn)),
        Subject:   userID.String(),
      }

      token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
      return token.SignedString([]byte(tokenSecret))
    }
  ```

- **For Session validation**
  Use the [jwt.ParseWithClaims](https://pkg.go.dev/github.com/golang-jwt/jwt/v5#ParseWithClaims) function to validate the signature of the JWT and extract the claims into a [*jwt.Token](https://pkg.go.dev/github.com/golang-jwt/jwt/v5#Token) struct.
  The `keyFunc` callback must return the same key type (`[]byte`) used when the token was signed. An error will be returned if the token is invalid or has expired.

  If all is well with the token, use the [token.Claims](https://pkg.go.dev/github.com/golang-jwt/jwt/v5#Claims) interface to get
  access to the user's id from the claims (which should be stored in the Subject field). Return the id as a uuid.UUID.

### NOTE

- JWTs are not encrypted. Anyone who has the token can read the data (like the expiry and the user id) in the token.
  This is why you should never store sensitive information in a JWT. It's just a way to authenticate a user.

- Use [JWT.io](https://www.jwt.io/) to inspect

- One of the main benefits of JWTs is that they're stateless.
  The server doesn't need to keep track of which users are logged in via JWT.
  The server just needs to issue a JWT to a user and the user can use that JWT to authenticate themselves.
  Statelessness is fast and scalable because your server doesn't need to consult a database to see if a user is currently logged in

  - However, that same benefit poses a potential problem. JWTs can't be revoked.
    If a user's JWT is stolen, there's no easy way to stop the JWT from being used. JWTs are just a signed string of text.

### Access Tokens & Refresh Tokens

Access tokens are used to authenticate a user to a server, and they provide access to protected resources.
Access tokens are:

- Stateless
- Short-lived (15m-24h)
- Irrevocable

They must be short-lived because they can't be revoked. The shorter the lifespan,
the more secure they are. Trouble is, this can create a poor user experience.
We don't want users to have to log in every 15 minutes.

The solution is Refresh Tokens

### Refresh Tokens

Refresh tokens don't provide access to resources directly, but they can be used to get new access tokens.
Refresh tokens are much longer lived, and importantly, they can be revoked. They are:

- Stateful
- Long-lived (24h-60d)
- Revocable

With this, we get the best of both worlds
