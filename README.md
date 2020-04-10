# iStone 

A toolkit to help you build an admin panel for your golang app.

## Features

- Building -> 🛠 RESTful API style
- Building -> 🛠 Router freamework powered by Gin.
- Building -> 🛠 Full support JWT.
- Building -> 🛠 Api docs powered by Swagger.
- Building -> 🛠 Database midddleware powered by GROM.
- Building -> 🚀 Fast: build a production admin panel app in ten minutes.
- Building -> 🎨 Theming: beautiful ui themes supported(more themes are coming.)
- Building -> 🔢 Plugins: many plugins to use.
- Building -> ✅ RBAC: out of box rbac auth system based on Casbin.
- Building -> ⚙️ Frameworks: support most of the go web frameworks.

## How to

[TODO]

## Code Structure

```text
.
├── LICENSE
├── Makefile
├── README-en.md
├── README.md
├── cmd             // commands
│   ├── api         // commands api
│   ├── cmd.go      // package cmd`s index
│   ├── conf        // package cmd`s config
│   ├── migrate     // package cmd`s migrate
│   └── version     // print iStone`s version
├── conf            // all configs for iStone
├── dev-logs        // development docs 
├── docs            // swagger docs configs
├── go.mod          
├── main.go
└── src             // main source files
    ├── admin       // source files for admin panel 
    ├── app         // source files for applications
    ├── core        // all the common and core source file
    ├── pkg         // iStone`s self packages
    └── tools       // utils source files  for iStone
```

## Contribution

git-flow

1. clone
2. checkout dev 
3. make new beranch  [feature-name]
4. coding
5. commit and PR


## Backers 

Your surpport will help me do better!

## Lisence

MIT 