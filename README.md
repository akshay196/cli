# rcloud-cli

# Currently supported
- clusters
  - create cluster of type import
      Using command(s): 
        rctl create cluster imported sample-imported-cluster -l sample-location
      Using file apply: 
        rctl apply -f cluster-config.yml
  - list clusters
      Using command(s): 
        rctl get cluster
        rctl get cluster sample-imported-cluster
  - download bootstrap (separate command)
      Using command(s): 
        rctl kubeconfig download --cluster sample-imported-cluster
- project
  - create project with basic information
      Using command(s): 
        rctl create p project1 --desc "Description of the project"
      Using file apply: 
        rctl apply -f project-config.yml
  - list projects
      Using command(s): 
        rctl get project
        rctl get project project1
- user
  - create user
      Using command(s):
        rctl create user john.doe@example.com
        rctl create user john.doe@example.com --console John, Doe
        rctl create user john.doe@example.com  --groups testingGroup, productionGroup --console John, Doe, 4089382091
      Using file apply:
        rctl apply -f user-config.yml
  - list users
      Using command(s):
        rctl get users
        rctl get user john.dow@example.com
- group
  - create group
      Using command(s):
        rctl create group sample-group --desc "Description of the group"
      Using file apply:
        rctl apply -f group-config.yml
  - list groups
      Using command(s):
        rctl get group
        rctl get group sample-group
- oidc
  - create oidc
      Using command(s):
        rctl create oidc github 721396hsad8721wjhad8 http://myownweburl.com/cb
      Using file apply:
        rctl apply -f oidc-config.yml
  - list oidc providers
      Using command(s):
        rctl get oidc
        rctl get oidc github
- groupassociation
  - update group association to projects and users
    Using command(s):
      rctl update groupassociation sample-group --associateproject sample-proj --roles PROJECT_READ_ONLY,INFRA_ADMIN
      rctl update groupassociation sample-group  --associateuser y --addusers example.user@company.co,example.user-two@company.co --removeusers example.user-three@company.co
Global Parameters
  -c, --config string    Customize cli config file
  -d, --debug            Enable debug logs
  -f, --file string      provide file with resource to be created
  -o, --output string    Print json, yaml or table output. Default is table (default "table")
  -p, --project string   provide a specific project context
  -v, --verbose          Verbose mode. A lot more information output.