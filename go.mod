module github.com/hashicorp/terraform-provider-nomad

go 1.17

exclude (
	github.com/Sirupsen/logrus v1.1.0
	github.com/Sirupsen/logrus v1.1.1
	github.com/Sirupsen/logrus v1.2.0
	github.com/Sirupsen/logrus v1.3.0
	github.com/Sirupsen/logrus v1.4.0
	github.com/Sirupsen/logrus v1.4.1
)

// Set to the version used by the Terraform Plugin SDK.
// https://github.com/hashicorp/terraform-plugin-sdk/blob/v1.17.2/go.mod#L47
replace github.com/spf13/afero => github.com/spf13/afero v1.2.2

require (
	github.com/dustin/go-humanize v1.0.0
	github.com/google/go-cmp v0.5.6
	github.com/hashicorp/go-cleanhttp v0.5.2
	github.com/hashicorp/go-multierror v1.1.1
	github.com/hashicorp/go-version v1.4.0
	github.com/hashicorp/nomad v1.3.1
	github.com/hashicorp/nomad/api v0.0.0-20220519231241-2b054e38e91a
	github.com/hashicorp/terraform-plugin-sdk v1.17.2
	github.com/hashicorp/vault v0.10.4
	github.com/stretchr/testify v1.7.1
)
