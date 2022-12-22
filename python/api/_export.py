from . import session

exportCmd = "export"
exportPos = "addrs"
exportFmt = "txt"
exportOpts = {
    "appearances": {"hotkey": "-p", "type": "switch"},
    "receipts": {"hotkey": "-r", "type": "switch"},
    "logs": {"hotkey": "-l", "type": "switch"},
    "traces": {"hotkey": "-t", "type": "switch"},
    "neighbors": {"hotkey": "-n", "type": "switch"},
    "accounting": {"hotkey": "-C", "type": "switch"},
    "statements": {"hotkey": "-A", "type": "switch"},
    "articulate": {"hotkey": "-a", "type": "switch"},
    "cache": {"hotkey": "-i", "type": "switch"},
    "cache_traces": {"hotkey": "-R", "type": "switch"},
    "count": {"hotkey": "-U", "type": "switch"},
    "first_record": {"hotkey": "-c", "type": "flag"},
    "max_records": {"hotkey": "-e", "type": "flag"},
    "relevant": {"hotkey": "", "type": "switch"},
    "emitter": {"hotkey": "", "type": "flag"},
    "topic": {"hotkey": "", "type": "flag"},
    "asset": {"hotkey": "", "type": "flag"},
    "flow": {"hotkey": "-f", "type": "flag"},
    "factory": {"hotkey": "-y", "type": "switch"},
    "unripe": {"hotkey": "-u", "type": "switch"},
    "first_block": {"hotkey": "-F", "type": "flag"},
    "last_block": {"hotkey": "-L", "type": "flag"},
    "fmt": {"hotkey": "-x", "type": "flag"},
    "verbose:": {"hotkey": "-v", "type": "switch"},
    "help": {"hotkey": "-h", "type": "switch"},
}

def export(self):
    ret = self.toUrl(exportCmd, exportPos, exportFmt, exportOpts)
    url = 'http://localhost:8080/' + ret[1]
    if ret[0] == 'json':
        return session.get(url).json()
    return session.get(url).text
