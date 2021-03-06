// Copyright 2014 Simon Zimmermann. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package stathat registers the "stathat" storage type, storing stats
// in StatHat.

package stathat

import (
	"fmt"

	"github.com/simonz05/carry"
	"github.com/simonz05/carry/config"
)

type stathatStorage struct {
	key string
	w   carry.StatsWriter
}

func (s *stathatStorage) String() string {
	return fmt.Sprintf("\"stathat\" storage for %q", s.key)
}

func newFromConfig(conf *config.Config) (carry.Storage, error) {
	var w carry.StatsWriter

	k := conf.Stathat.Key
	w = NewStathatWriter(k)

	if conf.Periodic {
		w = carry.NewPeriodicWriter(w)
	}

	return &stathatStorage{
		key: k,
		w:   w,
	}, nil
}

func init() {
	carry.RegisterStorageConstructor("stathat", carry.StorageConstructor(newFromConfig))
}

// compile check to verify stathat implements carry.ShutdownStorage
var _ carry.ShutdownStorage = &stathatStorage{}
