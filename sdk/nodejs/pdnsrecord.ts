// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class PDNSRecord extends pulumi.CustomResource {
    /**
     * Get an existing PDNSRecord resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): PDNSRecord {
        return new PDNSRecord(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'powerdns:index:PDNSRecord';

    /**
     * Returns true if the given object is an instance of PDNSRecord.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PDNSRecord {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PDNSRecord.__pulumiType;
    }

    public readonly name!: pulumi.Output<string>;
    public readonly records!: pulumi.Output<string[]>;
    /**
     * For A and AAAA records, if true, create corresponding PTR.
     */
    public readonly set_ptr!: pulumi.Output<boolean | undefined>;
    public readonly ttl!: pulumi.Output<number>;
    public readonly type!: pulumi.Output<string>;
    public readonly zone!: pulumi.Output<string>;

    /**
     * Create a PDNSRecord resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PDNSRecordArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.name === undefined) && !opts.urn) {
                throw new Error("Missing required property 'name'");
            }
            if ((!args || args.records === undefined) && !opts.urn) {
                throw new Error("Missing required property 'records'");
            }
            if ((!args || args.ttl === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ttl'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            if ((!args || args.zone === undefined) && !opts.urn) {
                throw new Error("Missing required property 'zone'");
            }
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["records"] = args ? args.records : undefined;
            resourceInputs["set_ptr"] = args ? args.set_ptr : undefined;
            resourceInputs["ttl"] = args ? args.ttl : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["zone"] = args ? args.zone : undefined;
        } else {
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["records"] = undefined /*out*/;
            resourceInputs["set_ptr"] = undefined /*out*/;
            resourceInputs["ttl"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
            resourceInputs["zone"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(PDNSRecord.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a PDNSRecord resource.
 */
export interface PDNSRecordArgs {
    name: pulumi.Input<string>;
    records: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * For A and AAAA records, if true, create corresponding PTR.
     */
    set_ptr?: pulumi.Input<boolean>;
    ttl: pulumi.Input<number>;
    type: pulumi.Input<string>;
    zone: pulumi.Input<string>;
}
