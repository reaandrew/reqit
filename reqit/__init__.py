import sys
import json

from ruamel import yaml
import requests
import humanize


def print_row(message, value):
    print("{0:>30}: {1:<}".format(message, value))


def print_header(title):
    print("\033[1m{}\033[0m".format(title))


def main():
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

        print_header("Response Info")
        print_row("Method", response.request.method)
        print_row("Status", response.status_code)
        print_row(
            "Elapsed time", "{} ms".format(int(response.elapsed.total_seconds() * 1000))
        )
        print_row("Size", humanize.naturalsize(len(response.content)))
        print_row("Encoding", response.encoding)
        print_header("Request Headers")
        for key, value in response.request.headers.items():
            print_row(key, value)
        print_header("Response Headers")
        for key, value in response.headers.items():
            print_row(key, value)

        print_header("Cookies")
        for key, value in response.cookies.items():
            print_row(key, value)

        print_header("Data")
        print(response.content)


if __name__ == "__main__":
    sys.exit(main())
