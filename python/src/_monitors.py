#
# This file was generated with makeClass --sdk. Do not edit it.
#
from . import session

monitorsCmd = "monitors"
monitorsPos = "addrs"
monitorsFmt = "json"
monitorsOpts = {
    "clean": {"hotkey": "", "type": "switch"},
    "delete": {"hotkey": "", "type": "switch"},
    "undelete": {"hotkey": "", "type": "switch"},
    "remove": {"hotkey": "", "type": "switch"},
    "decache": {"hotkey": "", "type": "switch"},
    "list": {"hotkey": "", "type": "switch"},
    "watch": {"hotkey": "", "type": "switch"},
    "sleep": {"hotkey": "-s", "type": "flag"},
    "firstBlock": {"hotkey": "-F", "type": "flag"},
    "lastBlock": {"hotkey": "-L", "type": "flag"},
    "fmt": {"hotkey": "-x", "type": "flag"},
    "verbose:": {"hotkey": "-v", "type": "switch"},
    "help": {"hotkey": "-h", "type": "switch"},
}

def monitors(self):
    ret = self.toUrl(monitorsCmd, monitorsPos, monitorsFmt, monitorsOpts)
    url = 'http://localhost:8080/' + ret[1]
    if ret[0] == 'json':
        return session.get(url).json()
    return session.get(url).text

