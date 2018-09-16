#!/usr/bin/osascript

if application "Spotify" is running
  tell application "Spotify"
    if (player state as string) is "playing" then
        set nowPlaying to "▶"
    else
        set nowPlaying to ""
    end if
    if repeating then
        set isRepeat to "[R]"
    else
        set isRepeat to ""
    end if
    if shuffling then
        set isShuffle to "[S]"
    else
        set isShuffle to ""
    end if
    set currentArtist to artist of current track as string
    set currentTrack to name of current track as string
    return currentArtist & currentTrack & " - " & " " & nowPlaying & isRepeat & " " & isShuffle
  end tell
else
  return "💤"
end if
