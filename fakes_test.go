//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package fakesuserstudy

import (
	"net/http"
	"testing"
)

// docs https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v5@v5.1.0-beta.1#readme-fakes
// see https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/azcore@v1.7.0-beta.2/fake for general docs on fakes

func Test_VirtualMachinesClient_Get(t *testing.T) {
	// write a fake for VirtualMachinesClient.Get that includes the following data

	const (
		// the name of the VM returned from VirtualMachinesClient.Get
		vmName = "virtualmachine1"

		// the resource ID of the VM returned from VirtualMachinesClient.Get
		resourceID = "/fake/resource/id"
	)
}

func Test_VirtualMachinesClient_Get_error(t *testing.T) {
	// write a fake for VirtualMachinesClient.Get that includes the following data

	const (
		// the HTTP status code of the failed request
		httpError = http.StatusBadRequest

		// the error code of the failed request
		errorCode = "ErrorResourceNotFound"
	)
}

func Test_VirtualMachinesClient_BeginCreateOrUpdate(t *testing.T) {
	// write a fake for VirtualMachinesClient.BeginCreateOrUpdate that includes the following data

	const (
		// the name of the VM returned when the long-running operation completes
		vmName = "virtualmachine1"

		// the resource ID of the VM returned when the long-running operation completes
		resourceID = "/fake/resource/id"
	)
}

func Test_VirtualMachinesClient_BeginCreateOrUpdate_error(t *testing.T) {
	// write a fake for VirtualMachinesClient.BeginCreateOrUpdate that includes the following data

	const (
		// the name of the VM returned when the long-running operation completes
		vmName = "virtualmachine1"

		// the resource ID of the VM returned when the long-running operation completes
		resourceID = "/fake/resource/id"
	)
}
