# procfs_dumper

It is a tool to dump procfs.

## Example

```sh
root@40bb5830d366:/tmp# procfs_dump -p 1 | jq .
{
  "PID": 1,
  "CmdLine": [
    "bash"
  ],
  "Comm": "bash",
  "Cwd": "/tmp",
  "Environ": [
    "PATH=/go/bin:/usr/local/go/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin",
    "HOSTNAME=40bb5830d366"
  ],
  "Executable": "/bin/bash",
  "FileDescriptorTargets": [
    "/dev/pts/0"
  ],
  "IO": {
    "RChar": 50454852,
    "WChar": 19745589,
    "SyscR": 56355,
    "SyscW": 5019,
    "ReadBytes": 2650112,
    "WriteBytes": 9211904,
    "CancelledWriteBytes": 77824
  },
  :
  :
```

## Easy install and extract

```sh
URL=$(curl -s https://api.github.com/repos/takumakume/procfs_dumper/releases/latest | grep "browser_download_url.*gz" | cut -d '"' -f 4) && curl -L -o procfs_dumper.tar.gz $URL && tar zxvf procfs_dumper.tar.gz
```

## Usage

### Dump a process

```sh
procfs_dump -p <PID> | jq .
```

```json
{
  "PID": 1,
  "CmdLine": [
    "bash"
  ],
  "Comm": "bash",
  :
  : (snip)
```

### Dump all processes

```sh
procfs_dump -P | jq .
```

```json
[
  {
    "PID": 1,
    "CmdLine": [
      "bash"
    ],
    "Comm": "bash",
    :
    : (snip)
  {
    "PID": 2,
    "CmdLine": [
      "bash"
    ],
    "Comm": "bash",
    :
    : (snip)
:
: (snip)
```

## TODO

- Output to file.
- Descrive information other than process. ( `/proc/hoge` )

## Credit

This software includes the work that is distributed in the Apache License 2.0.,

made using [prometheus/procfs](https://github.com/prometheus/procfs).
Thank you.

If you want to dump more information, contribute to [prometheus/procfs](https://github.com/prometheus/procfs).
