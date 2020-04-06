package common

// All tests are be linked[1] to the integreatly-test-cases[2] repo by using the same ID
// 1. https://gitlab.cee.redhat.com/integreatly-qe/integreatly-test-cases#how-to-automate-a-test-case-and-link-it-back
// 2. https://gitlab.cee.redhat.com/integreatly-qe/integreatly-test-cases
var (
	ALL_TESTS = []TestCase{
		// Add all the tests that should be executed in both e2e and osd suites here.
		// It is an array so the tests will be executed in the same order as they defined here.
		{"Verify CRD Exists", TestIntegreatlyCRDExists},
	}

	AFTER_INSTALL_TESTS = []TestCase{
		// Test to be execute after RHMI has been installed
		{"USer SSO permissions", TestUserSsoPermissions},
	}
)
