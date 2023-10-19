---
title: PowerDNS
meta_desc: Provides an overview of the PowerDNS Authoritative Provider for Pulumi.
layout: overview
---

## PowerDNS Authoritative Provider

The PowerDNS provider is used manipulate DNS records supported by PowerDNS server. The provider needs to be configured
with the proper credentials before it can be used. It supports both the [legacy API](https://doc.powerdns.com/3/httpapi/api_spec/) and the new [version 1 API](https://doc.powerdns.com/md/httpapi/api_spec/), however resources may need to be configured differently.

Use the navigation to the left to read about the available resources or take look at the examples below.

## Example

{{< chooser language "typescript,python,go,csharp" >}}
{{% choosable language typescript %}}

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
 
{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi
import pulumi_powerdns as pdns

test = pdns.Zone("test", name="foobar.com.", kind="master", account="admin")

record = pdns.Record(
    "record",
    zone="foobar.com.",
    name="test.foobar.com.",
    type="A",
    records=["10.0.0.0", "10.0.0.1"],
    ttl=300,
)
```

{{% /choosable %}}
{{% choosable language go %}}

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

{{% /choosable %}}
{{% choosable language csharp %}}

```csharp
using System.Collections.Generic;
using Pulumi;
using Pulumi.Powerdns;

return await Deployment.RunAsync(() =>
{
   // Add your resources here
   // e.g. var resource = new Resource("name", new ResourceArgs { });
   var zone = new Zone("foobar", new ZoneArgs {
      Name = "foobar.com.",
      Kind = "master",
      Account = "admin",
   });

   var record = new Record("record", new RecordArgs {
      Name = "test.foobar.com.",
      Zone = zone.Name,
      Type = "A",
      Records = new List<string>(){"10.0.0.0", "10.0.0.2"},
      Ttl = 300,
   });

   // Export outputs here
   return new Dictionary<string, object?>
   {
      ["outputKey"] = "outputValue"
   };
});
```

{{% /choosable %}}

{{< /chooser >}}