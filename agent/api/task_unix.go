// +build !windows

// Copyright 2014-2016 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//	http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package api

import (
	"github.com/fsouza/go-dockerclient"
)

const (
	portBindingHostIP = "0.0.0.0"

	//memorySwappinessDefault is the expected default value for this platform. This is used in task_windows.go
	//and is maintained here for unix default. Also used for testing
	memorySwappinessDefault = 0
)

func (task *Task) adjustForPlatform() {}

func getCanonicalPath(path string) string { return path }

// CgroupEnabled checks whether cgroups need be enabled at the task level
func (task *Task) CgroupEnabled() bool {
	//stub for now
	return false
}

func (task *Task) platformHostConfigOverride(hostConfig *docker.HostConfig) {}
