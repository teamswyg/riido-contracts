import path from "node:path";
import { fileURLToPath } from "node:url";

const here = path.dirname(fileURLToPath(import.meta.url));
export const root = path.resolve(here, "../..");
export const goldenDir = path.join(root, "golden");
export const currentDir = path.join(root, "test-results/golden-current");
export const diffDir = path.join(root, "test-results/golden-diff");

export function imagePath(dir, scenario) {
  return path.join(dir, `${scenario.name}.png`);
}
