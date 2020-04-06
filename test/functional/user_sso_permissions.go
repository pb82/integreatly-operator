package functional

import (
	"fmt"
	"net/http"
	"testing"
	"common/user_rhmi_developer_permissions"

	"github.com/integr8ly/integreatly-operator/test/resources"

	"k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	dynclient "sigs.k8s.io/controller-runtime/pkg/client"
)

type TestingContext struct {
	Client          dynclient.Client
	KubeConfig      *rest.Config
	KubeClient      kubernetes.Interface
	ExtensionClient *clientset.Clientset
	HttpClient      *http.Client
	SelfSignedCerts bool
}

func TestUserSsoPermissions(t *testing.T, ctx *TestingContext) {

	// get console master url
	rhmi, err := user_rhmi_developer_permissions.getRHMI(ctx.Client)
	if err != nil {
		t.Fatalf("error getting RHMI CR: %v", err)
	}
	masterURL := rhmi.Spec.MasterURL

	if err := resources.DoAuthOpenshiftUser(fmt.Sprintf("%s/auth/login", masterURL), "test-user-1", "Password1", ctx.HttpClient, "testing-idp"); err != nil {
		t.Fatalf("error occured trying to get token : %v", err)
	}

}
