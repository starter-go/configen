package v4

// Source 表示一个 go-source-file
type Source struct {
	OwnerPackage *Package
	Path         string
	Name         string
}
