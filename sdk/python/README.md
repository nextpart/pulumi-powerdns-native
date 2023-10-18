# Pulumi PowerDNS Native Provider

This repository is a contains a pulumi native provider for the PowerDNS Authoritative Nameserver. This allow the creation of zones and records.

> **NOTE:** This provider is experimental. Not save for production use.

## Installing

This package is available for several languages/platforms:

### Node.js (JavaScript/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

```bash
npm install @nextpart/powerdns
```

or `yarn`:

```bash
yarn add @nextpart/powerdns
```

### Python

To use from Python, install using `pip`:

```bash
pip install pulumi-powerdns
```

### Go

To use from Go, use `go get` to grab the latest version of the library:

```bash
go get github.com/nextpart/pulumi-powerdns-native/sdk/go/...
```

### .NET

To use from .NET, install using `dotnet add package`:

```bash
dotnet add package Pulumi.Aci
```

## Example

### Node.js (JavaScript/TypeScript)

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as powerdns from "@nextpart/powerdns";

const zone = new powerdns.Zone("foobar", { name: "foobar.com.", kind: "master", account: "admin"});

const record = new powerdns.Record("test", { 
    zone: "foobar.com.", type: "A", name: "test.foobar.com.", ttl: 300, records : ["10.0.0.1", "10.0.0.2", "10.0.0.3", "10.0.0.4"]
})

const record2 = new powerdns.Record("foo", { 
    zone: "foobar.com.", type: "A", name: "foo.foobar.com.", ttl: 300, records : ["10.0.0.1", "10.0.0.3", "10.0.0.4"]
})
```
 
### Python

```python
import pulumi
import pulumi_powerdns as pdns

test = pdns.Zone("test", name="fifo.com.", kind="master", account="admin")

record = pdns.Record(
    "record",
    zone="fifo.com.",
    name="test.fifo.com.",
    type="A",
    records=["10.0.0.0", "10.0.0.1"],
    ttl=300,
)
```

### Go

```go
import (
	"fmt"
	pdns "github.com/nextpart/pulumi-powerdns-native/sdk/go/powerdns"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
        _, err := pdns.NewZone(ctx, "foobar.com", &pdns.ZoneArgs{
            Name:    pulumi.String("foobar.com."),
            Kind:    pulumi.String("master"),
            Account: pulumi.String("admin"),
        })

        if err != nil {
            return err
        }

		_, err = pdns.NewRecord(ctx, "test.foobar.com.", &pdns.RecordArgs{
			Name: pulumi.String("test.foobar.com."),
			Zone: pulumi.String("foobar.com."),
			Type: pulumi.String("A"),
			Records: pulumi.ToStringArray([]string{"10.0.0.1", "10.0.0.2", "10.0.3.0"}),
			Ttl: pulumi.Int(300),

		})
		return err
	})
}
```

### .NET

```csharp
using Pulumi;
using Pulumi.Powerdns;

class Powerdns : Stack
{
    public Powerdns()
    {
    }
}
```

### Configuration

The following configuration points are available for the `powerdns` provider:
* `powerdns:url` (environment: `POWERDNS_URL`) - URL of the powerdns api endpoint
* `powerdns:key` (environment: `POWERDNS_KEY`) - The API key for the powerdns api endpoint

## Development
### Prerequisites

Ensure the following tools are installed and present in your `$PATH`:

* [`pulumictl`](https://github.com/pulumi/pulumictl#installation)
* [Go 1.21](https://golang.org/dl/) or 1.latest
* [NodeJS](https://nodejs.org/en/) 14.x.  We recommend using [nvm](https://github.com/nvm-sh/nvm) to manage NodeJS installations.
* [Yarn](https://yarnpkg.com/)
* [TypeScript](https://www.typescriptlang.org/)
* [Python](https://www.python.org/downloads/) (called as `python3`).  For recent versions of MacOS, the system-installed version is fine.
* [.NET](https://dotnet.microsoft.com/download)

#### Build the provider and install the plugin

   ```bash
   $ make build install
   ```
   
This will:

1. Create the SDK codegen binary and place it in a `./bin` folder (gitignored)
2. Create the provider binary and place it in the `./bin` folder (gitignored)
3. Generate the dotnet, Go, Node, and Python SDKs and place them in the `./sdk` folder
4. Install the provider on your machine.

#### Test against the example
   
```bash
$ cd examples/simple
$ yarn link @nextpart/powerdns
$ yarn install
$ pulumi stack init test
$ pulumi up
```

Now that you have completed all of the above steps, you have a working provider that generates a random string for you.

#### A brief repository overview

You now have:

1. A `provider/` folder containing the building and implementation logic
    1. `cmd/pulumi-resource-powerdns/main.go` - holds the provider's sample implementation logic.
2. `deployment-templates` - a set of files to help you around deployment and publication
3. `sdk` - holds the generated code libraries created by `pulumi-gen-powerdns/main.go`
4. `examples` a folder of Pulumi programs to try locally and/or use in CI.
5. A `Makefile` and this `README`.

#### Additional Details

This repository depends on the pulumi-go-provider library. For more details on building providers, please check
the [Pulumi Go Provider docs](https://github.com/pulumi/pulumi-go-provider).

### Build Examples

Create an example program using the resources defined in your provider, and place it in the `examples/` folder.

You can now repeat the steps for [build, install, and test](#test-against-the-example).

## References

Other resources/examples for implementing providers:
* [Pulumi Command provider](https://github.com/pulumi/pulumi-command/blob/master/provider/pkg/provider/provider.go)
* [Pulumi Go Provider repository](https://github.com/pulumi/pulumi-go-provider)
