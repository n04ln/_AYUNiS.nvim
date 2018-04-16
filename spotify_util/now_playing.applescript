#!/usr/bin/osascript

if application "Spotify" is running
  tell application "Spotify"
    if (player state as string) is "playing" then
        set nowPlaying to "♫ "
    else
        set nowPlaying to ""
    end if
    set currentArtist to artist of current track as string
    set currentTrack to name of current track as string
    return nowPlaying & currentTrack & " - " & currentArtist
  end tell
else
  return "no music"
end if
