package functional

import (
	"fmt"
	"testing"

	"github.com/integr8ly/integreatly-operator/test/resources"
)

const (
	TestingIDPRealm                = "testing-idp"
	defaultDedicatedAdminName      = "customer-admin"
	defaultNumberOfTestUsers       = 2
	defaultNumberOfDedicatedAdmins = 2
	defaultSecret                  = "rhmiForeva"
	DefaultTestUserName            = "test-user"
	DefaultPassword                = "Password1"
)

func TestUserSsoPermissions(t *testing.T, ctx *TestingContext) {

	// get console master url
	rhmi, err := shared_functions.getRHMI(ctx.Client)
	if err != nil {
		t.Fatalf("error getting RHMI CR: %v", err)
	}
	masterURL := rhmi.Spec.MasterURL

	if err := resources.DoAuthOpenshiftUser(fmt.Sprintf("%s/auth/login", masterURL), "test-user-1", "Password1", ctx.HttpClient, "testing-idp"); err != nil {
		t.Fatalf("error occured trying to get token : %v", err)
	}

}
