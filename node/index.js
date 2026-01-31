
let coldStart = true;
const initTime = Date.now();

exports.handler = async () => {
  const start = Date.now();

  console.log("Language=Node.js");
  console.log(`ColdStart=${coldStart}`);
  console.log(`InitDurationMs=${Date.now() - initTime}`);

  coldStart = false;

  console.log(`ExecutionDurationMs=${Date.now() - start}`);
};
