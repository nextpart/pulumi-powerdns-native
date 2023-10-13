package powerdns

import (
	"fmt"
	"strings"

	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi/sdk/v3/go/common/diag"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

type Record struct{}

type RecordArgs struct {
	// Fields projected into Pulumi must be public and hava a `pulumi:"..."` tag.
	// The pulumi tag doesn't need to match the field name, but its generally a
	// good idea.
	Zone    string   `pulumi:"zone"`
	Name    string   `pulumi:"name"`
	Type    string   `pulumi:"type"`
	TTL     int      `pulumi:"ttl"`
	Records []string `pulumi:"records"`
	SetPtr  *bool    `pulumi:"setptr,optional"`
}

// Each resource has a state, describing the fields that exist on the created resource.
type RecordState struct {
	RecId string `pulumi:"recId"`
	// It is generally a good idea to embed args in outputs, but it isn't strictly necessary.
	RecordArgs
	// Here we define a required output called result.
	Result string `pulumi:"result"`
}

// Implementing Annotate lets you provide descriptions and default values for resources and they will
// be visible in the provider's schema and the generated SDKs.
func (r *Record) Annotate(a infer.Annotator) {
	a.Describe(&r, "Test")
}

// The following statement is not required. It is a type assertion to indicate to Go that Command
// implements the following interfaces. If the function signature doesn't match or isn't implemented,
// we get nice compile time errors at this location.

var _ = (infer.CustomRead[RecordArgs, RecordState])((*Record)(nil))
var _ = (infer.CustomDelete[RecordState])((*Record)(nil))
var _ = (infer.CustomCheck[RecordArgs])((*Record)(nil))

func (r *Record) Create(ctx p.Context, name string, input RecordArgs, preview bool) (string, RecordState, error) {
	c := infer.GetConfig[Config](ctx)
	state := RecordState{RecordArgs: input}

	rrSet := ResourceRecordSet{
		Name: input.Name,
		Type: input.Type,
		TTL:  input.TTL,
	}

	zone := input.Zone
	ttl := input.TTL
	recs := input.Records
	setPtr := false

	if input.SetPtr != nil {
		setPtr = *input.SetPtr
	}

	if len(recs) > 0 {
		records := make([]RecordApi, 0, len(recs))
		for _, recContent := range recs {
			records = append(records,
				RecordApi{
					Name:    rrSet.Name,
					Type:    rrSet.Type,
					TTL:     ttl,
					Content: recContent,
					SetPtr:  setPtr,
				})
		}

		rrSet.Records = records

		ctx.Log(diag.Debug, fmt.Sprintf("Creating PowerDNS Record: %#v", rrSet))

		if preview {
			return name, state, nil
		}

		recID, err := c.Client.ReplaceRecordSet(zone, rrSet)
		if err != nil {
			return name, state, fmt.Errorf("Failed to create PowerDNS Record: %s", err)
		}

		state.RecId = recID
		ctx.Log(diag.Info, fmt.Sprintf("Created PowerDNS Record with ID: %s", recID))
		state.Result = fmt.Sprintf("Created PowerDNS Record with ID: %s", recID)
	}
	return state.RecId, state, nil
}

func (r *Record) Read(ctx p.Context, id string, input RecordArgs, state RecordState) (string, RecordArgs, RecordState, error) {
	c := infer.GetConfig[Config](ctx)

	ctx.Log(diag.Debug, fmt.Sprintf("Reading PowerDNS Record: %s", state.RecId))
	records, err := c.Client.ListRecordsByID(state.RecordArgs.Zone, state.RecId)
	if err != nil {
		return "", RecordArgs{}, RecordState{}, fmt.Errorf("Couldn't fetch PowerDNS Record: %s", err)
	}

	recs := make([]string, 0, len(records))
	for _, r := range records {
		recs = append(recs, r.Content)
	}
	input.Records = recs

	if len(records) > 0 {
		input.TTL = records[0].TTL
	}

	state.RecordArgs = input

	return state.RecId, input, state, nil
}

func (r *Record) Delete(ctx p.Context, id string, props RecordState) (err error) {
	c := infer.GetConfig[Config](ctx)

	ctx.Log(diag.Info, fmt.Sprintf("Deleting PowerDNS Record: %s", props.RecId))
	err = c.Client.DeleteRecordSetByID(props.RecordArgs.Zone, props.RecId)

	if err != nil && !strings.Contains(err.Error(), "Not Found") {
		return fmt.Errorf("Error deleting PowerDNS Record: %s", err)
	}

	return nil
}

func (r *Record) Check(ctx p.Context, name string, oldInputs resource.PropertyMap, newInputs resource.PropertyMap) (RecordArgs, []p.CheckFailure, error) {
	ctx.Log(diag.Debug,"Check func for records was called")
	if _, ok := newInputs["records"]; ok {
		if _, ok := newInputs["records"].V.([]resource.PropertyValue); ok {
			recs := newInputs["records"].V.([]resource.PropertyValue)
			for _, rec := range recs {
				if len(strings.Trim(rec.V.(string), " ")) == 0 {
					ctx.Log(diag.Warning,"One or more values in 'records' contain empty '' value(s)")
				}
			}
			if !(len(recs) > 0) {
				return RecordArgs{}, nil, fmt.Errorf("'records' must not be empty")
			}
		}
	}
	return infer.DefaultCheck[RecordArgs](newInputs)
}