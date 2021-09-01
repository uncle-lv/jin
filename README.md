# jin

![license](https://img.shields.io/github/license/uncle-lv/jin)  ![stars](https://img.shields.io/github/stars/uncle-lv/jin)   ![issues](https://img.shields.io/github/issues/uncle-lv/jin) ![forks](https://img.shields.io/github/forks/uncle-lv/jin) 

The **jin** is a simplified version of the gin web framework that can help you quickly understand the core principles of a web framework.

If this project is helpful for you, please star the project, thank you!!!

## Functions

- Support GET, POST, PUT, DELETE, PATCH, HEAD, OPTIONS
- Querystring parameters
- JSON, XML rendering
- Multipart/Urlencoded Form
- Get the parameters in Path
- Grouping routes
- Middleware
- Template
- Panic Recover
- Set and get a cookie
- Upload single file
- Upload multiple files
- Redirects
- A simple logger

## Diagrams

### Class Diagram

![class diagram](https://cdn.jsdelivr.net/gh/uncle-lv/PicX-image-hosting@main/jin/class_diagram.4nq0xbzu2cu0.svg)

## Usage

**jin** is built step by step, and I have used git tag to manage every milestone version. You can view all important version by typing the command `git tag`, and get the version you want by typing the command `git checkout <tagname>`. You can view more details by checking commit logs.

Example:

```bash
$ git tag
step1-a-simple-framework
step2-add-the-context-for-framework
step3-dynamic-routing
step4-router-group
step5-middleware
step6-HTML-template
step7-panic-recovery

$ git checkout step1-a-simple-framework
```

## Contributions

Any contributions you make are greatly appreciated.

## Lisence

[GPL-2.0](https://github.com/uncle-lv/jin/blob/main/LICENSE)

## Thanks

This project is based on [geektutu's](https://github.com/geektutu) [gee-web](https://github.com/geektutu/7days-golang/tree/master/gee-web). That's a wonderful project.I suggest you reading this [bolg](https://geektutu.com/post/gee.html) first.