SifuConf
========
[![Build Status](https://travis-ci.org/pradeepg26/sifuconf.svg?branch=master)](https://travis-ci.org/pradeepg26/sifuconf) [![CircleCI](https://circleci.com/gh/pradeepg26/sifuconf.svg?style=svg)](https://circleci.com/gh/pradeepg26/sifuconf)

```
.
├── service_a
│   ├── _defaults.sifu
│   ├── prod
│   │   ├── _prod.defaults.sifu
│   │   ├── prod.autodesk.sifu
│   │   ├── prod.microsoft.sifu
│   │   └── prod.skype.sifu
│   ├── stage
│   │   ├── _stage.defaults.sifu
│   │   ├── stage.autodesk.sifu
│   │   ├── stage.microsoft.sifu
│   │   └── stage.skype.sifu
│   ├── qa
│   │   ├── _qa.defaults.sifu
│   │   ├── qa.qa01.sifu
│   │   ├── qa.qa02.sifu
│   │   └── qa.qa03.sifu
│   └── dev
│       ├── _dev.defaults.sifu
│       ├── dev.adam.sifu
│       ├── dev.matt.sifu
│       └── dev.tony.sifu
├── service_b
├── service_c
└── service_d
```

```
service_c/_defaults.sifu
-------------
import "service_a" _
// imports service_a/_defaults.sifu into this config file

import "service_b" alternate
// imports service_b/_defaults.sifu into this config file
// but all the keys will be prefixed with "alternate."

// This is a comment
// This is how you set a config
port = 8080 // integer
featureY.enable = false // boolean
featureY.scaling_factor = 0.5 // float
featureY.display_name = "Awesome Feature" // String

hosts = ["host1", "host2", "host3"] // Lists

multiline_list = [
  "e1",
  "e2",
  "e3"
] // Multiline lists

// These configs cannot be overridden
final enable.featureX = true
final es.index.name = "idx_hello"

// These configs are required to be defined in sub configs
required phase = string

// Define a secret config
secret mysql_password
secret encryption_key
```

```
service_c/prod/_prod.defaults.sifu
------------------
import "service_a" _
// imports service_a/prod/_prod.defaults.sifu

import "service_a/skype" skype
// imports service_a/prod/prod.skype.sifu into the "skype" namespace

// Must declare that you're overriding a config
override port = 8090
override featureY.enable = true
// Finalize an override
final override featureY.scaling_factor = 0.6
final override phase = "prod"

```
