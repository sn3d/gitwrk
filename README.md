# GitWrk
GitWrk is a small commandline application that helps you with monthly reports by extracting working hours directly from git repository.

![](gitwrk.gif)

## How it works

The idea is very simple. We're following very simple convetion that every commit we spent some times contain this information:

```
git commit -m “feat: Add feature X to module Y” -m “Spent 4h15m” 
```

For better insights and transparency we adapted [semantic commit message](https://gist.github.com/joshbuchea/6f47e86d2510bce28f8e7f42ae84c716) convention the `gitwrk` can process as well.

Then you can easily create monthly report for given contributor by typing:

```
> gitwrk --last-month --author zdenko.vrabel@unravela.com
```


## Motivation

As freelance developers or outsourced developers, we need to report our work hours. For years we used tools like Jira or some weird 3rd party tools for logging our hours. Having 3 different tools for reporting our work is always exhausting and error prone. Because we love transparency, we want to provide good data.

This approach has several advantages: 
- We’re lazy engineers.  We don’t like to handle 3 different tools for one think. Git is the ultimate tool for our daily work. 
- Everything in one place. Multiple tools leads to disconnected informations. 
- Better transparency. Customer exactly know how much of time we spent on what.
- We can analyze where are our bottlenecks, where we’re burning too many of our work time.

Programming isn't 100% of our daily work of course. Sometimes you have a meeting, sometimes some administrative work to do, but programming is still a majority.The `--allow-empty` flag allows you to track also non-coding work. 

## Installation

### Install on Ubuntu 

Download the .DEB package and install it

```
wget https://github.com/unravela/gitwrk/releases/download/v1.0.6/gitwrk_1.0.6_linux_64-bit.deb
dpkg -i ./gitwrk_1.0.6_linux_64-bit.deb
```

Untar the file and copy to the right place.

```
tar -xzvf ./gitwrk-1.0-linux-amd64.tar.gz -C /tmp/
sudo mv /tmp/gitwrk /usr/local/bin
```

Test if app is running

```
gitwrk --help
```

### Install on Windows

- Download the gitwrk [archive](https://github.com/unravela/gitwrk/releases/download/v1.0/gitwrk-1.0-win-amd64.zip). 
- Unzip it and run the `gitwrk.exe`


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