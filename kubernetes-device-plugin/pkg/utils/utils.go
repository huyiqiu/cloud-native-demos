package utils

import (
	pluginapi "k8s.io/kubelet/pkg/apis/deviceplugin/v1beta1"
)

func Diff(new_list, old_list []*pluginapi.Device) bool {
	if len(new_list) != len(old_list) {
		return true
	}
	m := make(map[string]int)
	for _, new := range new_list {
		m[new.ID] ++
	}
	for _, old := range old_list {
		if m[old.ID] < 1 {
			return true
		}
	}
	return false
}