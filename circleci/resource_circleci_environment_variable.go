package circleci

import (
	"fmt"
	circleci "github.com/edahlseng/terraform-provider-circleci/circleci/circleci-go"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceCircleCiEnvironmentVariable() *schema.Resource {
	return &schema.Resource{
		Create: resourceCircleCiEnvironmentVariableCreate,
		Read:   resourceCircleCiEnvironmentVariableRead,
		Delete: resourceCircleCiEnvironmentVariableDelete,

		Schema: map[string]*schema.Schema{
			"project_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"value": {
				Type:      schema.TypeString,
				Required:  true,
				ForceNew:  true,
				Sensitive: true,
			},
		},
	}
}

func resourceCircleCiEnvironmentVariableCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*circleci.Client)

	environmentVariableInput := &circleci.EnvironmentVariable{
		ProjectId: d.Get("project_id").(string),
		Name:      d.Get("name").(string),
		Value:     d.Get("value").(string),
	}

	environmentVariable, _, err := client.EnvironmentVariables.Create(environmentVariableInput)
	if err != nil {
		return fmt.Errorf("Error creating CircleCI environment variable: %s", err)
	}

	if environmentVariable == nil {
		return fmt.Errorf("Error creating CircleCI environment variable: environment variable is nil")
	}

	d.SetId(fmt.Sprintf("%s/%s", environmentVariable.ProjectId, environmentVariable.Name))
	d.Set("project_id", environmentVariable.ProjectId)
	d.Set("name", environmentVariable.Name)
	d.Set("value", environmentVariable.Value)
	d.Set("value_masked", environmentVariable.ValueMasked)

	return nil
}

func resourceCircleCiEnvironmentVariableRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*circleci.Client)

	environmentVariableInput := &circleci.EnvironmentVariable{
		ProjectId: d.Get("project_id").(string),
		Name:      d.Get("name").(string),
	}

	environmentVariable, _, err := client.EnvironmentVariables.Read(environmentVariableInput)
	if err != nil {
		return fmt.Errorf("Error reading CircleCI environment variable: %s", err)
	}

	if environmentVariable == nil {
		return fmt.Errorf("Error creating CircleCI environment variable: environment variable is nil")
	}

	d.Set("name", environmentVariable.Name)
	d.Set("value_masked", environmentVariable.ValueMasked)

	return nil
}

func resourceCircleCiEnvironmentVariableDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*circleci.Client)

	environmentVariableInput := &circleci.EnvironmentVariable{
		ProjectId:   d.Get("project_id").(string),
		Name:        d.Get("name").(string),
		Value:       "",
		ValueMasked: "",
	}

	_, err := client.EnvironmentVariables.Delete(environmentVariableInput)
	if err != nil {
		return fmt.Errorf("Error deleting environment variable: %s", err)
	}

	return nil
}
