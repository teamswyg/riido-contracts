package deviceprincipal

type PrincipalKind string

const (
	UserPrincipal   PrincipalKind = "UserPrincipal"
	DevicePrincipal PrincipalKind = "DevicePrincipal"
)

func PrincipalKinds() []PrincipalKind {
	return []PrincipalKind{UserPrincipal, DevicePrincipal}
}
