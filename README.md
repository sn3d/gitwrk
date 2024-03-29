# GitWrk
[![Release](https://img.shields.io/github/release/unravela/gitwrk.svg?style=flat-square)](https://github.com/goreleaser/goreleaser/releases/latest)
[![Software License](https://img.shields.io/github/license/unravela/gitwrk?style=flat-square)](/LICENSE.md)
[![Build](https://img.shields.io/github/workflow/status/unravela/gitwrk/build/master?style=flat-square)](/actions?query=workflow%3Abuild)

GitWrk is a small commandline application that helps you with monthly reports by extracting working hours directly from git repository.

![](assets/gitwrk.gif)

## How it works

The idea is very simple. We're following very simple convention. Every commit we spent some times contain this information
via `spent` [trailer line](https://git-scm.com/docs/git-interpret-trailers).

```
git commit -m “feat: Add feature X to module Y” -m “spent: 4h15m” 
```

Another allowed form is non-trailer convention without `:`:

```
git commit -m “feat: Add feature X to module Y” -m “Spent 1h20m,7h30m” 
```

For better insights and transparency we adapted [conventional commits](https://www.conventionalcommits.org/en/v1.0.0/).
Then you can easily create monthly report for given contributor by typing:

```
> gitwrk --last-month --author zdenko.vrabel@unravela.com
```

## Motivation

As freelance developers or outsourced developers, we need to report our work hours. For years we used tools like Jira or some weird 3rd party tools for logging our hours. Having 3 different tools for reporting our work is always exhausting and error prone. Because we love transparency, we want to provide good data.

This approach has several advantages: 
- We’re lazy engineers.  We don’t like to handle 3 different tools for one think. Git is the ultimate tool for our daily work. 
- Everything in one place. Multiple tools leads to disconnected information. 
- Better transparency. Customer exactly know how much of time we spent on what.
- We can analyze where are our bottlenecks, where we’re burning too many of our work time.

Programming isn't 100% of our daily work of course. Sometimes you have a meeting, sometimes some administrative work to do, but programming is still a majority.The `--allow-empty` flag allows you to track also non-coding work. 

Example of meeting:
```
git commit --allow-empty -m "meet: planning for next sprint" -m "spent: 45m"
```

## Installation

### Install on Ubuntu (Snap)

You can install the application easily via Snap

```
snap install gitwrk
```

Test if app is installed

```
gitwrk --help
```

### Install on Linux (DEB)
Download the .DEB package and install it

```
wget https://github.com/unravela/gitwrk/releases/download/v1.0.8/gitwrk_1.0.8_linux_64-bit.deb
dpkg -i ./gitwrk_1.0.8_linux_64-bit.deb
```

### Install on Linux (RPM)
Download the .RPM package and install it

```
wget https://github.com/unravela/gitwrk/releases/download/v1.0.8/gitwrk_1.0.8_linux_64-bit.rpm
rpm -U ./gitwrk_1.0.8_linux_64-bit.rpm
```


### Install on Windows (Scoop)

First, ensure the [scoop](https://scoop.sh/) is present in your environment. If not install it.

Run commands
```
scoop bucket add unravela https://github.com/unravela/scoop-bucket
scoop install gitwrk
```

### Install on MacOS (Homebrew)

If you have Homebrew present in your environment, run command:
```
brew install unravela/tap/gitwrk
```

### Build from source code

If you have Go (version 1.13) installed on your system, you can use command:
```
go get github.com/unravela/gitwrk
```

This command will download and install `gitwrk` into your `GOPATH/bin` folder.

## How to use

The best way how to explore what gitwrk offers you is by help page:

```
    gitwrk --help
```

By default, the `gitwrk` will create report for all users for all commits. If you want to create report for concrete contributor in current month, you can use combination of flags `-current-month` and `--author`.

### Time frames

```
    gitwrk --current-month --author me@company.com
```

You can also request report for last finished month:

```
    gitwrk --last-month --author me@company.com
```

If you wish report for some time window e.g. since November of 2019 to Jaunary 2020, you can use flags `--since` and `--till`:

```
    gitwrk --since 2019-11-01 --till 2020-01-31
```

### Semantic commit message

The `gitwrk` is also supporting filtering by type and scope of [semantic commit message](https://gist.github.com/joshbuchea/6f47e86d2510bce28f8e7f42ae84c716). You can use flags `--type` and `--scope` for filtering. For example you want to know how many hours you spent on documentation:

```
    gitwrk --type docs --author me@company.com
```

Or how many hours spent developers on `module-a` last month:

```
    gitwrk --scope module-a --last-month
```

### Data in message

The `gitwrk` relly and parse commit message and extract some more data.
As you might see, the required field in every commit is `spent`. But there are
more.

**spent** - how much of time you spent on commit/task. You can use simple
format e.g. `1h45m`.

**date** - normally the date of the worklog is same as commit's date. Sometimes
this is not desired and you want to log some past work. You can use
`date: 2021-12-24` to override commit's date for worklog.

### JSON and CSV output

The `gitwrk` can produce also JSON or CSV reports. This is usefull if you want to automatize and export your work hours to external systems. All you need is tell the gitwrk the output format via `--output` or `-o` flag.

```
    gitwrk --last-month --author me@copmany.com -o json
```

This command will produce output:
```
[
        {
                "when": "2020-02-11T17:27:24+01:00",
                "author": "me@company.com",
                "scm_type": "docs",
                "scm_scope": "",
                "spent": "30m0s",
                "spent_minutes": 30
        },
        {
                "when": "2020-02-10T17:12:04+01:00",
                "author": "me@gmail.com",
                "scm_type": "docs",
                "scm_scope": "",
                "spent": "1h10m0s",
                "spent_minutes": 70
        }
]
```
