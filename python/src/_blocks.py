# Copyright 2016, 2024 The TrueBlocks Authors. All rights reserved.
# Use of this source code is governed by a license that can
# be found in the LICENSE file.

"""
This file was auto generated. DO NOT EDIT.
"""

from . import session

blocksCmd = "blocks"
blocksPos = "blocks"
blocksFmt = "json"
blocksOpts = {
    "hashes": {"hotkey": "-e", "type": "switch"},
    "uncles": {"hotkey": "-c", "type": "switch"},
    "traces": {"hotkey": "-t", "type": "switch"},
    "uniq": {"hotkey": "-u", "type": "switch"},
    "flow": {"hotkey": "-f", "type": "flag"},
    "logs": {"hotkey": "-l", "type": "switch"},
    "emitter": {"hotkey": "-m", "type": "flag"},
    "topic": {"hotkey": "-B", "type": "flag"},
    "withdrawals": {"hotkey": "-i", "type": "switch"},
    "articulate": {"hotkey": "-a", "type": "switch"},
    "count": {"hotkey": "-U", "type": "switch"},
    "cacheTxs": {"hotkey": "-X", "type": "switch"},
    "cacheTraces": {"hotkey": "-R", "type": "switch"},
    "chain": {"hotkey": "", "type": "flag"},
    "noHeader": {"hotkey": "", "type": "switch"},
    "cache": {"hotkey": "-o", "type": "switch"},
    "decache": {"hotkey": "-D", "type": "switch"},
    "ether": {"hotkey": "-H", "type": "switch"},
    "fmt": {"hotkey": "-x", "type": "flag"},
}

def blocks(self):
    ret = self.toUrl(blocksCmd, blocksPos, blocksFmt, blocksOpts)
    url = 'http://localhost:8080/' + ret[1]
    if ret[0] == 'json':
        return session.get(url).json()
    return session.get(url).text
