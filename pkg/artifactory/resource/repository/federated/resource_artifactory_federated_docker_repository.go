package federated

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/jfrog/terraform-provider-artifactory/v6/pkg/artifactory/resource/repository"
	"github.com/jfrog/terraform-provider-artifactory/v6/pkg/artifactory/resource/repository/local"
	"github.com/jfrog/terraform-provider-shared/packer"
	"github.com/jfrog/terraform-provider-shared/predicate"
	"github.com/jfrog/terraform-provider-shared/util"
)

type DockerFederatedRepositoryParams struct {
	local.DockerLocalRepositoryParams
	Members []Member `hcl:"member" json:"members"`
}

func ResourceArtifactoryFederatedDockerV2Repository() *schema.Resource {
	packageType := "docker"

	dockerV2FederatedSchema := util.MergeMaps(
		local.DockerV2LocalSchema,
		memberSchema,
		repository.RepoLayoutRefSchema(rclass, packageType),
	)

	var unpackFederatedDockerRepository = func(data *schema.ResourceData) (interface{}, string, error) {
		repo := DockerFederatedRepositoryParams{
			DockerLocalRepositoryParams: local.UnpackLocalDockerV2Repository(data, rclass),
			Members:                     unpackMembers(data),
		}
		return repo, repo.Id(), nil
	}

	var packDockerMembers = func(repo interface{}, d *schema.ResourceData) error {
		members := repo.(*DockerFederatedRepositoryParams).Members
		return packMembers(members, d)
	}

	pkr := packer.Compose(
		packer.Universal(
			predicate.All(
				predicate.NoClass,
				predicate.Ignore("member", "terraform_type"),
			),
		),
		packDockerMembers,
	)

	constructor := func() (interface{}, error) {
		return &DockerFederatedRepositoryParams{
			DockerLocalRepositoryParams: local.DockerLocalRepositoryParams{
				RepositoryBaseParams: local.RepositoryBaseParams{
					PackageType: packageType,
					Rclass:      rclass,
				},
			},
		}, nil
	}

	return repository.MkResourceSchema(dockerV2FederatedSchema, pkr, unpackFederatedDockerRepository, constructor)
}

func ResourceArtifactoryFederatedDockerV1Repository() *schema.Resource {
	packageType := "docker"

	dockerFederatedSchema := util.MergeMaps(
		local.DockerV1LocalSchema,
		memberSchema,
		repository.RepoLayoutRefSchema(rclass, packageType),
	)

	var unpackFederatedDockerRepository = func(data *schema.ResourceData) (interface{}, string, error) {
		repo := DockerFederatedRepositoryParams{
			DockerLocalRepositoryParams: local.UnpackLocalDockerV1Repository(data, rclass),
			Members:                     unpackMembers(data),
		}
		return repo, repo.Id(), nil
	}

	var packDockerMembers = func(repo interface{}, d *schema.ResourceData) error {
		members := repo.(*DockerFederatedRepositoryParams).Members
		return packMembers(members, d)
	}

	pkr := packer.Compose(
		packer.Universal(
			predicate.Ignore("class", "rclass", "member", "terraform_type"),
		),
		packDockerMembers,
	)

	constructor := func() (interface{}, error) {
		return &DockerFederatedRepositoryParams{
			DockerLocalRepositoryParams: local.DockerLocalRepositoryParams{
				RepositoryBaseParams: local.RepositoryBaseParams{
					PackageType: packageType,
					Rclass:      rclass,
				},
			},
		}, nil
	}

	return repository.MkResourceSchema(dockerFederatedSchema, pkr, unpackFederatedDockerRepository, constructor)
}
