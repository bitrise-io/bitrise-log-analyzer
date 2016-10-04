# bitrise-log-analyzer

Bitrise Log Analyzer tool


## Install

### If you have `Go` installed

```
go get -u github.com/bitrise-tools/bitrise-log-analyzer
```

## Usage

### Step infos

```
bitrise-log-analyzer stepinfos /path/to/bitrise.log
```

If all you want is the step run times (how long each step took to complete):

```
bitrise-log-analyzer stepinfos /path/to/bitrise.log --only-times
```

## Development

### Editor

If you want to quickly iterate on the `editor/www` files, just delete
the `editor/rice-box.go` file.

Once you're done with your changes run `bitrise run assets-precompile`
to re-compile `editor/rice-box.go`, which will prepare the content of `editor/www`
to be embedded in the binary after a `go install`.

#### Embed the editor's web content

To embed / refresh the editor's web content (`editor/www`) run:

```
bitrise run assets-precompile
go install
```

This will produce a new binary where the content of `editor/www` is embedded
in the binary.
