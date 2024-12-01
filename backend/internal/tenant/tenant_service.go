package tenant

// ValidateCompanyName checks if the provided company name is valid and not already in use.
func ValidateCompanyName(subdomain Tenant) (bool, error)

// GenerateSubdomain generates a unique subdomain based on the company name.
func GenerateSubdomain(subdomain Tenant) (string, error)

// SaveSubdomain associates the generated subdomain with the user in the database.
func SaveSubdomain(subdomain Tenant) error

// GetSubdomain retrieves the subdomain for a given user.
func GetSubdomain(subdomain Tenant) (string, error)
