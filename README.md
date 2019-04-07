CircleCI Terraform Provider
===========================

Setup
-----

```hcl
provider "circleci" {
  token = "<CircleCI Token>" // Will fallback to CIRCLECI_TOKEN environment variable if not explicitly specified
}
```

Resources
---------

### circleci_project

For reference, see [CircleCI's Projects and Builds Documentation](https://circleci.com/docs/2.0/project-build/).

#### Example Usage:

```hcl
resource "circleci_project" "example" {
  vcs_type = "github"
  username = "me"
  name     = "example"
}
```

#### Argument Reference:

The following arguments are supported:

* vcs_type (Required) - The version control system type that the project users. Can be either `github` or `bitbucket`.
* username (Required) - The username that owns the project in the version control system.
* name (Required) - The name of the project.

#### Attributes Reference

In addition to all arguments above, the following attributes are exported:

* id - A string with the format `<vcs_type>/<username>/<name>`.

#### Import

Instances can be imported using the id, e.g.

```shell
terraform import circleci_project.example github/me/example
```

Building The Provider
---------------------

```shell
make build
```
