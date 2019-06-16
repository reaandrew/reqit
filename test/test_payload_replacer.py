import unittest
import re
from faker import Faker


class SomeTestCase(unittest.TestCase):
    def test_it_does_something(self):

        fake = Faker("en_GB")
        data = 'Something: "{{fake:name}}"'

        matches = re.findall("{{([^}]+)}}", data)
        for match in matches:
            fields = match.split(":")
            field_type = fields[0]
            field_value = fields[1]

            if field_type == "fake":
                value = getattr(fake, field_value)()
                print(data.replace("{{{{{}}}}}".format(match), value))
            self.assertEqual(1, 2)
