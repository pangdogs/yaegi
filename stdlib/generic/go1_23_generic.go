//go:build go1.23
// +build go1.23

package generic

import (
	_ "embed"
)

//go:embed go1_23_maps_maps.go.txt
var go1_23_maps_maps string

//go:embed go1_23_slices_math_bits.go.txt
var go1_23_slices_math_bits string

//go:embed go1_23_slices_cmp.go.txt
var go1_23_slices_cmp string

//go:embed go1_23_slices_iter.go.txt
var go1_23_slices_iter string

//go:embed go1_23_slices_zsortanyfunc.go.txt
var go1_23_slices_zsortanyfunc string

//go:embed go1_23_slices_zsortordered.go.txt
var go1_23_slices_zsortordered string

//go:embed go1_23_slices_sort.go.txt
var go1_23_slices_sort string

//go:embed go1_23_slices_slices.go.txt
var go1_23_slices_slices string

// Sources contains the list of generic packages source strings.
var Sources = [...]string{
	go1_23_maps_maps,
	go1_23_slices_math_bits,
	go1_23_slices_cmp,
	go1_23_slices_iter,
	go1_23_slices_zsortanyfunc,
	go1_23_slices_zsortordered,
	go1_23_slices_sort,
	go1_23_slices_slices,
}
