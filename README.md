ConvSun
=========

This CLI-tool uses **Platform.sh** config files *(routes.yaml, services.yaml and .platform.app.yaml(s)/applications.yaml)* to generate the **Upsun** config file *(config.yaml)*  

> [!CAUTION]
> **This project is owned by the Upsun Advocacy team. It is in the early stage of development [experimental] and only intended to be used with caution by Upsun customers/community.   <br /><br />This project is not supported by Upsun and does not qualify for Support plans. Use this repository at your own risk, it is provided without guarantee or warranty!** 
> Don’t hesitate to join our [Discord](https://discord.com/invite/platformsh) to share your thoughts about this project.

> **WARNING: This tool handles classic 'multi-app...' cases but has not been tested for snowflake cases.**

#### Install

Download the last binary in [release section](https://github.com/upsun/convsun/releases).
Extract it and enjoy!

#### Syntax
```
Usage of convsun:
      --src string          Source project path to convert (default "./")
      --dst string          Destination project path where converted
      --mount_type string   Change 'Local' mount to upsun compatible mode : storage or instance. (default "storage")
  -v, --verbose             Enable verbose mode
```

#### Sample
`$ convsun --src=tests/convert"`
