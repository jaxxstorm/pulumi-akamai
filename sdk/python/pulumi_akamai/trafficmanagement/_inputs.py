# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GtmASmapAssignmentArgs',
    'GtmASmapDefaultDatacenterArgs',
    'GtmCidrmapAssignmentArgs',
    'GtmCidrmapDefaultDatacenterArgs',
    'GtmDatacenterDefaultLoadObjectArgs',
    'GtmGeomapAssignmentArgs',
    'GtmGeomapDefaultDatacenterArgs',
    'GtmPropertyLivenessTestArgs',
    'GtmPropertyLivenessTestHttpHeaderArgs',
    'GtmPropertyStaticRrSetArgs',
    'GtmPropertyTrafficTargetArgs',
    'GtmResourceResourceInstanceArgs',
]

@pulumi.input_type
class GtmASmapAssignmentArgs:
    def __init__(__self__, *,
                 as_numbers: pulumi.Input[Sequence[pulumi.Input[int]]],
                 datacenter_id: pulumi.Input[int],
                 nickname: pulumi.Input[str]):
        pulumi.set(__self__, "as_numbers", as_numbers)
        pulumi.set(__self__, "datacenter_id", datacenter_id)
        pulumi.set(__self__, "nickname", nickname)

    @property
    @pulumi.getter(name="asNumbers")
    def as_numbers(self) -> pulumi.Input[Sequence[pulumi.Input[int]]]:
        return pulumi.get(self, "as_numbers")

    @as_numbers.setter
    def as_numbers(self, value: pulumi.Input[Sequence[pulumi.Input[int]]]):
        pulumi.set(self, "as_numbers", value)

    @property
    @pulumi.getter(name="datacenterId")
    def datacenter_id(self) -> pulumi.Input[int]:
        return pulumi.get(self, "datacenter_id")

    @datacenter_id.setter
    def datacenter_id(self, value: pulumi.Input[int]):
        pulumi.set(self, "datacenter_id", value)

    @property
    @pulumi.getter
    def nickname(self) -> pulumi.Input[str]:
        return pulumi.get(self, "nickname")

    @nickname.setter
    def nickname(self, value: pulumi.Input[str]):
        pulumi.set(self, "nickname", value)


@pulumi.input_type
class GtmASmapDefaultDatacenterArgs:
    def __init__(__self__, *,
                 datacenter_id: pulumi.Input[int],
                 nickname: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "datacenter_id", datacenter_id)
        if nickname is not None:
            pulumi.set(__self__, "nickname", nickname)

    @property
    @pulumi.getter(name="datacenterId")
    def datacenter_id(self) -> pulumi.Input[int]:
        return pulumi.get(self, "datacenter_id")

    @datacenter_id.setter
    def datacenter_id(self, value: pulumi.Input[int]):
        pulumi.set(self, "datacenter_id", value)

    @property
    @pulumi.getter
    def nickname(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "nickname")

    @nickname.setter
    def nickname(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "nickname", value)


@pulumi.input_type
class GtmCidrmapAssignmentArgs:
    def __init__(__self__, *,
                 datacenter_id: pulumi.Input[int],
                 nickname: pulumi.Input[str],
                 blocks: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        pulumi.set(__self__, "datacenter_id", datacenter_id)
        pulumi.set(__self__, "nickname", nickname)
        if blocks is not None:
            pulumi.set(__self__, "blocks", blocks)

    @property
    @pulumi.getter(name="datacenterId")
    def datacenter_id(self) -> pulumi.Input[int]:
        return pulumi.get(self, "datacenter_id")

    @datacenter_id.setter
    def datacenter_id(self, value: pulumi.Input[int]):
        pulumi.set(self, "datacenter_id", value)

    @property
    @pulumi.getter
    def nickname(self) -> pulumi.Input[str]:
        return pulumi.get(self, "nickname")

    @nickname.setter
    def nickname(self, value: pulumi.Input[str]):
        pulumi.set(self, "nickname", value)

    @property
    @pulumi.getter
    def blocks(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "blocks")

    @blocks.setter
    def blocks(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "blocks", value)


@pulumi.input_type
class GtmCidrmapDefaultDatacenterArgs:
    def __init__(__self__, *,
                 datacenter_id: pulumi.Input[int],
                 nickname: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "datacenter_id", datacenter_id)
        if nickname is not None:
            pulumi.set(__self__, "nickname", nickname)

    @property
    @pulumi.getter(name="datacenterId")
    def datacenter_id(self) -> pulumi.Input[int]:
        return pulumi.get(self, "datacenter_id")

    @datacenter_id.setter
    def datacenter_id(self, value: pulumi.Input[int]):
        pulumi.set(self, "datacenter_id", value)

    @property
    @pulumi.getter
    def nickname(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "nickname")

    @nickname.setter
    def nickname(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "nickname", value)


@pulumi.input_type
class GtmDatacenterDefaultLoadObjectArgs:
    def __init__(__self__, *,
                 load_object: Optional[pulumi.Input[str]] = None,
                 load_object_port: Optional[pulumi.Input[int]] = None,
                 load_servers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        if load_object is not None:
            pulumi.set(__self__, "load_object", load_object)
        if load_object_port is not None:
            pulumi.set(__self__, "load_object_port", load_object_port)
        if load_servers is not None:
            pulumi.set(__self__, "load_servers", load_servers)

    @property
    @pulumi.getter(name="loadObject")
    def load_object(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "load_object")

    @load_object.setter
    def load_object(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "load_object", value)

    @property
    @pulumi.getter(name="loadObjectPort")
    def load_object_port(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "load_object_port")

    @load_object_port.setter
    def load_object_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "load_object_port", value)

    @property
    @pulumi.getter(name="loadServers")
    def load_servers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "load_servers")

    @load_servers.setter
    def load_servers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "load_servers", value)


@pulumi.input_type
class GtmGeomapAssignmentArgs:
    def __init__(__self__, *,
                 datacenter_id: pulumi.Input[int],
                 nickname: pulumi.Input[str],
                 countries: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        pulumi.set(__self__, "datacenter_id", datacenter_id)
        pulumi.set(__self__, "nickname", nickname)
        if countries is not None:
            pulumi.set(__self__, "countries", countries)

    @property
    @pulumi.getter(name="datacenterId")
    def datacenter_id(self) -> pulumi.Input[int]:
        return pulumi.get(self, "datacenter_id")

    @datacenter_id.setter
    def datacenter_id(self, value: pulumi.Input[int]):
        pulumi.set(self, "datacenter_id", value)

    @property
    @pulumi.getter
    def nickname(self) -> pulumi.Input[str]:
        return pulumi.get(self, "nickname")

    @nickname.setter
    def nickname(self, value: pulumi.Input[str]):
        pulumi.set(self, "nickname", value)

    @property
    @pulumi.getter
    def countries(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "countries")

    @countries.setter
    def countries(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "countries", value)


@pulumi.input_type
class GtmGeomapDefaultDatacenterArgs:
    def __init__(__self__, *,
                 datacenter_id: pulumi.Input[int],
                 nickname: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "datacenter_id", datacenter_id)
        if nickname is not None:
            pulumi.set(__self__, "nickname", nickname)

    @property
    @pulumi.getter(name="datacenterId")
    def datacenter_id(self) -> pulumi.Input[int]:
        return pulumi.get(self, "datacenter_id")

    @datacenter_id.setter
    def datacenter_id(self, value: pulumi.Input[int]):
        pulumi.set(self, "datacenter_id", value)

    @property
    @pulumi.getter
    def nickname(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "nickname")

    @nickname.setter
    def nickname(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "nickname", value)


@pulumi.input_type
class GtmPropertyLivenessTestArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 test_interval: pulumi.Input[int],
                 test_object: pulumi.Input[str],
                 test_object_protocol: pulumi.Input[str],
                 test_timeout: pulumi.Input[float],
                 answers_required: Optional[pulumi.Input[bool]] = None,
                 disable_nonstandard_port_warning: Optional[pulumi.Input[bool]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 error_penalty: Optional[pulumi.Input[float]] = None,
                 http_error3xx: Optional[pulumi.Input[bool]] = None,
                 http_error4xx: Optional[pulumi.Input[bool]] = None,
                 http_error5xx: Optional[pulumi.Input[bool]] = None,
                 http_headers: Optional[pulumi.Input[Sequence[pulumi.Input['GtmPropertyLivenessTestHttpHeaderArgs']]]] = None,
                 peer_certificate_verification: Optional[pulumi.Input[bool]] = None,
                 recursion_requested: Optional[pulumi.Input[bool]] = None,
                 request_string: Optional[pulumi.Input[str]] = None,
                 resource_type: Optional[pulumi.Input[str]] = None,
                 response_string: Optional[pulumi.Input[str]] = None,
                 ssl_client_certificate: Optional[pulumi.Input[str]] = None,
                 ssl_client_private_key: Optional[pulumi.Input[str]] = None,
                 test_object_password: Optional[pulumi.Input[str]] = None,
                 test_object_port: Optional[pulumi.Input[int]] = None,
                 test_object_username: Optional[pulumi.Input[str]] = None,
                 timeout_penalty: Optional[pulumi.Input[float]] = None):
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "test_interval", test_interval)
        pulumi.set(__self__, "test_object", test_object)
        pulumi.set(__self__, "test_object_protocol", test_object_protocol)
        pulumi.set(__self__, "test_timeout", test_timeout)
        if answers_required is not None:
            pulumi.set(__self__, "answers_required", answers_required)
        if disable_nonstandard_port_warning is not None:
            pulumi.set(__self__, "disable_nonstandard_port_warning", disable_nonstandard_port_warning)
        if disabled is not None:
            pulumi.set(__self__, "disabled", disabled)
        if error_penalty is not None:
            pulumi.set(__self__, "error_penalty", error_penalty)
        if http_error3xx is not None:
            pulumi.set(__self__, "http_error3xx", http_error3xx)
        if http_error4xx is not None:
            pulumi.set(__self__, "http_error4xx", http_error4xx)
        if http_error5xx is not None:
            pulumi.set(__self__, "http_error5xx", http_error5xx)
        if http_headers is not None:
            pulumi.set(__self__, "http_headers", http_headers)
        if peer_certificate_verification is not None:
            pulumi.set(__self__, "peer_certificate_verification", peer_certificate_verification)
        if recursion_requested is not None:
            pulumi.set(__self__, "recursion_requested", recursion_requested)
        if request_string is not None:
            pulumi.set(__self__, "request_string", request_string)
        if resource_type is not None:
            pulumi.set(__self__, "resource_type", resource_type)
        if response_string is not None:
            pulumi.set(__self__, "response_string", response_string)
        if ssl_client_certificate is not None:
            pulumi.set(__self__, "ssl_client_certificate", ssl_client_certificate)
        if ssl_client_private_key is not None:
            pulumi.set(__self__, "ssl_client_private_key", ssl_client_private_key)
        if test_object_password is not None:
            pulumi.set(__self__, "test_object_password", test_object_password)
        if test_object_port is not None:
            pulumi.set(__self__, "test_object_port", test_object_port)
        if test_object_username is not None:
            pulumi.set(__self__, "test_object_username", test_object_username)
        if timeout_penalty is not None:
            pulumi.set(__self__, "timeout_penalty", timeout_penalty)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="testInterval")
    def test_interval(self) -> pulumi.Input[int]:
        return pulumi.get(self, "test_interval")

    @test_interval.setter
    def test_interval(self, value: pulumi.Input[int]):
        pulumi.set(self, "test_interval", value)

    @property
    @pulumi.getter(name="testObject")
    def test_object(self) -> pulumi.Input[str]:
        return pulumi.get(self, "test_object")

    @test_object.setter
    def test_object(self, value: pulumi.Input[str]):
        pulumi.set(self, "test_object", value)

    @property
    @pulumi.getter(name="testObjectProtocol")
    def test_object_protocol(self) -> pulumi.Input[str]:
        return pulumi.get(self, "test_object_protocol")

    @test_object_protocol.setter
    def test_object_protocol(self, value: pulumi.Input[str]):
        pulumi.set(self, "test_object_protocol", value)

    @property
    @pulumi.getter(name="testTimeout")
    def test_timeout(self) -> pulumi.Input[float]:
        return pulumi.get(self, "test_timeout")

    @test_timeout.setter
    def test_timeout(self, value: pulumi.Input[float]):
        pulumi.set(self, "test_timeout", value)

    @property
    @pulumi.getter(name="answersRequired")
    def answers_required(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "answers_required")

    @answers_required.setter
    def answers_required(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "answers_required", value)

    @property
    @pulumi.getter(name="disableNonstandardPortWarning")
    def disable_nonstandard_port_warning(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "disable_nonstandard_port_warning")

    @disable_nonstandard_port_warning.setter
    def disable_nonstandard_port_warning(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disable_nonstandard_port_warning", value)

    @property
    @pulumi.getter
    def disabled(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "disabled")

    @disabled.setter
    def disabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disabled", value)

    @property
    @pulumi.getter(name="errorPenalty")
    def error_penalty(self) -> Optional[pulumi.Input[float]]:
        return pulumi.get(self, "error_penalty")

    @error_penalty.setter
    def error_penalty(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "error_penalty", value)

    @property
    @pulumi.getter(name="httpError3xx")
    def http_error3xx(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "http_error3xx")

    @http_error3xx.setter
    def http_error3xx(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "http_error3xx", value)

    @property
    @pulumi.getter(name="httpError4xx")
    def http_error4xx(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "http_error4xx")

    @http_error4xx.setter
    def http_error4xx(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "http_error4xx", value)

    @property
    @pulumi.getter(name="httpError5xx")
    def http_error5xx(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "http_error5xx")

    @http_error5xx.setter
    def http_error5xx(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "http_error5xx", value)

    @property
    @pulumi.getter(name="httpHeaders")
    def http_headers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['GtmPropertyLivenessTestHttpHeaderArgs']]]]:
        return pulumi.get(self, "http_headers")

    @http_headers.setter
    def http_headers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['GtmPropertyLivenessTestHttpHeaderArgs']]]]):
        pulumi.set(self, "http_headers", value)

    @property
    @pulumi.getter(name="peerCertificateVerification")
    def peer_certificate_verification(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "peer_certificate_verification")

    @peer_certificate_verification.setter
    def peer_certificate_verification(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "peer_certificate_verification", value)

    @property
    @pulumi.getter(name="recursionRequested")
    def recursion_requested(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "recursion_requested")

    @recursion_requested.setter
    def recursion_requested(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "recursion_requested", value)

    @property
    @pulumi.getter(name="requestString")
    def request_string(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "request_string")

    @request_string.setter
    def request_string(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "request_string", value)

    @property
    @pulumi.getter(name="resourceType")
    def resource_type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "resource_type")

    @resource_type.setter
    def resource_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_type", value)

    @property
    @pulumi.getter(name="responseString")
    def response_string(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "response_string")

    @response_string.setter
    def response_string(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "response_string", value)

    @property
    @pulumi.getter(name="sslClientCertificate")
    def ssl_client_certificate(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "ssl_client_certificate")

    @ssl_client_certificate.setter
    def ssl_client_certificate(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ssl_client_certificate", value)

    @property
    @pulumi.getter(name="sslClientPrivateKey")
    def ssl_client_private_key(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "ssl_client_private_key")

    @ssl_client_private_key.setter
    def ssl_client_private_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ssl_client_private_key", value)

    @property
    @pulumi.getter(name="testObjectPassword")
    def test_object_password(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "test_object_password")

    @test_object_password.setter
    def test_object_password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "test_object_password", value)

    @property
    @pulumi.getter(name="testObjectPort")
    def test_object_port(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "test_object_port")

    @test_object_port.setter
    def test_object_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "test_object_port", value)

    @property
    @pulumi.getter(name="testObjectUsername")
    def test_object_username(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "test_object_username")

    @test_object_username.setter
    def test_object_username(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "test_object_username", value)

    @property
    @pulumi.getter(name="timeoutPenalty")
    def timeout_penalty(self) -> Optional[pulumi.Input[float]]:
        return pulumi.get(self, "timeout_penalty")

    @timeout_penalty.setter
    def timeout_penalty(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "timeout_penalty", value)


@pulumi.input_type
class GtmPropertyLivenessTestHttpHeaderArgs:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[str]] = None):
        if name is not None:
            pulumi.set(__self__, "name", name)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class GtmPropertyStaticRrSetArgs:
    def __init__(__self__, *,
                 rdatas: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 ttl: Optional[pulumi.Input[int]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        if rdatas is not None:
            pulumi.set(__self__, "rdatas", rdatas)
        if ttl is not None:
            pulumi.set(__self__, "ttl", ttl)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def rdatas(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "rdatas")

    @rdatas.setter
    def rdatas(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "rdatas", value)

    @property
    @pulumi.getter
    def ttl(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "ttl")

    @ttl.setter
    def ttl(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "ttl", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class GtmPropertyTrafficTargetArgs:
    def __init__(__self__, *,
                 datacenter_id: Optional[pulumi.Input[int]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 handout_cname: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 servers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 weight: Optional[pulumi.Input[float]] = None):
        if datacenter_id is not None:
            pulumi.set(__self__, "datacenter_id", datacenter_id)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if handout_cname is not None:
            pulumi.set(__self__, "handout_cname", handout_cname)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if servers is not None:
            pulumi.set(__self__, "servers", servers)
        if weight is not None:
            pulumi.set(__self__, "weight", weight)

    @property
    @pulumi.getter(name="datacenterId")
    def datacenter_id(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "datacenter_id")

    @datacenter_id.setter
    def datacenter_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "datacenter_id", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="handoutCname")
    def handout_cname(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "handout_cname")

    @handout_cname.setter
    def handout_cname(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "handout_cname", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def servers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "servers")

    @servers.setter
    def servers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "servers", value)

    @property
    @pulumi.getter
    def weight(self) -> Optional[pulumi.Input[float]]:
        return pulumi.get(self, "weight")

    @weight.setter
    def weight(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "weight", value)


@pulumi.input_type
class GtmResourceResourceInstanceArgs:
    def __init__(__self__, *,
                 datacenter_id: pulumi.Input[int],
                 load_object: Optional[pulumi.Input[str]] = None,
                 load_object_port: Optional[pulumi.Input[int]] = None,
                 load_servers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 use_default_load_object: Optional[pulumi.Input[bool]] = None):
        pulumi.set(__self__, "datacenter_id", datacenter_id)
        if load_object is not None:
            pulumi.set(__self__, "load_object", load_object)
        if load_object_port is not None:
            pulumi.set(__self__, "load_object_port", load_object_port)
        if load_servers is not None:
            pulumi.set(__self__, "load_servers", load_servers)
        if use_default_load_object is not None:
            pulumi.set(__self__, "use_default_load_object", use_default_load_object)

    @property
    @pulumi.getter(name="datacenterId")
    def datacenter_id(self) -> pulumi.Input[int]:
        return pulumi.get(self, "datacenter_id")

    @datacenter_id.setter
    def datacenter_id(self, value: pulumi.Input[int]):
        pulumi.set(self, "datacenter_id", value)

    @property
    @pulumi.getter(name="loadObject")
    def load_object(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "load_object")

    @load_object.setter
    def load_object(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "load_object", value)

    @property
    @pulumi.getter(name="loadObjectPort")
    def load_object_port(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "load_object_port")

    @load_object_port.setter
    def load_object_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "load_object_port", value)

    @property
    @pulumi.getter(name="loadServers")
    def load_servers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "load_servers")

    @load_servers.setter
    def load_servers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "load_servers", value)

    @property
    @pulumi.getter(name="useDefaultLoadObject")
    def use_default_load_object(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "use_default_load_object")

    @use_default_load_object.setter
    def use_default_load_object(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "use_default_load_object", value)


