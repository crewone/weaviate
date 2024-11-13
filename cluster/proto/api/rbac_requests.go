//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2024 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

package api

import (
	"github.com/weaviate/weaviate/entities/models"
)

type CreateRolesRequest struct {
	Roles []*models.Role
}

type DeleteRolesRequest struct {
	Roles []string
}

type RemovePermissionsRequest struct {
	Role        string
	Permissions []*models.Permission
}

type AddRolesForUsersRequest struct {
	User  string
	Roles []string
}

type RevokeRolesForUserRequest struct {
	User  string
	Roles []string
}
