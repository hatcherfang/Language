import argparse
description = 'This is used for argparse lib test'
version = '1.0'
def valid_server(server):
    try:
        return str(server)
    except ValueError:
        msg = 'Invalide server url: "{0}"'.format(server)
        raise argparse.ArgumentTypeError(msg)

parser = argparse.ArgumentParser(description=description)
parser.add_argument('-v', '--version', action='version', version=version)
parser.add_argument('-s', '--server', help='server ip address',
                    type=valid_server, required=True)
parser.add_argument('-a', action='store', dest='collection', default='None',
                    help='add values to collection', required=True)
print parser.parse_args()
