#!/usr/bin/env python

# get the unicode data .txt file from 
# http://source.icu-project.org/repos/icu/icu/trunk/source/data/unidata/UnicodeData.txt

# pass this file to this script as an argument 
# to generate unicode.go

import fileinput
import re

emoticon_match = re.compile('^1F[3456][0-9ABCDEF]{2}')


print """package main

var codepoints = map[string]string{
"""
for line in fileinput.input():
    if emoticon_match.match(line) :
        (cp,desc) = line.split(';')[0:2]
        desc = desc.lower().replace(' ','-')
        print '\t"{0}" : "\U{1:0>8}",'.format(desc,cp)
print "}"
