import sys
import json
import re
import os
import jinja2

from ruamel import yaml
import requests

from reqit.models import Result
from reqit.printers import StdOutPrinter


def expand(value):
    return jinja2.Template(str(value)).render(os.environ)


def existing():
    with open(sys.argv[1], "r+") as yaml_stream:
        data = list(yaml.load_all(yaml_stream, Loader=yaml.RoundTripLoader))
        properties = data[0]
        payload = None
        try:
            payload = json.dumps(data[1])
        except IndexError:
            pass

        for header, header_value in properties["headers"].items():
            properties["headers"][header] = expand(header_value)

        try:
            response = requests.request(
                expand(properties["method"]),
                url=expand(properties["url"]),
                headers=properties["headers"],
                data=expand(payload),
                verify=properties["verify"] == "true",
            )

            result = Result(response)
            printer = StdOutPrinter()
            printer.print(result)
        except requests.exceptions.ConnectionError:
            print("Connection refused")


def main():
    existing()


if __name__ == "__main__":
    sys.exit(main())
