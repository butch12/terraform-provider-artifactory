---
subcategory: "Remote Repositories"
---
# Artifactory Remote Repository Resource

Provides a remote Helm repository. 
Official documentation can be found [here](https://www.jfrog.com/confluence/display/JFROG/Kubernetes+Helm+Chart+Repositories).


## Example Usage

```hcl
resource "artifactory_remote_helm_repository" "helm-remote" {
  key                             = "helm-remote-foo25"
  url                             = "https://repo.chartcenter.io/"
  helm_charts_base_url            = "https://foo.com"
  external_dependencies_enabled   = true
  external_dependencies_patterns  = [
    "**github.com**"
  ]
}
```

## Argument Reference

Arguments have a one to one mapping with the [JFrog API](https://www.jfrog.com/confluence/display/RTF/Repository+Configuration+JSON).
The following arguments are supported, along with the [common list of arguments for the remote repositories](remote.md):

* `key` - (Required) A mandatory identifier for the repository that must be unique. It cannot begin with a number or
  contain spaces or special characters.
* `description` - (Optional)
* `notes` - (Optional)
* `url` - (Required) The remote repo URL.
* `helm_charts_base_url` - (Optional) No documentation is available. Hopefully you know what this means.
* `external_dependencies_enabled` - (Optional) When set, external dependencies are rewritten. `External Dependency Rewrite` in the UI.
* `external_dependencies_patterns` - (Optional) An allow list of Ant-style path patterns that determine which remote VCS roots Artifactory will
  follow to download remote modules from, when presented with 'go-import' meta tags in the remote repository response.
  Default value in the UI is empty. This attribute must be set together with `external_dependencies_enabled = true`.

## Import

Remote repositories can be imported using their name, e.g.
```
$ terraform import artifactory_remote_helm_repository.helm-remote helm-remote
```
