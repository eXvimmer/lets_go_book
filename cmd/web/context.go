package main

type contextKey string

// to remove the risk of naming collisions
const isAuthenticatedContextKey = contextKey("isAuthenticated")
