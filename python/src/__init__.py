import os
import requests
import sys
import logging

session = requests.Session()
session.params = {}

class chifra():
    def __init__(self):
        pass

    from ._usage import isValidCommand, usage
    from ._utils import toUrl

    from ._abis import abis
    from ._blocks import blocks
    from ._chunks import chunks
    from ._config import config
    // from ._daemon import daemon
    // from ._explore import explore
    from ._export import export
    from ._init import init
    from ._list import list
    from ._logs import logs
    from ._monitors import monitors
    from ._names import names
    from ._receipts import receipts
    from ._scrape import scrape
    from ._slurp import slurp
    from ._state import state
    from ._tokens import tokens
    from ._traces import traces
    from ._transactions import transactions
    from ._when import when

    def dispatch(self):
        logging.getLogger().setLevel(logging.INFO)
        if self.isValidCommand() == False:
            self.usage("The first argument must be one of the valid commands.")

        response = ""
        match sys.argv[1]:
            case 'abis':
                return self.abis()
            case 'blocks':
                return self.blocks()
            case 'chunks':
                return self.chunks()
            case 'config':
                return self.config()
            case 'daemon':
                response = self.daemon()
            case 'explore':
                response = self.explore()
            case 'export':
                response = self.export()
            case 'init':
                response = self.init()
            case 'list':
                return self.list()
            case 'logs':
                response = self.logs()
            case 'monitors':
                response = self.monitors()
            case 'names':
                return self.names()
            case 'receipts':
                response = self.receipts()
            case 'scrape':
                response = self.scrape()
            case 'slurp':
                response = self.slurp()
            case 'state':
                response = self.state()
            case 'tokens':
                response = self.tokens()
            case 'traces':
                response = self.traces()
            case 'transactions':
                response = self.transactions()
            case 'when':
                return self.when()

    def cmdLine(self):
        ret = "chifra"
        for i, arg in enumerate(sys.argv):
            if i > 0:
                ret += (" " + arg)
        return ret
