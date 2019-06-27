import sys
import json
import re

from ruamel import yaml
import requests

from reqit.models import Result
from reqit.printers import StdOutPrinter
from reqit.security import generate_keys


def existing():
    with open(sys.argv[1], "r+") as yaml_stream:
        data = list(yaml.load_all(yaml_stream, Loader=yaml.RoundTripLoader))
        properties = data[0]
        payload = None
        try:
            payload = data[1]
        except IndexError:
            pass

        response = requests.request(
            properties["method"],
            url=properties["url"],
            headers=properties["headers"],
            data=payload,
        )

        if payload is not None:
            matches = re.findall("\{\{[\w]+\}\}", str(payload))
            print(matches)

        result = Result(response)
        printer = StdOutPrinter()
        printer.print(result)


def main():
    import argparse

    parser = argparse.ArgumentParser(description="Process some integers.")
    parser.add_argument(
        "file",
        help="an integer for the accumulator",
    )

    args = parser.parse_args()
    print(args.accumulate(args.integers))

    existing()


if __name__ == "__main__":
    sys.exit(main())
