from . import session

transactionsCmd = "transactions"
transactionsPos = "addrs"
transactionsFmt = "txt"
transactionsOpts = {
    "articulate": {"hotkey": "-a", "type": "switch"},
    "trace": {"hotkey": "-t", "type": "switch"},
    "uniq": {"hotkey": "-u", "type": "switch"},
    "flow": {"hotkey": "-f", "type": "flag"},
    "account_for": {"hotkey": "-A", "type": "flag"},
    "cache": {"hotkey": "-o", "type": "switch"},
    "fmt": {"hotkey": "-x", "type": "flag"},
    "verbose:": {"hotkey": "-v", "type": "switch"},
    "help": {"hotkey": "-h", "type": "switch"},
}

def transactions(self):
    ret = self.toUrl(transactionsCmd, transactionsPos, transactionsFmt, transactionsOpts)
    url = 'http://localhost:8080/' + ret[1]
    if ret[0] == 'json':
        return session.get(url).json()
    return session.get(url).text
