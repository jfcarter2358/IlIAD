"""
Destroy script for provider resources
"""

import sys
import logging

if __name__ == '__main__':
    if len(sys.argv) != 2:
        logging.fatal(f'Invalid number of arguments passed, expected 1, got {len(sys.argv) - 1}')
    conf = sys.argv[1]
    pass
