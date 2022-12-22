from . import session

stateCmd = "state"
statePos = "addrs"
stateFmt = "txt"
stateOpts = {
    "parts": {"hotkey": "-p", "type": "flag"},
    "changes": {"hotkey": "-c", "type": "switch"},
    "no_zero": {"hotkey": "-n", "type": "switch"},
    "fmt": {"hotkey": "-x", "type": "flag"},
    "verbose:": {"hotkey": "-v", "type": "switch"},
    "help": {"hotkey": "-h", "type": "switch"},
}

def state(self):
    ret = self.toUrl(stateCmd, statePos, stateFmt, stateOpts)
    url = 'http://localhost:8080/' + ret[1]
    if ret[0] == 'json':
        return session.get(url).json()
    return session.get(url).text
