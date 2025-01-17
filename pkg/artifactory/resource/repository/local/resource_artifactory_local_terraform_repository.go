package local

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/jfrog/terraform-provider-artifactory/v6/pkg/artifactory/resource/repository"
	"github.com/jfrog/terraform-provider-shared/packer"
	"github.com/jfrog/terraform-provider-shared/util"
)

func GetTerraformLocalSchema(registryType string) map[string]*schema.Schema {
	return util.MergeMaps(
		BaseLocalRepoSchema,
		repository.RepoLayoutRefSchema("local", "terraform_"+registryType),
	)
}

func UnpackLocalTerraformRepository(data *schema.ResourceData, rclass string, registryType string) RepositoryBaseParams {
	repo := UnpackBaseRepo(rclass, data, "terraform_"+registryType)
	repo.TerraformType = registryType

	return repo
}

func ResourceArtifactoryLocalTerraformRepository(registryType string) *schema.Resource {

	var unpackLocalTerraformRepository = func(data *schema.ResourceData) (interface{}, string, error) {
		repo := UnpackLocalTerraformRepository(data, rclass, registryType)
		return repo, repo.Id(), nil
	}

	terraformLocalSchema := GetTerraformLocalSchema(registryType)

	constructor := func() (interface{}, error) {
		return &RepositoryBaseParams{
			PackageType: "terraform_" + registryType,
			Rclass:      rclass,
		}, nil
	}

	return repository.MkResourceSchema(
		terraformLocalSchema,
		packer.Default(terraformLocalSchema),
		unpackLocalTerraformRepository,
		constructor,
	)
}
