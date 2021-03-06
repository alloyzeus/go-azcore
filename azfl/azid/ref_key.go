package azid

import (
	"github.com/alloyzeus/go-azfl/azfl/azob"
)

// RefKey defines the contract for all its concrete implementations.
//
// A RefKey is analogous to path in URL.
//
//     https://example.com/stores/12345/items/456789
//
// The part `/stores/12345/items/456789` is item's ref-key, while 456789 is the
// ID.
//
// An ID is always scoped, while a ref-key could be used to distinctively
// identify an instance of an entity system-wide.
//
//     https://example.com/stores/StoreABC/items/456789
//     https://example.com/stores/StoreXYZ/items/456789
//
// In above example, two items from two different stores have the same ID but
// their ref-keys will not be the same. This system allows each store to scale
// independently without significantly affecting each others.
//
//TODO: RefKey is ValueObject.
type RefKey interface {
	azob.Equatable

	AZRefKey()

	BinFieldMarshalable
	BinMarshalable
	TextMarshalable

	// // Returns an array of the hosts' ref-keys.
	// Hosts() []RefKey

	// AZIDNum returns only the IDNum-part of this ref-key.
	AZIDNum() IDNum
}

// RefKeyFromString is a function which creates an instance of RefKey
// from an input string.
type RefKeyFromString func(refKeyString string) (RefKey, Error)
