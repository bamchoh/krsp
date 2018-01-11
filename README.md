# krsp
krsp is an application to kill an application with the passed pid

# usage
krsp just kill an application with passing pid

```
> krsp <pid>
```

Here is example. krsp is useful by using `bamchoh/kick` together.
```
> kick ruby -e "loop { print 1 }"
4644
> krsp 4644

```

I checked only on windows. sorry.

# why don't we use `taskkill`?

`taskkill` is slower than krsp for me.

# license

MIT

# author

Yoshihiko Yamanaka (a.k.a bamchoh)
