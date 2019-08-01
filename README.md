# search : Search returns integer(s) in string

## Overview [![GoDoc](https://godoc.org/github.com/SuruthiGK/search?status.svg)](https://godoc.org/github.com/SuruthiGK/search)

Search Api is built on Gin Framework.
Request to the app can be made through POSTMAN or cURL.

## Install

```
go get github.com/SuruthiGK/search
```

## Example

```
Method: POST
Header: Content-Type:application/json

Sample Input:
{
    "data": "H@2RM6A6a8n"
}

Sample Output:
[
    "2",
    "6",
    "6",
    "8"
]

cURL Request:

cURL -X POST \
  http://0.0.0.0:5000/api/search/integers \
  -H 'Cache-Control: no-cache' \
  -H 'Content-Type: application/json' \
  -H 'Postman-Token: a2fb1086-2e25-40a3-b742-1968f6caaeb6' \
  -d '{
	"data": "H@2RM6A6a8n"
}'

Response:
["2","6","6","8"]
```

