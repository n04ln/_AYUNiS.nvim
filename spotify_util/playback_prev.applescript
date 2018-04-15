#!/usr/bin/osascript

if application "Spotify" is running
  tell application "Spotify"
    previous track
  end tell
end if
