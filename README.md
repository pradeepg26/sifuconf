```
.
├── defaults.sifu
├── prod
│   ├── prod.defaults.sifu
│   └── deployments
│       ├── prod.autodesk.sifu
│       ├── prod.bskyb.sifu
│       └── prod.skypec.sifu
├── stage
│   ├── stage.defaults.sifu
│   └── deployments
│       ├── stage.autodesk.sifu
│       ├── stage.bskyb.sifu
│       └── stage.skypec.sifu
├── qa
│   ├── qa.defaults.sifu
│   └── deployments
│       ├── qa.qa01.sifu
│       ├── qa.qa02.sifu
│       └── qa.qa03.sifu
└── dev
    ├── dev.defaults.sifu
    └── deployments
        ├── dev.adam.sifu
        ├── dev.matt.sifu
        └── dev.tony.sifu
```

```
defaults.sifu
-------------
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
abstract phase

// Define a secret config
secret mysql_password = "mysql_password_for_user1"
secret encryption_key = "keys/key1"
```

```
prod.defaults.sifu
------------------
// Must declare that you're overriding a config
override port = 8090
override featureY.enable = true
// Finalize an override
final override featureY.scaling_factor = 0.6

// Don't need to declare override for abstract overrides
final phase = prod // Also finalize the config

```
