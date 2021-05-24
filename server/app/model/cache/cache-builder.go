package cache

import "github.com/obonobo/express-vpn-updater/server/config"

type Builder struct {
	bucket    string
	sourceUrl string
}

func CacheBuilder() *Builder {
	return &Builder{}
}

func (b *Builder) WithBucket(bucket string) *Builder {
	b.bucket = bucket
	return b
}

func (b *Builder) WithSourceUrl(url string) *Builder {
	b.sourceUrl = url
	return b
}

func (b *Builder) Build() Cache {
	return NewCache(b.Bucket(), b.SourceUrl())
}

func (b *Builder) Bucket() string {
	return defaultIfEmpty(b.bucket, func() string {
		return config.Get().Bucket
	})
}

func (b *Builder) SourceUrl() string {
	return defaultIfEmpty(b.sourceUrl, func() string {
		return config.Get().Url
	})
}

func defaultIfEmpty(this string, orGet func() string) string {
	if this == "" {
		return orGet()
	}
	return this
}
