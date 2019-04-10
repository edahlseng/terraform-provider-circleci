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

### circleci_ssh_key

For reference, see [CircleCI's documentation on Adding an SSH Key](https://circleci.com/docs/2.0/add-ssh-key/).

#### Example Usage:

```hcl
resource "tls_private_key" "example_key" {
  algorithm = "RSA"
  rsa_bits  = "4096"
}

resource "circleci_ssh_key" "example_ssh_key" {
  project_id      = "${circleci_project.example.id}"
  hostname        = "example.com"
  private_key     = "${tls.private_key.example_key.private_key_pem}"
  fingerprint_md5 = "${tls.private_key.example_key.public_key_fingerprint_md5}"
}
```

#### Argument Reference:

The following arguments are supported:

* project_id (Required) - The ID of the project (in `<vcs_type>/<username>/<name>` format).
* hostname (Required) - The hostname for the key.
* private_key (Required) - The private key to use for the given hostname.
* fingerprint_md5 (Required) - The MD5 fingerprint of the given private_key.

#### Attributes Reference

In addition to all arguments above, the following attributes are exported:

* id - A string with the format `<vcs_type>/<username>/<name>/<hostname>/<fingerprint_md5>`.

#### Import

The `circleci_ssh_key` resource does not support importing.

Building The Provider
---------------------

```shell
make build
```
