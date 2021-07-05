#!/bin/bash
find . -name '*.sh' | sed 's#.*/##' | cut -d '.' -f1 

