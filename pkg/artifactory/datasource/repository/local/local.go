package local

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/jfrog/terraform-provider-artifactory/v6/pkg/artifactory/resource/repository"
	"github.com/jfrog/terraform-provider-shared/packer"
)

func MkRepoReadDataSource(pack packer.PackFunc, construct repository.Constructor) schema.ReadContextFunc {
	return func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
		repo, err := construct()
		if err != nil {
			return diag.FromErr(err)
		}

		key := d.Get("key").(string)
		// repo must be a pointer
		_, err = m.(*resty.Client).R().
			SetResult(repo).
			SetPathParam("key", key).
			Get(repository.RepositoriesEndpoint)

		if err != nil {
			return diag.FromErr(err)
		}

		d.SetId(key)

		return diag.FromErr(pack(repo, d))
	}
}
