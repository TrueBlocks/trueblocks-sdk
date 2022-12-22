from . import session

tracesCmd = "traces"
tracesPos = "addrs"
tracesFmt = "txt"
tracesOpts = {
    "articulate": {"hotkey": "-a", "type": "switch"},
    "filter": {"hotkey": "-f", "type": "flag"},
    "statediff": {"hotkey": "-d", "type": "switch"},
    "count": {"hotkey": "-U", "type": "switch"},
    "fmt": {"hotkey": "-x", "type": "flag"},
    "verbose:": {"hotkey": "-v", "type": "switch"},
    "help": {"hotkey": "-h", "type": "switch"},
}

def traces(self):
    ret = self.toUrl(tracesCmd, tracesPos, tracesFmt, tracesOpts)
    url = 'http://localhost:8080/' + ret[1]
    if ret[0] == 'json':
        return session.get(url).json()
    return session.get(url).text
