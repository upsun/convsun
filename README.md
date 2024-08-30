ConvSun
=========

This CLI-tool uses **Platform.sh** config files *(routes.yaml, services.yaml and .platform.app.yaml(s)/applications.yaml)* to generate the **Upsun** config file *(config.yaml)*  
> **WARNING : This tool handles classic 'multi-app...' cases but has not been tested for snowflack cases.**

#### Install

Download the last binary in [release section](https://github.com/upsun/convsun/releases).
Extract it and enjoy !

#### Syntax
```
Usage of convsun:
      --src string          Source project path to convert (default "./")
      --dst string          Destination project path where converted
      --mount_type string   Change 'Local' mount to upsun compatible mode : storage or instance. (default "storage")
  -v, --verbose             Enable verbose mode
```

#### Sample
`$ upsun_convert --src=tests/convert"`