# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['VrackDedicatedServerInterfaceArgs', 'VrackDedicatedServerInterface']

@pulumi.input_type
class VrackDedicatedServerInterfaceArgs:
    def __init__(__self__, *,
                 interface_id: pulumi.Input[str],
                 service_name: pulumi.Input[str]):
        """
        The set of arguments for constructing a VrackDedicatedServerInterface resource.
        :param pulumi.Input[str] interface_id: The id of dedicated server network interface.
        :param pulumi.Input[str] service_name: The id of the vrack. If omitted,
               the `OVH_VRACK_SERVICE` environment variable is used.
        """
        pulumi.set(__self__, "interface_id", interface_id)
        pulumi.set(__self__, "service_name", service_name)

    @property
    @pulumi.getter(name="interfaceId")
    def interface_id(self) -> pulumi.Input[str]:
        """
        The id of dedicated server network interface.
        """
        return pulumi.get(self, "interface_id")

    @interface_id.setter
    def interface_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "interface_id", value)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Input[str]:
        """
        The id of the vrack. If omitted,
        the `OVH_VRACK_SERVICE` environment variable is used.
        """
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "service_name", value)


@pulumi.input_type
class _VrackDedicatedServerInterfaceState:
    def __init__(__self__, *,
                 interface_id: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering VrackDedicatedServerInterface resources.
        :param pulumi.Input[str] interface_id: The id of dedicated server network interface.
        :param pulumi.Input[str] service_name: The id of the vrack. If omitted,
               the `OVH_VRACK_SERVICE` environment variable is used.
        """
        if interface_id is not None:
            pulumi.set(__self__, "interface_id", interface_id)
        if service_name is not None:
            pulumi.set(__self__, "service_name", service_name)

    @property
    @pulumi.getter(name="interfaceId")
    def interface_id(self) -> Optional[pulumi.Input[str]]:
        """
        The id of dedicated server network interface.
        """
        return pulumi.get(self, "interface_id")

    @interface_id.setter
    def interface_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "interface_id", value)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> Optional[pulumi.Input[str]]:
        """
        The id of the vrack. If omitted,
        the `OVH_VRACK_SERVICE` environment variable is used.
        """
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_name", value)


class VrackDedicatedServerInterface(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 interface_id: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Attach a Dedicated Server Network Interface to a VRack.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_ovh as ovh

        vdsi = ovh.VrackDedicatedServerInterface("vdsi",
            interface_id="67890",
            service_name="12345")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] interface_id: The id of dedicated server network interface.
        :param pulumi.Input[str] service_name: The id of the vrack. If omitted,
               the `OVH_VRACK_SERVICE` environment variable is used.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: VrackDedicatedServerInterfaceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Attach a Dedicated Server Network Interface to a VRack.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_ovh as ovh

        vdsi = ovh.VrackDedicatedServerInterface("vdsi",
            interface_id="67890",
            service_name="12345")
        ```

        :param str resource_name: The name of the resource.
        :param VrackDedicatedServerInterfaceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(VrackDedicatedServerInterfaceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 interface_id: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = VrackDedicatedServerInterfaceArgs.__new__(VrackDedicatedServerInterfaceArgs)

            if interface_id is None and not opts.urn:
                raise TypeError("Missing required property 'interface_id'")
            __props__.__dict__["interface_id"] = interface_id
            if service_name is None and not opts.urn:
                raise TypeError("Missing required property 'service_name'")
            __props__.__dict__["service_name"] = service_name
        super(VrackDedicatedServerInterface, __self__).__init__(
            'ovh:index/vrackDedicatedServerInterface:VrackDedicatedServerInterface',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            interface_id: Optional[pulumi.Input[str]] = None,
            service_name: Optional[pulumi.Input[str]] = None) -> 'VrackDedicatedServerInterface':
        """
        Get an existing VrackDedicatedServerInterface resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] interface_id: The id of dedicated server network interface.
        :param pulumi.Input[str] service_name: The id of the vrack. If omitted,
               the `OVH_VRACK_SERVICE` environment variable is used.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _VrackDedicatedServerInterfaceState.__new__(_VrackDedicatedServerInterfaceState)

        __props__.__dict__["interface_id"] = interface_id
        __props__.__dict__["service_name"] = service_name
        return VrackDedicatedServerInterface(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="interfaceId")
    def interface_id(self) -> pulumi.Output[str]:
        """
        The id of dedicated server network interface.
        """
        return pulumi.get(self, "interface_id")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Output[str]:
        """
        The id of the vrack. If omitted,
        the `OVH_VRACK_SERVICE` environment variable is used.
        """
        return pulumi.get(self, "service_name")
