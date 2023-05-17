#
# This file was generated with makeClass --sdk. Do not edit it.
#
from . import session

listCmd = "list"
listPos = "addrs"
listFmt = "json"
listOpts = {
    "count": {"hotkey": "-U", "type": "switch"},
    "appearances": {"hotkey": "-p", "type": "switch"},
    "silent": {"hotkey": "", "type": "switch"},
    "noZero": {"hotkey": "-n", "type": "switch"},
    "unripe": {"hotkey": "-u", "type": "switch"},
    "firstRecord": {"hotkey": "-c", "type": "flag"},
    "maxRecords": {"hotkey": "-e", "type": "flag"},
    "firstBlock": {"hotkey": "-F", "type": "flag"},
    "lastBlock": {"hotkey": "-L", "type": "flag"},
    "bounds": {"hotkey": "", "type": "switch"},
    "fmt": {"hotkey": "-x", "type": "flag"},
    "verbose:": {"hotkey": "-v", "type": "switch"},
    "help": {"hotkey": "-h", "type": "switch"},
}

def list(self):
    ret = self.toUrl(listCmd, listPos, listFmt, listOpts)
    url = 'http://localhost:8080/' + ret[1]
    if ret[0] == 'json':
        return session.get(url).json()
    return session.get(url).text

