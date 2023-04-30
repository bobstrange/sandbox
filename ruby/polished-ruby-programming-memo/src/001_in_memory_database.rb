# frozen_string_literal: true

album_infos = 100.times.flat_map do |i|
  10.times.map do |j|
    # Album name, track number, artist name
    ["Album #{i}", j, "Artist #{j}"]
  end
end

# Pattern 1
# store album artists and album track artists separately
# Memory usage: not good Search speed: good
album_artists = {}
album_track_artists = {}
album_infos.each do |album, track, artist|
  (album_artists[album] ||= []) << artist
  (album_track_artists[[album, track]] ||= []) << artist
end
album_artists.each_value(&:uniq!)

lookup = lambda do |album, track = nil|
  if track
    album_track_artists[[album, track]]
  else
    album_artists[album]
  end
end

puts album_artists.size

# Pattern 2
# store album artists and album track artists together
# Memory usage: good
# Search speed: not good. If you want to search by only album name, we have to
# generate artist list for all tracks in the album.
albums = {}
album_infos.each do |album, track, artist|
  ((albums[album] ||= {})[track] ||= []) << artist
end

lookup_album = lambda do |album, track = nil|
  if track
    albums.dig(album, track)
  else
    albums[album].map do |_, artist|
      artist
    end.uniq
  end
end
