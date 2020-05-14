package circleci

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
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
			"base_url": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("CIRCLECI_BASE_URL", "https://circleci.com/api/v1.1/"),
				Description: "Base URL for a CircleCI API, v1.1",
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
		BaseURL:   d.Get("base_url").(string),
	}

	client := config.NewClient()

	return client, nil
}
