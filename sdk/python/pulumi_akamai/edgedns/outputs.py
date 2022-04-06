# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'DnsZoneTsigKey',
]

@pulumi.output_type
class DnsZoneTsigKey(dict):
    def __init__(__self__, *,
                 algorithm: str,
                 name: str,
                 secret: str):
        """
        :param str algorithm: The hashing algorithm.
        :param str name: The key name.
        :param str secret: String known between transfer endpoints.
        """
        pulumi.set(__self__, "algorithm", algorithm)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "secret", secret)

    @property
    @pulumi.getter
    def algorithm(self) -> str:
        """
        The hashing algorithm.
        """
        return pulumi.get(self, "algorithm")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The key name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def secret(self) -> str:
        """
        String known between transfer endpoints.
        """
        return pulumi.get(self, "secret")


