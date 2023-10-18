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
	Url     		string   `pulumi:"url"`
	Key 			string   `pulumi:"key" provider:"secret"`
	Insecure 	   	*bool 	 `pulumi:"insecure,optional"`
	Logging  	   	*bool    `pulumi:"logging,optional"`
	Client         	*PDNSClient
}

var _ = (infer.Annotated)((*Config)(nil))

func (c *Config) Annotate(a infer.Annotator) {
	a.Describe(&c.Url, "The api endpoint of the powerdns server")
	a.Describe(&c.Key, "The access key for API operations")
	a.Describe(&c.Insecure, `Explicitly allow the provider to perform "insecure" SSL requests. If omitted, default value is "false"`)
	a.SetDefault(&c.Url, "", "POWERDNS_URL")
	a.SetDefault(&c.Key, "", "POWERDNS_KEY")
	//a.SetDefault(&c.Insecure, true, "API_INSECURE")
}

var _ = (infer.CustomCheck[*Config])((*Config)(nil))

// workaround for https://github.com/pulumi/pulumi-go-provider/issues/110
func (c *Config) Check(ctx p.Context, name string, oldInputs, newInputs resource.PropertyMap) (*Config, []p.CheckFailure, error) {
	failures := []p.CheckFailure{}
	if newInputs["url"].IsString() {
		c.Url = newInputs["url"].StringValue()
	} else {
		failures = append(failures, p.CheckFailure{
			Property: "url",
			Reason:  "Field does not exist or is empty",
		})
	}

	if newInputs["key"].IsString() {
		c.Key = newInputs["key"].StringValue()
	} else {
		failures = append(failures, p.CheckFailure{
			Property: "key",
			Reason:  "Field does not exist or is empty",
		})
	}

	if newInputs["insecure"].IsBool() {
		insecure := newInputs["insecure"].BoolValue()
		c.Insecure = &insecure
	}
	if newInputs["logging"].IsBool() {
		logging := newInputs["logging"].BoolValue()
		c.Logging = &logging
	}

	return c, failures, nil
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

	// TODO: Create insecure tls context when filed is enabled
	if insecure {
		client, _ := NewClient(c.Url, c.Key, nil, true, "2", 10)
		c.Client = client
	} else {
		client, _ := NewClient(c.Url, c.Key, nil, true, "2", 10)
		c.Client = client
	}

	return nil
}