# Usage

The indicator uses [wttr API](https://github.com/chubin/wttr.in)

By default used ip location. Set the environment variable `CITY_WEATHER` to override.

## Signs

`r` - rain

`s` - snow

## Sample

Updated every 30 minutes.

```
"custom/weather": {
    "max-length": 15,
    "return-type": "json",
    "format": "WX:{}",
    "exec": "$HOME/.config/waybar/scripts/weather",
    "interval": 1800
}
```

# Preview

![preview](/assets/preview.png)
