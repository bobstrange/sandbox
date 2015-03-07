# AppStore http://machacks.whiskykilo.com/hacks/app-store.html

# To list all available updates
sudo softwareupdate -l

# To download and install all updates available for my osx
sudo softwareupdate -i -a

# Set my timezone
sudo systemsetup -settimezone Asia/Tokyo

# OSX http://machacks.whiskykilo.com/hacks/os-x.html

# Finder http://machacks.whiskykilo.com/hacks/finder.html
# Enable Tap-to-Click
defaults write com.apple.driver.AppleBluetoothMultitouch.trackpad Clicking -bool true
defaults -currentHost write NSGlobalDomain com.apple.mouse.tapBehavior -int 1
defaults write NSGlobalDomain com.apple.mouse.tapBehavior -int 1

# Stop iTunes from Popping Up with the Media Keys
launchctl unload -w /System/Library/LaunchAgents/com.apple.rcd.plist 2> /dev/null

# Change Dock Tile Size
defaults write com.apple.dock tilesize -integer 24

# Kill the Dashboard
defaults write com.apple.dashboard mcx-disabled -boolean YES
killall Dock

# Don't Hide ~/Library
chflags nohidden ~/Library

# Show all Filename Extensions
defaults write NSGlobalDomain AppleShowAllExtensions -bool true

# Show Path Bar
defaults write com.apple.finder ShowPathbar -bool true

# Search Current Folder by Default
defaults write com.apple.finder FXDefaultSearchScope -string "SCcf"

# Disable Empty Trash Warning
defaults write com.apple.finder WarnOnEmptyTrash -bool false

# Show all Hard Drives, Servers, and Removable Media
defaults write com.apple.finder ShowExternalHardDrivesOnDesktop -bool true
defaults write com.apple.finder ShowHardDrivesOnDesktop -bool true
defaults write com.apple.finder ShowMountedServersOnDesktop -bool true
defaults write com.apple.finder ShowRemovableMediaOnDesktop -bool true

# Prevent Writing .DS_Store Files on Network Drives
defaults write com.apple.desktopservices DSDontWriteNetworkStores

# Security http://machacks.whiskykilo.com/hacks/security.html
defaults write com.apple.screensaver askForPassword -int 1
defaults write com.apple.screensaver askForPasswordDelay -int 0

# Notification Center http://machacks.whiskykilo.com/hacks/notification-center.html
# Disable Notification Center
launchctl unload -w /System/Library/LaunchAgents/com.apple.notificationcenterui.plist