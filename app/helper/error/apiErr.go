package error

import "errors"

var InternalServerError = errors.New("Internal server error")

// user
var ErrUserNotFound = errors.New("User not found")

// role
var ErrRoleNotFound = errors.New("No Role found")

// menu
var ErrMenuOrPathAlreadyExists = errors.New("menu already or path exists")
var ErrMenuNotFound = errors.New("Menu not found")

// permission
var ErrPermissionNotFound = errors.New("No Permission found")
var ErrPermissionNotFoundForUser = errors.New("No permissions found for user")

// subscription
var ErrSubscriptionCategoryAlreadyExists = errors.New("Subscription already or path exists")
var ErrSubscriptionCategoryNotFound = errors.New("No subscription category found")
var ErrCategoryBenefitNotFound = errors.New("No category benefit found")
var ErrCategoryBenefitNotEmpty = errors.New("Subscription category benefits not empty")
