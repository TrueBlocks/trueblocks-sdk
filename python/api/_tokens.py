from . import session

tokensCmd = "tokens"
tokensPos = "addrs"
tokensFmt = "txt"
tokensOpts = {
    "parts": {"hotkey": "-p", "type": "flag"},
    "by_acct": {"hotkey": "-b", "type": "switch"},
    "no_zero": {"hotkey": "-n", "type": "switch"},
    "fmt": {"hotkey": "-x", "type": "flag"},
    "verbose:": {"hotkey": "-v", "type": "switch"},
    "help": {"hotkey": "-h", "type": "switch"},
}

def tokens(self):
    ret = self.toUrl(tokensCmd, tokensPos, tokensFmt, tokensOpts)
    url = 'http://localhost:8080/' + ret[1]
    if ret[0] == 'json':
        return session.get(url).json()
    return session.get(url).text
