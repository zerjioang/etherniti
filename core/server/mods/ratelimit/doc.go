package ratelimit

/*

HTTP Headers and Response Codes

Use the HTTP headers in order to understand where the application is at for a given rate limit, on the method that was just utilized.

x-rate-limit-limit: the rate limit ceiling for that given endpoint
x-rate-limit-remaining: the number of requests left for the 15 minute window
x-rate-limit-reset: the remaining window before the rate limit resets, in UTC epoch seconds

When an application exceeds the rate limit for a given standard API endpoint, the API will return a HTTP 429 “Too Many Requests” response code, and the following error will be returned in the response body:

{
  "code": 429,
  "msg": "Too Many Requests",
  "details": "rate limit reached"
}

*/
