let coldStart = true;
const initTime = Date.now();

exports.handler = async (event) => {
  const start = Date.now();

  // Print full SNS event
  console.log("SNS Event:");
  console.log(JSON.stringify(event, null, 2));

  // Print message(s)
  if (event.Records) {
    for (const record of event.Records) {
      console.log("SNS Message:", record.Sns?.Message);
    }
  }

  console.log("Language=Node.js");
  console.log(`ColdStart=${coldStart}`);
  console.log(`InitDurationMs=${Date.now() - initTime}`);

  coldStart = false;

  console.log(`ExecutionDurationMs=${Date.now() - start}`);
};
