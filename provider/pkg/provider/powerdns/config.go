package powerdns

import (
	"os"
	"strconv"

	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

type Config struct {
	Version			string   `pulumi:"version"`
	ApiEndpoint     string   `pulumi:"apiEndpoint"`
	ApiKey 			string   `pulumi:"apiKey" provider:"secret"`
	Insecure 	   	*bool 	 `pulumi:"insecure,optional"`
	Logging  	   	*bool    `pulumi:"logging,optional"`
	Client         	*PDNSClient
}

var _ = (infer.Annotated)((*Config)(nil))

func (c *Config) Annotate(a infer.Annotator) {
	a.Describe(&c.ApiEndpoint, "The api endpoint of the powerdns server")
	a.Describe(&c.ApiKey, "The access key for API operations")
	a.Describe(&c.Insecure, `Explicitly allow the provider to perform "insecure" SSL requests. If omitted, default value is "false"`)
	
	//a.SetDefault(&c.Insecure, true, "API_INSECURE")
}

var _ = (infer.CustomCheck[*Config])((*Config)(nil))

// workaround for https://github.com/pulumi/pulumi-go-provider/issues/110
func (c *Config) Check(ctx p.Context, name string, oldInputs, newInputs resource.PropertyMap) (*Config, []p.CheckFailure, error) {
	c.ApiEndpoint = newInputs["apiEndpoint"].StringValue()
	c.ApiKey = newInputs["ApiKey"].StringValue()
	if newInputs["insecure"].IsBool() {
		insecure := newInputs["insecure"].BoolValue()
		c.Insecure = &insecure
	}
	if newInputs["logging"].IsBool() {
		logging := newInputs["logging"].BoolValue()
		c.Logging = &logging
	}

	return c, []p.CheckFailure{}, nil
}

var _ = (infer.CustomConfigure)((*Config)(nil))

func (c *Config) Configure(ctx p.Context) error {
	var insecure bool
	if c.Insecure == nil {
		insecureStr := os.Getenv("POWERDNS_INSECURE")
		if insecureStr == "" {
			insecure = true
		} else {
			insecure, _ = strconv.ParseBool(insecureStr)
		}
	} else {
		insecure = *c.Insecure
	}

	if insecure {
		client, _ := NewClient(c.ApiEndpoint, c.ApiKey, nil, true, "2", 10)
		c.Client = client
	} else {
		client, _ := NewClient(c.ApiEndpoint, c.ApiKey, nil, true, "2", 10)
		c.Client = client
	}

	return nil
}