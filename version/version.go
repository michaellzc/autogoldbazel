package version

import (
	goversion "github.com/hashicorp/go-version"
)

type Version struct {
	semver *goversion.Version
}

func New(v string) (*Version, error) {
	version, err := goversion.NewSemver(v)
	if err != nil {
		return nil, err
	}
	return &Version{semver: version}, nil
}
