# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from . import _utilities
import typing
# Export this package's modules as members:
from .provider import *
from .record import *
from .zone import *

# Make subpackages available:
if typing.TYPE_CHECKING:
    import pulumi_powerdns.config as __config
    config = __config
else:
    config = _utilities.lazy_import('pulumi_powerdns.config')

_utilities.register(
    resource_modules="""
[
 {
  "pkg": "powerdns",
  "mod": "index",
  "fqn": "pulumi_powerdns",
  "classes": {
   "powerdns:index:Record": "Record",
   "powerdns:index:Zone": "Zone"
  }
 }
]
""",
    resource_packages="""
[
 {
  "pkg": "powerdns",
  "token": "pulumi:providers:powerdns",
  "fqn": "pulumi_powerdns",
  "class": "Provider"
 }
]
"""
)
