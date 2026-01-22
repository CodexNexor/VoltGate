package auth

import (
	"fmt"

	"github.com/voltgate/voltgate/v6/internal/interfaces"
)

// ProjectSelectionError is returned when the user must pick from multiple GCP projects.
type ProjectSelectionError struct {
	Email    string
	Projects []interfaces.GCPProjectProjects
}

func (e *ProjectSelectionError) Error() string {
	if e == nil {
		return "voltgate auth: project selection required"
	}
	return fmt.Sprintf("voltgate auth: project selection required for %s", e.Email)
}

// ProjectsDisplay returns the projects list for caller presentation.
func (e *ProjectSelectionError) ProjectsDisplay() []interfaces.GCPProjectProjects {
	if e == nil {
		return nil
	}
	return e.Projects
}

// EmailRequiredError indicates that the calling context must provide an email or alias.
type EmailRequiredError struct {
	Prompt string
}

func (e *EmailRequiredError) Error() string {
	if e == nil || e.Prompt == "" {
		return "voltgate auth: email is required"
	}
	return e.Prompt
}
