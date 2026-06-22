import fs from "node:fs/promises";
import pixelmatch from "pixelmatch";
import { PNG } from "pngjs";
import { diffDir, goldenDir, imagePath } from "./paths.mjs";

const maxDiffRatio = 0.05;

export async function compareAll(scenarios, currentDir) {
  const failures = [];
  for (const scenario of scenarios) {
    const result = await compareScenario(scenario, currentDir);
    console.log(`${result.ok ? "ok" : "diff"} ${scenario.name} ${result.ratio.toFixed(4)}`);
    if (!result.ok) failures.push(result);
  }
  if (failures.length > 0) throw new Error(`${failures.length} golden screenshot(s) changed`);
}

async function compareScenario(scenario, currentDir) {
  const goldenPath = imagePath(goldenDir, scenario);
  const currentPath = imagePath(currentDir, scenario);
  const [golden, current] = normalizePair(
    PNG.sync.read(await fs.readFile(goldenPath)),
    PNG.sync.read(await fs.readFile(currentPath)),
  );
  const diff = new PNG({ width: golden.width, height: golden.height });
  const changed = pixelmatch(golden.data, current.data, diff.data, golden.width, golden.height, { threshold: 0.12 });
  const ratio = changed / (golden.width * golden.height);
  if (ratio > 0) await fs.writeFile(imagePath(diffDir, scenario), PNG.sync.write(diff));
  return { ok: ratio <= maxDiffRatio, ratio };
}

function normalizePair(a, b) {
  const width = Math.max(a.width, b.width);
  const height = Math.max(a.height, b.height);
  return [padImage(a, width, height), padImage(b, width, height)];
}

function padImage(source, width, height) {
  if (source.width === width && source.height === height) return source;
  const out = new PNG({ width, height });
  for (let i = 0; i < out.data.length; i += 4) {
    out.data[i] = 255;
    out.data[i + 1] = 255;
    out.data[i + 2] = 255;
    out.data[i + 3] = 255;
  }
  PNG.bitblt(source, out, 0, 0, source.width, source.height, 0, 0);
  return out;
}
