---
title: PowerDNS Installation & Configuration
meta_desc: Information on how to install the PowerDNS provider.
layout: installation
---

## Installation

The PowerDNS provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@nextpart/powerdns`](https://www.npmjs.com/package/@nextpart/powerdns)
* Python: [`pulumi_powerdns`](https://pypi.org/project/pulumi-powerdns/)
* Go: [`github.com/nextpart/pulumi-powerdns-native/sdk/go/powerdns`](https://pkg.go.dev/github.com/nextpart/pulumi-powerdns-native/sdk)
* .NET: [`Pulumi.Powerdns`](https://www.nuget.org/packages/Pulumi.Powerdns)

## Setup

To provision resources with the Pulumi PowerDNS provider, you need to have correct url and key configuration.

## Configuration Options

Use `pulumi config set powerdns:<option> (--secret)`.

| Option | Environment Variable | Required/Optional | Description |
|-----|------|------|----|
| `url` | `POWERDNS_URL` | Required | URL of the PowerDNS Authoritative Server api endpoint  |
| `key`| `POWERDNS_KEY` | Required (Secret) | Api key for the PowerDNS Authoritative Server |
| `insecure` | `POWERDNS_INSECURE` | Optional | Allow insecure HTTPS client |
| `logging` | `POWERDNS_LOGGING` | Optional | Enable debug logging |