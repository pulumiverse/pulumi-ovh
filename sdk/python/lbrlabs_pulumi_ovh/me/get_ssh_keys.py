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
    'GetSshKeysResult',
    'AwaitableGetSshKeysResult',
    'get_ssh_keys',
]

@pulumi.output_type
class GetSshKeysResult:
    """
    A collection of values returned by getSshKeys.
    """
    def __init__(__self__, id=None, names=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if names and not isinstance(names, list):
            raise TypeError("Expected argument 'names' to be a list")
        pulumi.set(__self__, "names", names)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def names(self) -> Sequence[str]:
        """
        The list of the names of all the SSH keys.
        """
        return pulumi.get(self, "names")


class AwaitableGetSshKeysResult(GetSshKeysResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSshKeysResult(
            id=self.id,
            names=self.names)


def get_ssh_keys(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSshKeysResult:
    """
    Use this data source to retrieve list of names of the account's SSH keys.
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('ovh:Me/getSshKeys:getSshKeys', __args__, opts=opts, typ=GetSshKeysResult).value

    return AwaitableGetSshKeysResult(
        id=__ret__.id,
        names=__ret__.names)