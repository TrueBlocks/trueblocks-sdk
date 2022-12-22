from . import session

slurpCmd = "slurp"
slurpPos = "addrs"
slurpFmt = "txt"
slurpOpts = {
    "types": {"hotkey": "-t", "type": "flag"},
    "appearances": {"hotkey": "-p", "type": "switch"},
    "fmt": {"hotkey": "-x", "type": "flag"},
    "verbose:": {"hotkey": "-v", "type": "switch"},
    "help": {"hotkey": "-h", "type": "switch"},
}

def slurp(self):
    ret = self.toUrl(slurpCmd, slurpPos, slurpFmt, slurpOpts)
    url = 'http://localhost:8080/' + ret[1]
    if ret[0] == 'json':
        return session.get(url).json()
    return session.get(url).text
