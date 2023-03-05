# Sample

Updated every 30 minutes. By default used ip location, set CITY_WEATHER os env or city's variable to override. CITY_WEATHER has the highest priority.

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
