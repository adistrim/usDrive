locals {
  envfile = {
    for line in split("\n", file(".env")) : split("=", line)[0] => regex("=(.*)", line)[0]
    if !startswith(line, "#") && length(split("=", line)) > 1
  }
}

data "external_schema" "gorm" {
  program = [
    "go",
    "run",
    "./atlas_loader/main.go"
  ]
}

env "default" {
  src = data.external_schema.gorm.url

  url = local.envfile["DATABASE_URL"]

  dev = "docker://postgres/15/dev"

  migration {
    dir = "file://migrations"
  }

  format {
    migrate {
      diff = "{{ range .Changes }}{{ .Cmd }};{{ end }}"
    }
  }
}
