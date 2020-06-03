package resolver

// Resolver is a interface for metadata resolvers
type Resolver interface {
	Resolve(string) (Metadata, error)
}

// Metadata of novel games
type Metadata struct {
	Title       string
	Thumbnail   string
	IsR18       bool
	IsNovelGame bool
}
