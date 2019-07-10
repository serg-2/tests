#!/usr/bin/python
from distutils.core import setup
from distutils.extension import Extension
from Cython.Distutils import build_ext

ext_modules = [
#    Extension("mymodule2",  ["mymodule2.py"]),
    Extension("test_functions",  ["test_functions.py"])

#   ... all your modules that need be compiled ...

]

setup(
    name = 'My Program Name',
    cmdclass = {'build_ext': build_ext},
    ext_modules = ext_modules
)

