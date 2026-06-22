import fs from "node:fs/promises";
import { captureAll } from "./golden/capture.mjs";
import { compareAll } from "./golden/compare.mjs";
import { scenarios } from "./golden/scenarios.mjs";
import { startServer } from "./golden/server.mjs";
import { currentDir, diffDir, goldenDir, imagePath } from "./golden/paths.mjs";

const update = process.argv.includes("--update");

await prepareDirs();
const server = await startServer();
try {
  await captureAll(server.url, scenarios, currentDir);
  if (update) {
    await Promise.all(scenarios.map((scenario) => fs.copyFile(imagePath(currentDir, scenario), imagePath(goldenDir, scenario))));
    console.log(`updated ${scenarios.length} golden screenshot(s)`);
  } else {
    await compareAll(scenarios, currentDir);
  }
} finally {
  await server.stop();
}

async function prepareDirs() {
  await fs.mkdir(goldenDir, { recursive: true });
  await fs.rm(currentDir, { recursive: true, force: true });
  await fs.rm(diffDir, { recursive: true, force: true });
  await fs.mkdir(currentDir, { recursive: true });
  await fs.mkdir(diffDir, { recursive: true });
}
