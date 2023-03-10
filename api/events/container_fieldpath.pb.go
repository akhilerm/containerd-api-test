// Code generated by protoc-gen-go-fieldpath. DO NOT EDIT.
// source: github.com/containerd/containerd/api/events/container.proto
package events

import (
	v2 "github.com/containerd/typeurl/v2"
	strings "strings"
)

// Field returns the value for the given fieldpath as a string, if defined.
// If the value is not defined, the second value will be false.
func (m *ContainerCreate) Field(fieldpath []string) (string, bool) {
	if len(fieldpath) == 0 {
		return "", false
	}
	switch fieldpath[0] {
	case "id":
		return string(m.ID), len(m.ID) > 0
	case "image":
		return string(m.Image), len(m.Image) > 0
	case "runtime":
		// NOTE(stevvooe): This is probably not correct in many cases.
		// We assume that the target message also implements the Field
		// method, which isn't likely true in a lot of cases.
		//
		// If you have a broken build and have found this comment,
		// you may be closer to a solution.
		if m.Runtime == nil {
			return "", false
		}
		return m.Runtime.Field(fieldpath[1:])
	}
	return "", false
}

// Field returns the value for the given fieldpath as a string, if defined.
// If the value is not defined, the second value will be false.
func (m *ContainerCreate_Runtime) Field(fieldpath []string) (string, bool) {
	if len(fieldpath) == 0 {
		return "", false
	}
	switch fieldpath[0] {
	case "name":
		return string(m.Name), len(m.Name) > 0
	case "options":
		decoded, err := v2.UnmarshalAny(m.Options)
		if err != nil {
			return "", false
		}
		adaptor, ok := decoded.(interface{ Field([]string) (string, bool) })
		if !ok {
			return "", false
		}
		return adaptor.Field(fieldpath[1:])
	}
	return "", false
}

// Field returns the value for the given fieldpath as a string, if defined.
// If the value is not defined, the second value will be false.
func (m *ContainerUpdate) Field(fieldpath []string) (string, bool) {
	if len(fieldpath) == 0 {
		return "", false
	}
	switch fieldpath[0] {
	case "id":
		return string(m.ID), len(m.ID) > 0
	case "image":
		return string(m.Image), len(m.Image) > 0
	case "labels":
		// Labels fields have been special-cased by name. If this breaks,
		// add better special casing to fieldpath plugin.
		if len(m.Labels) == 0 {
			return "", false
		}
		value, ok := m.Labels[strings.Join(fieldpath[1:], ".")]
		return value, ok
	case "snapshot_key":
		return string(m.SnapshotKey), len(m.SnapshotKey) > 0
	}
	return "", false
}

// Field returns the value for the given fieldpath as a string, if defined.
// If the value is not defined, the second value will be false.
func (m *ContainerDelete) Field(fieldpath []string) (string, bool) {
	if len(fieldpath) == 0 {
		return "", false
	}
	switch fieldpath[0] {
	case "id":
		return string(m.ID), len(m.ID) > 0
	}
	return "", false
}
