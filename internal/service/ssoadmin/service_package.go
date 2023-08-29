// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ssoadmin

import (
	"context"

	aws_sdkv1 "github.com/aws/aws-sdk-go/aws"
	request_sdkv1 "github.com/aws/aws-sdk-go/aws/request"
	ssoadmin_sdkv1 "github.com/aws/aws-sdk-go/service/ssoadmin"
	"github.com/hashicorp/aws-sdk-go-base/v2/awsv1shim/v2/tfawserr"
)

// CustomizeConn customizes a new AWS SDK for Go v1 client for this service package's AWS API.
func (p *servicePackage) CustomizeConn(ctx context.Context, conn *ssoadmin_sdkv1.SSOAdmin) (*ssoadmin_sdkv1.SSOAdmin, error) {
	// Reference: https://github.com/hashicorp/terraform-provider-aws/issues/19215.
	conn.Handlers.Retry.PushBack(func(r *request_sdkv1.Request) {
		switch err := r.Error; r.Operation.Name {
		case "AttachCustomerManagedPolicyReferenceToPermissionSet", "DetachCustomerManagedPolicyReferenceFromPermissionSet",
			"AttachManagedPolicyToPermissionSet", "DetachManagedPolicyFromPermissionSet",
			"PutPermissionsBoundaryToPermissionSet", "DeletePermissionsBoundaryFromPermissionSet",
			"ProvisionPermissionSet":
			if tfawserr.ErrCodeEquals(err, ssoadmin_sdkv1.ErrCodeConflictException, ssoadmin_sdkv1.ErrCodeThrottlingException) {
				r.Retryable = aws_sdkv1.Bool(true)
			}
		}
	})

	return conn, nil
}
