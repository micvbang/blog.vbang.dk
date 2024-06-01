---
layout: post
title:  "Data exploration using VIM"
date:   2024-06-01 19:08:00 +0200
tags: [data-exploration, vim, cli]
---

I've used vim and/or vim bindings for the better part of 10 years. But apparently there's this tiny piece of magic that has completely escaped me all this time.

About half a year ago I received a tip from a good friend (thanks Jörn ❤️) that I kind of forgot about and never took the time to actually try out.

Then, this week I had to do a bunch of random data exploration and, luckily, it somehow jumped back into my brain. Just this week I've saved countless hours looking through gigabytes and gigabytes of sketchy data from the Danish Business Authority. Public data is _awesome_, but the quality of that data? Often not so much :( 

Anyway. The tip is this: you can use the vim command (is it called that?) `:%! [cmd]` to invoke CLI programs on data in vim's buffer. That's it. It's crazy powerful and I love it.

I've made a 3.5 minute screencast of how this looks in practice.
If you don't want to watch the video, I'll give you a short description of how it works below.

<script src="https://asciinema.org/a/662066.js" id="asciicast-662066" async="true"></script>

For example, let's say you have the following in your buffer:

```text
9msAkRIqstFcQAdfpvFZqgWGPBbReNS
3JEFbIfJuIGZBZodTONfnzyCykPtsBR
4KdSIqYYlDEIxpGiHFbRpqiZsFlgLxL
7UNqzFGgxEkzfzWLdTSKabDsUtTcSDs
5IqHRWKquwsekkritCxsnInXbsPeLvx
2ZdEuPTvYKFXNpOkhOytByqaDUQRSQI
0UreGiTTUnRxxrtNtaBfNYfbDhDlKwJ
1aOaHMrQzwGFjFtmwcPwdTfKVwteivR
6abgfdynLiidyiSBPUVMbkhKEsJMNVy
4doltlrfrOLmkuvCdVyJzqZRGkCOzkD
```

Running `:%! sort` in vim will pipe the data from our buffer into sort and put it back in the buffer:

```text
0UreGiTTUnRxxrtNtaBfNYfbDhDlKwJ
1aOaHMrQzwGFjFtmwcPwdTfKVwteivR
2ZdEuPTvYKFXNpOkhOytByqaDUQRSQI
3JEFbIfJuIGZBZodTONfnzyCykPtsBR
4doltlrfrOLmkuvCdVyJzqZRGkCOzkD
4KdSIqYYlDEIxpGiHFbRpqiZsFlgLxL
5IqHRWKquwsekkritCxsnInXbsPeLvx
6abgfdynLiidyiSBPUVMbkhKEsJMNVy
7UNqzFGgxEkzfzWLdTSKabDsUtTcSDs
9msAkRIqstFcQAdfpvFZqgWGPBbReNS
```

We can continue doing this as much as we like, using all of our normal CLI tools, e.g. `:%! grep "sekkrit"`

```text
IqHRWKquwsekkritCxsnInXbsPeLvx
```

And, the best part, because we're in vim, we can undo and redo all of the commands that we run, retry failed commands (remember to add those pesky quotes around spaces for grep!!), search and replace, and the list goes on. You're only limited by your imagination and the tools you have available on the CLI.

So, now you also know. Go spread the word!
