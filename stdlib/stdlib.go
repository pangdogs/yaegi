//go:build go1.21

// Package stdlib provides wrappers of standard library packages to be imported natively in Yaegi.
package stdlib

import "reflect"

// Symbols variable stores the map of stdlib symbols per package.
var Symbols = map[string]map[string]reflect.Value{}

// MapTypes variable contains a map of functions which have an interface{} as parameter but
// do something special if the parameter implements a given interface.
var MapTypes = map[reflect.Value][]reflect.Type{}

func init() {
	Symbols["github.com/pangdogs/yaegi/stdlib/stdlib"] = map[string]reflect.Value{
		"Symbols": reflect.ValueOf(Symbols),
	}
	Symbols["."] = map[string]reflect.Value{
		"MapTypes": reflect.ValueOf(MapTypes),
	}
}

// Provide access to go standard library (http://golang.org/pkg/)
// go list std | grep -v internal | grep -v '\.' | grep -v unsafe | grep -v syscall

//go:generate ../internal/cmd/extract/extract archive/tar
//go:generate ../internal/cmd/extract/extract archive/zip
//go:generate ../internal/cmd/extract/extract bufio
//go:generate ../internal/cmd/extract/extract bytes
//go:generate ../internal/cmd/extract/extract cmp
//go:generate ../internal/cmd/extract/extract compress/bzip2
//go:generate ../internal/cmd/extract/extract compress/flate
//go:generate ../internal/cmd/extract/extract compress/gzip
//go:generate ../internal/cmd/extract/extract compress/lzw
//go:generate ../internal/cmd/extract/extract compress/zlib
//go:generate ../internal/cmd/extract/extract container/heap
//go:generate ../internal/cmd/extract/extract container/list
//go:generate ../internal/cmd/extract/extract container/ring
//go:generate ../internal/cmd/extract/extract context
//go:generate ../internal/cmd/extract/extract crypto
//go:generate ../internal/cmd/extract/extract crypto/aes
//go:generate ../internal/cmd/extract/extract crypto/cipher
//go:generate ../internal/cmd/extract/extract crypto/des
//go:generate ../internal/cmd/extract/extract crypto/dsa
//go:generate ../internal/cmd/extract/extract crypto/ecdh
//go:generate ../internal/cmd/extract/extract crypto/ecdsa
//go:generate ../internal/cmd/extract/extract crypto/ed25519
//go:generate ../internal/cmd/extract/extract crypto/elliptic
//go:generate ../internal/cmd/extract/extract crypto/hmac
//go:generate ../internal/cmd/extract/extract crypto/md5
//go:generate ../internal/cmd/extract/extract crypto/rand
//go:generate ../internal/cmd/extract/extract crypto/rc4
//go:generate ../internal/cmd/extract/extract crypto/rsa
//go:generate ../internal/cmd/extract/extract crypto/sha1
//go:generate ../internal/cmd/extract/extract crypto/sha256
//go:generate ../internal/cmd/extract/extract crypto/sha512
//go:generate ../internal/cmd/extract/extract crypto/subtle
//go:generate ../internal/cmd/extract/extract crypto/tls
//go:generate ../internal/cmd/extract/extract crypto/x509
//go:generate ../internal/cmd/extract/extract crypto/x509/pkix
//go:generate ../internal/cmd/extract/extract database/sql
//go:generate ../internal/cmd/extract/extract database/sql/driver
//go:generate ../internal/cmd/extract/extract debug/buildinfo
//go:generate ../internal/cmd/extract/extract debug/dwarf
//go:generate ../internal/cmd/extract/extract debug/elf
//go:generate ../internal/cmd/extract/extract debug/gosym
//go:generate ../internal/cmd/extract/extract debug/macho
//go:generate ../internal/cmd/extract/extract debug/pe
//go:generate ../internal/cmd/extract/extract debug/plan9obj
//go:generate ../internal/cmd/extract/extract encoding
//go:generate ../internal/cmd/extract/extract encoding/ascii85
//go:generate ../internal/cmd/extract/extract encoding/asn1
//go:generate ../internal/cmd/extract/extract encoding/base32
//go:generate ../internal/cmd/extract/extract encoding/base64
//go:generate ../internal/cmd/extract/extract encoding/binary
//go:generate ../internal/cmd/extract/extract encoding/csv
//go:generate ../internal/cmd/extract/extract encoding/gob
//go:generate ../internal/cmd/extract/extract encoding/hex
//go:generate ../internal/cmd/extract/extract encoding/json
//go:generate ../internal/cmd/extract/extract encoding/pem
//go:generate ../internal/cmd/extract/extract encoding/xml
//go:generate ../internal/cmd/extract/extract errors
//go:generate ../internal/cmd/extract/extract expvar
//go:generate ../internal/cmd/extract/extract flag
//go:generate ../internal/cmd/extract/extract fmt
//go:generate ../internal/cmd/extract/extract go/ast
//go:generate ../internal/cmd/extract/extract go/build
//go:generate ../internal/cmd/extract/extract go/build/constraint
//go:generate ../internal/cmd/extract/extract go/constant
//go:generate ../internal/cmd/extract/extract go/doc
//go:generate ../internal/cmd/extract/extract go/doc/comment
//go:generate ../internal/cmd/extract/extract go/format
//go:generate ../internal/cmd/extract/extract go/importer
//go:generate ../internal/cmd/extract/extract go/parser
//go:generate ../internal/cmd/extract/extract go/printer
//go:generate ../internal/cmd/extract/extract go/scanner
//go:generate ../internal/cmd/extract/extract go/token
//go:generate ../internal/cmd/extract/extract go/types
//go:generate ../internal/cmd/extract/extract hash
//go:generate ../internal/cmd/extract/extract hash/adler32
//go:generate ../internal/cmd/extract/extract hash/crc32
//go:generate ../internal/cmd/extract/extract hash/crc64
//go:generate ../internal/cmd/extract/extract hash/fnv
//go:generate ../internal/cmd/extract/extract hash/maphash
//go:generate ../internal/cmd/extract/extract html
//go:generate ../internal/cmd/extract/extract html/template
//go:generate ../internal/cmd/extract/extract image
//go:generate ../internal/cmd/extract/extract image/color
//go:generate ../internal/cmd/extract/extract image/color/palette
//go:generate ../internal/cmd/extract/extract image/draw
//go:generate ../internal/cmd/extract/extract image/gif
//go:generate ../internal/cmd/extract/extract image/jpeg
//go:generate ../internal/cmd/extract/extract image/png
//go:generate ../internal/cmd/extract/extract index/suffixarray
//go:generate ../internal/cmd/extract/extract io
//go:generate ../internal/cmd/extract/extract io/fs
//go:generate ../internal/cmd/extract/extract io/ioutil
//go:generate ../internal/cmd/extract/extract log
//go:generate ../internal/cmd/extract/extract log/slog
//go:generate ../internal/cmd/extract/extract log/syslog
//go:generate ../internal/cmd/extract/extract maps
//go:generate ../internal/cmd/extract/extract math
//go:generate ../internal/cmd/extract/extract math/big
//go:generate ../internal/cmd/extract/extract math/bits
//go:generate ../internal/cmd/extract/extract math/cmplx
//go:generate ../internal/cmd/extract/extract math/rand
//go:generate ../internal/cmd/extract/extract mime
//go:generate ../internal/cmd/extract/extract mime/multipart
//go:generate ../internal/cmd/extract/extract mime/quotedprintable
//go:generate ../internal/cmd/extract/extract net
//go:generate ../internal/cmd/extract/extract net/http
//go:generate ../internal/cmd/extract/extract net/http/cgi
//go:generate ../internal/cmd/extract/extract net/http/cookiejar
//go:generate ../internal/cmd/extract/extract net/http/fcgi
//go:generate ../internal/cmd/extract/extract net/http/httptest
//go:generate ../internal/cmd/extract/extract net/http/httptrace
//go:generate ../internal/cmd/extract/extract net/http/httputil
//go:generate ../internal/cmd/extract/extract net/http/pprof
//go:generate ../internal/cmd/extract/extract net/mail
//go:generate ../internal/cmd/extract/extract net/netip
//go:generate ../internal/cmd/extract/extract net/rpc
//go:generate ../internal/cmd/extract/extract net/rpc/jsonrpc
//go:generate ../internal/cmd/extract/extract net/smtp
//go:generate ../internal/cmd/extract/extract net/textproto
//go:generate ../internal/cmd/extract/extract net/url
//go:generate ../internal/cmd/extract/extract os
//go:generate ../internal/cmd/extract/extract os/signal
//go:generate ../internal/cmd/extract/extract os/user
//go:generate ../internal/cmd/extract/extract path
//go:generate ../internal/cmd/extract/extract path/filepath
//go:generate ../internal/cmd/extract/extract reflect
//go:generate ../internal/cmd/extract/extract regexp
//go:generate ../internal/cmd/extract/extract regexp/syntax
//go:generate ../internal/cmd/extract/extract runtime
//go:generate ../internal/cmd/extract/extract runtime/debug
//go:generate ../internal/cmd/extract/extract runtime/metrics
//go:generate ../internal/cmd/extract/extract runtime/pprof
//go:generate ../internal/cmd/extract/extract runtime/trace
//go:generate ../internal/cmd/extract/extract slices
//go:generate ../internal/cmd/extract/extract sort
//go:generate ../internal/cmd/extract/extract strconv
//go:generate ../internal/cmd/extract/extract strings
//go:generate ../internal/cmd/extract/extract sync
//go:generate ../internal/cmd/extract/extract sync/atomic
//go:generate ../internal/cmd/extract/extract testing
//go:generate ../internal/cmd/extract/extract testing/fstest
//go:generate ../internal/cmd/extract/extract testing/iotest
//go:generate ../internal/cmd/extract/extract testing/quick
//go:generate ../internal/cmd/extract/extract testing/slogtest
//go:generate ../internal/cmd/extract/extract text/scanner
//go:generate ../internal/cmd/extract/extract text/tabwriter
//go:generate ../internal/cmd/extract/extract text/template
//go:generate ../internal/cmd/extract/extract text/template/parse
//go:generate ../internal/cmd/extract/extract time
//go:generate ../internal/cmd/extract/extract unicode
//go:generate ../internal/cmd/extract/extract unicode/utf16
//go:generate ../internal/cmd/extract/extract unicode/utf8
