package circleci

import (
	"fmt"
	circleci "github.com/edahlseng/terraform-provider-circleci/circleci/circleci-go"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceCircleCiProject() *schema.Resource {
	return &schema.Resource{
		Create: resourceCircleCiProjectCreate,
		Read:   resourceCircleCiProjectRead,
		Delete: resourceCircleCiProjectDelete,

		Schema: map[string]*schema.Schema{
			"vcs_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"username": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceCircleCiProjectCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*circleci.Client)

	projectInput := &circleci.Project{
		VcsType:  d.Get("vcs_type").(string),
		Username: d.Get("username").(string),
		Name:     d.Get("name").(string),
	}

	project, _, err := client.Projects.Create(projectInput)
	if err != nil {
		return fmt.Errorf("Error creating CircleCI project: %s", err)
	}

	if project == nil {
		return fmt.Errorf("Error creating CircleCI project: project is nil")
	}

	d.SetId(circleci.ProjectIdFromProject(*project))
	d.Set("vcs_type", project.VcsType)
	d.Set("username", project.Username)
	d.Set("name", project.Name)

	return nil
}

func resourceCircleCiProjectRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*circleci.Client)

	project, _, err := client.Projects.Read(d.Id())
	if err != nil {
		return fmt.Errorf("Error reading CircleCI project: %s", err)
	}

	if project == nil {
		return fmt.Errorf("Error creating CircleCI project: project is nil")
	}

	d.Set("vcs_type", project.VcsType)
	d.Set("username", project.Username)
	d.Set("name", project.Name)

	return nil
}

func resourceCircleCiProjectDelete(d *schema.ResourceData, meta interface{}) error {
	// The delete step does nothing, as it's not possible to stop building a CircleCI
	// project via the API.

	return nil
}
