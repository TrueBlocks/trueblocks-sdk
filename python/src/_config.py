#
# This file was generated with makeClass --sdk. Do not edit it.
#
from . import session

configCmd = "config"
configPos = "modes"
configFmt = "json"
configOpts = {
    "module": {"hotkey": "", "type": "flag"},
    "details": {"hotkey": "-d", "type": "switch"},
    "types": {"hotkey": "-t", "type": "flag"},
    "depth": {"hotkey": "-p", "type": "flag"},
    "terse": {"hotkey": "-e", "type": "switch"},
    "firstBlock": {"hotkey": "-F", "type": "flag"},
    "lastBlock": {"hotkey": "-L", "type": "flag"},
    "fmt": {"hotkey": "-x", "type": "flag"},
    "verbose:": {"hotkey": "-v", "type": "switch"},
    "help": {"hotkey": "-h", "type": "switch"},
}

def config(self):
    ret = self.toUrl(configCmd, configPos, configFmt, configOpts)
    url = 'http://localhost:8080/' + ret[1]
    if ret[0] == 'json':
        return session.get(url).json()
    return session.get(url).text

