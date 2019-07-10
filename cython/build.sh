#!/bin/bash
python compile.py build_ext --inplace

rm -rf build
rm test_functions.c 
