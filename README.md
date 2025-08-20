```
                __       __          __
   ____ _____  / /______/ /_  ____ _/ /
  / __ `/ __ \/ __/ ___/ __ \/ __ `/ / 
 / /_/ / /_/ / /_/ /__/ / / / /_/ /_/  
 \__, /\____/\__/\___/_/ /_/\__,_(_)   
/____/
```

**Real** simple system fetch with the looks..? (is this a hyperland reference??)

![sample](/assets/nixos.png)

<details>
<summary><b>More screenshots</b></summary>
<img src="/assets/debian.png" alt="debian demo" />
<img src="/assets/arch.png" alt="arch demo" />
<img src="/assets/void.png" alt="void demo" />
<img src="/assets/gentoo.png" alt="gentoo demo" />
<img src="/assets/bazzite.png" alt="bazzite demo" />
</details>

### Features

1. 0 external dependencies.
2. Dedicated to public domain. see [LICENSE](LICENSE).
3. Balanced between look and performance.
4. `NO_COLOR` spec compliant.
6. Easy to use.
7. UNIX only.
8. Unit colors change depending on the usage. (e.g: if you are using too much ram, it will appear **red**)
9. Customizable.

### Why

I've seen some "minimal" system fetch with a boatload of deps
just for displaying basic system info, ridiculous.

So i decided to make gotcha (and [novofetch](https://github.com/yehorovye/novofetch), another minimal
system fetch in V).

### Valid env variables

* `NO_COLOR` - Returns the output with absolutely 0 ansi colors. (default: 0)

### Config variables

* `DISABLE` - Fields to disable, separated with commas. (default: nil)
* `DIVIDER` - string to use as divider between user and data. (default: -)
* MOUNTS

By default, gotcha searches for a file called "config" on the same dir as the binary.
alternatively you can pass flag "config" with the dir of the config to override it.

# Installing

This program isn't officially uploaded as a package anywhere,
instead, i encourage you to build it yourself or using
[the latest binary](https://github.com/yehorovye/gotcha/releases).

You are free to distribute this as a package.

# License

public domain, do whatever you want.
