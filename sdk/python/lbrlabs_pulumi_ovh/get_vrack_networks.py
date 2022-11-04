# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetVrackNetworksResult',
    'AwaitableGetVrackNetworksResult',
    'get_vrack_networks',
    'get_vrack_networks_output',
]

@pulumi.output_type
class GetVrackNetworksResult:
    """
    A collection of values returned by getVrackNetworks.
    """
    def __init__(__self__, id=None, results=None, service_name=None, subnet=None, vlan_id=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if results and not isinstance(results, list):
            raise TypeError("Expected argument 'results' to be a list")
        pulumi.set(__self__, "results", results)
        if service_name and not isinstance(service_name, str):
            raise TypeError("Expected argument 'service_name' to be a str")
        pulumi.set(__self__, "service_name", service_name)
        if subnet and not isinstance(subnet, str):
            raise TypeError("Expected argument 'subnet' to be a str")
        pulumi.set(__self__, "subnet", subnet)
        if vlan_id and not isinstance(vlan_id, int):
            raise TypeError("Expected argument 'vlan_id' to be a int")
        pulumi.set(__self__, "vlan_id", vlan_id)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def results(self) -> Sequence[int]:
        """
        The list of vrack network ids.
        """
        return pulumi.get(self, "results")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> str:
        return pulumi.get(self, "service_name")

    @property
    @pulumi.getter
    def subnet(self) -> Optional[str]:
        return pulumi.get(self, "subnet")

    @property
    @pulumi.getter(name="vlanId")
    def vlan_id(self) -> Optional[int]:
        return pulumi.get(self, "vlan_id")


class AwaitableGetVrackNetworksResult(GetVrackNetworksResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVrackNetworksResult(
            id=self.id,
            results=self.results,
            service_name=self.service_name,
            subnet=self.subnet,
            vlan_id=self.vlan_id)


def get_vrack_networks(service_name: Optional[str] = None,
                       subnet: Optional[str] = None,
                       vlan_id: Optional[int] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetVrackNetworksResult:
    """
    Use this data source to get the list of Vrack network ids available for your IPLoadbalancer associated with your OVHcloud account.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    lb_networks = ovh.get_vrack_networks(service_name="XXXXXX",
        subnet="10.0.0.0/24")
    ```


    :param str service_name: The internal name of your IP load balancing
    :param str subnet: Filters networks on the subnet.
    :param int vlan_id: Filters networks on the vlan id.
    """
    __args__ = dict()
    __args__['serviceName'] = service_name
    __args__['subnet'] = subnet
    __args__['vlanId'] = vlan_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('ovh:index/getVrackNetworks:getVrackNetworks', __args__, opts=opts, typ=GetVrackNetworksResult).value

    return AwaitableGetVrackNetworksResult(
        id=__ret__.id,
        results=__ret__.results,
        service_name=__ret__.service_name,
        subnet=__ret__.subnet,
        vlan_id=__ret__.vlan_id)


@_utilities.lift_output_func(get_vrack_networks)
def get_vrack_networks_output(service_name: Optional[pulumi.Input[str]] = None,
                              subnet: Optional[pulumi.Input[Optional[str]]] = None,
                              vlan_id: Optional[pulumi.Input[Optional[int]]] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetVrackNetworksResult]:
    """
    Use this data source to get the list of Vrack network ids available for your IPLoadbalancer associated with your OVHcloud account.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    lb_networks = ovh.get_vrack_networks(service_name="XXXXXX",
        subnet="10.0.0.0/24")
    ```


    :param str service_name: The internal name of your IP load balancing
    :param str subnet: Filters networks on the subnet.
    :param int vlan_id: Filters networks on the vlan id.
    """
    ...