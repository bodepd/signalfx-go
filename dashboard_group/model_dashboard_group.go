/*
 * Dashboard Groups API
 *
 * Dashboard groups let you collect dashboards with common characteristics in one place in the web UI, so you can view them together or in sequence. ## Dashboard group membership A dashboard can belong to only one dashboard group, but a group can contain multiple dashboards. However, user default dashboard groups can only contain the user's default dashboard. <br> When you create a new dashboard group, the system generates a new dashboard for it. Similarly, when you create a new dashboard, the system creates a new dashboard group for it. Because a dashboard can only belong to one group, you can't add existing dashboards to a new group (the dashboards already belong to a group). To add an existing dashboard to a new group, create the group and then change the dashboard group of the dashboards to the new group. <br> You can add a new dashboard to an existing group at any time. ## Cloning dashboards into different groups You can also clone existing dashboards into another group. Use this feature to test dashboards in your user dashboard group before cloning them into a production group. You can also use this feature to customize an existing dashboard, by cloning it into your user group and then modifying it. ## Deleting dashboard groups When you delete a dashboard group, the system deletes the dashboards in the group and the charts that those dashboards contain.<br> **Note:** The system doesn't do this for dashboards. When you delete a dashboard, the system orphans its charts, but it doesn't delete them. <br> You can assign a dashboard group to one or more teams in your organization. The groups then appear on the team's landing page in the web UI. ## Viewing a dashboard group To view a dashboard group you create using the API, navigate to the following URL:<br> `https://app.<REALM>.signalfx.com/#/dashboard/<GROUP_ID>` <br> For `<GROUP_ID>`, substitute the value of the dashboard group ID. In the web UI, you see the dashboard group name, and underneath it the dashboard names displayed as tabs. ## Dashboard group authorizations By default, all users can edit or delete dashboard groups. If your organization has the \"write permissions\" feature enabled, your administrator can limit editing and deletion of specific dashboard groups to individual users or teams or both. This feature helps prevent unauthorized or accidental modifications to dashboard groups. Administrators can always modify write permissions, even for dashboard groups which they don't have permission to edit. This lets administrators escalate their access for any dashboard group. When a user deletes a dashboard group, the system deletes the group's dashboards without regard to the *dashboard* permissions. Only the dashgroup group permissions are considered.
 *
 * API version: 3.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package dashboard_group

type DashboardGroup struct {
	AuthorizedWriters *AuthorizedWriters `json:"authorizedWriters,omitempty"`
	// The dashboard group creation date and time, in the form of a Unix time value (milliseconds since the Unix epoch 1970-01-01 00:00:00 UTC+0). The system sets this value, and you can't modify it.
	Created int64 `json:"created,omitempty"`
	// SignalFx-assigned user ID of the user that created the dashboard group. If the system created this dashboard group, the value is \"AAAAAAAAAA\". The system sets this value, and you can't modify it.
	Creator          string             `json:"creator,omitempty"`
	DashboardConfigs []*DashboardConfig `json:"dashboardConfigs,omitempty"`
	// Array of dashboard IDs. The system adds the specified dashboards to the dashboard group you're creating. If you omit the property, the system creates a new dashboard and assigns it to the new dashboard group.
	Dashboards []string `json:"dashboards,omitempty"`
	// Description of the dashboard group. This value appears in the tooltip for the dashboard group on the Dashboards page in the web UI.
	Description string `json:"description,omitempty"`
	// The SignalFx-assigned ID for this dashboard.
	Id string `json:"id,omitempty"`

	ImportQualifiers []*ImportQualifier `json:"importQualifiers,omitempty"`
	// The last time the dashboard group was updated, in the form of a Unix timestamp (milliseconds since the Unix epoch 1970-01-01 00:00:00 UTC+0) This value is \"read-only\".
	LastUpdated int64 `json:"lastUpdated,omitempty"`
	// SignalFx-assigned ID of the last user who updated the dashboard group. If the last update was by the system, the value is \"AAAAAAAAAA\". This value is \"read-only\".
	LastUpdatedBy string `json:"lastUpdatedBy,omitempty"`
	// Name of the dashboard group. This value identifies the dashboard group in the web UI. It appears on the dashboards page and in the catalog. It also appears at the top left corner of the screen whenever you're viewing a dashboard that the group contains.
	Name string `json:"name"`
	// Array of existing team object ID. The dashboard group appears on the team landing page for any team in the list.
	Teams []string `json:"teams,omitempty"`
}
