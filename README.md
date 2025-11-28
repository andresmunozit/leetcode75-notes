# Leetcode 75 Notes
Exercises taken from the Leetcode 75 interview problem compilation. Each exercise folder has the
following files:
```sh
$ tree ,<exercise-name> 
,<exercise-name>
├── index.js        # Solution implementation
├── package.json    # Contains execution scripts
├── README.md       # Problem description
└── test.js         # Test cases implementation

```

## Usage
To streamline the setup process for each exercise, you can automatically inject the common scripts
defined in `leetcode-node-scripts.json` into the `package.json` of each exercise directory using the
`add-scripts.sh` bash script.

### Injecting Scripts Automatically
1. Ensure you have a `package.json` file in your exercise directory (create one with `npm init -y`
if necessary).
2. Use the `add-scripts.sh` script to inject the scripts from `leetcode-node-scripts.json` into your
`package.json`:

```sh
# Execute from the repository root
$ ./add-scripts.sh <path-to-exercise>/package.json

```

After running the script, your package.json will include the following scripts from 
`leetcode-node-scripts.json`:
```json
{
  "start": "node index.js",
  "debug": "node --inspect-brk index.js",
  "test": "node test.js"
}

```

### Understanding the Scripts
- `start`: Runs the `index.js` file to start the program. Use this when you want to run you
 solution.
- `debug`: Starts the program in debug mode with `--inspect-brk`, allowing you to attach a debugger
and step through the code.
- `test`: Runs the `test.js` file to execute all test cases associated with the exercise, ensuring
your solution meets the specified requirements.

## Support for Golang
When implementing a golang solution the project tree structure changes:
```
000-example-exercise-name
|---js
    |--- (js structure defined above)
|---go
    |--- solution.go
    |--- solution.test.go
    
```

## Knowledge Base
There is a directory called knowledge base, this will contain a structured version of relevant
pieces of knowledge learn during the problem resolution, for example binary search.

Inside of the `kb` directory, there is a `README.md` file that is an index to other md files with
specific pieces of information.
