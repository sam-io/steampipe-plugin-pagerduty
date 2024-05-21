package pagerduty

import (
	"context"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tablePagerDutyIncidentCustomField(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "pagerduty_incident_custom_field",
		Description: "Custom fields for an incident",
		List: &plugin.ListConfig{
			Hydrate: listPagerDutyIncidentCustomFields,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "incident_id",
					Require: plugin.Required,
				},
			},
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Description: "An unique identifier of the log entry.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("id"),
			},
			{
				Name:        "incident_id",
				Description: "An unique identifier of the queried incident.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromQual("incident_id"),
			},
			{
				Name:        "name",
				Description: "The name of the field.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("name"),
			},
			{
				Name:        "display_name",
				Description: "The human-readable name of the field.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("display_name"),
			},
			{
				Name:        "field_type",
				Description: "The type of data this field contains.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("field_type"),
			},
			{
				Name:        "data_type",
				Description: "The kind of data the custom field is allowed to contain.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("data_type"),
			},
			{
				Name:        "description",
				Description: "A description of the data this field contains.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("description"),
			},
			{
				Name:        "value",
				Description: "Valuer of the field.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("value"),
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: "Title of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ID"),
			},
		},
	}
}

//// LIST FUNCTION

func listPagerDutyIncidentCustomFields(ctx context.Context, queryData *plugin.QueryData, hydrateData *plugin.HydrateData) (interface{}, error) {
	client, err := getSessionConfig(ctx, queryData)
	if err != nil {
		plugin.Logger(ctx).Error("pagerduty_incident_custom_fields.listPagerDutyIncidentCustomFields", "connection_error", err)
		return nil, err
	}

	incidentID := queryData.EqualsQuals["incident_id"].GetStringValue()

	plugin.Logger(ctx).Trace("pagerduty_incident_custom_fields.listPagerDutyIncidentCustomFields", incidentID)

	resp, err := client.GetIncidentCustomFields(ctx, incidentID)
	if err != nil {
		return nil, err
	}

	for _, sd := range (resp["custom_fields"]).([]interface{}) {
		queryData.StreamListItem(ctx, sd)
	}

	return nil, nil
}
