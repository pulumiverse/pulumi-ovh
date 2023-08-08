# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'GetCapabilitiesResult',
    'AwaitableGetCapabilitiesResult',
    'get_capabilities',
    'get_capabilities_output',
]

@pulumi.output_type
class GetCapabilitiesResult:
    """
    A collection of values returned by getCapabilities.
    """
    def __init__(__self__, engines=None, flavors=None, id=None, options=None, plans=None, service_name=None):
        if engines and not isinstance(engines, list):
            raise TypeError("Expected argument 'engines' to be a list")
        pulumi.set(__self__, "engines", engines)
        if flavors and not isinstance(flavors, list):
            raise TypeError("Expected argument 'flavors' to be a list")
        pulumi.set(__self__, "flavors", flavors)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if options and not isinstance(options, list):
            raise TypeError("Expected argument 'options' to be a list")
        pulumi.set(__self__, "options", options)
        if plans and not isinstance(plans, list):
            raise TypeError("Expected argument 'plans' to be a list")
        pulumi.set(__self__, "plans", plans)
        if service_name and not isinstance(service_name, str):
            raise TypeError("Expected argument 'service_name' to be a str")
        pulumi.set(__self__, "service_name", service_name)

    @property
    @pulumi.getter
    def engines(self) -> Sequence['outputs.GetCapabilitiesEngineResult']:
        """
        Database engines available.
        """
        return pulumi.get(self, "engines")

    @property
    @pulumi.getter
    def flavors(self) -> Sequence['outputs.GetCapabilitiesFlavorResult']:
        """
        Flavors available.
        """
        return pulumi.get(self, "flavors")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def options(self) -> Sequence['outputs.GetCapabilitiesOptionResult']:
        """
        Options available.
        """
        return pulumi.get(self, "options")

    @property
    @pulumi.getter
    def plans(self) -> Sequence['outputs.GetCapabilitiesPlanResult']:
        """
        Plans available.
        """
        return pulumi.get(self, "plans")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> str:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "service_name")


class AwaitableGetCapabilitiesResult(GetCapabilitiesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetCapabilitiesResult(
            engines=self.engines,
            flavors=self.flavors,
            id=self.id,
            options=self.options,
            plans=self.plans,
            service_name=self.service_name)


def get_capabilities(service_name: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetCapabilitiesResult:
    """
    Use this data source to get information about capabilities of a public cloud project.


    :param str service_name: The id of the public cloud project. If omitted,
           the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
    """
    __args__ = dict()
    __args__['serviceName'] = service_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('ovh:CloudProjectDatabase/getCapabilities:getCapabilities', __args__, opts=opts, typ=GetCapabilitiesResult).value

    return AwaitableGetCapabilitiesResult(
        engines=pulumi.get(__ret__, 'engines'),
        flavors=pulumi.get(__ret__, 'flavors'),
        id=pulumi.get(__ret__, 'id'),
        options=pulumi.get(__ret__, 'options'),
        plans=pulumi.get(__ret__, 'plans'),
        service_name=pulumi.get(__ret__, 'service_name'))


@_utilities.lift_output_func(get_capabilities)
def get_capabilities_output(service_name: Optional[pulumi.Input[str]] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetCapabilitiesResult]:
    """
    Use this data source to get information about capabilities of a public cloud project.


    :param str service_name: The id of the public cloud project. If omitted,
           the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
    """
    ...
