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
    'GetMePaymentmeanBankAccountResult',
    'AwaitableGetMePaymentmeanBankAccountResult',
    'get_me_paymentmean_bank_account',
    'get_me_paymentmean_bank_account_output',
]

@pulumi.output_type
class GetMePaymentmeanBankAccountResult:
    """
    A collection of values returned by getMePaymentmeanBankAccount.
    """
    def __init__(__self__, default=None, description=None, description_regexp=None, id=None, state=None, use_default=None, use_oldest=None):
        if default and not isinstance(default, bool):
            raise TypeError("Expected argument 'default' to be a bool")
        pulumi.set(__self__, "default", default)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if description_regexp and not isinstance(description_regexp, str):
            raise TypeError("Expected argument 'description_regexp' to be a str")
        pulumi.set(__self__, "description_regexp", description_regexp)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)
        if use_default and not isinstance(use_default, bool):
            raise TypeError("Expected argument 'use_default' to be a bool")
        pulumi.set(__self__, "use_default", use_default)
        if use_oldest and not isinstance(use_oldest, bool):
            raise TypeError("Expected argument 'use_oldest' to be a bool")
        pulumi.set(__self__, "use_oldest", use_oldest)

    @property
    @pulumi.getter
    def default(self) -> bool:
        """
        a boolean which tells if the retrieved bank account
        is marked as the default payment mean
        """
        return pulumi.get(self, "default")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        the description attribute of the bank account
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="descriptionRegexp")
    def description_regexp(self) -> Optional[str]:
        return pulumi.get(self, "description_regexp")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def state(self) -> str:
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="useDefault")
    def use_default(self) -> Optional[bool]:
        return pulumi.get(self, "use_default")

    @property
    @pulumi.getter(name="useOldest")
    def use_oldest(self) -> Optional[bool]:
        return pulumi.get(self, "use_oldest")


class AwaitableGetMePaymentmeanBankAccountResult(GetMePaymentmeanBankAccountResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetMePaymentmeanBankAccountResult(
            default=self.default,
            description=self.description,
            description_regexp=self.description_regexp,
            id=self.id,
            state=self.state,
            use_default=self.use_default,
            use_oldest=self.use_oldest)


def get_me_paymentmean_bank_account(description_regexp: Optional[str] = None,
                                    state: Optional[str] = None,
                                    use_default: Optional[bool] = None,
                                    use_oldest: Optional[bool] = None,
                                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetMePaymentmeanBankAccountResult:
    """
    Use this data source to retrieve information about a bank account
    payment mean associated with an OVH account.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    ba = ovh.get_me_paymentmean_bank_account(use_default=True)
    ```


    :param str description_regexp: a regexp used to filter bank accounts 
           on their `description` attributes.
    :param str state: Filter bank accounts on their `state` attribute.
           Can be "blockedForIncidents", "valid", "pendingValidation"
    :param bool use_default: Retrieve bank account marked as default payment mean.
    :param bool use_oldest: Retrieve oldest bank account.
           project.
    """
    __args__ = dict()
    __args__['descriptionRegexp'] = description_regexp
    __args__['state'] = state
    __args__['useDefault'] = use_default
    __args__['useOldest'] = use_oldest
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('ovh:index/getMePaymentmeanBankAccount:getMePaymentmeanBankAccount', __args__, opts=opts, typ=GetMePaymentmeanBankAccountResult).value

    return AwaitableGetMePaymentmeanBankAccountResult(
        default=__ret__.default,
        description=__ret__.description,
        description_regexp=__ret__.description_regexp,
        id=__ret__.id,
        state=__ret__.state,
        use_default=__ret__.use_default,
        use_oldest=__ret__.use_oldest)


@_utilities.lift_output_func(get_me_paymentmean_bank_account)
def get_me_paymentmean_bank_account_output(description_regexp: Optional[pulumi.Input[Optional[str]]] = None,
                                           state: Optional[pulumi.Input[Optional[str]]] = None,
                                           use_default: Optional[pulumi.Input[Optional[bool]]] = None,
                                           use_oldest: Optional[pulumi.Input[Optional[bool]]] = None,
                                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetMePaymentmeanBankAccountResult]:
    """
    Use this data source to retrieve information about a bank account
    payment mean associated with an OVH account.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    ba = ovh.get_me_paymentmean_bank_account(use_default=True)
    ```


    :param str description_regexp: a regexp used to filter bank accounts 
           on their `description` attributes.
    :param str state: Filter bank accounts on their `state` attribute.
           Can be "blockedForIncidents", "valid", "pendingValidation"
    :param bool use_default: Retrieve bank account marked as default payment mean.
    :param bool use_oldest: Retrieve oldest bank account.
           project.
    """
    ...
