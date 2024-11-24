package main //todo fix
// Create a subdomain on nimbus.app.
import "github.com/anirudhi89/nimbus/tenant"


// ValidateCompanyName checks if the provided company name is valid and not already in use.
func ValidateCompanyName(subdomain tenant) (bool, error) 

// GenerateSubdomain generates a unique subdomain based on the company name.
func GenerateSubdomain(subdomain tenant) (string, error)

// SaveSubdomain associates the generated subdomain with the user in the database.
func SaveSubdomain(subdomain tenant) error

// GetSubdomain retrieves the subdomain for a given user.
func GetSubdomain(subdomain tenant) (string, error)
