//go:generate cproto
//go:generate svcdec -type OncallInfoServer
//go:generate svcdec -type RotationAdminServer
//go:generate goimports -w .

package rotangapi
