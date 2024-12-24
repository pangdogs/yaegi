//go:build go1.23
// +build go1.23

package generic

import (
	_ "embed"
)

//go:embed go1_23_cmp_cmp.go.txt
var cmpSource string

////go:embed go1_23_iter_iter.go.txt
//var iterSource string

//go:embed go1_23_maps_maps.go.txt
var mapsSource string

//go:embed go1_23_slices_slices.go.txt
var slicesSource string

//go:embed go1_23_slices_sort.go.txt
var slicesSortSource string

//go:embed go1_23_slices_zsortanyfunc.go.txt
var slicesZsortanyfuncSource string

//go:embed go1_23_slices_zsortordered.go.txt
var slicesZsortorderedSource string

////go:embed go1_23_sync_atomic_type.go.txt
//var syncAtomicTypeSource string

//go:embed go1_23_sync_oncefunc.go.txt
var syncOncefuncSource string

////go:embed go1_23_unique_handle.go.txt
//var uniqueHandle string
//
////go:embed go1_23_unique_clone.go.txt
//var uniqueClone string

// Sources contains the list of generic packages source strings.
var Sources = [...]string{
	cmpSource,
	//iterSource,
	mapsSource,
	slicesZsortanyfuncSource,
	slicesZsortorderedSource,
	slicesSortSource,
	slicesSource,
	//syncAtomicTypeSource,
	syncOncefuncSource,
	//uniqueHandle,
	//uniqueClone,
}
