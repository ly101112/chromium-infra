// Copyright 2023 The Chromium Authors
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package datastorage

import (
	"context"
	"errors"
	"fmt"

	"cloud.google.com/go/datastore"
	"go.chromium.org/luci/common/logging"
)

var (
	ErrInsufficientArgs  = errors.New("insufficent arguments")
	ErrConnection        = errors.New("connection error")
	ErrInvalidKey        = errors.New("invalid key")
	ErrInvalidEntityType = errors.New("invalid entity type")
	ErrEntityNotFound    = errors.New("entity not found for given key")
	ErrInternal          = errors.New("internal error")
)

type DataStoreClient struct {
	datastoreClient *datastore.Client
}

// NewDataStoreClient function creates a new Datastore client and
// returns it if successful. On failure, returns an error.
// Note that this client is specifically designed to communicate
// with Google's Datastore service.
func NewDataStoreClient(ctx context.Context, cloudProject string) (*DataStoreClient, error) {
	datastoreClient, err := datastore.NewClient(ctx, cloudProject)
	if err != nil {
		logging.Errorf(ctx, "Datastore connection cannot be established: %s", err)
		return nil, ErrConnection
	}
	c := DataStoreClient{
		datastoreClient: datastoreClient,
	}
	return &c, nil
}

// Get takes in the Entity name, entity key, options (in case the key needs ancestor)
// and an empty struct reference. Returns an error if the fetch was unsuccessful.
// Otherwise copies the entity data to the result argument which should be an
// empty struct reference.
//
// Important Notes:
// 1. This function can accept 4 or 6 arguments depending on whether the entity in question
// requires an ancestor to be found or not.
// 2. The 4th and 6th (if present) arguments can be of type string or number because key
// can either be a NameKey or an IdKey.
//
// Example usage:
// str := A{}
// Example 1: dsClient.Get(ctx, &str, "EntityA", "EntityAKey", "AncestorEntity", "AncestorEntityKey")
// Example 2: dsClient.Get(ctx, &str, "EntityA", 123, "AncestorEntity", 345)
// Example 3: dsClient.Get(ctx, &str, "EntityA", "k")
func (c DataStoreClient) Get(ctx context.Context, result interface{}, entityName string, key interface{}, options ...interface{}) error {
	if !(len(options) == 0 || len(options) == 2) {
		return fmt.Errorf("%s: Expected 4 or 6 arguments but got %d", ErrInsufficientArgs, len(options)+4)
	}

	var ancestorKeyLiteral *datastore.Key
	if len(options) == 2 {
		var ancestorEntityName string
		if e, isStringType := options[0].(string); isStringType {
			ancestorEntityName = e
		} else {
			return fmt.Errorf("%s: Ancestor entity name should be of type string", ErrInvalidEntityType.Error())
		}

		switch options[1].(type) {
		case string:
			ancestorKeyLiteral = datastore.NameKey(ancestorEntityName, options[1].(string), nil)
		case int64:
			ancestorKeyLiteral = datastore.IDKey(ancestorEntityName, options[1].(int64), nil)
		default:
			return fmt.Errorf("%s: Ancestor key should be of type string or int64", ErrInvalidKey.Error())
		}
	}

	var entityKeyLiteral *datastore.Key
	switch k := key.(type) {
	case string:
		entityKeyLiteral = datastore.NameKey(entityName, k, ancestorKeyLiteral)
	case int64:
		entityKeyLiteral = datastore.IDKey(entityName, k, ancestorKeyLiteral)
	default:
		return fmt.Errorf("%s: Entity key should be a string or int64", ErrInvalidKey.Error())
	}

	err := c.datastoreClient.Get(ctx, entityKeyLiteral, result)
	if err != nil {
		if err == datastore.ErrInvalidEntityType {
			return fmt.Errorf("%s: The result argument is likely an invalid type", ErrInvalidEntityType)
		}
		if err == datastore.ErrInvalidKey {
			return ErrInvalidKey
		}
		if err == datastore.ErrNoSuchEntity {
			return ErrEntityNotFound
		}
		logging.Errorf(ctx, "Error fetching %s: %s", entityName, err)
		return ErrInternal
	}
	return nil
}
