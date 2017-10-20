// Copyright © 2017 The Things Network Foundation, distributed under the MIT license (see LICENSE file)

// Package encoding defines interfaces shared by other packages that
// convert data to and from byte-level and textual representations.
// Packages that check for these interfaces include encoding/gob,
// encoding/json, and encoding/xml. As a result, implementing an
// interface once can make a type useful in multiple encodings.
// Standard types that implement these interfaces include time.Time and net.IP.
// The interfaces come in pairs that produce and consume encoded data.
// This package is an extension of standard library 'encoding'
// package with LoRaWAN encoding interfaces added.
package encoding

// LoRaWANMarshaler is the interface implemented by an object that can
// marshal itself into LoRaWAN form.
//
// MarshalLoRaWAN encodes the receiver into LoRaWAN form and returns the result.
type LoRaWANMarshaler interface {
	MarshalLoRaWAN() (data []byte, err error)
}

// LoRaWANAppender is the interface implemented by an object that can append LoRaWAN representation of itself to a byte slice.
//
// AppendLoRaWAN encodes the receiver into LoRaWAN form, appends it to supplied byte slice and returns the result.
type LoRaWANAppender interface {
	AppendLoRaWAN(dst []byte) (data []byte, err error)
}

// LoRaWANUnmarshaler is the interface implemented by an object that can
// unmarshal a LoRaWAN representation of itself.
//
// UnmarshalLoRaWAN must be able to decode the form generated by MarshalLoRaWAN.
// UnmarshalLoRaWAN must copy the data if it wishes to retain the data
// after returning.
type LoRaWANUnmarshaler interface {
	UnmarshalLoRaWAN(data []byte) error
}
