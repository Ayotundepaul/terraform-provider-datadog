---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "datadog_csm_threats_agent_rules Data Source - terraform-provider-datadog"
subcategory: ""
description: |-
  Use this data source to retrieve information about existing Agent rules.
---

# datadog_csm_threats_agent_rules (Data Source)

Use this data source to retrieve information about existing Agent rules.



<!-- schema generated by tfplugindocs -->
## Schema

### Read-Only

- `agent_rules` (List of Object) List of Agent rules (see [below for nested schema](#nestedatt--agent_rules))
- `agent_rules_ids` (List of String) List of IDs for the Agent rules.
- `id` (String) The ID of this resource.

<a id="nestedatt--agent_rules"></a>
### Nested Schema for `agent_rules`

Read-Only:

- `description` (String)
- `enabled` (Boolean)
- `expression` (String)
- `id` (String)
- `name` (String)
