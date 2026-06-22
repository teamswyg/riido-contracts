import { spawn } from "node:child_process";

export async function startServer() {
  const port = 5174;
  const child = spawn("npm", ["run", "dev", "--", "--host", "127.0.0.1", "--port", String(port)], {
    detached: process.platform !== "win32",
    stdio: ["ignore", "pipe", "pipe"],
  });
  child.stdout.on("data", (chunk) => process.stdout.write(chunk));
  child.stderr.on("data", (chunk) => process.stderr.write(chunk));
  const url = `http://127.0.0.1:${port}`;
  await waitForServer(url);
  return { url, stop: () => stopServer(child) };
}

async function waitForServer(url) {
  const deadline = Date.now() + 15000;
  while (Date.now() < deadline) {
    try {
      const response = await fetch(url);
      if (response.ok) return;
    } catch {
      await delay(250);
    }
  }
  throw new Error(`contract UI server did not start: ${url}`);
}

function delay(ms) {
  return new Promise((resolve) => setTimeout(resolve, ms));
}

async function stopServer(child) {
  if (child.exitCode !== null || child.killed) return;
  const terminated = waitForExit(child, 5000);
  terminate(child, "SIGTERM");
  if (!(await terminated)) {
    const killed = waitForExit(child, 5000);
    terminate(child, "SIGKILL");
    await killed;
  }
}

function terminate(child, signal) {
  try {
    if (process.platform === "win32") {
      child.kill(signal);
    } else {
      process.kill(-child.pid, signal);
    }
  } catch (error) {
    if (error.code !== "ESRCH") throw error;
  }
}

function waitForExit(child, timeoutMs) {
  if (child.exitCode !== null) return Promise.resolve(true);
  return new Promise((resolve) => {
    const timeout = setTimeout(() => done(false), timeoutMs);
    const onExit = () => done(true);
    child.once("exit", onExit);
    function done(exited) {
      clearTimeout(timeout);
      child.off("exit", onExit);
      resolve(exited);
    }
  });
}
