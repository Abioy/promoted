# README

## Introduction
This package provides some functions to check if promoted methods from an embedded struct are overridden.
It might be useful to check if a middleware/module, like ones in router/services, with a embedded struct that implement a interface has the default behaviors or overwrites them, and then do only the different parts.

## Functions
### Promoted
```go
func Promoted(t reflect.Type, method string) (bool, bool)
```
Promoted returns if a method is promoted or not.

### PromotedFrom
```go
func PromotedFrom(curr, base reflect.Type, method string) bool
```
PromotedFrom returns if a method is promoted from base or not, assume that
base did have the specified method.

### GetOverriddenMethods
```go
func GetOverriddenMethods(impl, base, api any) []string
```
GetOverriddenMethods returns all methods declare in api that are overridden
in impl and original promoted from base. Assume that impl and base both
implement interface api. It panic if api is not an interface.

## NOTICE
The package depends on the implement of promoted method in golang.
