import * as powerdns from "@pulumi/powerdns";

const zone = new powerdns.Zone("foobar", { name: "foobar.com.", kind: "master", account: "admin"});

const record = new powerdns.Record("test", { 
    zone: "foobar.com.", type: "A", name: "test.foobar.com.", ttl: 300, records : ["10.0.0.1", "10.0.0.2", "10.0.0.3", "10.0.0.4"]
})

const record2 = new powerdns.Record("foo", { 
    zone: "foobar.com.", type: "A", name: "foo.foobar.com.", ttl: 300, records : ["10.0.0.1", "", "10.0.0.3", "10.0.0.4"]
})