package router

import "net/http"

type MethodHandlers map[string]http.Handler
