# path_fixer

A simple utility to rearrange segments of a unix PATH environment variable
based on predefined priorities. These priorities have been defined in the
tools source:

1. Paths prefixed `/User/` or `/home/`
2. Any unranked paths
3. Paths prefixed `/bin/` or `/sbin/`
4. Paths prefixed `/usr/local/`
5. Paths prefixed `/usr/bin/` or `/usr/sbin/`

This is written to solve the specific problem where a global `/etc/zprofile`
adds a path that takes higher priority than one defined in `~/.zshenv` as the
latter loads first.
