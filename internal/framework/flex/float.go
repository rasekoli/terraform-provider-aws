// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package flex

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Float64ToFramework converts a float64 pointer to a Framework Float64 value.
// A nil float64 pointer is converted to a null Float64.
func Float64ToFramework(_ context.Context, v *float64) types.Float64 {
	if v == nil {
		return types.Float64Null()
	}

	return types.Float64Value(aws.ToFloat64(v))
}

// Float64ToFrameworkLegacy converts a float64 pointer to a Framework Float64 value.
// A nil float64 pointer is converted to a zero float64.
func Float64ToFrameworkLegacy(_ context.Context, v *float64) types.Float64 {
	return types.Float64Value(aws.ToFloat64(v))
}
