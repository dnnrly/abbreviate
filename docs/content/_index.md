---
title: "abbreviate"
---

# abbreviate

> Documentation for **abbreviate** - a handy CLI tool for abbreviating strings.


Way back when, I was fixing up a microservice and deploying in a CI/CD pipeline. We had a neat pipeline that would deploy feature branches for you automatically but there were a few problems with it, all around the fact that the CNAME that it created for this new deployment used the raw branch name:

1. Route53 CNAMES can't have '/' - which makes this awkward if you're using Gitflow or something similar
2. There is a character limit so once you've got rid of the '/' then it's quite easy to blow this limit when you consider that a common branch name words such as "feature" is 7 characters long.

Wouldn't it be nice if we could automatically detect common words and replace them with a sensible abbreviation? And what about removing non-compliant characters?

Abbreviate hopes to do exactly this!

## The basic CNAME example

```bash
$ abbreviate snake -s '-' application-dev-feature/new-strategy
app-dev-feat-new-stg
```
## Shorter variable names

Imagine you're converting some "Enterprise Standard" Java to something a little terser:

```bash
$ abbreviate camel visibleSubscribers
visSub
```

## Convert between case styles

Perhaps you don't actually need to shorten a string, just change it to your new favourite style. As long as the string is shorter than the maximum length, it will only change cases:

```bash
$ /abbreviate pascal -m 99 visibleSubscribers
VisibleSubscribers
```

