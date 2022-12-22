from . import session

daemonCmd = "daemon"
daemonPos = "addrs"
daemonFmt = "txt"
daemonOpts = {
    "port": {"hotkey": "-p", "type": "flag"},
    "scrape": {"hotkey": "-s", "type": "flag"},
    "monitor": {"hotkey": "-m", "type": "switch"},
    "api": {"hotkey": "-a", "type": "flag"},
    "fmt": {"hotkey": "-x", "type": "flag"},
    "verbose:": {"hotkey": "-v", "type": "switch"},
    "help": {"hotkey": "-h", "type": "switch"},
}

def daemon(self):
    ret = self.toUrl(daemonCmd, daemonPos, daemonFmt, daemonOpts)
    url = 'http://localhost:8080/' + ret[1]
    if ret[0] == 'json':
        return session.get(url).json()
    return session.get(url).text
