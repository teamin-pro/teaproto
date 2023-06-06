install:
	pip install mkdocs

docs: install
	mkdocs build

serve: install
	mkdocs serve

gh-deploy: install
	mkdocs gh-deploy
