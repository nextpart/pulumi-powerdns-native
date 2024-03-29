// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Immutable;

namespace Pulumi.Powerdns
{
    public static class Config
    {
        [global::System.Diagnostics.CodeAnalysis.SuppressMessage("Microsoft.Design", "IDE1006", Justification = 
        "Double underscore prefix used to avoid conflicts with variable names.")]
        private sealed class __Value<T>
        {
            private readonly Func<T> _getter;
            private T _value = default!;
            private bool _set;

            public __Value(Func<T> getter)
            {
                _getter = getter;
            }

            public T Get() => _set ? _value : _getter();

            public void Set(T value)
            {
                _value = value;
                _set = true;
            }
        }

        private static readonly global::Pulumi.Config __config = new global::Pulumi.Config("powerdns");

        private static readonly __Value<bool?> _insecure = new __Value<bool?>(() => __config.GetBoolean("insecure"));
        /// <summary>
        /// Explicitly allow the provider to perform "insecure" SSL requests. If omitted, default value is "false"
        /// </summary>
        public static bool? Insecure
        {
            get => _insecure.Get();
            set => _insecure.Set(value);
        }

        private static readonly __Value<string?> _key = new __Value<string?>(() => __config.Get("key") ?? Utilities.GetEnv("POWERDNS_KEY") ?? "");
        /// <summary>
        /// The access key for API operations
        /// </summary>
        public static string? Key
        {
            get => _key.Get();
            set => _key.Set(value);
        }

        private static readonly __Value<bool?> _logging = new __Value<bool?>(() => __config.GetBoolean("logging"));
        public static bool? Logging
        {
            get => _logging.Get();
            set => _logging.Set(value);
        }

        private static readonly __Value<string?> _url = new __Value<string?>(() => __config.Get("url") ?? Utilities.GetEnv("POWERDNS_URL") ?? "");
        /// <summary>
        /// The api endpoint of the powerdns server
        /// </summary>
        public static string? Url
        {
            get => _url.Get();
            set => _url.Set(value);
        }

        private static readonly __Value<string?> _version = new __Value<string?>(() => __config.Get("version"));
        public static string? Version
        {
            get => _version.Get();
            set => _version.Set(value);
        }

    }
}
