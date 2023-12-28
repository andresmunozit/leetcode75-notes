# Leetcode 75 Notes
Exercises taken from the Leetcode 75 interview problem compilation. Each exercise folder has the
following files:
```sh
$ tree 283-move-zeroes 
283-move-zeroes     # Exercise name
├── index.js        # Solution implementation
├── package.json    # Contains execution scripts
├── README.md       # Problem description
└── test.js         # Test cases implementation

```

## `package.json` scripts
You can use the following scripts on each exercise:
```json
{
  // ...
  "scripts": {
    "start": "node index.js",               // Starts the program
    "debug": "node --inspect-brk index.js", // Starts the program in debug mode
    "test": "node test.js"                  // Run the test suite for this exercise
  },
  // ...
}

```
