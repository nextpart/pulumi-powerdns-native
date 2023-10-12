package powerdns

import (
	"fmt"
	"net"
	"strconv"
	"strings"

	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi/sdk/v3/go/common/diag"
)

type Zone struct{}

type ZoneArgs struct {
	// Fields projected into Pulumi must be public and hava a `pulumi:"..."` tag.
	// The pulumi tag doesn't need to match the field name, but its generally a
	// good idea.
	Name        string   `pulumi:"name"`
	Kind        string   `pulumi:"kind"`
	Account     string   `pulumi:"account,optional"`
	Nameservers []string `pulumi:"nameservers,optional"`
	Masters     []string `pulumi:"masters,optional"`
	SoaEditAPI  string   `pulumi:"soaEditAPI,optional"`
}

// Each resource has a state, describing the fields that exist on the created resource.
type ZoneState struct {
	ZoneId string `pulumi:"zoneId"`
	// It is generally a good idea to embed args in outputs, but it isn't strictly necessary.
	ZoneArgs
	// Here we define a required output called result.
	Result string `pulumi:"result"`
}

// The following statement is not required. It is a type assertion to indicate to Go that Command
// implements the following interfaces. If the function signature doesn't match or isn't implemented,
// we get nice compile time errors at this location.

var _ = (infer.CustomRead[ZoneArgs, ZoneState])((*Zone)(nil))
var _ = (infer.CustomUpdate[ZoneArgs, ZoneState])((*Zone)(nil))
var _ = (infer.CustomDelete[ZoneState])((*Zone)(nil))

func (*Zone) Create(ctx p.Context, id string, input ZoneArgs, preview bool) (name string, state ZoneState, err error) {
	c := infer.GetConfig[Config](ctx)
	state = ZoneState{ZoneArgs: input}

	var nameservers []string
	for _, nameserver := range input.Nameservers {
		nameservers = append(nameservers, nameserver)
	}

	var masters []string
	for _, masterIPPort := range input.Masters {
		splitIPPort := strings.Split(masterIPPort, ":")
		// if there are more elements
		if len(splitIPPort) > 2 {
			return "", ZoneState{}, fmt.Errorf("more than one colon in <ip>:<port> string")
		}
		// when there are exactly 2 elements in list, assume second is port and check the port range
		if len(splitIPPort) == 2 {
			port, err := strconv.Atoi(splitIPPort[1])
			if err != nil {
				return "", ZoneState{}, fmt.Errorf("Error converting port value in masters atribute")
			}
			if port < 1 || port > 65535 {
				return "", ZoneState{}, fmt.Errorf("Invalid port value in masters atribute")
			}
		}
		// no matter if string contains just IP or IP:port pair, the first element in split list will be IP
		masterIP := splitIPPort[0]
		if net.ParseIP(masterIP) == nil {
			return "", ZoneState{}, fmt.Errorf("values in masters list attribute must be valid IPs")
		}
		masters = append(masters, masterIPPort)
	}

	zoneInfo := ZoneInfo{
		Name:        input.Name,
		Kind:        input.Kind,
		Account:     input.Account,
		Nameservers: nameservers,
		SoaEditAPI:  input.SoaEditAPI,
	}

	if len(masters) != 0 {
		if strings.EqualFold(zoneInfo.Kind, "Slave") {
			zoneInfo.Masters = masters
		} else {
			return "", ZoneState{}, fmt.Errorf("masters attribute is supported only for Slave kind")
		}
	}

	if preview {
		return name, state, nil
	}

	createdZoneInfo, err := c.Client.CreateZone(zoneInfo)
	if err != nil {
		return "", ZoneState{}, err
	}

	state.ZoneId = createdZoneInfo.ID
	state.Result = fmt.Sprintf("Created zone with id: %s", zoneInfo.ID)
	return state.ZoneId, state, nil
}

func (*Zone) Read(ctx p.Context, id string, input ZoneArgs, state ZoneState) (string, ZoneArgs, ZoneState, error) {
	ctx.Log(diag.Debug, fmt.Sprintf("Reading PowerDNS Zone: %s", state.ZoneId))
	c := infer.GetConfig[Config](ctx)
	zoneInfo, err := c.Client.GetZone(state.ZoneId)
	if err != nil {
		return "", ZoneArgs{}, ZoneState{}, fmt.Errorf("Couldn't fetch PowerDNS Zone: %s", err)
	}

	input.Name = zoneInfo.Name
	input.Kind = zoneInfo.Kind
	input.Account = zoneInfo.Account
	input.SoaEditAPI = zoneInfo.SoaEditAPI

	if zoneInfo.Kind != "Slave" {
		nameservers, err := c.Client.ListRecordsInRRSet(zoneInfo.Name, zoneInfo.Name, "NS")
		if err != nil {
			return "", ZoneArgs{}, ZoneState{}, fmt.Errorf("couldn't fetch zone %s nameservers from PowerDNS: %v", zoneInfo.Name, err)
		}

		var zoneNameservers []string
		for _, nameserver := range nameservers {
			zoneNameservers = append(zoneNameservers, nameserver.Content)
		}

		input.Nameservers = zoneNameservers
	}

	if strings.EqualFold(zoneInfo.Kind, "Slave") {
		input.Masters = zoneInfo.Masters
	}

	state.ZoneArgs = input

	return id, input, state, nil
}

func (*Zone) Update(ctx p.Context, id string, olds ZoneState, news ZoneArgs, preview bool) (ZoneState, error) {
	c := infer.GetConfig[Config](ctx)
	state := ZoneState{ZoneArgs: news}
	ctx.Log(diag.Debug, fmt.Sprintf("Updating PowerDNS Zone: %s", olds.ZoneArgs.Name))

	if preview {
		return state, nil
	}

	zoneInfo := ZoneInfoUpd{}
	if !isEqual(olds.ZoneArgs, state.ZoneArgs) {
		zoneInfo.Name = state.ZoneArgs.Name
		zoneInfo.Kind = state.ZoneArgs.Kind
		zoneInfo.Account = state.ZoneArgs.Account
		zoneInfo.SoaEditAPI = state.ZoneArgs.SoaEditAPI

		err := c.Client.UpdateZone(state.ZoneArgs.Name, zoneInfo)
		return state, err
	}
	return state, nil
}

func (*Zone) Delete(ctx p.Context, id string, props ZoneState) (err error) {
	c := infer.GetConfig[Config](ctx)

	ctx.Log(diag.Info, fmt.Sprintf("Deleting PowerDNS Zone: %s", props.ZoneArgs.Name))
	err = c.Client.DeleteZone(props.ZoneArgs.Name)

	if err != nil {
		return fmt.Errorf("Error deleting PowerDNS Zone: %s", err)
	}
	return nil
}

func isEqual(A, B interface{}) bool {
	// Find out the type of A & B is ZoneArgs or not
	if _, ok := A.(*ZoneArgs); ok {
		if _, ok := B.(*ZoneArgs); ok {
			if A.(*ZoneArgs).Kind == B.(*ZoneArgs).Kind && A.(*ZoneArgs).Account == B.(*ZoneArgs).Account && A.(*ZoneArgs).SoaEditAPI == B.(*ZoneArgs).SoaEditAPI {
				return true
			}
		}
	}
	return false
}
