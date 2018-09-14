# Router


[![MIT License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/zpatrick/router/blob/master/LICENSE)
[![Go Doc](https://godoc.org/github.com/zpatrick/router?status.svg)](https://godoc.org/github.com/zpatrick/router)

## Usage
Please see the [example](/example) directory for a working example. 
```go
func main() {
       rm := router.RouteMap{
                "/products": router.MethodHandlers{
                        http.MethodGet:  http.HandlerFunc(ListProducts),
                        http.MethodPost: http.HandlerFunc(AddProduct),
                },
                "/products/:productID": router.MethodHandlers{
                        http.MethodGet:    http.HandlerFunc(GetProduct),
                        http.MethodDelete: http.HandlerFunc(DeleteProduct),
                },
        }

        r := router.NewRouter(rm.VariableMatch())
        http.ListenAndServe(":8000", r)
}
```

### Route Matching
The [RouteMap](https://godoc.org/github.com/zpatrick/router#RouteMap) object is essentially just a way to organize `http.Handlers` in a simple, readable way.
This object has helper functions to convert each `http.Handler` into a [HandlerMatcher](https://godoc.org/github.com/zpatrick/router#HandlerMatcher):

```go
type HandlerMatcher func(r *http.Request) (handler http.Handler, matchFound bool)
```

This package currently has four built-in `HandlerMatchers`. 
Each requires that the `request.Method` exactly match the specified method in the `RouteMap`, 
and each uses different methods to match the `request.URL.Path`: 
* [Glob](https://godoc.org/github.com/zpatrick/router#NewGlobHandlerMatcher) - returns a match if the `request.URL.Path` [glob matches](https://godoc.org/github.com/ryanuber/go-glob#Glob) the pattern used in the `RouteMap`. 
* [Regex](https://godoc.org/github.com/zpatrick/router#NewRegexHandlerMatcher) - returns a match if the `request.URL.Path` [regex matches](https://golang.org/pkg/regexp/#Regexp.MatchString) the pattern used in the `RouteMap`.
* [String](https://godoc.org/github.com/zpatrick/router#NewStringHandlerMatcher) - returns a match if the `request.URL.Path` exactly matches the pattern used in the `RouteMap`.
* [Variable](https://godoc.org/github.com/zpatrick/router#NewVariableHandlerMatcher) - returns a match if the `request.URL.Path` [variable matches](https://godoc.org/github.com/zpatrick/router#NewVariableHandlerMatcher) the pattern used in the `RouteMap`.

### Path Variables
Path variables can be fetched using [Segments](https://godoc.org/github.com/zpatrick/router#Segments). 
Segments are just sections in a url's path delimited by the `/` character.  
For example, the segments for `/product/p123` are `[]string{"product", "p123"}`.
```go
func GetProduct(w http.ResponseWriter, r *http.Request) {
  productID := router.Segment(r.URL.Path, 1)
  ...
}
```

There are helper functions for [integer segments](https://godoc.org/github.com/zpatrick/router#IntSegment):
```go
func GetProduct(w http.ResponseWriter, r *http.Request) {
  productID, err := router.IntSegment(r.URL.Path, 1)
  ...
}

```

### Middleware
[Middleware](https://godoc.org/github.com/zpatrick/router#Middleware) adds functionality to a `http.Handler`:
```go
type Middleware func(http.Handler) http.Handler
```

This package currently has the following middleware:
* [Logging](https://godoc.org/github.com/zpatrick/router#LoggingMiddleware)
* [BasicAuth](https://godoc.org/github.com/zpatrick/router#BasicAuthMiddleware)

Middleware can be applied to a [RouteMap](https://godoc.org/github.com/zpatrick/router#RouteMap.ApplyMiddleware):
```go
rm := router.RouteMap{}
rm.ApplyMiddleware(router.LoggingMiddleware(), router.BasicAuthMiddleware("user", "pass"))
```
