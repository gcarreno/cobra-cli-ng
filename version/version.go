package version

import (
	"fmt"
)

type Version struct {
	Major int
	Minor int
	Patch int
}

func NewVersion() *Version {
	return &Version{
		Major: 0,
		Minor: 0,
		Patch: 1,
	}
}

func (v *Version) AsString() string {
	return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
}

func (v *Version) AsSemver() string {
	return fmt.Sprintf("v%d.%d.%d", v.Major, v.Minor, v.Patch)
}
