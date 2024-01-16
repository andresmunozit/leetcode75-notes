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
// leetcode-node-scripts.json
{
  "start": "node index.js",
  "debug": "node --inspect-brk index.js",
  "test": "node test.js"
}

```

Where:
- `start`: Starts the program
- `debug`: Starts the program in debug mode
- `test`: Run the test suite for this exercise

### Populating Automatically the `scripts` Key of `package.json` files
Execute the following script, in this example the `package.json` file has been created beforehand
inside the directory `./1004-max-consecutive-ones-iii` using `npm init -y`:
```sh
# Execute from the repository root
$ ./add-scripts.sh ./1004-max-consecutive-ones-iii/package.json

```
