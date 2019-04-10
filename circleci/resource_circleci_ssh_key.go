package circleci

import (
	"fmt"
	circleci "github.com/edahlseng/terraform-provider-circleci/circleci/circleci-go"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceCircleCiSshKey() *schema.Resource {
	return &schema.Resource{
		Create: resourceCircleCiSshKeyCreate,
		Read:   resourceCircleCiSshKeyRead,
		Delete: resourceCircleCiSshKeyDelete,

		Schema: map[string]*schema.Schema{
			"project_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"hostname": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"private_key": {
				Type:      schema.TypeString,
				Required:  true,
				ForceNew:  true,
				Sensitive: true,
			},
			"fingerprint_md5": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceCircleCiSshKeyCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*circleci.Client)

	sshKeyInput := &circleci.SshKey{
		ProjectId:      d.Get("project_id").(string),
		Hostname:       d.Get("hostname").(string),
		PrivateKey:     d.Get("private_key").(string),
		FingerprintMd5: d.Get("fingerprint_md5").(string),
	}

	sshKey, _, err := client.SshKeys.Create(sshKeyInput)
	if err != nil {
		return fmt.Errorf("Error creating CircleCI SSH key: %s", err)
	}

	if sshKey == nil {
		return fmt.Errorf("Error creating CircleCI SSH key: SSH key is nil")
	}

	d.SetId(fmt.Sprintf("%s/%s/%s", sshKey.ProjectId, sshKey.Hostname, sshKey.FingerprintMd5))
	d.Set("project_id", sshKey.ProjectId)
	d.Set("hostname", sshKey.Hostname)
	d.Set("private_key", sshKey.PrivateKey)
	d.Set("fingerprint_md5", sshKey.FingerprintMd5)

	return nil
}

func resourceCircleCiSshKeyRead(d *schema.ResourceData, meta interface{}) error {
	// There's not much that we could do during read that's useful, as the CircleCI
	// API doesn't allow for reading of any SSH key information (though we could
	// maybe read something from the settings route, but it's not clear that would
	// have any benefit)

	return nil
}

func resourceCircleCiSshKeyDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*circleci.Client)

	sshKeyInput := &circleci.SshKey{
		ProjectId:      d.Get("project_id").(string),
		Hostname:       d.Get("hostname").(string),
		PrivateKey:     d.Get("private_key").(string),
		FingerprintMd5: d.Get("fingerprint_md5").(string),
	}

	_, err := client.SshKeys.Delete(sshKeyInput)
	if err != nil {
		return fmt.Errorf("Error deleting SSH key: %s", err)
	}

	return nil
}
