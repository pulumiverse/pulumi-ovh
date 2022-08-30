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
    'GetIpLoadbalancingVrackNetworkResult',
    'AwaitableGetIpLoadbalancingVrackNetworkResult',
    'get_ip_loadbalancing_vrack_network',
    'get_ip_loadbalancing_vrack_network_output',
]

@pulumi.output_type
class GetIpLoadbalancingVrackNetworkResult:
    """
    A collection of values returned by getIpLoadbalancingVrackNetwork.
    """
    def __init__(__self__, display_name=None, id=None, nat_ip=None, service_name=None, subnet=None, vlan=None, vrack_network_id=None):
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if nat_ip and not isinstance(nat_ip, str):
            raise TypeError("Expected argument 'nat_ip' to be a str")
        pulumi.set(__self__, "nat_ip", nat_ip)
        if service_name and not isinstance(service_name, str):
            raise TypeError("Expected argument 'service_name' to be a str")
        pulumi.set(__self__, "service_name", service_name)
        if subnet and not isinstance(subnet, str):
            raise TypeError("Expected argument 'subnet' to be a str")
        pulumi.set(__self__, "subnet", subnet)
        if vlan and not isinstance(vlan, int):
            raise TypeError("Expected argument 'vlan' to be a int")
        pulumi.set(__self__, "vlan", vlan)
        if vrack_network_id and not isinstance(vrack_network_id, int):
            raise TypeError("Expected argument 'vrack_network_id' to be a int")
        pulumi.set(__self__, "vrack_network_id", vrack_network_id)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        """
        Human readable name for your vrack network
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="natIp")
    def nat_ip(self) -> str:
        """
        An IP block used as a pool of IPs by this Load Balancer to connect to the servers in this private network. The blck must be in the private network and reserved for the Load Balancer
        """
        return pulumi.get(self, "nat_ip")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> str:
        return pulumi.get(self, "service_name")

    @property
    @pulumi.getter
    def subnet(self) -> str:
        """
        IP block of the private network in the vRack
        """
        return pulumi.get(self, "subnet")

    @property
    @pulumi.getter
    def vlan(self) -> int:
        """
        VLAN of the private network in the vRack. 0 if the private network is not in a VLAN
        """
        return pulumi.get(self, "vlan")

    @property
    @pulumi.getter(name="vrackNetworkId")
    def vrack_network_id(self) -> int:
        return pulumi.get(self, "vrack_network_id")


class AwaitableGetIpLoadbalancingVrackNetworkResult(GetIpLoadbalancingVrackNetworkResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetIpLoadbalancingVrackNetworkResult(
            display_name=self.display_name,
            id=self.id,
            nat_ip=self.nat_ip,
            service_name=self.service_name,
            subnet=self.subnet,
            vlan=self.vlan,
            vrack_network_id=self.vrack_network_id)


def get_ip_loadbalancing_vrack_network(service_name: Optional[str] = None,
                                       vrack_network_id: Optional[int] = None,
                                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetIpLoadbalancingVrackNetworkResult:
    """
    Use this data source to get the details of Vrack network available for your IPLoadbalancer associated with your OVH account.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    lb_network = ovh.get_ip_loadbalancing_vrack_network(service_name="XXXXXX",
        vrack_network_id="yyy")
    ```


    :param str service_name: The internal name of your IP load balancing
    :param int vrack_network_id: Internal Load Balancer identifier of the vRack private network
    """
    __args__ = dict()
    __args__['serviceName'] = service_name
    __args__['vrackNetworkId'] = vrack_network_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('ovh:index/getIpLoadbalancingVrackNetwork:getIpLoadbalancingVrackNetwork', __args__, opts=opts, typ=GetIpLoadbalancingVrackNetworkResult).value

    return AwaitableGetIpLoadbalancingVrackNetworkResult(
        display_name=__ret__.display_name,
        id=__ret__.id,
        nat_ip=__ret__.nat_ip,
        service_name=__ret__.service_name,
        subnet=__ret__.subnet,
        vlan=__ret__.vlan,
        vrack_network_id=__ret__.vrack_network_id)


@_utilities.lift_output_func(get_ip_loadbalancing_vrack_network)
def get_ip_loadbalancing_vrack_network_output(service_name: Optional[pulumi.Input[str]] = None,
                                              vrack_network_id: Optional[pulumi.Input[int]] = None,
                                              opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetIpLoadbalancingVrackNetworkResult]:
    """
    Use this data source to get the details of Vrack network available for your IPLoadbalancer associated with your OVH account.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    lb_network = ovh.get_ip_loadbalancing_vrack_network(service_name="XXXXXX",
        vrack_network_id="yyy")
    ```


    :param str service_name: The internal name of your IP load balancing
    :param int vrack_network_id: Internal Load Balancer identifier of the vRack private network
    """
    ...
