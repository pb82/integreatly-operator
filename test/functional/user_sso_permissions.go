package functional

import (
	"fmt"
	"testing"
	"common"
	"github.com/integr8ly/integreatly-operator/test/resources"
	"net/http"
	"net/url"
	"strings"
)

func TestUserSsoPermissions(t *testing.T, ctx *TestingContext) {

	//err := t.DoAuthOpenshiftUser(, "test-user-1", t.DefaultPassword)

	err := resources.DoAuthOpenshiftUser(t.http., "test-user-1", common.DefaultPassword, ctx., common.TestingIDPRealm)
	if err != nil {
		fmt.Fatalf("error: %e", err)
	}
	
	// if err != nil {
	// 	fmt.Fatalf("error: %e", err)
	// }

	// err := resources.DoAuthOpenshiftUser("","test-user-1", ctx.)
	// err := resources.DoAuthOpenshiftUser(masterURL, "test-user-1", common.DefaultPassword, ctx., common.TestingIDPRealm)
	// if err != nil {
	// 	fmt.Fatalf("error: %e", err)

}
