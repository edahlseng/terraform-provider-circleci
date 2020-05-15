package circleci

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"token": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("CIRCLECI_TOKEN", nil),
				Description: "Authentication token for a CircleCI user",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"circleci_project":              resourceCircleCiProject(),
			"circleci_ssh_key":              resourceCircleCiSshKey(),
			"circleci_environment_variable": resourceCircleCiEnvironmentVariable(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		AuthToken: d.Get("token").(string),
	}

	client := config.NewClient()

	return client, nil
}
