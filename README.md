# netdata-desktop

This is a thin desktop wrapper for [netdata](https://github.com/netdata/netdata). It utilizes [webview](https://github.com/zserge/webview) library to manipulate and render the website available on address [127.0.0.1:19999](http://127.0.0.1:19999) by default.

## Quickstart

Make sure you have [netdata](https://github.com/netdata/netdata) installed first. 
`netdata-desktop` can be easily compiled using the Makefile (or manually) with Go 1.11.4 and later. 

## Dependencies

There is only a single dependency to run this application, I've used [webview](https://github.com/zserge/webview) to browse the website served by `netdata`.

## Config

While the application itself is just a thin wrapper around a website, you're still able to configure the connection string by providing it over command line:

```sh
netdata-desktop --conn=http://127.0.0.1:19999/
```

## Alternatives

Google Chrome (and derivatives) provides this exact functionality simply by navigating to **More Tools** > **Create Shortcut...** in menu. This opens up a modal window asking for a shortcut name, but the important part here is: **Open as window**. This will effectively replicate the behavior this app was aimed for.

## Copying

Read COPYING.md for licensing information.
