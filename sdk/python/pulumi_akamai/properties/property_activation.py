# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['PropertyActivationArgs', 'PropertyActivation']

@pulumi.input_type
class PropertyActivationArgs:
    def __init__(__self__, *,
                 contacts: pulumi.Input[Sequence[pulumi.Input[str]]],
                 version: pulumi.Input[int],
                 activation_id: Optional[pulumi.Input[str]] = None,
                 auto_acknowledge_rule_warnings: Optional[pulumi.Input[bool]] = None,
                 network: Optional[pulumi.Input[str]] = None,
                 note: Optional[pulumi.Input[str]] = None,
                 property: Optional[pulumi.Input[str]] = None,
                 property_id: Optional[pulumi.Input[str]] = None,
                 rule_errors: Optional[pulumi.Input[Sequence[pulumi.Input['PropertyActivationRuleErrorArgs']]]] = None,
                 rule_warnings: Optional[pulumi.Input[Sequence[pulumi.Input['PropertyActivationRuleWarningArgs']]]] = None):
        """
        The set of arguments for constructing a PropertyActivation resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] contacts: One or more email addresses to send activation status changes to.
        :param pulumi.Input[int] version: The property version to activate. Previously this field was optional. It now depends on the `Property` resource to identify latest instead of calculating it locally.  This association helps keep the dependency tree properly aligned. To always use the latest version, enter this value `{resource}.{resource identifier}.{field name}`. Using the example code above, the entry would be `akamai_property.example.latest_version` since we want the value of the `latest_version` attribute in the `Property` resource labeled `example`.
        :param pulumi.Input[str] activation_id: The ID given to the activation event while it's in progress.
        :param pulumi.Input[bool] auto_acknowledge_rule_warnings: Whether the activation should proceed despite any warnings. By default set to `true`.
        :param pulumi.Input[str] network: Akamai network to activate on, either `STAGING` or `PRODUCTION`. `STAGING` is the default.
        :param pulumi.Input[str] note: A log message you can assign to the activation request.
        :param pulumi.Input[str] property: - (Deprecated) Replaced by `property_id`. Maintained for legacy purposes.
        :param pulumi.Input[str] property_id: - (Required) The property's unique identifier, including the `prp_` prefix.
        """
        pulumi.set(__self__, "contacts", contacts)
        pulumi.set(__self__, "version", version)
        if activation_id is not None:
            pulumi.set(__self__, "activation_id", activation_id)
        if auto_acknowledge_rule_warnings is not None:
            pulumi.set(__self__, "auto_acknowledge_rule_warnings", auto_acknowledge_rule_warnings)
        if network is not None:
            pulumi.set(__self__, "network", network)
        if note is not None:
            pulumi.set(__self__, "note", note)
        if property is not None:
            warnings.warn("""The setting \"property\" has been deprecated.""", DeprecationWarning)
            pulumi.log.warn("""property is deprecated: The setting \"property\" has been deprecated.""")
        if property is not None:
            pulumi.set(__self__, "property", property)
        if property_id is not None:
            pulumi.set(__self__, "property_id", property_id)
        if rule_errors is not None:
            pulumi.set(__self__, "rule_errors", rule_errors)
        if rule_warnings is not None:
            warnings.warn("""Rule warnings will not be set in state anymore""", DeprecationWarning)
            pulumi.log.warn("""rule_warnings is deprecated: Rule warnings will not be set in state anymore""")
        if rule_warnings is not None:
            pulumi.set(__self__, "rule_warnings", rule_warnings)

    @property
    @pulumi.getter
    def contacts(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        One or more email addresses to send activation status changes to.
        """
        return pulumi.get(self, "contacts")

    @contacts.setter
    def contacts(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "contacts", value)

    @property
    @pulumi.getter
    def version(self) -> pulumi.Input[int]:
        """
        The property version to activate. Previously this field was optional. It now depends on the `Property` resource to identify latest instead of calculating it locally.  This association helps keep the dependency tree properly aligned. To always use the latest version, enter this value `{resource}.{resource identifier}.{field name}`. Using the example code above, the entry would be `akamai_property.example.latest_version` since we want the value of the `latest_version` attribute in the `Property` resource labeled `example`.
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: pulumi.Input[int]):
        pulumi.set(self, "version", value)

    @property
    @pulumi.getter(name="activationId")
    def activation_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID given to the activation event while it's in progress.
        """
        return pulumi.get(self, "activation_id")

    @activation_id.setter
    def activation_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "activation_id", value)

    @property
    @pulumi.getter(name="autoAcknowledgeRuleWarnings")
    def auto_acknowledge_rule_warnings(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the activation should proceed despite any warnings. By default set to `true`.
        """
        return pulumi.get(self, "auto_acknowledge_rule_warnings")

    @auto_acknowledge_rule_warnings.setter
    def auto_acknowledge_rule_warnings(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "auto_acknowledge_rule_warnings", value)

    @property
    @pulumi.getter
    def network(self) -> Optional[pulumi.Input[str]]:
        """
        Akamai network to activate on, either `STAGING` or `PRODUCTION`. `STAGING` is the default.
        """
        return pulumi.get(self, "network")

    @network.setter
    def network(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "network", value)

    @property
    @pulumi.getter
    def note(self) -> Optional[pulumi.Input[str]]:
        """
        A log message you can assign to the activation request.
        """
        return pulumi.get(self, "note")

    @note.setter
    def note(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "note", value)

    @property
    @pulumi.getter(name="propertyId")
    def property_id(self) -> Optional[pulumi.Input[str]]:
        """
        - (Required) The property's unique identifier, including the `prp_` prefix.
        """
        return pulumi.get(self, "property_id")

    @property_id.setter
    def property_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "property_id", value)

    @property
    @pulumi.getter(name="ruleErrors")
    def rule_errors(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['PropertyActivationRuleErrorArgs']]]]:
        return pulumi.get(self, "rule_errors")

    @rule_errors.setter
    def rule_errors(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['PropertyActivationRuleErrorArgs']]]]):
        pulumi.set(self, "rule_errors", value)

    @property
    @pulumi.getter(name="ruleWarnings")
    def rule_warnings(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['PropertyActivationRuleWarningArgs']]]]:
        return pulumi.get(self, "rule_warnings")

    @rule_warnings.setter
    def rule_warnings(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['PropertyActivationRuleWarningArgs']]]]):
        pulumi.set(self, "rule_warnings", value)

    @property
    @pulumi.getter
    def property(self) -> Optional[pulumi.Input[str]]:
        """
        - (Deprecated) Replaced by `property_id`. Maintained for legacy purposes.
        """
        return pulumi.get(self, "property")

    @property.setter
    def property(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "property", value)


@pulumi.input_type
class _PropertyActivationState:
    def __init__(__self__, *,
                 activation_id: Optional[pulumi.Input[str]] = None,
                 auto_acknowledge_rule_warnings: Optional[pulumi.Input[bool]] = None,
                 contacts: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 errors: Optional[pulumi.Input[str]] = None,
                 network: Optional[pulumi.Input[str]] = None,
                 note: Optional[pulumi.Input[str]] = None,
                 property: Optional[pulumi.Input[str]] = None,
                 property_id: Optional[pulumi.Input[str]] = None,
                 rule_errors: Optional[pulumi.Input[Sequence[pulumi.Input['PropertyActivationRuleErrorArgs']]]] = None,
                 rule_warnings: Optional[pulumi.Input[Sequence[pulumi.Input['PropertyActivationRuleWarningArgs']]]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[int]] = None,
                 warnings: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering PropertyActivation resources.
        :param pulumi.Input[str] activation_id: The ID given to the activation event while it's in progress.
        :param pulumi.Input[bool] auto_acknowledge_rule_warnings: Whether the activation should proceed despite any warnings. By default set to `true`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] contacts: One or more email addresses to send activation status changes to.
        :param pulumi.Input[str] errors: The contents of `errors` field returned by the API. For more information see [Errors](https://developer.akamai.com/api/core_features/property_manager/v1.html#errors) in the PAPI documentation.
        :param pulumi.Input[str] network: Akamai network to activate on, either `STAGING` or `PRODUCTION`. `STAGING` is the default.
        :param pulumi.Input[str] note: A log message you can assign to the activation request.
        :param pulumi.Input[str] property: - (Deprecated) Replaced by `property_id`. Maintained for legacy purposes.
        :param pulumi.Input[str] property_id: - (Required) The property's unique identifier, including the `prp_` prefix.
        :param pulumi.Input[str] status: The property version's activation status on the selected network.
        :param pulumi.Input[int] version: The property version to activate. Previously this field was optional. It now depends on the `Property` resource to identify latest instead of calculating it locally.  This association helps keep the dependency tree properly aligned. To always use the latest version, enter this value `{resource}.{resource identifier}.{field name}`. Using the example code above, the entry would be `akamai_property.example.latest_version` since we want the value of the `latest_version` attribute in the `Property` resource labeled `example`.
        :param pulumi.Input[str] warnings: The contents of `warnings` field returned by the API. For more information see [Errors](https://developer.akamai.com/api/core_features/property_manager/v1.html#errors) in the PAPI documentation.
        """
        if activation_id is not None:
            pulumi.set(__self__, "activation_id", activation_id)
        if auto_acknowledge_rule_warnings is not None:
            pulumi.set(__self__, "auto_acknowledge_rule_warnings", auto_acknowledge_rule_warnings)
        if contacts is not None:
            pulumi.set(__self__, "contacts", contacts)
        if errors is not None:
            pulumi.set(__self__, "errors", errors)
        if network is not None:
            pulumi.set(__self__, "network", network)
        if note is not None:
            pulumi.set(__self__, "note", note)
        if property is not None:
            warnings.warn("""The setting \"property\" has been deprecated.""", DeprecationWarning)
            pulumi.log.warn("""property is deprecated: The setting \"property\" has been deprecated.""")
        if property is not None:
            pulumi.set(__self__, "property", property)
        if property_id is not None:
            pulumi.set(__self__, "property_id", property_id)
        if rule_errors is not None:
            pulumi.set(__self__, "rule_errors", rule_errors)
        if rule_warnings is not None:
            warnings.warn("""Rule warnings will not be set in state anymore""", DeprecationWarning)
            pulumi.log.warn("""rule_warnings is deprecated: Rule warnings will not be set in state anymore""")
        if rule_warnings is not None:
            pulumi.set(__self__, "rule_warnings", rule_warnings)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if version is not None:
            pulumi.set(__self__, "version", version)
        if warnings is not None:
            pulumi.set(__self__, "warnings", warnings)

    @property
    @pulumi.getter(name="activationId")
    def activation_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID given to the activation event while it's in progress.
        """
        return pulumi.get(self, "activation_id")

    @activation_id.setter
    def activation_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "activation_id", value)

    @property
    @pulumi.getter(name="autoAcknowledgeRuleWarnings")
    def auto_acknowledge_rule_warnings(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the activation should proceed despite any warnings. By default set to `true`.
        """
        return pulumi.get(self, "auto_acknowledge_rule_warnings")

    @auto_acknowledge_rule_warnings.setter
    def auto_acknowledge_rule_warnings(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "auto_acknowledge_rule_warnings", value)

    @property
    @pulumi.getter
    def contacts(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        One or more email addresses to send activation status changes to.
        """
        return pulumi.get(self, "contacts")

    @contacts.setter
    def contacts(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "contacts", value)

    @property
    @pulumi.getter
    def errors(self) -> Optional[pulumi.Input[str]]:
        """
        The contents of `errors` field returned by the API. For more information see [Errors](https://developer.akamai.com/api/core_features/property_manager/v1.html#errors) in the PAPI documentation.
        """
        return pulumi.get(self, "errors")

    @errors.setter
    def errors(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "errors", value)

    @property
    @pulumi.getter
    def network(self) -> Optional[pulumi.Input[str]]:
        """
        Akamai network to activate on, either `STAGING` or `PRODUCTION`. `STAGING` is the default.
        """
        return pulumi.get(self, "network")

    @network.setter
    def network(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "network", value)

    @property
    @pulumi.getter
    def note(self) -> Optional[pulumi.Input[str]]:
        """
        A log message you can assign to the activation request.
        """
        return pulumi.get(self, "note")

    @note.setter
    def note(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "note", value)

    @property
    @pulumi.getter(name="propertyId")
    def property_id(self) -> Optional[pulumi.Input[str]]:
        """
        - (Required) The property's unique identifier, including the `prp_` prefix.
        """
        return pulumi.get(self, "property_id")

    @property_id.setter
    def property_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "property_id", value)

    @property
    @pulumi.getter(name="ruleErrors")
    def rule_errors(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['PropertyActivationRuleErrorArgs']]]]:
        return pulumi.get(self, "rule_errors")

    @rule_errors.setter
    def rule_errors(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['PropertyActivationRuleErrorArgs']]]]):
        pulumi.set(self, "rule_errors", value)

    @property
    @pulumi.getter(name="ruleWarnings")
    def rule_warnings(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['PropertyActivationRuleWarningArgs']]]]:
        return pulumi.get(self, "rule_warnings")

    @rule_warnings.setter
    def rule_warnings(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['PropertyActivationRuleWarningArgs']]]]):
        pulumi.set(self, "rule_warnings", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The property version's activation status on the selected network.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input[int]]:
        """
        The property version to activate. Previously this field was optional. It now depends on the `Property` resource to identify latest instead of calculating it locally.  This association helps keep the dependency tree properly aligned. To always use the latest version, enter this value `{resource}.{resource identifier}.{field name}`. Using the example code above, the entry would be `akamai_property.example.latest_version` since we want the value of the `latest_version` attribute in the `Property` resource labeled `example`.
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "version", value)

    @property
    @pulumi.getter
    def warnings(self) -> Optional[pulumi.Input[str]]:
        """
        The contents of `warnings` field returned by the API. For more information see [Errors](https://developer.akamai.com/api/core_features/property_manager/v1.html#errors) in the PAPI documentation.
        """
        return pulumi.get(self, "warnings")

    @warnings.setter
    def warnings(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "warnings", value)

    @property
    @pulumi.getter
    def property(self) -> Optional[pulumi.Input[str]]:
        """
        - (Deprecated) Replaced by `property_id`. Maintained for legacy purposes.
        """
        return pulumi.get(self, "property")

    @property.setter
    def property(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "property", value)


warnings.warn("""akamai.properties.PropertyActivation has been deprecated in favor of akamai.PropertyActivation""", DeprecationWarning)


class PropertyActivation(pulumi.CustomResource):
    warnings.warn("""akamai.properties.PropertyActivation has been deprecated in favor of akamai.PropertyActivation""", DeprecationWarning)

    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 activation_id: Optional[pulumi.Input[str]] = None,
                 auto_acknowledge_rule_warnings: Optional[pulumi.Input[bool]] = None,
                 contacts: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 network: Optional[pulumi.Input[str]] = None,
                 note: Optional[pulumi.Input[str]] = None,
                 property: Optional[pulumi.Input[str]] = None,
                 property_id: Optional[pulumi.Input[str]] = None,
                 rule_errors: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PropertyActivationRuleErrorArgs']]]]] = None,
                 rule_warnings: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PropertyActivationRuleWarningArgs']]]]] = None,
                 version: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        The `PropertyActivation` resource lets you activate a property version. An activation deploys the version to either the Akamai staging or production network. You can activate a specific version multiple times if you need to.

        Before activating on production, activate on staging first. This way you can detect any problems in staging before your changes progress to production.

        ## Example Usage

        Basic usage:

        ```python
        import pulumi
        import pulumi_akamai as akamai

        email = "user@example.org"
        rule_format = "v2020-03-04"
        example = akamai.Property("example",
            product_id="prd_SPM",
            contract_id=var["contractid"],
            group_id=var["groupid"],
            hostnames={
                "example.org": "example.org.edgesuite.net",
                "www.example.org": "example.org.edgesuite.net",
                "sub.example.org": "sub.example.org.edgesuite.net",
            },
            rule_format=rule_format,
            rules=(lambda path: open(path).read())(f"{path['module']}/main.json"))
        example_staging = akamai.PropertyActivation("exampleStaging",
            property_id=example.id,
            contacts=[email],
            version=example.latest_version,
            note="Sample activation")
        example_prod = akamai.PropertyActivation("exampleProd",
            property_id=example.id,
            network="PRODUCTION",
            version=3,
            contacts=[email],
            opts=pulumi.ResourceOptions(depends_on=[example_staging]))
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] activation_id: The ID given to the activation event while it's in progress.
        :param pulumi.Input[bool] auto_acknowledge_rule_warnings: Whether the activation should proceed despite any warnings. By default set to `true`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] contacts: One or more email addresses to send activation status changes to.
        :param pulumi.Input[str] network: Akamai network to activate on, either `STAGING` or `PRODUCTION`. `STAGING` is the default.
        :param pulumi.Input[str] note: A log message you can assign to the activation request.
        :param pulumi.Input[str] property: - (Deprecated) Replaced by `property_id`. Maintained for legacy purposes.
        :param pulumi.Input[str] property_id: - (Required) The property's unique identifier, including the `prp_` prefix.
        :param pulumi.Input[int] version: The property version to activate. Previously this field was optional. It now depends on the `Property` resource to identify latest instead of calculating it locally.  This association helps keep the dependency tree properly aligned. To always use the latest version, enter this value `{resource}.{resource identifier}.{field name}`. Using the example code above, the entry would be `akamai_property.example.latest_version` since we want the value of the `latest_version` attribute in the `Property` resource labeled `example`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PropertyActivationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The `PropertyActivation` resource lets you activate a property version. An activation deploys the version to either the Akamai staging or production network. You can activate a specific version multiple times if you need to.

        Before activating on production, activate on staging first. This way you can detect any problems in staging before your changes progress to production.

        ## Example Usage

        Basic usage:

        ```python
        import pulumi
        import pulumi_akamai as akamai

        email = "user@example.org"
        rule_format = "v2020-03-04"
        example = akamai.Property("example",
            product_id="prd_SPM",
            contract_id=var["contractid"],
            group_id=var["groupid"],
            hostnames={
                "example.org": "example.org.edgesuite.net",
                "www.example.org": "example.org.edgesuite.net",
                "sub.example.org": "sub.example.org.edgesuite.net",
            },
            rule_format=rule_format,
            rules=(lambda path: open(path).read())(f"{path['module']}/main.json"))
        example_staging = akamai.PropertyActivation("exampleStaging",
            property_id=example.id,
            contacts=[email],
            version=example.latest_version,
            note="Sample activation")
        example_prod = akamai.PropertyActivation("exampleProd",
            property_id=example.id,
            network="PRODUCTION",
            version=3,
            contacts=[email],
            opts=pulumi.ResourceOptions(depends_on=[example_staging]))
        ```

        :param str resource_name: The name of the resource.
        :param PropertyActivationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PropertyActivationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 activation_id: Optional[pulumi.Input[str]] = None,
                 auto_acknowledge_rule_warnings: Optional[pulumi.Input[bool]] = None,
                 contacts: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 network: Optional[pulumi.Input[str]] = None,
                 note: Optional[pulumi.Input[str]] = None,
                 property: Optional[pulumi.Input[str]] = None,
                 property_id: Optional[pulumi.Input[str]] = None,
                 rule_errors: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PropertyActivationRuleErrorArgs']]]]] = None,
                 rule_warnings: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PropertyActivationRuleWarningArgs']]]]] = None,
                 version: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        pulumi.log.warn("""PropertyActivation is deprecated: akamai.properties.PropertyActivation has been deprecated in favor of akamai.PropertyActivation""")
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PropertyActivationArgs.__new__(PropertyActivationArgs)

            __props__.__dict__["activation_id"] = activation_id
            __props__.__dict__["auto_acknowledge_rule_warnings"] = auto_acknowledge_rule_warnings
            if contacts is None and not opts.urn:
                raise TypeError("Missing required property 'contacts'")
            __props__.__dict__["contacts"] = contacts
            __props__.__dict__["network"] = network
            __props__.__dict__["note"] = note
            if property is not None and not opts.urn:
                warnings.warn("""The setting \"property\" has been deprecated.""", DeprecationWarning)
                pulumi.log.warn("""property is deprecated: The setting \"property\" has been deprecated.""")
            __props__.__dict__["property"] = property
            __props__.__dict__["property_id"] = property_id
            __props__.__dict__["rule_errors"] = rule_errors
            if rule_warnings is not None and not opts.urn:
                warnings.warn("""Rule warnings will not be set in state anymore""", DeprecationWarning)
                pulumi.log.warn("""rule_warnings is deprecated: Rule warnings will not be set in state anymore""")
            __props__.__dict__["rule_warnings"] = rule_warnings
            if version is None and not opts.urn:
                raise TypeError("Missing required property 'version'")
            __props__.__dict__["version"] = version
            __props__.__dict__["errors"] = None
            __props__.__dict__["status"] = None
            __props__.__dict__["warnings"] = None
        super(PropertyActivation, __self__).__init__(
            'akamai:properties/propertyActivation:PropertyActivation',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            activation_id: Optional[pulumi.Input[str]] = None,
            auto_acknowledge_rule_warnings: Optional[pulumi.Input[bool]] = None,
            contacts: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            errors: Optional[pulumi.Input[str]] = None,
            network: Optional[pulumi.Input[str]] = None,
            note: Optional[pulumi.Input[str]] = None,
            property: Optional[pulumi.Input[str]] = None,
            property_id: Optional[pulumi.Input[str]] = None,
            rule_errors: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PropertyActivationRuleErrorArgs']]]]] = None,
            rule_warnings: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PropertyActivationRuleWarningArgs']]]]] = None,
            status: Optional[pulumi.Input[str]] = None,
            version: Optional[pulumi.Input[int]] = None,
            warnings: Optional[pulumi.Input[str]] = None) -> 'PropertyActivation':
        """
        Get an existing PropertyActivation resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] activation_id: The ID given to the activation event while it's in progress.
        :param pulumi.Input[bool] auto_acknowledge_rule_warnings: Whether the activation should proceed despite any warnings. By default set to `true`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] contacts: One or more email addresses to send activation status changes to.
        :param pulumi.Input[str] errors: The contents of `errors` field returned by the API. For more information see [Errors](https://developer.akamai.com/api/core_features/property_manager/v1.html#errors) in the PAPI documentation.
        :param pulumi.Input[str] network: Akamai network to activate on, either `STAGING` or `PRODUCTION`. `STAGING` is the default.
        :param pulumi.Input[str] note: A log message you can assign to the activation request.
        :param pulumi.Input[str] property: - (Deprecated) Replaced by `property_id`. Maintained for legacy purposes.
        :param pulumi.Input[str] property_id: - (Required) The property's unique identifier, including the `prp_` prefix.
        :param pulumi.Input[str] status: The property version's activation status on the selected network.
        :param pulumi.Input[int] version: The property version to activate. Previously this field was optional. It now depends on the `Property` resource to identify latest instead of calculating it locally.  This association helps keep the dependency tree properly aligned. To always use the latest version, enter this value `{resource}.{resource identifier}.{field name}`. Using the example code above, the entry would be `akamai_property.example.latest_version` since we want the value of the `latest_version` attribute in the `Property` resource labeled `example`.
        :param pulumi.Input[str] warnings: The contents of `warnings` field returned by the API. For more information see [Errors](https://developer.akamai.com/api/core_features/property_manager/v1.html#errors) in the PAPI documentation.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PropertyActivationState.__new__(_PropertyActivationState)

        __props__.__dict__["activation_id"] = activation_id
        __props__.__dict__["auto_acknowledge_rule_warnings"] = auto_acknowledge_rule_warnings
        __props__.__dict__["contacts"] = contacts
        __props__.__dict__["errors"] = errors
        __props__.__dict__["network"] = network
        __props__.__dict__["note"] = note
        __props__.__dict__["property"] = property
        __props__.__dict__["property_id"] = property_id
        __props__.__dict__["rule_errors"] = rule_errors
        __props__.__dict__["rule_warnings"] = rule_warnings
        __props__.__dict__["status"] = status
        __props__.__dict__["version"] = version
        __props__.__dict__["warnings"] = warnings
        return PropertyActivation(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="activationId")
    def activation_id(self) -> pulumi.Output[str]:
        """
        The ID given to the activation event while it's in progress.
        """
        return pulumi.get(self, "activation_id")

    @property
    @pulumi.getter(name="autoAcknowledgeRuleWarnings")
    def auto_acknowledge_rule_warnings(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether the activation should proceed despite any warnings. By default set to `true`.
        """
        return pulumi.get(self, "auto_acknowledge_rule_warnings")

    @property
    @pulumi.getter
    def contacts(self) -> pulumi.Output[Sequence[str]]:
        """
        One or more email addresses to send activation status changes to.
        """
        return pulumi.get(self, "contacts")

    @property
    @pulumi.getter
    def errors(self) -> pulumi.Output[str]:
        """
        The contents of `errors` field returned by the API. For more information see [Errors](https://developer.akamai.com/api/core_features/property_manager/v1.html#errors) in the PAPI documentation.
        """
        return pulumi.get(self, "errors")

    @property
    @pulumi.getter
    def network(self) -> pulumi.Output[Optional[str]]:
        """
        Akamai network to activate on, either `STAGING` or `PRODUCTION`. `STAGING` is the default.
        """
        return pulumi.get(self, "network")

    @property
    @pulumi.getter
    def note(self) -> pulumi.Output[Optional[str]]:
        """
        A log message you can assign to the activation request.
        """
        return pulumi.get(self, "note")

    @property
    @pulumi.getter(name="propertyId")
    def property_id(self) -> pulumi.Output[str]:
        """
        - (Required) The property's unique identifier, including the `prp_` prefix.
        """
        return pulumi.get(self, "property_id")

    @property
    @pulumi.getter(name="ruleErrors")
    def rule_errors(self) -> pulumi.Output[Sequence['outputs.PropertyActivationRuleError']]:
        return pulumi.get(self, "rule_errors")

    @property
    @pulumi.getter(name="ruleWarnings")
    def rule_warnings(self) -> pulumi.Output[Sequence['outputs.PropertyActivationRuleWarning']]:
        return pulumi.get(self, "rule_warnings")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The property version's activation status on the selected network.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def version(self) -> pulumi.Output[int]:
        """
        The property version to activate. Previously this field was optional. It now depends on the `Property` resource to identify latest instead of calculating it locally.  This association helps keep the dependency tree properly aligned. To always use the latest version, enter this value `{resource}.{resource identifier}.{field name}`. Using the example code above, the entry would be `akamai_property.example.latest_version` since we want the value of the `latest_version` attribute in the `Property` resource labeled `example`.
        """
        return pulumi.get(self, "version")

    @property
    @pulumi.getter
    def warnings(self) -> pulumi.Output[str]:
        """
        The contents of `warnings` field returned by the API. For more information see [Errors](https://developer.akamai.com/api/core_features/property_manager/v1.html#errors) in the PAPI documentation.
        """
        return pulumi.get(self, "warnings")

    @property
    @pulumi.getter
    def property(self) -> pulumi.Output[str]:
        """
        - (Deprecated) Replaced by `property_id`. Maintained for legacy purposes.
        """
        return pulumi.get(self, "property")

