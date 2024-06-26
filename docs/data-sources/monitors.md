---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "datadog_monitors Data Source - terraform-provider-datadog"
subcategory: ""
description: |-
  Use this data source to list several existing monitors for use in other resources.
---

# datadog_monitors (Data Source)

Use this data source to list several existing monitors for use in other resources.



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `monitor_tags_filter` (List of String) A list of monitor tags to limit the search. This filters on the tags set on the monitor itself.
- `name_filter` (String) A monitor name to limit the search.
- `tags_filter` (List of String) A list of tags to limit the search. This filters on the monitor scope.

### Read-Only

- `id` (String) The ID of this resource.
- `monitors` (List of Object) List of monitors (see [below for nested schema](#nestedatt--monitors))

<a id="nestedatt--monitors"></a>
### Nested Schema for `monitors`

Read-Only:

- `id` (Number)
- `name` (String)
- `type` (String)
