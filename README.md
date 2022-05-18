# Idea-CLI

## Intro
This is a CLI application built with go using the flag package. The CLI allows the user to read, create, edit, and delete entries in a list using
subcommands. 

## Subcommands

#### view 
    goIdeas view
Allows the user to view the current list.

#### new
    goIdeas new -I="Some Text"
Allows the user to add a new entry to the list. The new list item will be added to the list and the new list displayed.

#### edit
    goIdeas edit -i={Integer} -t="New Text"
Allows the user to edit an exiting entry in the list. The user specifies the index of the list entry and the new text. The newly edited list will then be 
displayed.

#### del
    goIdeas del -i={Integer}
Allows the user to delete an entry from the list. The indexs of the list entries will then be updated and the list will be displayed.

## Installation
To use the Idea-CLI you must first have golang installed on your system. Follow the steps at the [Golang website](https://go.dev/doc/install).

Clone this repo to your system.

Once go is installed run the following command.
    
    go build goIdeas.go
    
This will build the executable and you are now free to use the Idea-CLI.

