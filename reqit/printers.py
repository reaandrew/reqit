import humanize
import json


class StdOutPrinter:
    def print_row(self, message, value):
        if value is None:
            value = ""
        print("{0:>30}: {1:<}".format(message, value))

    def print_header(self, title):
        print()
        print("{}".format(title))
        print()

    def visit(self, response):
        self.print_header("Response Info")
        self.print_row("Method", response.request.method)
        self.print_row("Status", response.status_code)
        self.print_row(
            "Elapsed time", "{} ms".format(int(response.elapsed.total_seconds() * 1000))
        )
        self.print_row("Size", humanize.naturalsize(len(response.content)))
        self.print_row("Encoding", response.encoding)
        self.print_header("Request Headers")
        for key, value in response.request.headers.items():
            self.print_row(key, value)

        self.print_header("Response Headers")
        for key, value in response.headers.items():
            self.print_row(key, value)

        self.print_header("Cookies")
        if len(response.cookies) == 0:
            print("No cookies")

        for key, value in response.cookies.items():
            self.print_row(key, value)
        self.print_header("Data")

        try:
            parsed = json.loads(response.content.decode("utf-8"))
            print(json.dumps(parsed, indent=4, sort_keys=True))
        except:
            print(response.content.decode("utf-8"))

    def print(self, result):
        result.accept(self)
