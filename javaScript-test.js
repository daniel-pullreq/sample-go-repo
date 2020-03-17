
const fs = require('fs');
const readline = require('readline');

console.log('helloooooo');

var x = 12;

if (x===-0){
  console.log("Hello");
}

while(true){
}


console.log("A");
debugger;

let foo = function() {
  try {
    return 1;
  } catch (err) {
    return 2;
  } finally {
    return 3;
  }
};

try {
  doSomethingThatMightThrow();
} catch (e) {
  throw e;
}

try {
  doSomethingThatMightThrow();
} catch (e) {
  throw e;
} finally {
  cleanUp();
}

var roo = function rar() {};
roo = function rar() {};
obj.roo = function rar() {};
obj["roo"] = function rar() {};
var obj = { roo: function rar() {} };
({ ["roo"]: function rar() {} });

let newVar = Array(0, 1, 2);
