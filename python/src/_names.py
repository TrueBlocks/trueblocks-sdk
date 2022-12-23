#
# This file was generated with makeClass --sdk. Do not edit it.
#
from . import session

namesCmd = "names"
namesPos = "terms"
namesFmt = "json"
namesOpts = {
    "expand": {"hotkey": "-e", "type": "switch"},
    "matchCase": {"hotkey": "-m", "type": "switch"},
    "all": {"hotkey": "-l", "type": "switch"},
    "custom": {"hotkey": "-c", "type": "switch"},
    "prefund": {"hotkey": "-p", "type": "switch"},
    "named": {"hotkey": "-n", "type": "switch"},
    "addr": {"hotkey": "-a", "type": "switch"},
    "tags": {"hotkey": "-g", "type": "switch"},
    "toCustom": {"hotkey": "-u", "type": "switch"},
    "clean": {"hotkey": "-C", "type": "switch"},
    "autoname": {"hotkey": "-A", "type": "flag"},
    "create": {"hotkey": "", "type": "switch"},
    "update": {"hotkey": "", "type": "switch"},
    "delete": {"hotkey": "", "type": "switch"},
    "undelete": {"hotkey": "", "type": "switch"},
    "remove": {"hotkey": "", "type": "switch"},
    "fmt": {"hotkey": "-x", "type": "flag"},
    "verbose:": {"hotkey": "-v", "type": "switch"},
    "help": {"hotkey": "-h", "type": "switch"},
}

def names(self):
    ret = self.toUrl(namesCmd, namesPos, namesFmt, namesOpts)
    url = 'http://localhost:8080/' + ret[1]
    if ret[0] == 'json':
        return session.get(url).json()
    return session.get(url).text

