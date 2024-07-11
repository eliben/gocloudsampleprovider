This is a simple custom provider for the [Go CDK](https://github.com/google/go-cloud) `blob` API.
This custom provider is implemented "out-of-tree" - not part of the official Go CDK repository.

The provider is named `gocloudsampleprovider` and it's a clone of `fileblob`;
the URL is registered at `gcspfile://`.

See the `using-sample-provider` directory for an example of using this custom
provider. It's a separate Go module.
