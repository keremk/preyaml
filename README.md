# YAML preprocessor

This is a very simple YAML preprocessor. It takes an input of a template YAML file and then applies what is inside another file as key/value pairs to substitute and generate a new YAML file. The main use case is to keep secrets (e.g. oauth tokens) in a separate file so they are not checked into a git repo.

Usage: 

```
preyaml gen -t sample_template.yaml -d secrets.env -o sample.yaml
```

See the correspond