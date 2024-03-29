# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

import types

__config__ = pulumi.Config('powerdns')


class _ExportableConfig(types.ModuleType):
    @property
    def insecure(self) -> Optional[bool]:
        """
        Explicitly allow the provider to perform "insecure" SSL requests. If omitted, default value is "false"
        """
        return __config__.get_bool('insecure')

    @property
    def key(self) -> str:
        """
        The access key for API operations
        """
        return __config__.get('key') or (_utilities.get_env('POWERDNS_KEY') or '')

    @property
    def logging(self) -> Optional[bool]:
        return __config__.get_bool('logging')

    @property
    def url(self) -> str:
        """
        The api endpoint of the powerdns server
        """
        return __config__.get('url') or (_utilities.get_env('POWERDNS_URL') or '')

    @property
    def version(self) -> Optional[str]:
        return __config__.get('version')

