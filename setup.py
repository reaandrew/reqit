# -*- coding: utf-8 -*-

from setuptools import setup, setuptools
try: # for pip >= 10
    from pip._internal.req import parse_requirements
except ImportError: # for pip <= 9.0.3
    from pip.req import parse_requirements

install_reqs = parse_requirements("./requirements.txt", session=False)
reqs = [str(ir.req) for ir in install_reqs]


with open("README.md") as f:
    readme = f.read()

with open("LICENSE") as f:
    license = f.read()

setup(
    name="reqit",
    version="0.1.0",
    description="reqit",
    long_description=readme,
    author="Andrew Rea",
    author_email="email@andrewrea.co.uk",
    url="https://github.com/reaandrew/reqit",
    license=license,
    packages=setuptools.find_packages(),
    install_requires=reqs,
    include_package_data=True,
    classifiers=[
        "License :: OSI Approved :: MIT License",
        "Programming Language :: Python :: 3",
        "Programming Language :: Python :: 3.7",
    ],
    entry_points={"console_scripts": "reqit=reqit:main"},
)
