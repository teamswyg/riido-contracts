import { chromium } from "playwright";
import { imagePath } from "./paths.mjs";

export async function captureAll(url, scenarios, currentDir) {
  const browser = await chromium.launch();
  try {
    const page = await browser.newPage({ viewport: { width: 1440, height: 1100 } });
    await page.emulateMedia({ reducedMotion: "reduce" });
    for (const scenario of scenarios) {
      await captureScenario(page, url, scenario, imagePath(currentDir, scenario));
    }
  } finally {
    await browser.close();
  }
}

async function captureScenario(page, url, scenario, outPath) {
  await page.goto(url, { waitUntil: "networkidle" });
  if (scenario.query) {
    await page.getByLabel("Search").fill(scenario.query);
  }
  await page.screenshot({ fullPage: true, path: outPath });
}
