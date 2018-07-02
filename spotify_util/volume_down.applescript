#!/usr/bin/osascript

on max(x, y)
  if x ≥ y then
    return x
  end if
  return y
end max

tell application "Spotify" to set sound volume to (my max(sound volume - 10, 0))
