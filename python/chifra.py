#!/usr/bin/env python

import pprint
import sys
from api import chifra

obj = chifra().dispatch()
pprint.pprint(obj)
