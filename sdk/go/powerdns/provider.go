// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package powerdns

import (
	"context"
	"reflect"

	"errors"
	"github.com/nextpart/pulumi-powerdns-native/sdk/go/powerdns/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

type Provider struct {
	pulumi.ProviderResourceState

	// The api endpoint of the powerdns server
	ApiEndpoint pulumi.StringOutput `pulumi:"apiEndpoint"`
	// The access key for API operations
	ApiKey  pulumi.StringOutput `pulumi:"apiKey"`
	Version pulumi.StringOutput `pulumi:"version"`
}

// NewProvider registers a new resource with the given unique name, arguments, and options.
func NewProvider(ctx *pulumi.Context,
	name string, args *ProviderArgs, opts ...pulumi.ResourceOption) (*Provider, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiEndpoint == nil {
		return nil, errors.New("invalid value for required argument 'ApiEndpoint'")
	}
	if args.ApiKey == nil {
		return nil, errors.New("invalid value for required argument 'ApiKey'")
	}
	if args.Version == nil {
		return nil, errors.New("invalid value for required argument 'Version'")
	}
	if args.ApiKey != nil {
		args.ApiKey = pulumi.ToSecret(args.ApiKey).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"apiKey",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Provider
	err := ctx.RegisterResource("pulumi:providers:powerdns", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type providerArgs struct {
	// The api endpoint of the powerdns server
	ApiEndpoint string `pulumi:"apiEndpoint"`
	// The access key for API operations
	ApiKey string `pulumi:"apiKey"`
	// Explicitly allow the provider to perform "insecure" SSL requests. If omitted, default value is "false"
	Insecure *bool  `pulumi:"insecure"`
	Logging  *bool  `pulumi:"logging"`
	Version  string `pulumi:"version"`
}

// The set of arguments for constructing a Provider resource.
type ProviderArgs struct {
	// The api endpoint of the powerdns server
	ApiEndpoint pulumi.StringInput
	// The access key for API operations
	ApiKey pulumi.StringInput
	// Explicitly allow the provider to perform "insecure" SSL requests. If omitted, default value is "false"
	Insecure pulumi.BoolPtrInput
	Logging  pulumi.BoolPtrInput
	Version  pulumi.StringInput
}

func (ProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*providerArgs)(nil)).Elem()
}

type ProviderInput interface {
	pulumi.Input

	ToProviderOutput() ProviderOutput
	ToProviderOutputWithContext(ctx context.Context) ProviderOutput
}

func (*Provider) ElementType() reflect.Type {
	return reflect.TypeOf((**Provider)(nil)).Elem()
}

func (i *Provider) ToProviderOutput() ProviderOutput {
	return i.ToProviderOutputWithContext(context.Background())
}

func (i *Provider) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderOutput)
}

func (i *Provider) ToOutput(ctx context.Context) pulumix.Output[*Provider] {
	return pulumix.Output[*Provider]{
		OutputState: i.ToProviderOutputWithContext(ctx).OutputState,
	}
}

type ProviderOutput struct{ *pulumi.OutputState }

func (ProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Provider)(nil)).Elem()
}

func (o ProviderOutput) ToProviderOutput() ProviderOutput {
	return o
}

func (o ProviderOutput) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return o
}

func (o ProviderOutput) ToOutput(ctx context.Context) pulumix.Output[*Provider] {
	return pulumix.Output[*Provider]{
		OutputState: o.OutputState,
	}
}

// The api endpoint of the powerdns server
func (o ProviderOutput) ApiEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringOutput { return v.ApiEndpoint }).(pulumi.StringOutput)
}

// The access key for API operations
func (o ProviderOutput) ApiKey() pulumi.StringOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringOutput { return v.ApiKey }).(pulumi.StringOutput)
}

func (o ProviderOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringOutput { return v.Version }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderInput)(nil)).Elem(), &Provider{})
	pulumi.RegisterOutputType(ProviderOutput{})
}
