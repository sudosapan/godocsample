# Go Documentation Samples

This repository contains examples for documenting go code and result of what it ends up in the go doc.


## Basic
### New Line
Give one blank comment before the text to go to new line
#### Comment
```go
// These will be 
// shown in one line
//
// This starts in new line
```

#### Go doc output:
![newline](/assets/newline.png)

### Heading
Text on newline starting with capital letter and not ending in full stop is rendered as a heading.

#### Comment
```go
//
// Heading
//
// This is under the heading.
```

#### Go doc output:
![heading](/assets/heading.png)


## Package Documentation

Comments preceding the package directive in the go file go to the package documentation. 

There should be (vertical) gap between the comment and package directive.



