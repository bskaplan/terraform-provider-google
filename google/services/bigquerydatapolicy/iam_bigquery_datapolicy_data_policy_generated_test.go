// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package bigquerydatapolicy_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/hashicorp/terraform-provider-google/google/acctest"
	"github.com/hashicorp/terraform-provider-google/google/envvar"
)

func TestAccBigqueryDatapolicyDataPolicyIamBindingGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
		"role":          "roles/viewer",
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccBigqueryDatapolicyDataPolicyIamBinding_basicGenerated(context),
			},
			{
				ResourceName:      "google_bigquery_datapolicy_data_policy_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/dataPolicies/%s roles/viewer", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf_test_data_policy%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Test Iam Binding update
				Config: testAccBigqueryDatapolicyDataPolicyIamBinding_updateGenerated(context),
			},
			{
				ResourceName:      "google_bigquery_datapolicy_data_policy_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/dataPolicies/%s roles/viewer", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf_test_data_policy%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccBigqueryDatapolicyDataPolicyIamMemberGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
		"role":          "roles/viewer",
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				// Test Iam Member creation (no update for member, no need to test)
				Config: testAccBigqueryDatapolicyDataPolicyIamMember_basicGenerated(context),
			},
			{
				ResourceName:      "google_bigquery_datapolicy_data_policy_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/dataPolicies/%s roles/viewer user:admin@hashicorptest.com", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf_test_data_policy%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccBigqueryDatapolicyDataPolicyIamPolicyGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
		"role":          "roles/viewer",
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccBigqueryDatapolicyDataPolicyIamPolicy_basicGenerated(context),
				Check:  resource.TestCheckResourceAttrSet("data.google_bigquery_datapolicy_data_policy_iam_policy.foo", "policy_data"),
			},
			{
				ResourceName:      "google_bigquery_datapolicy_data_policy_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/dataPolicies/%s", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf_test_data_policy%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccBigqueryDatapolicyDataPolicyIamPolicy_emptyBinding(context),
			},
			{
				ResourceName:      "google_bigquery_datapolicy_data_policy_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/dataPolicies/%s", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf_test_data_policy%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccBigqueryDatapolicyDataPolicyIamMember_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_bigquery_datapolicy_data_policy" "data_policy" {
  location         = "us-central1"
  data_policy_id   = "tf_test_data_policy%{random_suffix}"
  policy_tag       = google_data_catalog_policy_tag.policy_tag.name
  data_policy_type = "COLUMN_LEVEL_SECURITY_POLICY"
}

resource "google_data_catalog_policy_tag" "policy_tag" {
  taxonomy     = google_data_catalog_taxonomy.taxonomy.id
  display_name = "Low security"
  description  = "A policy tag normally associated with low security items"
}

resource "google_data_catalog_taxonomy" "taxonomy" {
  region                 = "us-central1"
  display_name           = "taxonomy%{random_suffix}"
  description            = "A collection of policy tags"
  activated_policy_types = ["FINE_GRAINED_ACCESS_CONTROL"]
}

resource "google_bigquery_datapolicy_data_policy_iam_member" "foo" {
  project = google_bigquery_datapolicy_data_policy.data_policy.project
  location = google_bigquery_datapolicy_data_policy.data_policy.location
  data_policy_id = google_bigquery_datapolicy_data_policy.data_policy.data_policy_id
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccBigqueryDatapolicyDataPolicyIamPolicy_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_bigquery_datapolicy_data_policy" "data_policy" {
  location         = "us-central1"
  data_policy_id   = "tf_test_data_policy%{random_suffix}"
  policy_tag       = google_data_catalog_policy_tag.policy_tag.name
  data_policy_type = "COLUMN_LEVEL_SECURITY_POLICY"
}

resource "google_data_catalog_policy_tag" "policy_tag" {
  taxonomy     = google_data_catalog_taxonomy.taxonomy.id
  display_name = "Low security"
  description  = "A policy tag normally associated with low security items"
}

resource "google_data_catalog_taxonomy" "taxonomy" {
  region                 = "us-central1"
  display_name           = "taxonomy%{random_suffix}"
  description            = "A collection of policy tags"
  activated_policy_types = ["FINE_GRAINED_ACCESS_CONTROL"]
}

data "google_iam_policy" "foo" {
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
  }
}

resource "google_bigquery_datapolicy_data_policy_iam_policy" "foo" {
  project = google_bigquery_datapolicy_data_policy.data_policy.project
  location = google_bigquery_datapolicy_data_policy.data_policy.location
  data_policy_id = google_bigquery_datapolicy_data_policy.data_policy.data_policy_id
  policy_data = data.google_iam_policy.foo.policy_data
}

data "google_bigquery_datapolicy_data_policy_iam_policy" "foo" {
  project = google_bigquery_datapolicy_data_policy.data_policy.project
  location = google_bigquery_datapolicy_data_policy.data_policy.location
  data_policy_id = google_bigquery_datapolicy_data_policy.data_policy.data_policy_id
  depends_on = [
    google_bigquery_datapolicy_data_policy_iam_policy.foo
  ]
}
`, context)
}

func testAccBigqueryDatapolicyDataPolicyIamPolicy_emptyBinding(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_bigquery_datapolicy_data_policy" "data_policy" {
  location         = "us-central1"
  data_policy_id   = "tf_test_data_policy%{random_suffix}"
  policy_tag       = google_data_catalog_policy_tag.policy_tag.name
  data_policy_type = "COLUMN_LEVEL_SECURITY_POLICY"
}

resource "google_data_catalog_policy_tag" "policy_tag" {
  taxonomy     = google_data_catalog_taxonomy.taxonomy.id
  display_name = "Low security"
  description  = "A policy tag normally associated with low security items"
}

resource "google_data_catalog_taxonomy" "taxonomy" {
  region                 = "us-central1"
  display_name           = "taxonomy%{random_suffix}"
  description            = "A collection of policy tags"
  activated_policy_types = ["FINE_GRAINED_ACCESS_CONTROL"]
}

data "google_iam_policy" "foo" {
}

resource "google_bigquery_datapolicy_data_policy_iam_policy" "foo" {
  project = google_bigquery_datapolicy_data_policy.data_policy.project
  location = google_bigquery_datapolicy_data_policy.data_policy.location
  data_policy_id = google_bigquery_datapolicy_data_policy.data_policy.data_policy_id
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccBigqueryDatapolicyDataPolicyIamBinding_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_bigquery_datapolicy_data_policy" "data_policy" {
  location         = "us-central1"
  data_policy_id   = "tf_test_data_policy%{random_suffix}"
  policy_tag       = google_data_catalog_policy_tag.policy_tag.name
  data_policy_type = "COLUMN_LEVEL_SECURITY_POLICY"
}

resource "google_data_catalog_policy_tag" "policy_tag" {
  taxonomy     = google_data_catalog_taxonomy.taxonomy.id
  display_name = "Low security"
  description  = "A policy tag normally associated with low security items"
}

resource "google_data_catalog_taxonomy" "taxonomy" {
  region                 = "us-central1"
  display_name           = "taxonomy%{random_suffix}"
  description            = "A collection of policy tags"
  activated_policy_types = ["FINE_GRAINED_ACCESS_CONTROL"]
}

resource "google_bigquery_datapolicy_data_policy_iam_binding" "foo" {
  project = google_bigquery_datapolicy_data_policy.data_policy.project
  location = google_bigquery_datapolicy_data_policy.data_policy.location
  data_policy_id = google_bigquery_datapolicy_data_policy.data_policy.data_policy_id
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccBigqueryDatapolicyDataPolicyIamBinding_updateGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_bigquery_datapolicy_data_policy" "data_policy" {
  location         = "us-central1"
  data_policy_id   = "tf_test_data_policy%{random_suffix}"
  policy_tag       = google_data_catalog_policy_tag.policy_tag.name
  data_policy_type = "COLUMN_LEVEL_SECURITY_POLICY"
}

resource "google_data_catalog_policy_tag" "policy_tag" {
  taxonomy     = google_data_catalog_taxonomy.taxonomy.id
  display_name = "Low security"
  description  = "A policy tag normally associated with low security items"
}

resource "google_data_catalog_taxonomy" "taxonomy" {
  region                 = "us-central1"
  display_name           = "taxonomy%{random_suffix}"
  description            = "A collection of policy tags"
  activated_policy_types = ["FINE_GRAINED_ACCESS_CONTROL"]
}

resource "google_bigquery_datapolicy_data_policy_iam_binding" "foo" {
  project = google_bigquery_datapolicy_data_policy.data_policy.project
  location = google_bigquery_datapolicy_data_policy.data_policy.location
  data_policy_id = google_bigquery_datapolicy_data_policy.data_policy.data_policy_id
  role = "%{role}"
  members = ["user:admin@hashicorptest.com", "user:gterraformtest1@gmail.com"]
}
`, context)
}
