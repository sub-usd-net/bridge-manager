package service

import "errors"

var (
	errRead      = errors.New("error reading deposit info")
	errTransfer  = errors.New("error execute bridge transfer")
	errInvariant = errors.New("invariant error")
)
