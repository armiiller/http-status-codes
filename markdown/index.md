# HTTP Status Codes 🤠
A service for generating HTTP codes.

It's useful for testing monitoring services.

Just add the status code you want to the end of the URL, like this: [https://statuscode.app/200](https://statuscode.app/200)

If you want to delay a response add the query parameter **sleep** and specify the duration in milliseconds (maximum: 5000). Example: [https://statuscode.app/200?sleep=5000](https://statuscode.app/200?sleep=5000)

---
| Code | Category | Description |
| ---- | -------- | ----------- |
| 1xx | Informational | Request received, continuing process |
| 2xx | Success | The action was successfully received, understood, and accepted |
| 3xx | Redirection | Further action must be taken in order to complete the request |
| 4xx | Client Error | The request contains bad syntax or cannot be fulfilled |
| 5xx | Server Error | The server failed to fulfill an apparently valid request |

Please see the [IANA website](https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml) for all available codes and details.

---

## Supported Status Codes
| Code | Description |
| ---- | ----------- |
| [100](/100) | Continue |
| [101](/101) | Switching Protocols |
| [102](/102) | Processing |
| [103](/103) | Early Hints |
| [200](/200) | OK |
| [201](/201) | Created |
| [202](/202) | Accepted |
| [203](/203) | Non-Authoritative Information |
| [204](/204) | No Content|
| [205](/205) | Reset Content |
| [206](/206) | Partial Content |
| [207](/207) | Multi-Status |
| [208](/208) | Already Reported |
| [226](/226) | IM Used |
| [300](/300) | Multiple Choices |
| [301](/301) | Moved Permanently |
| [302](/302) | Found |
| [303](/303) | See Other |
| [304](/304) | Not Modified |
| [305](/305) | Use Proxy |
| [306](/306) | Switch Proxy |
| [307](/307) | Temporary Redirect |
| [308](/308) | Permanent Redirect |
| [400](/400) | Bad Request |
| [401](/401) | Unauthorized |
| [402](/402) | Payment Required |
| [403](/403) | Forbidden |
| [404](/404) | Not Found |
| [405](/405) | Method Not Allowed |
| [406](/406) | Not Acceptable |
| [407](/407) | Proxy Authentication Required |
| [408](/408) | Request Timeout |
| [409](/409) | Conflict |
| [410](/410) | Gone |
| [411](/411) | Length Required |
| [412](/412) | Precondition Failed |
| [413](/413) | Request Entity Too Large |
| [414](/414) | Request-URL Too Long |
| [415](/415) | Unsupported Media Type |
| [416](/416) | Requested Range Not Satisfiable |
| [417](/417) | Expectation Failed |
| [418](/418) | I'm a teapot |
| [421](/421) | Misdirected Request |
| [422](/422) | Unprocessable Entity |
| [423](/423) | Locked |
| [424](/424) | Failed Dependency |
| [425](/425) | Too Early |
| [426](/426) | Upgrade Required |
| [428](/428) | Precondition Required |
| [429](/429) | Too Many Requests |
| [431](/431) | Request Header Fields Too Large |
| [451](/451) | Unavailable For Legal Reasons |
| [500](/500) | Internal Server Error |
| [501](/501) | Not Implemented |
| [502](/502) | Bad Gateway |
| [503](/503) | Service Unavailable |
| [504](/504) | Gateway Timeout |
| [505](/505) | HTTP Version Not Supported |
| [506](/506) | Variant Also Negotiates |
| [507](/507) | Insufficient Storage |
| [508](/508) | Loop Detected |
| [510](/510) | Not Extended |
| [511](/511) | Network Authentication Required |

---
Created by [Austin Miller](https://austinmiller.dev), Hosted on [Fly.io](https://fly.io), Open Sourced on [Github](https://github.com/armiiller/http-status-codes).

We don't capture or store any data about the requests you make.