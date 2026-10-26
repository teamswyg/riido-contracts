import test from "node:test";
import assert from "node:assert/strict";
import { queryParamDefaults, requestURL } from "./requestModel.js";

test("read pagination survives URL and local proxy encoding", () => {
  const operation = { query_parameters: [{ name: "page", default: 1 }] };
  assert.deepEqual(queryParamDefaults(operation), { page: 1 });
  globalThis.window = { location: { hostname: "example.test" } };
  const url = requestURL("https://api.example.test", "/runs/{run_id}/reads", { run_id: "a/b" }, { page: "2", omitted: "" });
  assert.equal(url, "https://api.example.test/runs/a%2Fb/reads?page=2");
  window.location.hostname = "localhost";
  const proxy = requestURL("https://api.example.test", "/runs/{run_id}/reads", { run_id: "a/b" }, { page: "2" });
  assert.equal(new URL(proxy, "http://localhost").searchParams.get("url"), url);
  delete globalThis.window;
});
