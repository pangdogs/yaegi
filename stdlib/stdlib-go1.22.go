//go:build go1.22 && !go1.23

package stdlib

//go:generate ../internal/cmd/extract/extract go/version
//go:generate ../internal/cmd/extract/extract math/rand/v2
