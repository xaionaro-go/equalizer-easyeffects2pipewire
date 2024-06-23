Go to https://autoeq.app/, select the headphones you use, choose "EasyEffects" as the "equalizer app", download the file, and then run something like:
```
go install github.com/xaionaro-go/equalizer-easyeffects2pipewire/cmd/eq-ee2pw@latest
eq-ee2pw ~/Downloads/my-easyeffects-preset.txt /home/xaionaro/.config/pipewire/pipewire.conf.d/equalizer.conf
systemctl restart --user pipewire
```

That's it.
