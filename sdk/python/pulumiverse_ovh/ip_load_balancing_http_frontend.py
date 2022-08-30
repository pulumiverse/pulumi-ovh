# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['IpLoadBalancingHttpFrontendArgs', 'IpLoadBalancingHttpFrontend']

@pulumi.input_type
class IpLoadBalancingHttpFrontendArgs:
    def __init__(__self__, *,
                 port: pulumi.Input[str],
                 service_name: pulumi.Input[str],
                 zone: pulumi.Input[str],
                 allowed_sources: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 dedicated_ipfos: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 default_farm_id: Optional[pulumi.Input[int]] = None,
                 default_ssl_id: Optional[pulumi.Input[int]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 redirect_location: Optional[pulumi.Input[str]] = None,
                 ssl: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a IpLoadBalancingHttpFrontend resource.
        :param pulumi.Input[str] port: Port(s) attached to your frontend. Supports single port (numerical value),
               range (2 dash-delimited increasing ports) and comma-separated list of 'single port'
               and/or 'range'. Each port must be in the [1;49151] range
        :param pulumi.Input[str] service_name: The internal name of your IP load balancing
        :param pulumi.Input[str] zone: Zone where the frontend will be defined (ie. `gra`, `bhs` also supports `all`)
        :param pulumi.Input[Sequence[pulumi.Input[str]]] allowed_sources: Restrict IP Load Balancing access to these ip block. No restriction if null. List of IP blocks.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] dedicated_ipfos: Only attach frontend on these ip. No restriction if null. List of Ip blocks.
        :param pulumi.Input[int] default_farm_id: Default TCP Farm of your frontend
        :param pulumi.Input[int] default_ssl_id: Default ssl served to your customer
        :param pulumi.Input[bool] disabled: Disable your frontend. Default: 'false'
        :param pulumi.Input[str] display_name: Human readable name for your frontend, this field is for you
        :param pulumi.Input[str] redirect_location: Redirection HTTP'
        :param pulumi.Input[bool] ssl: SSL deciphering. Default: 'false'
        """
        pulumi.set(__self__, "port", port)
        pulumi.set(__self__, "service_name", service_name)
        pulumi.set(__self__, "zone", zone)
        if allowed_sources is not None:
            pulumi.set(__self__, "allowed_sources", allowed_sources)
        if dedicated_ipfos is not None:
            pulumi.set(__self__, "dedicated_ipfos", dedicated_ipfos)
        if default_farm_id is not None:
            pulumi.set(__self__, "default_farm_id", default_farm_id)
        if default_ssl_id is not None:
            pulumi.set(__self__, "default_ssl_id", default_ssl_id)
        if disabled is not None:
            pulumi.set(__self__, "disabled", disabled)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if redirect_location is not None:
            pulumi.set(__self__, "redirect_location", redirect_location)
        if ssl is not None:
            pulumi.set(__self__, "ssl", ssl)

    @property
    @pulumi.getter
    def port(self) -> pulumi.Input[str]:
        """
        Port(s) attached to your frontend. Supports single port (numerical value),
        range (2 dash-delimited increasing ports) and comma-separated list of 'single port'
        and/or 'range'. Each port must be in the [1;49151] range
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: pulumi.Input[str]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Input[str]:
        """
        The internal name of your IP load balancing
        """
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "service_name", value)

    @property
    @pulumi.getter
    def zone(self) -> pulumi.Input[str]:
        """
        Zone where the frontend will be defined (ie. `gra`, `bhs` also supports `all`)
        """
        return pulumi.get(self, "zone")

    @zone.setter
    def zone(self, value: pulumi.Input[str]):
        pulumi.set(self, "zone", value)

    @property
    @pulumi.getter(name="allowedSources")
    def allowed_sources(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Restrict IP Load Balancing access to these ip block. No restriction if null. List of IP blocks.
        """
        return pulumi.get(self, "allowed_sources")

    @allowed_sources.setter
    def allowed_sources(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "allowed_sources", value)

    @property
    @pulumi.getter(name="dedicatedIpfos")
    def dedicated_ipfos(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Only attach frontend on these ip. No restriction if null. List of Ip blocks.
        """
        return pulumi.get(self, "dedicated_ipfos")

    @dedicated_ipfos.setter
    def dedicated_ipfos(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "dedicated_ipfos", value)

    @property
    @pulumi.getter(name="defaultFarmId")
    def default_farm_id(self) -> Optional[pulumi.Input[int]]:
        """
        Default TCP Farm of your frontend
        """
        return pulumi.get(self, "default_farm_id")

    @default_farm_id.setter
    def default_farm_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "default_farm_id", value)

    @property
    @pulumi.getter(name="defaultSslId")
    def default_ssl_id(self) -> Optional[pulumi.Input[int]]:
        """
        Default ssl served to your customer
        """
        return pulumi.get(self, "default_ssl_id")

    @default_ssl_id.setter
    def default_ssl_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "default_ssl_id", value)

    @property
    @pulumi.getter
    def disabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Disable your frontend. Default: 'false'
        """
        return pulumi.get(self, "disabled")

    @disabled.setter
    def disabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disabled", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        Human readable name for your frontend, this field is for you
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="redirectLocation")
    def redirect_location(self) -> Optional[pulumi.Input[str]]:
        """
        Redirection HTTP'
        """
        return pulumi.get(self, "redirect_location")

    @redirect_location.setter
    def redirect_location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "redirect_location", value)

    @property
    @pulumi.getter
    def ssl(self) -> Optional[pulumi.Input[bool]]:
        """
        SSL deciphering. Default: 'false'
        """
        return pulumi.get(self, "ssl")

    @ssl.setter
    def ssl(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "ssl", value)


@pulumi.input_type
class _IpLoadBalancingHttpFrontendState:
    def __init__(__self__, *,
                 allowed_sources: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 dedicated_ipfos: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 default_farm_id: Optional[pulumi.Input[int]] = None,
                 default_ssl_id: Optional[pulumi.Input[int]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[str]] = None,
                 redirect_location: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 ssl: Optional[pulumi.Input[bool]] = None,
                 zone: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering IpLoadBalancingHttpFrontend resources.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] allowed_sources: Restrict IP Load Balancing access to these ip block. No restriction if null. List of IP blocks.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] dedicated_ipfos: Only attach frontend on these ip. No restriction if null. List of Ip blocks.
        :param pulumi.Input[int] default_farm_id: Default TCP Farm of your frontend
        :param pulumi.Input[int] default_ssl_id: Default ssl served to your customer
        :param pulumi.Input[bool] disabled: Disable your frontend. Default: 'false'
        :param pulumi.Input[str] display_name: Human readable name for your frontend, this field is for you
        :param pulumi.Input[str] port: Port(s) attached to your frontend. Supports single port (numerical value),
               range (2 dash-delimited increasing ports) and comma-separated list of 'single port'
               and/or 'range'. Each port must be in the [1;49151] range
        :param pulumi.Input[str] redirect_location: Redirection HTTP'
        :param pulumi.Input[str] service_name: The internal name of your IP load balancing
        :param pulumi.Input[bool] ssl: SSL deciphering. Default: 'false'
        :param pulumi.Input[str] zone: Zone where the frontend will be defined (ie. `gra`, `bhs` also supports `all`)
        """
        if allowed_sources is not None:
            pulumi.set(__self__, "allowed_sources", allowed_sources)
        if dedicated_ipfos is not None:
            pulumi.set(__self__, "dedicated_ipfos", dedicated_ipfos)
        if default_farm_id is not None:
            pulumi.set(__self__, "default_farm_id", default_farm_id)
        if default_ssl_id is not None:
            pulumi.set(__self__, "default_ssl_id", default_ssl_id)
        if disabled is not None:
            pulumi.set(__self__, "disabled", disabled)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if redirect_location is not None:
            pulumi.set(__self__, "redirect_location", redirect_location)
        if service_name is not None:
            pulumi.set(__self__, "service_name", service_name)
        if ssl is not None:
            pulumi.set(__self__, "ssl", ssl)
        if zone is not None:
            pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter(name="allowedSources")
    def allowed_sources(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Restrict IP Load Balancing access to these ip block. No restriction if null. List of IP blocks.
        """
        return pulumi.get(self, "allowed_sources")

    @allowed_sources.setter
    def allowed_sources(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "allowed_sources", value)

    @property
    @pulumi.getter(name="dedicatedIpfos")
    def dedicated_ipfos(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Only attach frontend on these ip. No restriction if null. List of Ip blocks.
        """
        return pulumi.get(self, "dedicated_ipfos")

    @dedicated_ipfos.setter
    def dedicated_ipfos(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "dedicated_ipfos", value)

    @property
    @pulumi.getter(name="defaultFarmId")
    def default_farm_id(self) -> Optional[pulumi.Input[int]]:
        """
        Default TCP Farm of your frontend
        """
        return pulumi.get(self, "default_farm_id")

    @default_farm_id.setter
    def default_farm_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "default_farm_id", value)

    @property
    @pulumi.getter(name="defaultSslId")
    def default_ssl_id(self) -> Optional[pulumi.Input[int]]:
        """
        Default ssl served to your customer
        """
        return pulumi.get(self, "default_ssl_id")

    @default_ssl_id.setter
    def default_ssl_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "default_ssl_id", value)

    @property
    @pulumi.getter
    def disabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Disable your frontend. Default: 'false'
        """
        return pulumi.get(self, "disabled")

    @disabled.setter
    def disabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disabled", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        Human readable name for your frontend, this field is for you
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[str]]:
        """
        Port(s) attached to your frontend. Supports single port (numerical value),
        range (2 dash-delimited increasing ports) and comma-separated list of 'single port'
        and/or 'range'. Each port must be in the [1;49151] range
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter(name="redirectLocation")
    def redirect_location(self) -> Optional[pulumi.Input[str]]:
        """
        Redirection HTTP'
        """
        return pulumi.get(self, "redirect_location")

    @redirect_location.setter
    def redirect_location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "redirect_location", value)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> Optional[pulumi.Input[str]]:
        """
        The internal name of your IP load balancing
        """
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_name", value)

    @property
    @pulumi.getter
    def ssl(self) -> Optional[pulumi.Input[bool]]:
        """
        SSL deciphering. Default: 'false'
        """
        return pulumi.get(self, "ssl")

    @ssl.setter
    def ssl(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "ssl", value)

    @property
    @pulumi.getter
    def zone(self) -> Optional[pulumi.Input[str]]:
        """
        Zone where the frontend will be defined (ie. `gra`, `bhs` also supports `all`)
        """
        return pulumi.get(self, "zone")

    @zone.setter
    def zone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zone", value)


class IpLoadBalancingHttpFrontend(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allowed_sources: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 dedicated_ipfos: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 default_farm_id: Optional[pulumi.Input[int]] = None,
                 default_ssl_id: Optional[pulumi.Input[int]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[str]] = None,
                 redirect_location: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 ssl: Optional[pulumi.Input[bool]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates a backend http server group (frontend) to be used by loadbalancing frontend(s)

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] allowed_sources: Restrict IP Load Balancing access to these ip block. No restriction if null. List of IP blocks.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] dedicated_ipfos: Only attach frontend on these ip. No restriction if null. List of Ip blocks.
        :param pulumi.Input[int] default_farm_id: Default TCP Farm of your frontend
        :param pulumi.Input[int] default_ssl_id: Default ssl served to your customer
        :param pulumi.Input[bool] disabled: Disable your frontend. Default: 'false'
        :param pulumi.Input[str] display_name: Human readable name for your frontend, this field is for you
        :param pulumi.Input[str] port: Port(s) attached to your frontend. Supports single port (numerical value),
               range (2 dash-delimited increasing ports) and comma-separated list of 'single port'
               and/or 'range'. Each port must be in the [1;49151] range
        :param pulumi.Input[str] redirect_location: Redirection HTTP'
        :param pulumi.Input[str] service_name: The internal name of your IP load balancing
        :param pulumi.Input[bool] ssl: SSL deciphering. Default: 'false'
        :param pulumi.Input[str] zone: Zone where the frontend will be defined (ie. `gra`, `bhs` also supports `all`)
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: IpLoadBalancingHttpFrontendArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a backend http server group (frontend) to be used by loadbalancing frontend(s)

        :param str resource_name: The name of the resource.
        :param IpLoadBalancingHttpFrontendArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(IpLoadBalancingHttpFrontendArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allowed_sources: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 dedicated_ipfos: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 default_farm_id: Optional[pulumi.Input[int]] = None,
                 default_ssl_id: Optional[pulumi.Input[int]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[str]] = None,
                 redirect_location: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 ssl: Optional[pulumi.Input[bool]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = IpLoadBalancingHttpFrontendArgs.__new__(IpLoadBalancingHttpFrontendArgs)

            __props__.__dict__["allowed_sources"] = allowed_sources
            __props__.__dict__["dedicated_ipfos"] = dedicated_ipfos
            __props__.__dict__["default_farm_id"] = default_farm_id
            __props__.__dict__["default_ssl_id"] = default_ssl_id
            __props__.__dict__["disabled"] = disabled
            __props__.__dict__["display_name"] = display_name
            if port is None and not opts.urn:
                raise TypeError("Missing required property 'port'")
            __props__.__dict__["port"] = port
            __props__.__dict__["redirect_location"] = redirect_location
            if service_name is None and not opts.urn:
                raise TypeError("Missing required property 'service_name'")
            __props__.__dict__["service_name"] = service_name
            __props__.__dict__["ssl"] = ssl
            if zone is None and not opts.urn:
                raise TypeError("Missing required property 'zone'")
            __props__.__dict__["zone"] = zone
        super(IpLoadBalancingHttpFrontend, __self__).__init__(
            'ovh:index/ipLoadBalancingHttpFrontend:IpLoadBalancingHttpFrontend',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            allowed_sources: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            dedicated_ipfos: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            default_farm_id: Optional[pulumi.Input[int]] = None,
            default_ssl_id: Optional[pulumi.Input[int]] = None,
            disabled: Optional[pulumi.Input[bool]] = None,
            display_name: Optional[pulumi.Input[str]] = None,
            port: Optional[pulumi.Input[str]] = None,
            redirect_location: Optional[pulumi.Input[str]] = None,
            service_name: Optional[pulumi.Input[str]] = None,
            ssl: Optional[pulumi.Input[bool]] = None,
            zone: Optional[pulumi.Input[str]] = None) -> 'IpLoadBalancingHttpFrontend':
        """
        Get an existing IpLoadBalancingHttpFrontend resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] allowed_sources: Restrict IP Load Balancing access to these ip block. No restriction if null. List of IP blocks.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] dedicated_ipfos: Only attach frontend on these ip. No restriction if null. List of Ip blocks.
        :param pulumi.Input[int] default_farm_id: Default TCP Farm of your frontend
        :param pulumi.Input[int] default_ssl_id: Default ssl served to your customer
        :param pulumi.Input[bool] disabled: Disable your frontend. Default: 'false'
        :param pulumi.Input[str] display_name: Human readable name for your frontend, this field is for you
        :param pulumi.Input[str] port: Port(s) attached to your frontend. Supports single port (numerical value),
               range (2 dash-delimited increasing ports) and comma-separated list of 'single port'
               and/or 'range'. Each port must be in the [1;49151] range
        :param pulumi.Input[str] redirect_location: Redirection HTTP'
        :param pulumi.Input[str] service_name: The internal name of your IP load balancing
        :param pulumi.Input[bool] ssl: SSL deciphering. Default: 'false'
        :param pulumi.Input[str] zone: Zone where the frontend will be defined (ie. `gra`, `bhs` also supports `all`)
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _IpLoadBalancingHttpFrontendState.__new__(_IpLoadBalancingHttpFrontendState)

        __props__.__dict__["allowed_sources"] = allowed_sources
        __props__.__dict__["dedicated_ipfos"] = dedicated_ipfos
        __props__.__dict__["default_farm_id"] = default_farm_id
        __props__.__dict__["default_ssl_id"] = default_ssl_id
        __props__.__dict__["disabled"] = disabled
        __props__.__dict__["display_name"] = display_name
        __props__.__dict__["port"] = port
        __props__.__dict__["redirect_location"] = redirect_location
        __props__.__dict__["service_name"] = service_name
        __props__.__dict__["ssl"] = ssl
        __props__.__dict__["zone"] = zone
        return IpLoadBalancingHttpFrontend(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="allowedSources")
    def allowed_sources(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        Restrict IP Load Balancing access to these ip block. No restriction if null. List of IP blocks.
        """
        return pulumi.get(self, "allowed_sources")

    @property
    @pulumi.getter(name="dedicatedIpfos")
    def dedicated_ipfos(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        Only attach frontend on these ip. No restriction if null. List of Ip blocks.
        """
        return pulumi.get(self, "dedicated_ipfos")

    @property
    @pulumi.getter(name="defaultFarmId")
    def default_farm_id(self) -> pulumi.Output[int]:
        """
        Default TCP Farm of your frontend
        """
        return pulumi.get(self, "default_farm_id")

    @property
    @pulumi.getter(name="defaultSslId")
    def default_ssl_id(self) -> pulumi.Output[int]:
        """
        Default ssl served to your customer
        """
        return pulumi.get(self, "default_ssl_id")

    @property
    @pulumi.getter
    def disabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Disable your frontend. Default: 'false'
        """
        return pulumi.get(self, "disabled")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[Optional[str]]:
        """
        Human readable name for your frontend, this field is for you
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def port(self) -> pulumi.Output[str]:
        """
        Port(s) attached to your frontend. Supports single port (numerical value),
        range (2 dash-delimited increasing ports) and comma-separated list of 'single port'
        and/or 'range'. Each port must be in the [1;49151] range
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter(name="redirectLocation")
    def redirect_location(self) -> pulumi.Output[Optional[str]]:
        """
        Redirection HTTP'
        """
        return pulumi.get(self, "redirect_location")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Output[str]:
        """
        The internal name of your IP load balancing
        """
        return pulumi.get(self, "service_name")

    @property
    @pulumi.getter
    def ssl(self) -> pulumi.Output[Optional[bool]]:
        """
        SSL deciphering. Default: 'false'
        """
        return pulumi.get(self, "ssl")

    @property
    @pulumi.getter
    def zone(self) -> pulumi.Output[str]:
        """
        Zone where the frontend will be defined (ie. `gra`, `bhs` also supports `all`)
        """
        return pulumi.get(self, "zone")
