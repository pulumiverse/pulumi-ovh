# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetM3dbNamespacesResult',
    'AwaitableGetM3dbNamespacesResult',
    'get_m3db_namespaces',
    'get_m3db_namespaces_output',
]

@pulumi.output_type
class GetM3dbNamespacesResult:
    """
    A collection of values returned by getM3dbNamespaces.
    """
    def __init__(__self__, cluster_id=None, id=None, namespace_ids=None, service_name=None):
        if cluster_id and not isinstance(cluster_id, str):
            raise TypeError("Expected argument 'cluster_id' to be a str")
        pulumi.set(__self__, "cluster_id", cluster_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if namespace_ids and not isinstance(namespace_ids, list):
            raise TypeError("Expected argument 'namespace_ids' to be a list")
        pulumi.set(__self__, "namespace_ids", namespace_ids)
        if service_name and not isinstance(service_name, str):
            raise TypeError("Expected argument 'service_name' to be a str")
        pulumi.set(__self__, "service_name", service_name)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> str:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="namespaceIds")
    def namespace_ids(self) -> Sequence[str]:
        """
        The list of namespaces ids of the M3DB cluster associated with the project.
        """
        return pulumi.get(self, "namespace_ids")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> str:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "service_name")


class AwaitableGetM3dbNamespacesResult(GetM3dbNamespacesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetM3dbNamespacesResult(
            cluster_id=self.cluster_id,
            id=self.id,
            namespace_ids=self.namespace_ids,
            service_name=self.service_name)


def get_m3db_namespaces(cluster_id: Optional[str] = None,
                        service_name: Optional[str] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetM3dbNamespacesResult:
    """
    Use this data source to get the list of namespaces of a M3DB cluster associated with a public cloud project.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    namespaces = ovh.CloudProject.get_m3db_namespaces(service_name="XXX",
        cluster_id="YYY")
    pulumi.export("namespaceIds", namespaces.namespace_ids)
    ```


    :param str cluster_id: Cluster ID
    :param str service_name: The id of the public cloud project. If omitted,
           the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
    """
    __args__ = dict()
    __args__['clusterId'] = cluster_id
    __args__['serviceName'] = service_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('ovh:CloudProject/getM3dbNamespaces:getM3dbNamespaces', __args__, opts=opts, typ=GetM3dbNamespacesResult).value

    return AwaitableGetM3dbNamespacesResult(
        cluster_id=__ret__.cluster_id,
        id=__ret__.id,
        namespace_ids=__ret__.namespace_ids,
        service_name=__ret__.service_name)


@_utilities.lift_output_func(get_m3db_namespaces)
def get_m3db_namespaces_output(cluster_id: Optional[pulumi.Input[str]] = None,
                               service_name: Optional[pulumi.Input[str]] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetM3dbNamespacesResult]:
    """
    Use this data source to get the list of namespaces of a M3DB cluster associated with a public cloud project.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    namespaces = ovh.CloudProject.get_m3db_namespaces(service_name="XXX",
        cluster_id="YYY")
    pulumi.export("namespaceIds", namespaces.namespace_ids)
    ```


    :param str cluster_id: Cluster ID
    :param str service_name: The id of the public cloud project. If omitted,
           the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
    """
    ...
