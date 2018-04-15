#!/usr/bin/osascript

if application "Spotify" is running
  tell application "Spotify"
    next track
  end tell
end if
