# envutil

Envutil contains small utility functions for dealing with environments and args, specifically:

* `envmap` provides access to environment variables as a `map[string]string` and can convert those maps back to `[]string`
* `splitarg` provides splitting of `os.Args` by `--` into arguments and arguments for another command
