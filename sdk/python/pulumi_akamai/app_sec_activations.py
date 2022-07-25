# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['AppSecActivationsArgs', 'AppSecActivations']

@pulumi.input_type
class AppSecActivationsArgs:
    def __init__(__self__, *,
                 config_id: pulumi.Input[int],
                 notification_emails: pulumi.Input[Sequence[pulumi.Input[str]]],
                 version: pulumi.Input[int],
                 activate: Optional[pulumi.Input[bool]] = None,
                 network: Optional[pulumi.Input[str]] = None,
                 note: Optional[pulumi.Input[str]] = None,
                 notes: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a AppSecActivations resource.
        :param pulumi.Input[int] config_id: . Unique identifier of the security configuration being activated. This is unchanged from previous versions.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] notification_emails: . JSON array containing the email addresses of the people to be notified when activation is complete. This is unchanged from previous versions.
        :param pulumi.Input[int] version: . Version number of the security configuration being activated. This can be a hard-coded version number (for example, **5**), or you can use the security configuration’s **latest_version** attribute (data.akamai_appsec_configuration.configuration.latest_version). If you do the latter, you’ll always activate the most recent version of the configuration. This argument applies only to versions 2.0.0 and later.
        :param pulumi.Input[bool] activate: . Set to **true** to activate the specified security configuration or set to **false** to deactivate the configuration. If not included, the security configuration is activated. This argument applies only to versions prior to 2.0.0.
        :param pulumi.Input[str] network: . Network on which activation will occur; if not included, activation takes place on the staging network. Allowed values are:
               * **PRODUCTION**
               * **STAGING**
        :param pulumi.Input[str] note: . Brief description of the activation or deactivation process. If no attributes have changed since the last time you called the **akamai_appsec_activations** resource, neither activation nor deactivation takes place. That's because something must be different in order to trigger these processes. Because of that, it's recommended that you always update the **note** argument. That ensures that the resource is called and that activation or deactivation occurs.
        :param pulumi.Input[str] notes: . Brief description of the activation or deactivation process. If no attributes have changed since the last time you called the **akamai_appsec_activations** resource, neither activation nor deactivation takes place. That's because something must be different in order to trigger one of these processes. Because of that, it's recommended that you always update the `notes` argument. Doing so ensures that the resource is called and activation or deactivation occurs. This argument applies only to versions prior to 2.0.0.
        """
        pulumi.set(__self__, "config_id", config_id)
        pulumi.set(__self__, "notification_emails", notification_emails)
        pulumi.set(__self__, "version", version)
        if activate is not None:
            warnings.warn("""The setting activate has been deprecated; \"terraform apply\" will always perform activation. (Use \"terraform destroy\" for deactivation.)""", DeprecationWarning)
            pulumi.log.warn("""activate is deprecated: The setting activate has been deprecated; \"terraform apply\" will always perform activation. (Use \"terraform destroy\" for deactivation.)""")
        if activate is not None:
            pulumi.set(__self__, "activate", activate)
        if network is not None:
            pulumi.set(__self__, "network", network)
        if note is not None:
            pulumi.set(__self__, "note", note)
        if notes is not None:
            warnings.warn("""The setting notes has been deprecated. Use \"note\" instead.""", DeprecationWarning)
            pulumi.log.warn("""notes is deprecated: The setting notes has been deprecated. Use \"note\" instead.""")
        if notes is not None:
            pulumi.set(__self__, "notes", notes)

    @property
    @pulumi.getter(name="configId")
    def config_id(self) -> pulumi.Input[int]:
        """
        . Unique identifier of the security configuration being activated. This is unchanged from previous versions.
        """
        return pulumi.get(self, "config_id")

    @config_id.setter
    def config_id(self, value: pulumi.Input[int]):
        pulumi.set(self, "config_id", value)

    @property
    @pulumi.getter(name="notificationEmails")
    def notification_emails(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        . JSON array containing the email addresses of the people to be notified when activation is complete. This is unchanged from previous versions.
        """
        return pulumi.get(self, "notification_emails")

    @notification_emails.setter
    def notification_emails(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "notification_emails", value)

    @property
    @pulumi.getter
    def version(self) -> pulumi.Input[int]:
        """
        . Version number of the security configuration being activated. This can be a hard-coded version number (for example, **5**), or you can use the security configuration’s **latest_version** attribute (data.akamai_appsec_configuration.configuration.latest_version). If you do the latter, you’ll always activate the most recent version of the configuration. This argument applies only to versions 2.0.0 and later.
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: pulumi.Input[int]):
        pulumi.set(self, "version", value)

    @property
    @pulumi.getter
    def activate(self) -> Optional[pulumi.Input[bool]]:
        """
        . Set to **true** to activate the specified security configuration or set to **false** to deactivate the configuration. If not included, the security configuration is activated. This argument applies only to versions prior to 2.0.0.
        """
        return pulumi.get(self, "activate")

    @activate.setter
    def activate(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "activate", value)

    @property
    @pulumi.getter
    def network(self) -> Optional[pulumi.Input[str]]:
        """
        . Network on which activation will occur; if not included, activation takes place on the staging network. Allowed values are:
        * **PRODUCTION**
        * **STAGING**
        """
        return pulumi.get(self, "network")

    @network.setter
    def network(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "network", value)

    @property
    @pulumi.getter
    def note(self) -> Optional[pulumi.Input[str]]:
        """
        . Brief description of the activation or deactivation process. If no attributes have changed since the last time you called the **akamai_appsec_activations** resource, neither activation nor deactivation takes place. That's because something must be different in order to trigger these processes. Because of that, it's recommended that you always update the **note** argument. That ensures that the resource is called and that activation or deactivation occurs.
        """
        return pulumi.get(self, "note")

    @note.setter
    def note(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "note", value)

    @property
    @pulumi.getter
    def notes(self) -> Optional[pulumi.Input[str]]:
        """
        . Brief description of the activation or deactivation process. If no attributes have changed since the last time you called the **akamai_appsec_activations** resource, neither activation nor deactivation takes place. That's because something must be different in order to trigger one of these processes. Because of that, it's recommended that you always update the `notes` argument. Doing so ensures that the resource is called and activation or deactivation occurs. This argument applies only to versions prior to 2.0.0.
        """
        return pulumi.get(self, "notes")

    @notes.setter
    def notes(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "notes", value)


@pulumi.input_type
class _AppSecActivationsState:
    def __init__(__self__, *,
                 activate: Optional[pulumi.Input[bool]] = None,
                 config_id: Optional[pulumi.Input[int]] = None,
                 network: Optional[pulumi.Input[str]] = None,
                 note: Optional[pulumi.Input[str]] = None,
                 notes: Optional[pulumi.Input[str]] = None,
                 notification_emails: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering AppSecActivations resources.
        :param pulumi.Input[bool] activate: . Set to **true** to activate the specified security configuration or set to **false** to deactivate the configuration. If not included, the security configuration is activated. This argument applies only to versions prior to 2.0.0.
        :param pulumi.Input[int] config_id: . Unique identifier of the security configuration being activated. This is unchanged from previous versions.
        :param pulumi.Input[str] network: . Network on which activation will occur; if not included, activation takes place on the staging network. Allowed values are:
               * **PRODUCTION**
               * **STAGING**
        :param pulumi.Input[str] note: . Brief description of the activation or deactivation process. If no attributes have changed since the last time you called the **akamai_appsec_activations** resource, neither activation nor deactivation takes place. That's because something must be different in order to trigger these processes. Because of that, it's recommended that you always update the **note** argument. That ensures that the resource is called and that activation or deactivation occurs.
        :param pulumi.Input[str] notes: . Brief description of the activation or deactivation process. If no attributes have changed since the last time you called the **akamai_appsec_activations** resource, neither activation nor deactivation takes place. That's because something must be different in order to trigger one of these processes. Because of that, it's recommended that you always update the `notes` argument. Doing so ensures that the resource is called and activation or deactivation occurs. This argument applies only to versions prior to 2.0.0.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] notification_emails: . JSON array containing the email addresses of the people to be notified when activation is complete. This is unchanged from previous versions.
        :param pulumi.Input[str] status: The results of the activation
        :param pulumi.Input[int] version: . Version number of the security configuration being activated. This can be a hard-coded version number (for example, **5**), or you can use the security configuration’s **latest_version** attribute (data.akamai_appsec_configuration.configuration.latest_version). If you do the latter, you’ll always activate the most recent version of the configuration. This argument applies only to versions 2.0.0 and later.
        """
        if activate is not None:
            warnings.warn("""The setting activate has been deprecated; \"terraform apply\" will always perform activation. (Use \"terraform destroy\" for deactivation.)""", DeprecationWarning)
            pulumi.log.warn("""activate is deprecated: The setting activate has been deprecated; \"terraform apply\" will always perform activation. (Use \"terraform destroy\" for deactivation.)""")
        if activate is not None:
            pulumi.set(__self__, "activate", activate)
        if config_id is not None:
            pulumi.set(__self__, "config_id", config_id)
        if network is not None:
            pulumi.set(__self__, "network", network)
        if note is not None:
            pulumi.set(__self__, "note", note)
        if notes is not None:
            warnings.warn("""The setting notes has been deprecated. Use \"note\" instead.""", DeprecationWarning)
            pulumi.log.warn("""notes is deprecated: The setting notes has been deprecated. Use \"note\" instead.""")
        if notes is not None:
            pulumi.set(__self__, "notes", notes)
        if notification_emails is not None:
            pulumi.set(__self__, "notification_emails", notification_emails)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter
    def activate(self) -> Optional[pulumi.Input[bool]]:
        """
        . Set to **true** to activate the specified security configuration or set to **false** to deactivate the configuration. If not included, the security configuration is activated. This argument applies only to versions prior to 2.0.0.
        """
        return pulumi.get(self, "activate")

    @activate.setter
    def activate(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "activate", value)

    @property
    @pulumi.getter(name="configId")
    def config_id(self) -> Optional[pulumi.Input[int]]:
        """
        . Unique identifier of the security configuration being activated. This is unchanged from previous versions.
        """
        return pulumi.get(self, "config_id")

    @config_id.setter
    def config_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "config_id", value)

    @property
    @pulumi.getter
    def network(self) -> Optional[pulumi.Input[str]]:
        """
        . Network on which activation will occur; if not included, activation takes place on the staging network. Allowed values are:
        * **PRODUCTION**
        * **STAGING**
        """
        return pulumi.get(self, "network")

    @network.setter
    def network(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "network", value)

    @property
    @pulumi.getter
    def note(self) -> Optional[pulumi.Input[str]]:
        """
        . Brief description of the activation or deactivation process. If no attributes have changed since the last time you called the **akamai_appsec_activations** resource, neither activation nor deactivation takes place. That's because something must be different in order to trigger these processes. Because of that, it's recommended that you always update the **note** argument. That ensures that the resource is called and that activation or deactivation occurs.
        """
        return pulumi.get(self, "note")

    @note.setter
    def note(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "note", value)

    @property
    @pulumi.getter
    def notes(self) -> Optional[pulumi.Input[str]]:
        """
        . Brief description of the activation or deactivation process. If no attributes have changed since the last time you called the **akamai_appsec_activations** resource, neither activation nor deactivation takes place. That's because something must be different in order to trigger one of these processes. Because of that, it's recommended that you always update the `notes` argument. Doing so ensures that the resource is called and activation or deactivation occurs. This argument applies only to versions prior to 2.0.0.
        """
        return pulumi.get(self, "notes")

    @notes.setter
    def notes(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "notes", value)

    @property
    @pulumi.getter(name="notificationEmails")
    def notification_emails(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        . JSON array containing the email addresses of the people to be notified when activation is complete. This is unchanged from previous versions.
        """
        return pulumi.get(self, "notification_emails")

    @notification_emails.setter
    def notification_emails(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "notification_emails", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The results of the activation
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input[int]]:
        """
        . Version number of the security configuration being activated. This can be a hard-coded version number (for example, **5**), or you can use the security configuration’s **latest_version** attribute (data.akamai_appsec_configuration.configuration.latest_version). If you do the latter, you’ll always activate the most recent version of the configuration. This argument applies only to versions 2.0.0 and later.
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "version", value)


class AppSecActivations(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 activate: Optional[pulumi.Input[bool]] = None,
                 config_id: Optional[pulumi.Input[int]] = None,
                 network: Optional[pulumi.Input[str]] = None,
                 note: Optional[pulumi.Input[str]] = None,
                 notes: Optional[pulumi.Input[str]] = None,
                 notification_emails: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 version: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        Create a AppSecActivations resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] activate: . Set to **true** to activate the specified security configuration or set to **false** to deactivate the configuration. If not included, the security configuration is activated. This argument applies only to versions prior to 2.0.0.
        :param pulumi.Input[int] config_id: . Unique identifier of the security configuration being activated. This is unchanged from previous versions.
        :param pulumi.Input[str] network: . Network on which activation will occur; if not included, activation takes place on the staging network. Allowed values are:
               * **PRODUCTION**
               * **STAGING**
        :param pulumi.Input[str] note: . Brief description of the activation or deactivation process. If no attributes have changed since the last time you called the **akamai_appsec_activations** resource, neither activation nor deactivation takes place. That's because something must be different in order to trigger these processes. Because of that, it's recommended that you always update the **note** argument. That ensures that the resource is called and that activation or deactivation occurs.
        :param pulumi.Input[str] notes: . Brief description of the activation or deactivation process. If no attributes have changed since the last time you called the **akamai_appsec_activations** resource, neither activation nor deactivation takes place. That's because something must be different in order to trigger one of these processes. Because of that, it's recommended that you always update the `notes` argument. Doing so ensures that the resource is called and activation or deactivation occurs. This argument applies only to versions prior to 2.0.0.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] notification_emails: . JSON array containing the email addresses of the people to be notified when activation is complete. This is unchanged from previous versions.
        :param pulumi.Input[int] version: . Version number of the security configuration being activated. This can be a hard-coded version number (for example, **5**), or you can use the security configuration’s **latest_version** attribute (data.akamai_appsec_configuration.configuration.latest_version). If you do the latter, you’ll always activate the most recent version of the configuration. This argument applies only to versions 2.0.0 and later.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AppSecActivationsArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a AppSecActivations resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param AppSecActivationsArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AppSecActivationsArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 activate: Optional[pulumi.Input[bool]] = None,
                 config_id: Optional[pulumi.Input[int]] = None,
                 network: Optional[pulumi.Input[str]] = None,
                 note: Optional[pulumi.Input[str]] = None,
                 notes: Optional[pulumi.Input[str]] = None,
                 notification_emails: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 version: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AppSecActivationsArgs.__new__(AppSecActivationsArgs)

            if activate is not None and not opts.urn:
                warnings.warn("""The setting activate has been deprecated; \"terraform apply\" will always perform activation. (Use \"terraform destroy\" for deactivation.)""", DeprecationWarning)
                pulumi.log.warn("""activate is deprecated: The setting activate has been deprecated; \"terraform apply\" will always perform activation. (Use \"terraform destroy\" for deactivation.)""")
            __props__.__dict__["activate"] = activate
            if config_id is None and not opts.urn:
                raise TypeError("Missing required property 'config_id'")
            __props__.__dict__["config_id"] = config_id
            __props__.__dict__["network"] = network
            __props__.__dict__["note"] = note
            if notes is not None and not opts.urn:
                warnings.warn("""The setting notes has been deprecated. Use \"note\" instead.""", DeprecationWarning)
                pulumi.log.warn("""notes is deprecated: The setting notes has been deprecated. Use \"note\" instead.""")
            __props__.__dict__["notes"] = notes
            if notification_emails is None and not opts.urn:
                raise TypeError("Missing required property 'notification_emails'")
            __props__.__dict__["notification_emails"] = notification_emails
            if version is None and not opts.urn:
                raise TypeError("Missing required property 'version'")
            __props__.__dict__["version"] = version
            __props__.__dict__["status"] = None
        super(AppSecActivations, __self__).__init__(
            'akamai:index/appSecActivations:AppSecActivations',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            activate: Optional[pulumi.Input[bool]] = None,
            config_id: Optional[pulumi.Input[int]] = None,
            network: Optional[pulumi.Input[str]] = None,
            note: Optional[pulumi.Input[str]] = None,
            notes: Optional[pulumi.Input[str]] = None,
            notification_emails: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            status: Optional[pulumi.Input[str]] = None,
            version: Optional[pulumi.Input[int]] = None) -> 'AppSecActivations':
        """
        Get an existing AppSecActivations resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] activate: . Set to **true** to activate the specified security configuration or set to **false** to deactivate the configuration. If not included, the security configuration is activated. This argument applies only to versions prior to 2.0.0.
        :param pulumi.Input[int] config_id: . Unique identifier of the security configuration being activated. This is unchanged from previous versions.
        :param pulumi.Input[str] network: . Network on which activation will occur; if not included, activation takes place on the staging network. Allowed values are:
               * **PRODUCTION**
               * **STAGING**
        :param pulumi.Input[str] note: . Brief description of the activation or deactivation process. If no attributes have changed since the last time you called the **akamai_appsec_activations** resource, neither activation nor deactivation takes place. That's because something must be different in order to trigger these processes. Because of that, it's recommended that you always update the **note** argument. That ensures that the resource is called and that activation or deactivation occurs.
        :param pulumi.Input[str] notes: . Brief description of the activation or deactivation process. If no attributes have changed since the last time you called the **akamai_appsec_activations** resource, neither activation nor deactivation takes place. That's because something must be different in order to trigger one of these processes. Because of that, it's recommended that you always update the `notes` argument. Doing so ensures that the resource is called and activation or deactivation occurs. This argument applies only to versions prior to 2.0.0.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] notification_emails: . JSON array containing the email addresses of the people to be notified when activation is complete. This is unchanged from previous versions.
        :param pulumi.Input[str] status: The results of the activation
        :param pulumi.Input[int] version: . Version number of the security configuration being activated. This can be a hard-coded version number (for example, **5**), or you can use the security configuration’s **latest_version** attribute (data.akamai_appsec_configuration.configuration.latest_version). If you do the latter, you’ll always activate the most recent version of the configuration. This argument applies only to versions 2.0.0 and later.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AppSecActivationsState.__new__(_AppSecActivationsState)

        __props__.__dict__["activate"] = activate
        __props__.__dict__["config_id"] = config_id
        __props__.__dict__["network"] = network
        __props__.__dict__["note"] = note
        __props__.__dict__["notes"] = notes
        __props__.__dict__["notification_emails"] = notification_emails
        __props__.__dict__["status"] = status
        __props__.__dict__["version"] = version
        return AppSecActivations(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def activate(self) -> pulumi.Output[Optional[bool]]:
        """
        . Set to **true** to activate the specified security configuration or set to **false** to deactivate the configuration. If not included, the security configuration is activated. This argument applies only to versions prior to 2.0.0.
        """
        return pulumi.get(self, "activate")

    @property
    @pulumi.getter(name="configId")
    def config_id(self) -> pulumi.Output[int]:
        """
        . Unique identifier of the security configuration being activated. This is unchanged from previous versions.
        """
        return pulumi.get(self, "config_id")

    @property
    @pulumi.getter
    def network(self) -> pulumi.Output[Optional[str]]:
        """
        . Network on which activation will occur; if not included, activation takes place on the staging network. Allowed values are:
        * **PRODUCTION**
        * **STAGING**
        """
        return pulumi.get(self, "network")

    @property
    @pulumi.getter
    def note(self) -> pulumi.Output[Optional[str]]:
        """
        . Brief description of the activation or deactivation process. If no attributes have changed since the last time you called the **akamai_appsec_activations** resource, neither activation nor deactivation takes place. That's because something must be different in order to trigger these processes. Because of that, it's recommended that you always update the **note** argument. That ensures that the resource is called and that activation or deactivation occurs.
        """
        return pulumi.get(self, "note")

    @property
    @pulumi.getter
    def notes(self) -> pulumi.Output[Optional[str]]:
        """
        . Brief description of the activation or deactivation process. If no attributes have changed since the last time you called the **akamai_appsec_activations** resource, neither activation nor deactivation takes place. That's because something must be different in order to trigger one of these processes. Because of that, it's recommended that you always update the `notes` argument. Doing so ensures that the resource is called and activation or deactivation occurs. This argument applies only to versions prior to 2.0.0.
        """
        return pulumi.get(self, "notes")

    @property
    @pulumi.getter(name="notificationEmails")
    def notification_emails(self) -> pulumi.Output[Sequence[str]]:
        """
        . JSON array containing the email addresses of the people to be notified when activation is complete. This is unchanged from previous versions.
        """
        return pulumi.get(self, "notification_emails")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The results of the activation
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def version(self) -> pulumi.Output[int]:
        """
        . Version number of the security configuration being activated. This can be a hard-coded version number (for example, **5**), or you can use the security configuration’s **latest_version** attribute (data.akamai_appsec_configuration.configuration.latest_version). If you do the latter, you’ll always activate the most recent version of the configuration. This argument applies only to versions 2.0.0 and later.
        """
        return pulumi.get(self, "version")

