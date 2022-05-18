# Go Workspace

## Overview 

This is an example repo to demonstrate the new workspace feature of Go 1.18.

## Workspace

Workspace is a feature which allows you to work with multiple modules in one workspace.   
It uses the `go.work` file to include multiple modules into one workspace.  
Here in this repo, we have two modules namely `sample` and `equation`.  

The `sample` module uses the `equation` module. Because we are using workspace feature, 
when the `sample` module resolves its dependency, it uses the one from `./equation` folder 
instead of the remote one. This allows us to develop multiple modules with ease.

Without the workspace feature, we have to push changes to `equation` module (stored in a separate repo) and then 
fetch the changes to the `sample` module, which is a length process.


