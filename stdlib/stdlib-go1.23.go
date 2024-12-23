//go:build go1.23

package stdlib

//go:generate ../internal/cmd/extract/extract go/version
//go:generate ../internal/cmd/extract/extract iter
//go:generate ../internal/cmd/extract/extract math/rand/v2
//go:generate ../internal/cmd/extract/extract unique
