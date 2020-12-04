# JWT Decoder

Small utility to decode JWT's quickly, and print the vital information (including expiration time).

## Install
```
go get -u github.com/smugcloud/jwt-cli
```

## Usage
```
 $ jwt-cli eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c
 {
    "sub": "1234567890",
    "name": "John Doe",
    "iat": 1516239022
 }
 Expires: 1969-12-31 16:00:00 -0800 PST
 Issued: 2018-01-17 17:30:22 -0800 PST
```