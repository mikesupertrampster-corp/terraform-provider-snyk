package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"os"
	"testing"
)

var testAccProviders map[string]*schema.Provider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider()
	testAccProviders = map[string]*schema.Provider{
		"snyk": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ *schema.Provider = Provider()
}

func testAccPreCheck(t *testing.T) {
	if err := os.Getenv("SNYK_API_KEY"); err == "" {
		t.Fatal("SNYK_API_KEY must be set for acceptance tests")
	}

	if err := os.Getenv("SNYK_ORG_ID"); err == "" {
		t.Fatal("SNYK_ORG_ID must be set for acceptance tests")
	}
}
