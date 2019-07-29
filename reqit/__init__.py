import sys
import json
import re
import os
import jinja2

from ruamel import yaml
from faker import Faker
import requests

from reqit.models import Result
from reqit.printers import StdOutPrinter


class FakerWrapper:
    def __init__(self, faker):
        self.wrappee = faker

    def __getattr__(self, attr):
        return getattr(self.wrappee, attr)()


def expand(value):
    return jinja2.Template(str(value)).render(
        os.environ, fake=FakerWrapper(Faker("en_GB"))
    )


def existing():
    with open(sys.argv[1], "r+") as yaml_stream:
        data = list(yaml.load_all(yaml_stream, Loader=yaml.RoundTripLoader))
        properties = data[0]["request"]
        payload = None
        try:
            payload = json.dumps(data[1])
        except:
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
