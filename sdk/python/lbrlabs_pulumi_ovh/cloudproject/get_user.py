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
    'GetUserResult',
    'AwaitableGetUserResult',
    'get_user',
    'get_user_output',
]

@pulumi.output_type
class GetUserResult:
    """
    A collection of values returned by getUser.
    """
    def __init__(__self__, creation_date=None, description=None, id=None, roles=None, service_name=None, status=None, user_id=None, username=None):
        if creation_date and not isinstance(creation_date, str):
            raise TypeError("Expected argument 'creation_date' to be a str")
        pulumi.set(__self__, "creation_date", creation_date)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if roles and not isinstance(roles, list):
            raise TypeError("Expected argument 'roles' to be a list")
        pulumi.set(__self__, "roles", roles)
        if service_name and not isinstance(service_name, str):
            raise TypeError("Expected argument 'service_name' to be a str")
        pulumi.set(__self__, "service_name", service_name)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if user_id and not isinstance(user_id, str):
            raise TypeError("Expected argument 'user_id' to be a str")
        pulumi.set(__self__, "user_id", user_id)
        if username and not isinstance(username, str):
            raise TypeError("Expected argument 'username' to be a str")
        pulumi.set(__self__, "username", username)

    @property
    @pulumi.getter(name="creationDate")
    def creation_date(self) -> str:
        """
        the date the user was created.
        """
        return pulumi.get(self, "creation_date")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        description of the role
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
    @pulumi.getter
    def roles(self) -> Sequence['outputs.GetUserRoleResult']:
        """
        A list of roles associated with the user.
        """
        return pulumi.get(self, "roles")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> str:
        return pulumi.get(self, "service_name")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        the status of the user. should be normally set to 'ok'.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> str:
        return pulumi.get(self, "user_id")

    @property
    @pulumi.getter
    def username(self) -> str:
        """
        the username generated for the user. This username can be used with
        the Openstack API.
        """
        return pulumi.get(self, "username")


class AwaitableGetUserResult(GetUserResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetUserResult(
            creation_date=self.creation_date,
            description=self.description,
            id=self.id,
            roles=self.roles,
            service_name=self.service_name,
            status=self.status,
            user_id=self.user_id,
            username=self.username)


def get_user(service_name: Optional[str] = None,
             user_id: Optional[str] = None,
             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetUserResult:
    """
    Get the user details of a previously created public cloud project user.


    :param str service_name: The ID of the public cloud project. If omitted,
           the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
    :param str user_id: The ID of a public cloud project's user.
    """
    __args__ = dict()
    __args__['serviceName'] = service_name
    __args__['userId'] = user_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('ovh:CloudProject/getUser:getUser', __args__, opts=opts, typ=GetUserResult).value

    return AwaitableGetUserResult(
        creation_date=__ret__.creation_date,
        description=__ret__.description,
        id=__ret__.id,
        roles=__ret__.roles,
        service_name=__ret__.service_name,
        status=__ret__.status,
        user_id=__ret__.user_id,
        username=__ret__.username)


@_utilities.lift_output_func(get_user)
def get_user_output(service_name: Optional[pulumi.Input[str]] = None,
                    user_id: Optional[pulumi.Input[str]] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetUserResult]:
    """
    Get the user details of a previously created public cloud project user.


    :param str service_name: The ID of the public cloud project. If omitted,
           the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
    :param str user_id: The ID of a public cloud project's user.
    """
    ...
