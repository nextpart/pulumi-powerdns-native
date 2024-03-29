# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

insecure: Optional[bool]
"""
Explicitly allow the provider to perform "insecure" SSL requests. If omitted, default value is "false"
"""

key: str
"""
The access key for API operations
"""

logging: Optional[bool]

url: str
"""
The api endpoint of the powerdns server
"""

version: Optional[str]

