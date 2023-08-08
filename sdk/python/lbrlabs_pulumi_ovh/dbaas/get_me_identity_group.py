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
    'GetMeIdentityGroupResult',
    'AwaitableGetMeIdentityGroupResult',
    'get_me_identity_group',
    'get_me_identity_group_output',
]

@pulumi.output_type
class GetMeIdentityGroupResult:
    """
    A collection of values returned by getMeIdentityGroup.
    """
    def __init__(__self__, creation=None, default_group=None, description=None, id=None, last_update=None, name=None, role=None):
        if creation and not isinstance(creation, str):
            raise TypeError("Expected argument 'creation' to be a str")
        pulumi.set(__self__, "creation", creation)
        if default_group and not isinstance(default_group, bool):
            raise TypeError("Expected argument 'default_group' to be a bool")
        pulumi.set(__self__, "default_group", default_group)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if last_update and not isinstance(last_update, str):
            raise TypeError("Expected argument 'last_update' to be a str")
        pulumi.set(__self__, "last_update", last_update)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if role and not isinstance(role, str):
            raise TypeError("Expected argument 'role' to be a str")
        pulumi.set(__self__, "role", role)

    @property
    @pulumi.getter
    def creation(self) -> str:
        """
        Creation date of this group.
        """
        return pulumi.get(self, "creation")

    @property
    @pulumi.getter(name="defaultGroup")
    def default_group(self) -> bool:
        """
        Is the group a default and immutable one.
        """
        return pulumi.get(self, "default_group")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        Group description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="lastUpdate")
    def last_update(self) -> str:
        """
        Date of the last update of this group.
        """
        return pulumi.get(self, "last_update")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def role(self) -> str:
        """
        Role associated with the group. Valid roles are ADMIN, REGULAR, UNPRIVILEGED, and NONE.
        """
        return pulumi.get(self, "role")


class AwaitableGetMeIdentityGroupResult(GetMeIdentityGroupResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetMeIdentityGroupResult(
            creation=self.creation,
            default_group=self.default_group,
            description=self.description,
            id=self.id,
            last_update=self.last_update,
            name=self.name,
            role=self.role)


def get_me_identity_group(name: Optional[str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetMeIdentityGroupResult:
    """
    Use this data source to retrieve information about an identity group.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    my_group = ovh.Dbaas.get_me_identity_group(name="my_group_name")
    ```


    :param str name: Group name.
    """
    __args__ = dict()
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('ovh:Dbaas/getMeIdentityGroup:getMeIdentityGroup', __args__, opts=opts, typ=GetMeIdentityGroupResult).value

    return AwaitableGetMeIdentityGroupResult(
        creation=pulumi.get(__ret__, 'creation'),
        default_group=pulumi.get(__ret__, 'default_group'),
        description=pulumi.get(__ret__, 'description'),
        id=pulumi.get(__ret__, 'id'),
        last_update=pulumi.get(__ret__, 'last_update'),
        name=pulumi.get(__ret__, 'name'),
        role=pulumi.get(__ret__, 'role'))


@_utilities.lift_output_func(get_me_identity_group)
def get_me_identity_group_output(name: Optional[pulumi.Input[str]] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetMeIdentityGroupResult]:
    """
    Use this data source to retrieve information about an identity group.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    my_group = ovh.Dbaas.get_me_identity_group(name="my_group_name")
    ```


    :param str name: Group name.
    """
    ...
