from . import session

exploreCmd = "explore"
explorePos = "addrs"
exploreFmt = "txt"
exploreOpts = {
    "local": {"hotkey": "-l", "type": "switch"},
    "google": {"hotkey": "-g", "type": "switch"},
    "fmt": {"hotkey": "-x", "type": "flag"},
    "verbose:": {"hotkey": "-v", "type": "switch"},
    "help": {"hotkey": "-h", "type": "switch"},
}

def explore(self):
    ret = self.toUrl(exploreCmd, explorePos, exploreFmt, exploreOpts)
    url = 'http://localhost:8080/' + ret[1]
    if ret[0] == 'json':
        return session.get(url).json()
    return session.get(url).text
