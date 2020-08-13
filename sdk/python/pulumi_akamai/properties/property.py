# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class Property(pulumi.CustomResource):
    account: pulumi.Output[str]
    """
    — the Account ID under which the property is created.
    """
    contacts: pulumi.Output[list]
    """
    — (Required) One or more email addresses to inform about activation changes.
    """
    contract: pulumi.Output[str]
    """
    — (Optional) The contract ID.
    """
    cp_code: pulumi.Output[str]
    """
    — (Optional) The CP Code id or name to use (or create). Required unless a [cpCode behavior](https://developer.akamai.com/api/core_features/property_manager/vlatest.html#cpcode) is present in the default rule.
    """
    edge_hostnames: pulumi.Output[dict]
    """
    — the final public hostname to edge hostname map
    """
    group: pulumi.Output[str]
    """
    — (Optional) The group ID.
    """
    hostnames: pulumi.Output[dict]
    """
    — (Required) A map of public hostnames to edge hostnames (e.g. `{"example.org" = "example.org.edgesuite.net"}`)
    """
    is_secure: pulumi.Output[bool]
    """
    — (Optional) Whether the property is a secure (Enhanced TLS) property or not.
    """
    name: pulumi.Output[str]
    """
    — (Required) The property name.
    """
    origins: pulumi.Output[list]
    """
    — (Optional) The property origin (an origin must be specified to activate a property, but may be defined in your rules block).

      * `cacheKeyHostname` (`str`) - — (Optional) The hostname uses for the cache key. (default: `ORIGIN_HOSTNAME`).
      * `compress` (`bool`) - — (Optional, boolean) Whether origin supports gzip compression (default: `false`).
      * `enableTrueClientIp` (`bool`) - — (Optional, boolean) Whether the X-True-Client-IP header should be sent to origin (default: `false`).
      * `forwardHostname` (`str`) - — (Optional) The value for the Hostname header sent to origin. (default: `ORIGIN_HOSTNAME`).
      * `hostname` (`str`) - — (Required) The origin hostname.
      * `port` (`float`) - — (Optional) The origin port to connect to (default: 80).
    """
    product: pulumi.Output[str]
    """
    — (Optional) The product ID. (Default: `prd_SPM` for Ion)
    """
    production_version: pulumi.Output[float]
    """
    — the current version of the property active on the production network.
    """
    rule_format: pulumi.Output[str]
    """
    — (Optional) The rule format to use ([more](https://developer.akamai.com/api/core_features/property_manager/v1.html#getruleformats)).
    """
    rules: pulumi.Output[str]
    """
    — (Required) A JSON encoded string of property rules (see: [`properties.PropertyRules`](https://www.terraform.io/docs/providers/akamai/d/property_rules.html))
    """
    rulessha: pulumi.Output[str]
    staging_version: pulumi.Output[float]
    """
    — the current version of the property active on the staging network.
    """
    variables: pulumi.Output[str]
    """
    — (Optional) A JSON encoded string of property manager variable definitions (see: [`properties.PropertyVariables`](https://www.terraform.io/docs/providers/akamai/r/property_variables.html))
    """
    version: pulumi.Output[float]
    """
    — the current version of the property config.
    """
    def __init__(__self__, resource_name, opts=None, contacts=None, contract=None, cp_code=None, group=None, hostnames=None, is_secure=None, name=None, origins=None, product=None, rule_format=None, rules=None, variables=None, __props__=None, __name__=None, __opts__=None):
        """
        The `properties.Property` resource represents an Akamai property configuration, allowing you to create,
        update, and activate properties on the Akamai platform.

        ## Example Usage
        ### Basic usage:

        ```python
        import pulumi
        import pulumi_akamai as akamai

        example = akamai.properties.Property("example",
            contacts=["user@example.org"],
            contract="ctr_####",
            cp_code="cpc_#####",
            group="grp_####",
            hostnames={
                "example.org": "example.org.edgesuite.net",
                "sub.example.org": "sub.example.org.edgesuite.net",
                "www.example.org": "example.org.edgesuite.net",
            },
            product="prd_SPM",
            rule_format="v2018-02-27",
            rules=data["local_file"]["terraform-demo"]["content"],
            variables=akamai_property_variables["origin"]["json"])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] contacts: — (Required) One or more email addresses to inform about activation changes.
        :param pulumi.Input[str] contract: — (Optional) The contract ID.
        :param pulumi.Input[str] cp_code: — (Optional) The CP Code id or name to use (or create). Required unless a [cpCode behavior](https://developer.akamai.com/api/core_features/property_manager/vlatest.html#cpcode) is present in the default rule.
        :param pulumi.Input[str] group: — (Optional) The group ID.
        :param pulumi.Input[dict] hostnames: — (Required) A map of public hostnames to edge hostnames (e.g. `{"example.org" = "example.org.edgesuite.net"}`)
        :param pulumi.Input[bool] is_secure: — (Optional) Whether the property is a secure (Enhanced TLS) property or not.
        :param pulumi.Input[str] name: — (Required) The property name.
        :param pulumi.Input[list] origins: — (Optional) The property origin (an origin must be specified to activate a property, but may be defined in your rules block).
        :param pulumi.Input[str] product: — (Optional) The product ID. (Default: `prd_SPM` for Ion)
        :param pulumi.Input[str] rule_format: — (Optional) The rule format to use ([more](https://developer.akamai.com/api/core_features/property_manager/v1.html#getruleformats)).
        :param pulumi.Input[str] rules: — (Required) A JSON encoded string of property rules (see: [`properties.PropertyRules`](https://www.terraform.io/docs/providers/akamai/d/property_rules.html))
        :param pulumi.Input[str] variables: — (Optional) A JSON encoded string of property manager variable definitions (see: [`properties.PropertyVariables`](https://www.terraform.io/docs/providers/akamai/r/property_variables.html))

        The **origins** object supports the following:

          * `cacheKeyHostname` (`pulumi.Input[str]`) - — (Optional) The hostname uses for the cache key. (default: `ORIGIN_HOSTNAME`).
          * `compress` (`pulumi.Input[bool]`) - — (Optional, boolean) Whether origin supports gzip compression (default: `false`).
          * `enableTrueClientIp` (`pulumi.Input[bool]`) - — (Optional, boolean) Whether the X-True-Client-IP header should be sent to origin (default: `false`).
          * `forwardHostname` (`pulumi.Input[str]`) - — (Optional) The value for the Hostname header sent to origin. (default: `ORIGIN_HOSTNAME`).
          * `hostname` (`pulumi.Input[str]`) - — (Required) The origin hostname.
          * `port` (`pulumi.Input[float]`) - — (Optional) The origin port to connect to (default: 80).
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if contacts is None:
                raise TypeError("Missing required property 'contacts'")
            __props__['contacts'] = contacts
            __props__['contract'] = contract
            __props__['cp_code'] = cp_code
            __props__['group'] = group
            if hostnames is None:
                raise TypeError("Missing required property 'hostnames'")
            __props__['hostnames'] = hostnames
            __props__['is_secure'] = is_secure
            __props__['name'] = name
            __props__['origins'] = origins
            __props__['product'] = product
            __props__['rule_format'] = rule_format
            __props__['rules'] = rules
            __props__['variables'] = variables
            __props__['account'] = None
            __props__['edge_hostnames'] = None
            __props__['production_version'] = None
            __props__['rulessha'] = None
            __props__['staging_version'] = None
            __props__['version'] = None
        super(Property, __self__).__init__(
            'akamai:properties/property:Property',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, account=None, contacts=None, contract=None, cp_code=None, edge_hostnames=None, group=None, hostnames=None, is_secure=None, name=None, origins=None, product=None, production_version=None, rule_format=None, rules=None, rulessha=None, staging_version=None, variables=None, version=None):
        """
        Get an existing Property resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account: — the Account ID under which the property is created.
        :param pulumi.Input[list] contacts: — (Required) One or more email addresses to inform about activation changes.
        :param pulumi.Input[str] contract: — (Optional) The contract ID.
        :param pulumi.Input[str] cp_code: — (Optional) The CP Code id or name to use (or create). Required unless a [cpCode behavior](https://developer.akamai.com/api/core_features/property_manager/vlatest.html#cpcode) is present in the default rule.
        :param pulumi.Input[dict] edge_hostnames: — the final public hostname to edge hostname map
        :param pulumi.Input[str] group: — (Optional) The group ID.
        :param pulumi.Input[dict] hostnames: — (Required) A map of public hostnames to edge hostnames (e.g. `{"example.org" = "example.org.edgesuite.net"}`)
        :param pulumi.Input[bool] is_secure: — (Optional) Whether the property is a secure (Enhanced TLS) property or not.
        :param pulumi.Input[str] name: — (Required) The property name.
        :param pulumi.Input[list] origins: — (Optional) The property origin (an origin must be specified to activate a property, but may be defined in your rules block).
        :param pulumi.Input[str] product: — (Optional) The product ID. (Default: `prd_SPM` for Ion)
        :param pulumi.Input[float] production_version: — the current version of the property active on the production network.
        :param pulumi.Input[str] rule_format: — (Optional) The rule format to use ([more](https://developer.akamai.com/api/core_features/property_manager/v1.html#getruleformats)).
        :param pulumi.Input[str] rules: — (Required) A JSON encoded string of property rules (see: [`properties.PropertyRules`](https://www.terraform.io/docs/providers/akamai/d/property_rules.html))
        :param pulumi.Input[float] staging_version: — the current version of the property active on the staging network.
        :param pulumi.Input[str] variables: — (Optional) A JSON encoded string of property manager variable definitions (see: [`properties.PropertyVariables`](https://www.terraform.io/docs/providers/akamai/r/property_variables.html))
        :param pulumi.Input[float] version: — the current version of the property config.

        The **origins** object supports the following:

          * `cacheKeyHostname` (`pulumi.Input[str]`) - — (Optional) The hostname uses for the cache key. (default: `ORIGIN_HOSTNAME`).
          * `compress` (`pulumi.Input[bool]`) - — (Optional, boolean) Whether origin supports gzip compression (default: `false`).
          * `enableTrueClientIp` (`pulumi.Input[bool]`) - — (Optional, boolean) Whether the X-True-Client-IP header should be sent to origin (default: `false`).
          * `forwardHostname` (`pulumi.Input[str]`) - — (Optional) The value for the Hostname header sent to origin. (default: `ORIGIN_HOSTNAME`).
          * `hostname` (`pulumi.Input[str]`) - — (Required) The origin hostname.
          * `port` (`pulumi.Input[float]`) - — (Optional) The origin port to connect to (default: 80).
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["account"] = account
        __props__["contacts"] = contacts
        __props__["contract"] = contract
        __props__["cp_code"] = cp_code
        __props__["edge_hostnames"] = edge_hostnames
        __props__["group"] = group
        __props__["hostnames"] = hostnames
        __props__["is_secure"] = is_secure
        __props__["name"] = name
        __props__["origins"] = origins
        __props__["product"] = product
        __props__["production_version"] = production_version
        __props__["rule_format"] = rule_format
        __props__["rules"] = rules
        __props__["rulessha"] = rulessha
        __props__["staging_version"] = staging_version
        __props__["variables"] = variables
        __props__["version"] = version
        return Property(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
