from . import session

whenCmd = "when"
whenPos = "addrs"
whenFmt = "txt"
whenOpts = {
    "list": {"hotkey": "-l", "type": "switch"},
    "timestamps": {"hotkey": "-t", "type": "switch"},
    "count": {"hotkey": "-U", "type": "switch"},
    "repair": {"hotkey": "-r", "type": "switch"},
    "check": {"hotkey": "-c", "type": "switch"},
    "update": {"hotkey": "", "type": "switch"},
    "fmt": {"hotkey": "-x", "type": "flag"},
    "verbose:": {"hotkey": "-v", "type": "switch"},
    "help": {"hotkey": "-h", "type": "switch"},
}

def when(self):
    ret = self.toUrl(whenCmd, whenPos, whenFmt, whenOpts)
    url = 'http://localhost:8080/' + ret[1]
    if ret[0] == 'json':
        return session.get(url).json()
    return session.get(url).text
