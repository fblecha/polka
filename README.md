[![GoDoc](https://godoc.org/github.com/PolkaBand/polka?status.svg)](https://godoc.org/github.com/PolkaBand/polka) [![Build Status](https://drone.io/github.com/PolkaBand/polka/status.png)](https://drone.io/github.com/PolkaBand/polka/latest)

***Totally experimental still - use at your own risk :-)***

<!-- TOC depth:6 withLinks:1 updateOnSave:1 orderedList:0 -->

- [Quick Start](#quick-start)
- [Examples](#examples)
- [Setup](#setup)
- [Polka Commands](#polka-commands)
<!-- /TOC -->


__The following is getting dangerously close to actually working___ This is README driven development, e.g. write how you want the readme to work first, then implement the code.

# Quick Start

This does not work yet...

```shell
$ brew tap INSERT_TAP
$ brew install polka
$ cd ~/code
$ polka new todo_app
$ cd todo_app
$ ???
```

# Examples


# Setup
What needs to be setup for polka to work?
* AWS Services
* AWS IAM users/roles/policies

## AWS Credentials

Currently, polka assumes your AWS credentials are configured.  E.g., it assumes:
* Your credentials exist at ~/.aws/credentials  ( See this [AWS blog post](http://blogs.aws.amazon.com/security/post/Tx3D6U6WSFGOK2H/A-New-and-Standardized-Way-to-Manage-Credentials-in-the-AWS-SDKs) on setting up your credentials )


# Polka Commands

```shell
$ polka

usage: polka [--version] [--help] <command> [<args>]

Available commands are:
    config      configure a global polka app item
    deploy      deploys a polka app to AWS
    doc         Generates documentation for your polka app
    generate    generates lambda functions for endpoints
    new         creates a new polka app
    version     Prints the Polka version

```
