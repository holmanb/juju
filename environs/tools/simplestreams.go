// Copyright 2013 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

// The tools package supports locating, parsing, and filtering Ubuntu tools metadata in simplestreams format.
// See http://launchpad.net/simplestreams and in particular the doc/README file in that project for more information
// about the file formats.
package tools

import (
	"fmt"

	"launchpad.net/juju-core/environs/simplestreams"
	"launchpad.net/juju-core/version"
)

func init() {
	simplestreams.RegisterStructTags(ToolsMetadata{})
}

const (
	ContentDownload = "content-download"
)

// This needs to be a var so we can override it for testing.
var DefaultBaseURL = "http://juju.canonical.com/tools"

// ToolsConstraint defines criteria used to find a tools metadata record.
type ToolsConstraint struct {
	simplestreams.LookupParams
	Version      version.Number
	MajorVersion int
	MinorVersion int
}

// NewVersionedToolsConstraint returns a ToolsConstraint for a tools with a specific version.
func NewVersionedToolsConstraint(vers string, params simplestreams.LookupParams) *ToolsConstraint {
	versNum := version.MustParse(vers)
	return &ToolsConstraint{LookupParams: params, Version: versNum}
}

// NewGeneralToolsConstraint returns a ToolsConstraint for tools with matching major/minor version numbers.
func NewGeneralToolsConstraint(majorVersion, minorVersion int, params simplestreams.LookupParams) *ToolsConstraint {
	return &ToolsConstraint{LookupParams: params, Version: version.Zero, MajorVersion: majorVersion, MinorVersion: minorVersion}
}

// Generates a string array representing product ids formed similarly to an ISCSI qualified name (IQN).
func (tc *ToolsConstraint) Ids() ([]string, error) {
	version, err := simplestreams.SeriesVersion(tc.Series)
	if err != nil {
		return nil, err
	}
	ids := make([]string, len(tc.Arches))
	for i, arch := range tc.Arches {
		ids[i] = fmt.Sprintf("com.ubuntu.juju:%s:%s", version, arch)
	}
	return ids, nil
}

// ToolsMetadata holds information about a particular tools tarball.
type ToolsMetadata struct {
	Release  string  `json:"release"`
	Version  string  `json:"version"`
	Arch     string  `json:"arch"`
	Size     float64 `json:"size"`
	Path     string  `json:"path"`
	FileType string  `json:"ftype"`
	SHA256   string  `json:"sha256"`
}

// Fetch returns a list of tools for the specified cloud matching the constraint.
// The base URL locations are as specified - the first location which has a file is the one used.
// Signed data is preferred, but if there is no signed data available and onlySigned is false,
// then unsigned data is used.
func Fetch(baseURLs []string, indexPath string, cons *ToolsConstraint, onlySigned bool) ([]*ToolsMetadata, error) {
	params := simplestreams.ValueParams{
		DataType:      ContentDownload,
		FilterFunc:    appendMatchingTools,
		ValueTemplate: ToolsMetadata{},
	}
	items, err := simplestreams.GetMaybeSignedMetadata(baseURLs, indexPath+simplestreams.SignedSuffix, cons, true, params)
	if (err != nil || len(items) == 0) && !onlySigned {
		items, err = simplestreams.GetMaybeSignedMetadata(baseURLs, indexPath+simplestreams.UnsignedSuffix, cons, false, params)
	}
	if err != nil {
		return nil, err
	}
	metadata := make([]*ToolsMetadata, len(items))
	for i, md := range items {
		metadata[i] = md.(*ToolsMetadata)
	}
	return metadata, nil
}

// appendMatchingTools updates matchingTools with tools metadata records from tools which belong to the
// specified series. If a tools record already exists in matchingTools, it is not overwritten.
func appendMatchingTools(matchingTools []interface{}, tools map[string]interface{}, cons simplestreams.LookupConstraint) []interface{} {
	toolsMap := make(map[string]*ToolsMetadata, len(matchingTools))
	for _, val := range matchingTools {
		tm := val.(*ToolsMetadata)
		toolsMap[fmt.Sprintf("%s-%s-%s", tm.Release, tm.Version, tm.Arch)] = tm
	}
	for _, val := range tools {
		tm := val.(*ToolsMetadata)
		consSeries := cons.Params().Series
		if consSeries != "" && consSeries != tm.Release {
			continue
		}
		if toolsConstraint, ok := cons.(*ToolsConstraint); ok {
			tmNumber := version.MustParse(tm.Version)
			if toolsConstraint.Version == version.Zero {
				if toolsConstraint.MajorVersion != tmNumber.Major {
					continue
				}
				if toolsConstraint.MinorVersion >= 0 && toolsConstraint.MinorVersion != tmNumber.Minor {
					continue
				}
			} else {
				if toolsConstraint.Version != tmNumber {
					continue
				}
			}
		}
		if _, ok := toolsMap[fmt.Sprintf("%s-%s-%s", tm.Release, tm.Version, tm.Arch)]; !ok {
			matchingTools = append(matchingTools, tm)
		}
	}
	return matchingTools
}

// GetLatestToolsMetadata is provided so it can be call by tests outside the tools package.
func GetLatestToolsMetadata(data []byte, cons *ToolsConstraint) ([]*ToolsMetadata, error) {
	metadata, err := simplestreams.ParseCloudMetadata(data, "products:1.0", "<unknown>", ToolsMetadata{})
	if err != nil {
		return nil, err
	}
	items, err := simplestreams.GetLatestMetadata(metadata, cons, appendMatchingTools)
	if err != nil {
		return nil, err
	}
	result := make([]*ToolsMetadata, len(items))
	for i, md := range items {
		result[i] = md.(*ToolsMetadata)
	}
	return result, nil
}
