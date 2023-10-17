// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Powerdns
{
    [PowerdnsResourceType("powerdns:index:Zone")]
    public partial class Zone : global::Pulumi.CustomResource
    {
        [Output("account")]
        public Output<string?> Account { get; private set; } = null!;

        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        [Output("masters")]
        public Output<ImmutableArray<string>> Masters { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("nameservers")]
        public Output<ImmutableArray<string>> Nameservers { get; private set; } = null!;

        [Output("result")]
        public Output<string> Result { get; private set; } = null!;

        [Output("soaEditAPI")]
        public Output<string?> SoaEditAPI { get; private set; } = null!;

        [Output("zoneId")]
        public Output<string> ZoneId { get; private set; } = null!;


        /// <summary>
        /// Create a Zone resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Zone(string name, ZoneArgs args, CustomResourceOptions? options = null)
            : base("powerdns:index:Zone", name, args ?? new ZoneArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Zone(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("powerdns:index:Zone", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/nextpart",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Zone resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Zone Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Zone(name, id, options);
        }
    }

    public sealed class ZoneArgs : global::Pulumi.ResourceArgs
    {
        [Input("account")]
        public Input<string>? Account { get; set; }

        [Input("kind", required: true)]
        public Input<string> Kind { get; set; } = null!;

        [Input("masters")]
        private InputList<string>? _masters;
        public InputList<string> Masters
        {
            get => _masters ?? (_masters = new InputList<string>());
            set => _masters = value;
        }

        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("nameservers")]
        private InputList<string>? _nameservers;
        public InputList<string> Nameservers
        {
            get => _nameservers ?? (_nameservers = new InputList<string>());
            set => _nameservers = value;
        }

        [Input("soaEditAPI")]
        public Input<string>? SoaEditAPI { get; set; }

        public ZoneArgs()
        {
        }
        public static new ZoneArgs Empty => new ZoneArgs();
    }
}
