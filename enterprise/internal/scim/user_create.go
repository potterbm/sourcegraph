package scim

import (
	"net/http"
	"strconv"
	"time"

	"github.com/elimity-com/scim"
	scimerrors "github.com/elimity-com/scim/errors"
	"github.com/sourcegraph/sourcegraph/internal/database"
	"github.com/sourcegraph/sourcegraph/internal/extsvc"
	"github.com/sourcegraph/sourcegraph/internal/types"
	"github.com/sourcegraph/sourcegraph/lib/errors"
)

// Create stores given attributes. Returns a resource with the attributes that are stored and a (new) unique identifier.
func (h *UserResourceHandler) Create(r *http.Request, attributes scim.ResourceAttributes) (scim.Resource, error) {
	// Extract external ID, primary email, username, and display name from attributes to variables
	optionalExternalID := getOptionalExternalID(attributes)
	primaryEmail := extractPrimaryEmail(attributes)
	if primaryEmail == "" {
		return scim.Resource{}, scimerrors.ScimErrorBadParams([]string{"emails missing"})
	}
	displayName := extractDisplayName(attributes)

	// Make sure the username is unique, then create user with/without an external account ID
	var user *types.User
	err := h.db.WithTransact(r.Context(), func(tx database.DB) error {
		uniqueUsername, err := getUniqueUsername(r.Context(), tx.Users(), extractStringAttribute(attributes, AttrUserName))
		if err != nil {
			return err
		}

		// Create user (with or without external ID)
		// TODO: Use NewSCIMUser instead of NewUser?
		newUser := database.NewUser{
			Email:           primaryEmail,
			Username:        uniqueUsername,
			DisplayName:     displayName,
			EmailIsVerified: true,
		}
		var externalID = ""
		if optionalExternalID.Present() {
			externalID = optionalExternalID.Value()
		}
		if externalID == "" {
			externalID = "no-external-id-" + primaryEmail
		}
		accountSpec := extsvc.AccountSpec{
			ServiceType: "scim",
			// TODO: provide proper service ID
			ServiceID: "TODO",
			AccountID: externalID,
		}

		accountData, err := toAccountData(attributes)
		if err != nil {
			return scimerrors.ScimError{Status: http.StatusInternalServerError, Detail: err.Error()}
		}
		user, err = h.db.UserExternalAccounts().CreateUserAndSave(r.Context(), newUser, accountSpec, accountData)

		if err != nil {
			if dbErr, ok := containsErrCannotCreateUserError(err); ok {
				code := dbErr.Code()
				if code == database.ErrorCodeUsernameExists || code == database.ErrorCodeEmailExists {
					return scimerrors.ScimError{Status: http.StatusConflict, Detail: err.Error()}
				}
			}
			return scimerrors.ScimError{Status: http.StatusInternalServerError, Detail: err.Error()}
		}
		return nil
	})
	if err != nil {
		return scim.Resource{}, err
	}

	var now = time.Now()

	return scim.Resource{
		ID:         strconv.Itoa(int(user.ID)),
		ExternalID: optionalExternalID,
		Attributes: attributes,
		Meta: scim.Meta{
			Created:      &now,
			LastModified: &now,
		},
	}, nil
}

// extractPrimaryEmail extracts the primary email address from the given attributes.
// Tries to get the (first) email address marked as primary, otherwise uses the first email address it finds.
func extractPrimaryEmail(attributes scim.ResourceAttributes) (primaryEmail string) {
	if attributes[AttrEmails] == nil {
		return
	}
	emails := attributes[AttrEmails].([]interface{})
	for _, emailRaw := range emails {
		email := emailRaw.(map[string]interface{})
		if email["primary"] == true {
			primaryEmail = email["value"].(string)
			break
		}
	}
	if primaryEmail == "" && len(emails) > 0 {
		primaryEmail = emails[0].(map[string]interface{})["value"].(string)
	}
	return
}

// extractDisplayName extracts the user's display name from the given attributes.
// Ii defaults to the username if no display name is available.
func extractDisplayName(attributes scim.ResourceAttributes) (displayName string) {
	if attributes[AttrDisplayName] != nil {
		displayName = attributes[AttrDisplayName].(string)
	} else if attributes[AttrName] != nil {
		name := attributes[AttrName].(map[string]interface{})
		if name[AttrNameFormatted] != nil {
			displayName = name[AttrNameFormatted].(string)
		} else if name[AttrNameGiven] != nil && name[AttrNameFamily] != nil {
			if name[AttrNameMiddle] != nil {
				displayName = name[AttrNameGiven].(string) + " " + name[AttrNameMiddle].(string) + " " + name[AttrNameFamily].(string)
			} else {
				displayName = name[AttrNameGiven].(string) + " " + name[AttrNameFamily].(string)
			}
		}
	}
	// Fallback to username
	if displayName == "" {
		displayName = attributes[AttrUserName].(string)
	}
	return
}

// containsErrCannotCreateUserError returns true if the given error contains at least one database.ErrCannotCreateUser.
// It also returns the first such error.
func containsErrCannotCreateUserError(err error) (database.ErrCannotCreateUser, bool) {
	if err == nil {
		return database.ErrCannotCreateUser{}, false
	}
	if _, ok := err.(database.ErrCannotCreateUser); ok {
		return err.(database.ErrCannotCreateUser), true
	}

	// Handle multiError
	if multiErr, ok := err.(errors.MultiError); ok {
		for _, err := range multiErr.Errors() {
			if _, ok := err.(database.ErrCannotCreateUser); ok {
				return err.(database.ErrCannotCreateUser), true
			}
		}
	}

	return database.ErrCannotCreateUser{}, false
}
