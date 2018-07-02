#!/usr/bin/osascript

if application "Spotify" is running
  tell application "Spotify"
    if (player state as string) is "playing" then
        set nowPlaying to "🎧"
    else
        set nowPlaying to ""
    end if
    if repeating then
        set isRepeat to "🔃"
    else
        set isRepeat to ""
    end if
    if shuffling then
        set isShuffle to "🔀"
    else
        set isShuffle to ""
    end if
    set currentArtist to artist of current track as string
    set currentTrack to name of current track as string
    return nowPlaying & currentTrack & " - " & currentArtist & " " & isRepeat & " " & isShuffle
  end tell
else
  return "no music"
end if
