// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Immutable;

namespace Pulumi.Akamai
{
    public static class Config
    {
        [System.Diagnostics.CodeAnalysis.SuppressMessage("Microsoft.Design", "IDE1006", Justification = 
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

        private static readonly Pulumi.Config __config = new Pulumi.Config("akamai");

        private static readonly __Value<string?> _appsecSection = new __Value<string?>(() => __config.Get("appsecSection"));
        public static string? AppsecSection
        {
            get => _appsecSection.Get();
            set => _appsecSection.Set(value);
        }

        private static readonly __Value<ImmutableArray<Pulumi.Akamai.Config.Types.Appsecs>> _appsecs = new __Value<ImmutableArray<Pulumi.Akamai.Config.Types.Appsecs>>(() => __config.GetObject<ImmutableArray<Pulumi.Akamai.Config.Types.Appsecs>>("appsecs"));
        public static ImmutableArray<Pulumi.Akamai.Config.Types.Appsecs> Appsecs
        {
            get => _appsecs.Get();
            set => _appsecs.Set(value);
        }

        private static readonly __Value<bool?> _cacheEnabled = new __Value<bool?>(() => __config.GetBoolean("cacheEnabled"));
        public static bool? CacheEnabled
        {
            get => _cacheEnabled.Get();
            set => _cacheEnabled.Set(value);
        }

        private static readonly __Value<Pulumi.Akamai.Config.Types.Config?> _config = new __Value<Pulumi.Akamai.Config.Types.Config?>(() => __config.GetObject<Pulumi.Akamai.Config.Types.Config>("config"));
        public static Pulumi.Akamai.Config.Types.Config? ConfigDetails
        {
            get => _config.Get();
            set => _config.Set(value);
        }

        private static readonly __Value<string?> _configSection = new __Value<string?>(() => __config.Get("configSection"));
        /// <summary>
        /// The section of the edgerc file to use for configuration
        /// </summary>
        public static string? ConfigSection
        {
            get => _configSection.Get();
            set => _configSection.Set(value);
        }

        private static readonly __Value<Pulumi.Akamai.Config.Types.Dns?> _dns = new __Value<Pulumi.Akamai.Config.Types.Dns?>(() => __config.GetObject<Pulumi.Akamai.Config.Types.Dns>("dns"));
        public static Pulumi.Akamai.Config.Types.Dns? Dns
        {
            get => _dns.Get();
            set => _dns.Set(value);
        }

        private static readonly __Value<string?> _dnsSection = new __Value<string?>(() => __config.Get("dnsSection"));
        public static string? DnsSection
        {
            get => _dnsSection.Get();
            set => _dnsSection.Set(value);
        }

        private static readonly __Value<string?> _edgerc = new __Value<string?>(() => __config.Get("edgerc"));
        public static string? Edgerc
        {
            get => _edgerc.Get();
            set => _edgerc.Set(value);
        }

        private static readonly __Value<Pulumi.Akamai.Config.Types.Gtm?> _gtm = new __Value<Pulumi.Akamai.Config.Types.Gtm?>(() => __config.GetObject<Pulumi.Akamai.Config.Types.Gtm>("gtm"));
        public static Pulumi.Akamai.Config.Types.Gtm? Gtm
        {
            get => _gtm.Get();
            set => _gtm.Set(value);
        }

        private static readonly __Value<string?> _gtmSection = new __Value<string?>(() => __config.Get("gtmSection"));
        public static string? GtmSection
        {
            get => _gtmSection.Get();
            set => _gtmSection.Set(value);
        }

        private static readonly __Value<string?> _networklistSection = new __Value<string?>(() => __config.Get("networklistSection"));
        public static string? NetworklistSection
        {
            get => _networklistSection.Get();
            set => _networklistSection.Set(value);
        }

        private static readonly __Value<ImmutableArray<Pulumi.Akamai.Config.Types.Networks>> _networks = new __Value<ImmutableArray<Pulumi.Akamai.Config.Types.Networks>>(() => __config.GetObject<ImmutableArray<Pulumi.Akamai.Config.Types.Networks>>("networks"));
        public static ImmutableArray<Pulumi.Akamai.Config.Types.Networks> Networks
        {
            get => _networks.Get();
            set => _networks.Set(value);
        }

        private static readonly __Value<string?> _papiSection = new __Value<string?>(() => __config.Get("papiSection"));
        public static string? PapiSection
        {
            get => _papiSection.Get();
            set => _papiSection.Set(value);
        }

        private static readonly __Value<Pulumi.Akamai.Config.Types.Property?> _property = new __Value<Pulumi.Akamai.Config.Types.Property?>(() => __config.GetObject<Pulumi.Akamai.Config.Types.Property>("property"));
        public static Pulumi.Akamai.Config.Types.Property? Property
        {
            get => _property.Get();
            set => _property.Set(value);
        }

        private static readonly __Value<string?> _propertySection = new __Value<string?>(() => __config.Get("propertySection"));
        public static string? PropertySection
        {
            get => _propertySection.Get();
            set => _propertySection.Set(value);
        }

        public static class Types
        {

             public class Appsecs
             {
                public string? AccessToken { get; set; } = null!;
                public string? AccountKey { get; set; } = null!;
                public string? ClientSecret { get; set; } = null!;
                public string? ClientToken { get; set; } = null!;
                public string? Host { get; set; } = null!;
                public int? MaxBody { get; set; }
            }

             public class Config
             {
                public string? AccessToken { get; set; } = null!;
                public string? AccountKey { get; set; } = null!;
                public string? ClientSecret { get; set; } = null!;
                public string? ClientToken { get; set; } = null!;
                public string? Host { get; set; } = null!;
                public int? MaxBody { get; set; }
            }

             public class Dns
             {
                public string? AccessToken { get; set; } = null!;
                public string? AccountKey { get; set; } = null!;
                public string? ClientSecret { get; set; } = null!;
                public string? ClientToken { get; set; } = null!;
                public string? Host { get; set; } = null!;
                public int? MaxBody { get; set; }
            }

             public class Gtm
             {
                public string? AccessToken { get; set; } = null!;
                public string? AccountKey { get; set; } = null!;
                public string? ClientSecret { get; set; } = null!;
                public string? ClientToken { get; set; } = null!;
                public string? Host { get; set; } = null!;
                public int? MaxBody { get; set; }
            }

             public class Networks
             {
                public string? AccessToken { get; set; } = null!;
                public string? AccountKey { get; set; } = null!;
                public string? ClientSecret { get; set; } = null!;
                public string? ClientToken { get; set; } = null!;
                public string? Host { get; set; } = null!;
                public int? MaxBody { get; set; }
            }

             public class Property
             {
                public string? AccessToken { get; set; } = null!;
                public string? AccountKey { get; set; } = null!;
                public string? ClientSecret { get; set; } = null!;
                public string? ClientToken { get; set; } = null!;
                public string? Host { get; set; } = null!;
                public int? MaxBody { get; set; }
            }
        }
    }
}
