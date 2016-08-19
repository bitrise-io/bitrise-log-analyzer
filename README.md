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
