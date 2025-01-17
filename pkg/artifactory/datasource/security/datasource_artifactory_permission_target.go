package security

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/jfrog/terraform-provider-artifactory/v6/pkg/artifactory/resource/security"
)

func DataSourceArtifactoryPermissionTarget() *schema.Resource {
	dataSourcePermissionTargetRead := func(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
		permissionTarget := new(security.PermissionTargetParams)
		targetName := d.Get("name").(string)
		_, err := m.(*resty.Client).R().SetResult(permissionTarget).Get(security.PermissionsEndPoint + targetName)

		if err != nil {
			return diag.FromErr(err)
		}

		d.SetId(permissionTarget.Name)

		return security.PackPermissionTarget(permissionTarget, d)
	}
	return &schema.Resource{
		ReadContext: dataSourcePermissionTargetRead,
		Schema:      security.BuildPermissionTargetSchema(),
		Description: "Provides the permission target data source. Contains information about a specific permission target.",
	}
}
