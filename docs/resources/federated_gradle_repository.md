---
subcategory: "Federated Repositories"
---
# Artifactory Federated Gradle Repository Resource

Creates a federated Gradle repository.

## Example Usage

```hcl
resource "artifactory_federated_gradle_repository" "terraform-federated-test-gradle-repo" {
  key       = "terraform-federated-test-gradle-repo"

  member {
    url     = "http://tempurl.org/artifactory/terraform-federated-test-gradle-repo"
    enabled = true
  }

  member {
    url     = "http://tempurl2.org/artifactory/terraform-federated-test-gradle-repo-2"
    enabled = true
  }
}
```

## Argument Reference

The following attributes are supported, along with the [list of attributes from the local Gradle repository](local_gradle_repository.md):

* `key` - (Required) the identity key of the repo.
* `member` - (Required) The list of Federated members and must contain this repository URL (configured base URL
  `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
  Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
  to set up Federated repositories correctly.
  * `url` - (Required) Full URL to ending with the repository name.
  * `enabled` - (Required) Represents the active state of the federated member. It is supported to change the enabled
    status of my own member. The config will be updated on the other federated members automatically.

## Import

Federated repositories can be imported using their name, e.g.
```
$ terraform import artifactory_federated_gradle_repository.terraform-federated-test-gradle-repo terraform-federated-test-gradle-repo
```
