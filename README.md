# blackfn_exporter
A Cloud Function/Lambda adaptation of the [blackbox_exporter](https://github.com/prometheus/blackbox_exporter). This allows deploying the prober with minimal infrastructure, which can be useful or cost-effective when deploying it in large numbers.

## Supported cloud providers
Currently AWS Lambda and Google Cloud Functions are supported. Adding new providers should be very simple - the existing ones are less than 20 lines of code each!

## Configuration
The function uses a configuration file on the same format used by the blackbox_exporter, documented [here](https://github.com/prometheus/blackbox_exporter/blob/master/CONFIGURATION.md). The configuration file should be named `blackbox_exporter.yml` and placed in the root of the deployment.

In addition, two environment variables are considered:
- `CONFIG_FILE`: To override the path to the config file.
- `LOG_LEVEL`: To set the log level of lines written to stderr.

# Building
## AWS
Create a deployment package:

```bash
GOOS=linux go build ./clouds/aws -o blackbox_fn
zip deployment.zip blackbox_fn blackbox_exporter.yml
```

Upload and configure the package as detailed in the [AWS blog](https://aws.amazon.com/blogs/compute/announcing-go-support-for-aws-lambda/).

## GCP
_TODO: GCP requires the function to be in the package root, which makes it hard to have multiple clouds in the same repo. Contributions on how to improve this setup is much appreciated!_

Deploy with `gcloud`:

```bash
cp ./clouds/gcp/function.go . # This is a hack to get around the package root requirements mentioned above
gcloud functions deploy blackbox --entry-point=BlackboxProbe --runtime=go111 --source=. --trigger=http
```
