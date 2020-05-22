# gin-jwt-middleware
[Reference - ENG](https://sosedoff.com/2014/12/21/gin-middleware.html)  
[Reference - KOR](https://bourbonkk.tistory.com/63)

## What is JWT?
JSON Web Token (JWT) more information: 
[http://self-issued.info/docs/draft-ietf-oauth-json-web-token.html](http://self-issued.info/docs/draft-ietf-oauth-json-web-token.html)

## Quick Start
#### Install package
```bash
$ go get github.com/YankeeTube/gin-jwt-middleware
```

#### In your gin application main.go, import the package
```go
import (
    "github.com/YankeeTube/gin-jwt-middleware"
)
```

#### Use the middleware
```go
// ApiMiddleware will add the db connection to the context
func ApiMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		privateBytes := strings.Replace(os.Getenv("PRIVATE_KEY"), "\\n", "\n", -1)
		block, _ := pem.Decode([]byte(privateBytes))
		privateKey, _ := x509.ParsePKCS1PrivateKey(block.Bytes)
		c.Set("PRIVKEY", privateKey)
		c.Next()
	}
}

api := gin.Default()
api.Use(ApiMiddleware)  // Private Key Settings of Context
api.Use(authenticate.TokenAuthMiddleware)  // Jwt Token Check
```