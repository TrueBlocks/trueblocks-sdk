from . import session

listCmd = "list"
listPos = "addrs"
listFmt = "txt"
listOpts = {
    "count": {"hotkey": "-U", "type": "switch"},
    "appearances": {"hotkey": "-p", "type": "switch"},
    "silent": {"hotkey": "", "type": "switch"},
    "no_zero": {"hotkey": "-n", "type": "switch"},
    "first_record": {"hotkey": "-c", "type": "flag"},
    "max_records": {"hotkey": "-e", "type": "flag"},
    "first_block": {"hotkey": "-F", "type": "flag"},
    "last_block": {"hotkey": "-L", "type": "flag"},
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
