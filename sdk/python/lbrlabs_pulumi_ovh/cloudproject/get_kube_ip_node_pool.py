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
from ._inputs import *

__all__ = [
    'GetKubeIpNodePoolResult',
    'AwaitableGetKubeIpNodePoolResult',
    'get_kube_ip_node_pool',
    'get_kube_ip_node_pool_output',
]

@pulumi.output_type
class GetKubeIpNodePoolResult:
    """
    A collection of values returned by getKubeIpNodePool.
    """
    def __init__(__self__, anti_affinity=None, autoscale=None, available_nodes=None, created_at=None, current_nodes=None, desired_nodes=None, flavor=None, flavor_name=None, id=None, kube_id=None, max_nodes=None, min_nodes=None, monthly_billed=None, name=None, project_id=None, service_name=None, size_status=None, status=None, template=None, up_to_date_nodes=None, updated_at=None):
        if anti_affinity and not isinstance(anti_affinity, bool):
            raise TypeError("Expected argument 'anti_affinity' to be a bool")
        pulumi.set(__self__, "anti_affinity", anti_affinity)
        if autoscale and not isinstance(autoscale, bool):
            raise TypeError("Expected argument 'autoscale' to be a bool")
        pulumi.set(__self__, "autoscale", autoscale)
        if available_nodes and not isinstance(available_nodes, int):
            raise TypeError("Expected argument 'available_nodes' to be a int")
        pulumi.set(__self__, "available_nodes", available_nodes)
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
        if current_nodes and not isinstance(current_nodes, int):
            raise TypeError("Expected argument 'current_nodes' to be a int")
        pulumi.set(__self__, "current_nodes", current_nodes)
        if desired_nodes and not isinstance(desired_nodes, int):
            raise TypeError("Expected argument 'desired_nodes' to be a int")
        pulumi.set(__self__, "desired_nodes", desired_nodes)
        if flavor and not isinstance(flavor, str):
            raise TypeError("Expected argument 'flavor' to be a str")
        pulumi.set(__self__, "flavor", flavor)
        if flavor_name and not isinstance(flavor_name, str):
            raise TypeError("Expected argument 'flavor_name' to be a str")
        pulumi.set(__self__, "flavor_name", flavor_name)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if kube_id and not isinstance(kube_id, str):
            raise TypeError("Expected argument 'kube_id' to be a str")
        pulumi.set(__self__, "kube_id", kube_id)
        if max_nodes and not isinstance(max_nodes, int):
            raise TypeError("Expected argument 'max_nodes' to be a int")
        pulumi.set(__self__, "max_nodes", max_nodes)
        if min_nodes and not isinstance(min_nodes, int):
            raise TypeError("Expected argument 'min_nodes' to be a int")
        pulumi.set(__self__, "min_nodes", min_nodes)
        if monthly_billed and not isinstance(monthly_billed, bool):
            raise TypeError("Expected argument 'monthly_billed' to be a bool")
        pulumi.set(__self__, "monthly_billed", monthly_billed)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)
        if service_name and not isinstance(service_name, str):
            raise TypeError("Expected argument 'service_name' to be a str")
        pulumi.set(__self__, "service_name", service_name)
        if size_status and not isinstance(size_status, str):
            raise TypeError("Expected argument 'size_status' to be a str")
        pulumi.set(__self__, "size_status", size_status)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if template and not isinstance(template, dict):
            raise TypeError("Expected argument 'template' to be a dict")
        pulumi.set(__self__, "template", template)
        if up_to_date_nodes and not isinstance(up_to_date_nodes, int):
            raise TypeError("Expected argument 'up_to_date_nodes' to be a int")
        pulumi.set(__self__, "up_to_date_nodes", up_to_date_nodes)
        if updated_at and not isinstance(updated_at, str):
            raise TypeError("Expected argument 'updated_at' to be a str")
        pulumi.set(__self__, "updated_at", updated_at)

    @property
    @pulumi.getter(name="antiAffinity")
    def anti_affinity(self) -> bool:
        """
        (Optional) should the pool use the anti-affinity feature. Default to `false`.
        """
        return pulumi.get(self, "anti_affinity")

    @property
    @pulumi.getter
    def autoscale(self) -> bool:
        """
        (Optional) Enable auto-scaling for the pool. Default to `false`.
        """
        return pulumi.get(self, "autoscale")

    @property
    @pulumi.getter(name="availableNodes")
    def available_nodes(self) -> int:
        """
        Number of nodes which are actually ready in the pool
        """
        return pulumi.get(self, "available_nodes")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> str:
        """
        Creation date
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="currentNodes")
    def current_nodes(self) -> int:
        """
        Number of nodes present in the pool
        """
        return pulumi.get(self, "current_nodes")

    @property
    @pulumi.getter(name="desiredNodes")
    def desired_nodes(self) -> int:
        """
        Number of nodes you desire in the pool
        """
        return pulumi.get(self, "desired_nodes")

    @property
    @pulumi.getter
    def flavor(self) -> str:
        """
        Flavor name
        """
        return pulumi.get(self, "flavor")

    @property
    @pulumi.getter(name="flavorName")
    def flavor_name(self) -> str:
        """
        a valid OVHcloud public cloud flavor ID in which the nodes will be started.
        Ex: "b2-7". Changing this value recreates the resource.
        You can find the list of flavor IDs: https://www.ovhcloud.com/fr/public-cloud/prices/
        """
        return pulumi.get(self, "flavor_name")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="kubeId")
    def kube_id(self) -> str:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "kube_id")

    @property
    @pulumi.getter(name="maxNodes")
    def max_nodes(self) -> int:
        """
        maximum number of nodes allowed in the pool.
        Setting `desired_nodes` over this value will raise an error.
        """
        return pulumi.get(self, "max_nodes")

    @property
    @pulumi.getter(name="minNodes")
    def min_nodes(self) -> int:
        """
        minimum number of nodes allowed in the pool.
        Setting `desired_nodes` under this value will raise an error.
        """
        return pulumi.get(self, "min_nodes")

    @property
    @pulumi.getter(name="monthlyBilled")
    def monthly_billed(self) -> bool:
        """
        (Optional) should the nodes be billed on a monthly basis. Default to `false`.
        """
        return pulumi.get(self, "monthly_billed")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        (Optional) The name of the nodepool.
        Changing this value recreates the resource.
        Warning: "_" char is not allowed!
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> str:
        """
        Project id
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> str:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "service_name")

    @property
    @pulumi.getter(name="sizeStatus")
    def size_status(self) -> str:
        """
        Status describing the state between number of nodes wanted and available ones
        """
        return pulumi.get(self, "size_status")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        Current status
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def template(self) -> Optional['outputs.GetKubeIpNodePoolTemplateResult']:
        return pulumi.get(self, "template")

    @property
    @pulumi.getter(name="upToDateNodes")
    def up_to_date_nodes(self) -> int:
        """
        Number of nodes with the latest version installed in the pool
        """
        return pulumi.get(self, "up_to_date_nodes")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> str:
        """
        Last update date
        """
        return pulumi.get(self, "updated_at")


class AwaitableGetKubeIpNodePoolResult(GetKubeIpNodePoolResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetKubeIpNodePoolResult(
            anti_affinity=self.anti_affinity,
            autoscale=self.autoscale,
            available_nodes=self.available_nodes,
            created_at=self.created_at,
            current_nodes=self.current_nodes,
            desired_nodes=self.desired_nodes,
            flavor=self.flavor,
            flavor_name=self.flavor_name,
            id=self.id,
            kube_id=self.kube_id,
            max_nodes=self.max_nodes,
            min_nodes=self.min_nodes,
            monthly_billed=self.monthly_billed,
            name=self.name,
            project_id=self.project_id,
            service_name=self.service_name,
            size_status=self.size_status,
            status=self.status,
            template=self.template,
            up_to_date_nodes=self.up_to_date_nodes,
            updated_at=self.updated_at)


def get_kube_ip_node_pool(kube_id: Optional[str] = None,
                          name: Optional[str] = None,
                          service_name: Optional[str] = None,
                          template: Optional[pulumi.InputType['GetKubeIpNodePoolTemplateArgs']] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetKubeIpNodePoolResult:
    """
    Use this data source to get a OVHcloud Managed Kubernetes node pool.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    nodepool = ovh.CloudProject.get_kube_ip_node_pool(service_name="XXXXXX",
        kube_id="xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxx",
        name="xxxxxx")
    pulumi.export("maxNodes", nodepool.max_nodes)
    ```


    :param str kube_id: The id of the managed kubernetes cluster.
    :param str name: The name of the node pool.
    :param str service_name: The id of the public cloud project. If omitted,
           the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
    """
    __args__ = dict()
    __args__['kubeId'] = kube_id
    __args__['name'] = name
    __args__['serviceName'] = service_name
    __args__['template'] = template
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('ovh:CloudProject/getKubeIpNodePool:getKubeIpNodePool', __args__, opts=opts, typ=GetKubeIpNodePoolResult).value

    return AwaitableGetKubeIpNodePoolResult(
        anti_affinity=__ret__.anti_affinity,
        autoscale=__ret__.autoscale,
        available_nodes=__ret__.available_nodes,
        created_at=__ret__.created_at,
        current_nodes=__ret__.current_nodes,
        desired_nodes=__ret__.desired_nodes,
        flavor=__ret__.flavor,
        flavor_name=__ret__.flavor_name,
        id=__ret__.id,
        kube_id=__ret__.kube_id,
        max_nodes=__ret__.max_nodes,
        min_nodes=__ret__.min_nodes,
        monthly_billed=__ret__.monthly_billed,
        name=__ret__.name,
        project_id=__ret__.project_id,
        service_name=__ret__.service_name,
        size_status=__ret__.size_status,
        status=__ret__.status,
        template=__ret__.template,
        up_to_date_nodes=__ret__.up_to_date_nodes,
        updated_at=__ret__.updated_at)


@_utilities.lift_output_func(get_kube_ip_node_pool)
def get_kube_ip_node_pool_output(kube_id: Optional[pulumi.Input[str]] = None,
                                 name: Optional[pulumi.Input[str]] = None,
                                 service_name: Optional[pulumi.Input[str]] = None,
                                 template: Optional[pulumi.Input[Optional[pulumi.InputType['GetKubeIpNodePoolTemplateArgs']]]] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetKubeIpNodePoolResult]:
    """
    Use this data source to get a OVHcloud Managed Kubernetes node pool.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    nodepool = ovh.CloudProject.get_kube_ip_node_pool(service_name="XXXXXX",
        kube_id="xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxx",
        name="xxxxxx")
    pulumi.export("maxNodes", nodepool.max_nodes)
    ```


    :param str kube_id: The id of the managed kubernetes cluster.
    :param str name: The name of the node pool.
    :param str service_name: The id of the public cloud project. If omitted,
           the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
    """
    ...
