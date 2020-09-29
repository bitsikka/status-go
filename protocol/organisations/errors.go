package organisations

import "errors"

var ErrChatNotFound = errors.New("chat not found")
var ErrCantRequestAccess = errors.New("can't request access")
var ErrInvalidOrganisationDescription = errors.New("invalid organisation description")
var ErrNotAdmin = errors.New("no admin privileges for this organisation")
