# HTTP Status Codes
A service for generating HTTP codes.

It's useful for testing monitoring services.

Just add the status code you want to the URL, like this [https:///200](/200)

Please see the [IANA website](https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml) for details.

If you want to delay a response add the query parameter *sleep* and specify the duration in milliseconds (maximum: 5000). Example: https://


| Code | Category | Description |
| ---- | -------- | ----------- |
| 1xx | Informational | Request received, continuing process |
| 2xx | Success | The action was successfully received, understood, and accepted |
| 3xx | Redirection | Further action must be taken in order to complete the request |
| 4xx | Client Error | The request contains bad syntax or cannot be fulfilled |
| 5xx | Server Error | The server failed to fulfill an apparently valid request |

---
Created by [Austin Miller](https://austinmiller.dev), Hosted on [Fly.io](https://fly.io), Open Sourced on [Github](https://github.com/armiiller/http-status-codes).
We don't capture or store any data about the requests you make.