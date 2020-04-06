package common

import (
	goctx "context"
	"fmt"
	"github.com/integr8ly/keycloak-client/pkg/common"
	keycloak "github.com/keycloak/keycloak-operator/pkg/apis/keycloak/v1alpha1"
	"k8s.io/apimachinery/pkg/types"
	"testing"

	"github.com/integr8ly/integreatly-operator/test/resources"
)

func TestUserSsoPermissions(t *testing.T, ctx *TestingContext) {
	if err := createTestingIDP(goctx.TODO(), ctx.Client, ctx.HttpClient, ctx.SelfSignedCerts); err != nil {
		t.Fatalf("error while creating testing idp: %v", err)
	}

	// get console master url
	rhmi, err := getRHMI(ctx.Client)
	if err != nil {
		t.Fatalf("error getting RHMI CR: %v", err)
	}
	masterURL := rhmi.Spec.MasterURL

	if err := resources.DoAuthOpenshiftUser(fmt.Sprintf("%s/auth/login", masterURL), "test-user-1", "Password1", ctx.HttpClient, "testing-idp"); err != nil {
		t.Fatalf("error occured trying to get token : %v", err)
	}

	kc := &keycloak.Keycloak{}
	selector := types.NamespacedName{
		Namespace: "redhat-rhmi-rhsso",
		Name:      "rhsso",
	}

	err = ctx.Client.Get(goctx.TODO(), selector, kc)
	if err != nil {
		t.Fatal(err)
	}

	keycloakFactory := common.LocalConfigKeycloakFactory{}
	client, err := keycloakFactory.AuthenticatedClient(*kc)
	if err != nil {
		t.Fatal(err)
	}

	realms, err := client.ListRealms()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(realms)
}
