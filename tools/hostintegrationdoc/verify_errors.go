package main

import "errors"

var (
	errInvalidDistribution    = errors.New("distribution channel validity failed")
	errInvalidProviderRouting = errors.New("provider routing status validity failed")
	errInvalidStoreManaged    = errors.New("store-managed classification failed")
)
