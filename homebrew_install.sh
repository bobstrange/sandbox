#!/bin/bash

set -e

# Install Xcode Command Line Tools
xcode-select --install

# Install HomeBrew
ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"

brew doctor
brew update
brew upgrade

# Install Brew Cask
brew tap phinze/cask
brew install brew-cask

# iTerm http://machacks.whiskykilo.com/apps/iterm.html
# iTerm2 Prompt on Quit
defaults write com.googlecode.iterm2 PromptOnQuit -bool false
