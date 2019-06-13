
.PHONY: install
install:
	pip install -e .

.PHONY: format
format:
	find . -name '*.py'|grep -v migrations|xargs autoflake --in-place --remove-all-unused-imports --remove-unused-variables
	autopep8 --in-place --aggressive --aggressive **/*.py

.PHONY: test
test:
	pytest --cov-report html --cov-report xml --cov-report annotate  --cov=reqit --html testreport.html --self-contained-html test/*

.PHONY: lint
lint:
	pylint --rcfile pylintrc **/*.py

.PHONY: build
build: install test lint test
