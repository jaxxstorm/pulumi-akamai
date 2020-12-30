# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = [
    'Appsecs',
    'Config',
    'Dns',
    'Gtm',
    'Property',
]

@pulumi.output_type
class Appsecs(dict):
    def __init__(__self__, *,
                 access_token: Optional[str] = None,
                 account_key: Optional[str] = None,
                 client_secret: Optional[str] = None,
                 client_token: Optional[str] = None,
                 host: Optional[str] = None,
                 max_body: Optional[int] = None):
        if access_token is not None:
            pulumi.set(__self__, "access_token", access_token)
        if account_key is not None:
            pulumi.set(__self__, "account_key", account_key)
        if client_secret is not None:
            pulumi.set(__self__, "client_secret", client_secret)
        if client_token is not None:
            pulumi.set(__self__, "client_token", client_token)
        if host is not None:
            pulumi.set(__self__, "host", host)
        if max_body is not None:
            pulumi.set(__self__, "max_body", max_body)

    @property
    @pulumi.getter(name="accessToken")
    def access_token(self) -> Optional[str]:
        return pulumi.get(self, "access_token")

    @property
    @pulumi.getter(name="accountKey")
    def account_key(self) -> Optional[str]:
        return pulumi.get(self, "account_key")

    @property
    @pulumi.getter(name="clientSecret")
    def client_secret(self) -> Optional[str]:
        return pulumi.get(self, "client_secret")

    @property
    @pulumi.getter(name="clientToken")
    def client_token(self) -> Optional[str]:
        return pulumi.get(self, "client_token")

    @property
    @pulumi.getter
    def host(self) -> Optional[str]:
        return pulumi.get(self, "host")

    @property
    @pulumi.getter(name="maxBody")
    def max_body(self) -> Optional[int]:
        return pulumi.get(self, "max_body")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class Config(dict):
    def __init__(__self__, *,
                 access_token: Optional[str] = None,
                 account_key: Optional[str] = None,
                 client_secret: Optional[str] = None,
                 client_token: Optional[str] = None,
                 host: Optional[str] = None,
                 max_body: Optional[int] = None):
        if access_token is not None:
            pulumi.set(__self__, "access_token", access_token)
        if account_key is not None:
            pulumi.set(__self__, "account_key", account_key)
        if client_secret is not None:
            pulumi.set(__self__, "client_secret", client_secret)
        if client_token is not None:
            pulumi.set(__self__, "client_token", client_token)
        if host is not None:
            pulumi.set(__self__, "host", host)
        if max_body is not None:
            pulumi.set(__self__, "max_body", max_body)

    @property
    @pulumi.getter(name="accessToken")
    def access_token(self) -> Optional[str]:
        return pulumi.get(self, "access_token")

    @property
    @pulumi.getter(name="accountKey")
    def account_key(self) -> Optional[str]:
        return pulumi.get(self, "account_key")

    @property
    @pulumi.getter(name="clientSecret")
    def client_secret(self) -> Optional[str]:
        return pulumi.get(self, "client_secret")

    @property
    @pulumi.getter(name="clientToken")
    def client_token(self) -> Optional[str]:
        return pulumi.get(self, "client_token")

    @property
    @pulumi.getter
    def host(self) -> Optional[str]:
        return pulumi.get(self, "host")

    @property
    @pulumi.getter(name="maxBody")
    def max_body(self) -> Optional[int]:
        return pulumi.get(self, "max_body")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class Dns(dict):
    def __init__(__self__, *,
                 access_token: Optional[str] = None,
                 account_key: Optional[str] = None,
                 client_secret: Optional[str] = None,
                 client_token: Optional[str] = None,
                 host: Optional[str] = None,
                 max_body: Optional[int] = None):
        if access_token is not None:
            pulumi.set(__self__, "access_token", access_token)
        if account_key is not None:
            pulumi.set(__self__, "account_key", account_key)
        if client_secret is not None:
            pulumi.set(__self__, "client_secret", client_secret)
        if client_token is not None:
            pulumi.set(__self__, "client_token", client_token)
        if host is not None:
            pulumi.set(__self__, "host", host)
        if max_body is not None:
            pulumi.set(__self__, "max_body", max_body)

    @property
    @pulumi.getter(name="accessToken")
    def access_token(self) -> Optional[str]:
        return pulumi.get(self, "access_token")

    @property
    @pulumi.getter(name="accountKey")
    def account_key(self) -> Optional[str]:
        return pulumi.get(self, "account_key")

    @property
    @pulumi.getter(name="clientSecret")
    def client_secret(self) -> Optional[str]:
        return pulumi.get(self, "client_secret")

    @property
    @pulumi.getter(name="clientToken")
    def client_token(self) -> Optional[str]:
        return pulumi.get(self, "client_token")

    @property
    @pulumi.getter
    def host(self) -> Optional[str]:
        return pulumi.get(self, "host")

    @property
    @pulumi.getter(name="maxBody")
    def max_body(self) -> Optional[int]:
        return pulumi.get(self, "max_body")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class Gtm(dict):
    def __init__(__self__, *,
                 access_token: Optional[str] = None,
                 account_key: Optional[str] = None,
                 client_secret: Optional[str] = None,
                 client_token: Optional[str] = None,
                 host: Optional[str] = None,
                 max_body: Optional[int] = None):
        if access_token is not None:
            pulumi.set(__self__, "access_token", access_token)
        if account_key is not None:
            pulumi.set(__self__, "account_key", account_key)
        if client_secret is not None:
            pulumi.set(__self__, "client_secret", client_secret)
        if client_token is not None:
            pulumi.set(__self__, "client_token", client_token)
        if host is not None:
            pulumi.set(__self__, "host", host)
        if max_body is not None:
            pulumi.set(__self__, "max_body", max_body)

    @property
    @pulumi.getter(name="accessToken")
    def access_token(self) -> Optional[str]:
        return pulumi.get(self, "access_token")

    @property
    @pulumi.getter(name="accountKey")
    def account_key(self) -> Optional[str]:
        return pulumi.get(self, "account_key")

    @property
    @pulumi.getter(name="clientSecret")
    def client_secret(self) -> Optional[str]:
        return pulumi.get(self, "client_secret")

    @property
    @pulumi.getter(name="clientToken")
    def client_token(self) -> Optional[str]:
        return pulumi.get(self, "client_token")

    @property
    @pulumi.getter
    def host(self) -> Optional[str]:
        return pulumi.get(self, "host")

    @property
    @pulumi.getter(name="maxBody")
    def max_body(self) -> Optional[int]:
        return pulumi.get(self, "max_body")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class Property(dict):
    def __init__(__self__, *,
                 access_token: Optional[str] = None,
                 account_key: Optional[str] = None,
                 client_secret: Optional[str] = None,
                 client_token: Optional[str] = None,
                 host: Optional[str] = None,
                 max_body: Optional[int] = None):
        if access_token is not None:
            pulumi.set(__self__, "access_token", access_token)
        if account_key is not None:
            pulumi.set(__self__, "account_key", account_key)
        if client_secret is not None:
            pulumi.set(__self__, "client_secret", client_secret)
        if client_token is not None:
            pulumi.set(__self__, "client_token", client_token)
        if host is not None:
            pulumi.set(__self__, "host", host)
        if max_body is not None:
            pulumi.set(__self__, "max_body", max_body)

    @property
    @pulumi.getter(name="accessToken")
    def access_token(self) -> Optional[str]:
        return pulumi.get(self, "access_token")

    @property
    @pulumi.getter(name="accountKey")
    def account_key(self) -> Optional[str]:
        return pulumi.get(self, "account_key")

    @property
    @pulumi.getter(name="clientSecret")
    def client_secret(self) -> Optional[str]:
        return pulumi.get(self, "client_secret")

    @property
    @pulumi.getter(name="clientToken")
    def client_token(self) -> Optional[str]:
        return pulumi.get(self, "client_token")

    @property
    @pulumi.getter
    def host(self) -> Optional[str]:
        return pulumi.get(self, "host")

    @property
    @pulumi.getter(name="maxBody")
    def max_body(self) -> Optional[int]:
        return pulumi.get(self, "max_body")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


