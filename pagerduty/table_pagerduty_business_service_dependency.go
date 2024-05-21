package pagerduty

import (
	"context"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tablePagerDutyBusinessServiceDependency(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "pagerduty_business_service_dependency",
		Description: "Get all immediate dependencies of any Business Service.",
		List: &plugin.ListConfig{
			Hydrate: listPagerDutyBusinessServiceDependencies,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "business_service_id",
					Require: plugin.Required,
				},
			},
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Description: "An unique identifier of the business service dependency.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("id"),
			},
			{
				Name:        "business_service_id",
				Description: "An unique identifier of the queried incident.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromQual("business_service_id"),
			},
			{
				Name:        "supporting_service",
				Description: "The reference to the service that supports the Business Service.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("supporting_service"),
			},
			{
				Name:        "dependent_service",
				Description: "The reference to the service that is dependent on the Business Service.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("dependent_service"),
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: "Title of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("id"),
			},
		},
	}
}

//// LIST FUNCTION

func listPagerDutyBusinessServiceDependencies(ctx context.Context, queryData *plugin.QueryData, hydrateData *plugin.HydrateData) (interface{}, error) {
	client, err := getSessionConfig(ctx, queryData)
	if err != nil {
		plugin.Logger(ctx).Error("pagerduty_business_service.hydrateBusinessServiceDependencies", "connection_error", err)
		return nil, err
	}

	businessServiceId := queryData.EqualsQuals["business_service_id"].GetStringValue()

	plugin.Logger(ctx).Trace("pagerduty_business_service.hydrateBusinessServiceDependencies", businessServiceId)

	resp, err := client.GetBusinessServiceDependencies(ctx, businessServiceId)
	if err != nil {
		return nil, err
	}

	for _, sd := range (resp["relationships"]).([]interface{}) {
		queryData.StreamListItem(ctx, sd)
	}

	return nil, nil
}
