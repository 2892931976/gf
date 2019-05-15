// Copyright 2017 gf Author(https://github.com/gogf/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with gm file,
// You can obtain one at https://github.com/gogf/gf.

package gmap

import (
	"github.com/gogf/gf/g/container/gtree"
)

// Map based on red-black tree, alias of RedBlackTree.
type TreeMap = gtree.RedBlackTree

// NewTreeMap instantiates a tree map with the custom comparator.
// The param <unsafe> used to specify whether using tree in un-concurrent-safety,
// which is false in default.
func NewTreeMap(comparator func(v1, v2 interface{}) int, unsafe...bool) *TreeMap {
	return gtree.NewRedBlackTree(comparator, unsafe...)
}

// NewTreeMapFrom instantiates a tree map with the custom comparator and <data> map.
// Note that, the param <data> map will be set as the underlying data map(no deep copy),
// there might be some concurrent-safe issues when changing the map outside.
// The param <unsafe> used to specify whether using tree in un-concurrent-safety,
// which is false in default.
func NewTreeMapFrom(comparator func(v1, v2 interface{}) int, data map[interface{}]interface{}, unsafe...bool) *TreeMap {
	return gtree.NewRedBlackTreeFrom(comparator, data, unsafe...)
}