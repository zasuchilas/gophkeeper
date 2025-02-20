package grpcserver

// prefix from proto file (package name)
const prefix = "/api."

// OpenRoutes contains methods that are open for all.
func OpenRoutes() map[string]bool {
	return map[string]bool{
		prefix + "UserV1/Register": true,
		prefix + "UserV1/Login":    true,
	}
}

// CommonRoutes contains methods that has valid jwt.
func CommonRoutes() map[string]bool {
	return map[string]bool{
		prefix + "SecretsV1/List":   true,
		prefix + "SecretsV1/Get":    true,
		prefix + "SecretsV1/Create": true,
		prefix + "SecretsV1/Update": true,
		prefix + "SecretsV1/Delete": true,
	}
}
