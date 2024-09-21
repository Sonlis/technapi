# technapi

Simple command line tool to synchronize ansible inventory hosts with technitium DNS records.

## Installation

Download the binary according to your operating system and CPU architecture from the release page, and set it
in the PATH.

## Usage

Export the necessary environment variables and run the binary:
```bash
./technapi
```

| Variable name | Description | Example |
| ------------- | ----------- | ------- |
| TECHNITIUM_URL | Url to reach the technitium instance, including the protocol scheme e.g. http and port | `http://localhost:5380` |
| TECHNITIUM_USER | User to login to technitium | `admin` |
| TECHNITIUM_PASSWORD | Password of the user used to login to technitium | `admin` |
| ANSIBLE_INV_PATH | Relative path, from where the script is executed, to the ansible inventory file | `inventory.yaml` |
| ZONE_CONF_PATH | Relative path, from where the script is executed, to the configuration of the zone holding the records | `zone-conf.yaml` |

## Integration with Ansible

To use with Ansible, use the following snippet in any playbooks: 
```yaml
- name: Update technitium records
  hosts: 127.0.0.1
  connection: local
  vars_files:
    - ./variables.yaml # Variable holding the encrypted password and other configuration variable
  tasks:
    - name: Update records
      ansible.builtin.shell: TECHNITIUM_URL={{ technitium_url }} TECHNITIUM_USER=admin TECHNITIUM_PASSWORD={{ technitium_password }} ANSIBLE_INV_PATH={{ ansible_inv_path }} ZONE_CONF_PATH={{ zone_conf_path }} <path/to/script>/technapi
```

To encrypt the password in the variable file, use ansible-vault:
```bash
echo <vault-password> password-file; ansible-vault encrypt_string --vault-password-file password-file <password> --name technitium_password; rm password-file
```

and set the output in the variable file.

The script supports one IP per host, as it is my use case. So it must follow the following format:
```yaml
oui:
  hosts:
    192.168.0.2:

non:
  hosts:
    192.168.0.12:
```

In this case, the script adds an A record for `oui` pointing to 192.168.0.2 and an A record for `non` pointing to 192.168.0.12
