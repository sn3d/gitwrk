# GitWrk
GitWrk is a small commandline application that helps you with monthly reports by extracting working hours directly from git repository.

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
