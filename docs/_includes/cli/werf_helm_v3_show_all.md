{% if include.header %}
{% assign header = include.header %}
{% else %}
{% assign header = "###" %}
{% endif %}

This command inspects a chart (directory, file, or URL) and displays all its content
(values.yaml, Charts.yaml, README)


{{ header }} Syntax

```shell
werf helm-v3 show all [CHART] [flags] [options]
```

{{ header }} Options

```shell
      --ca-file='':
            verify certificates of HTTPS-enabled servers using this CA bundle
      --cert-file='':
            identify HTTPS client using this SSL certificate file
      --devel=false:
            use development versions, too. Equivalent to version '>0.0.0-0'. If --version is set,   
            this is ignored
  -h, --help=false:
            help for all
      --key-file='':
            identify HTTPS client using this SSL key file
      --keyring='/home/distorhead/.gnupg/pubring.gpg':
            location of public keys used for verification
      --password='':
            chart repository password where to locate the requested chart
      --repo='':
            chart repository url where to locate the requested chart
      --username='':
            chart repository username where to locate the requested chart
      --verify=false:
            verify the package before installing it
      --version='':
            specify the exact chart version to install. If this is not specified, the latest        
            version is installed
```
