// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ConfigurationsColumns holds the columns for the "configurations" table.
	ConfigurationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "key", Type: field.TypeString},
		{Name: "value", Type: field.TypeString},
	}
	// ConfigurationsTable holds the schema information for the "configurations" table.
	ConfigurationsTable = &schema.Table{
		Name:       "configurations",
		Columns:    ConfigurationsColumns,
		PrimaryKey: []*schema.Column{ConfigurationsColumns[0]},
	}
	// EvaluationsColumns holds the columns for the "evaluations" table.
	EvaluationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "external_id", Type: field.TypeString, Nullable: true},
		{Name: "date", Type: field.TypeTime, Nullable: true},
		{Name: "evaluation_result", Type: field.TypeString},
		{Name: "evaluation_prompt_id", Type: field.TypeUUID},
		{Name: "response_id", Type: field.TypeUUID},
		{Name: "user_id", Type: field.TypeUUID},
	}
	// EvaluationsTable holds the schema information for the "evaluations" table.
	EvaluationsTable = &schema.Table{
		Name:       "evaluations",
		Columns:    EvaluationsColumns,
		PrimaryKey: []*schema.Column{EvaluationsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "evaluations_evaluation_prompts_evaluations",
				Columns:    []*schema.Column{EvaluationsColumns[4]},
				RefColumns: []*schema.Column{EvaluationPromptsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "evaluations_responses_evaluations",
				Columns:    []*schema.Column{EvaluationsColumns[5]},
				RefColumns: []*schema.Column{ResponsesColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "evaluations_users_evaluations",
				Columns:    []*schema.Column{EvaluationsColumns[6]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// EvaluationPromptsColumns holds the columns for the "evaluation_prompts" table.
	EvaluationPromptsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "prompt", Type: field.TypeString},
	}
	// EvaluationPromptsTable holds the schema information for the "evaluation_prompts" table.
	EvaluationPromptsTable = &schema.Table{
		Name:       "evaluation_prompts",
		Columns:    EvaluationPromptsColumns,
		PrimaryKey: []*schema.Column{EvaluationPromptsColumns[0]},
	}
	// RequestsColumns holds the columns for the "requests" table.
	RequestsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "external_id", Type: field.TypeString, Nullable: true},
		{Name: "from", Type: field.TypeString, Nullable: true},
		{Name: "subject", Type: field.TypeString, Nullable: true},
		{Name: "body", Type: field.TypeString},
		{Name: "date", Type: field.TypeTime, Nullable: true},
	}
	// RequestsTable holds the schema information for the "requests" table.
	RequestsTable = &schema.Table{
		Name:       "requests",
		Columns:    RequestsColumns,
		PrimaryKey: []*schema.Column{RequestsColumns[0]},
	}
	// ResponsesColumns holds the columns for the "responses" table.
	ResponsesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "external_id", Type: field.TypeString, Nullable: true},
		{Name: "from", Type: field.TypeString, Nullable: true},
		{Name: "subject", Type: field.TypeString, Nullable: true},
		{Name: "body", Type: field.TypeString},
		{Name: "date", Type: field.TypeTime, Nullable: true},
		{Name: "request_id", Type: field.TypeUUID},
		{Name: "scenario_id", Type: field.TypeUUID},
	}
	// ResponsesTable holds the schema information for the "responses" table.
	ResponsesTable = &schema.Table{
		Name:       "responses",
		Columns:    ResponsesColumns,
		PrimaryKey: []*schema.Column{ResponsesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "responses_requests_responses",
				Columns:    []*schema.Column{ResponsesColumns[6]},
				RefColumns: []*schema.Column{RequestsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "responses_scenarios_responses",
				Columns:    []*schema.Column{ResponsesColumns[7]},
				RefColumns: []*schema.Column{ScenariosColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// ScenariosColumns holds the columns for the "scenarios" table.
	ScenariosColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeString},
		{Name: "external_id", Type: field.TypeString, Nullable: true},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "systemprompt", Type: field.TypeString, Nullable: true},
		{Name: "date", Type: field.TypeTime, Nullable: true},
	}
	// ScenariosTable holds the schema information for the "scenarios" table.
	ScenariosTable = &schema.Table{
		Name:       "scenarios",
		Columns:    ScenariosColumns,
		PrimaryKey: []*schema.Column{ScenariosColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeString},
		{Name: "external_id", Type: field.TypeString, Unique: true},
		{Name: "type", Type: field.TypeString, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ConfigurationsTable,
		EvaluationsTable,
		EvaluationPromptsTable,
		RequestsTable,
		ResponsesTable,
		ScenariosTable,
		UsersTable,
	}
)

func init() {
	EvaluationsTable.ForeignKeys[0].RefTable = EvaluationPromptsTable
	EvaluationsTable.ForeignKeys[1].RefTable = ResponsesTable
	EvaluationsTable.ForeignKeys[2].RefTable = UsersTable
	ResponsesTable.ForeignKeys[0].RefTable = RequestsTable
	ResponsesTable.ForeignKeys[1].RefTable = ScenariosTable
}
