import { defineConfig } from "vite";

const allowedHosts = new Set([
  "development.ai-api.riido.io",
  "testnet.ai-api.riido.io",
  "api.riido.io",
]);

export default defineConfig({
  plugins: [riidoSaasProxy()],
});

function riidoSaasProxy() {
  return {
    name: "riido-saas-proxy",
    configureServer(server) {
      server.middlewares.use("/__riido_saas_proxy", async (req, res) => {
        const target = parseTarget(req.url);
        if (!target) return send(res, 400, "invalid target");
        const init = { method: req.method, headers: forwardHeaders(req.headers) };
        const body = await readBody(req);
        if (body.length > 0) init.body = body;
        const upstream = await fetch(target, init);
        res.statusCode = upstream.status;
        res.setHeader("content-type", upstream.headers.get("content-type") ?? "text/plain");
        res.end(Buffer.from(await upstream.arrayBuffer()));
      });
    },
  };
}

function parseTarget(rawURL = "") {
  try {
    const requestURL = new URL(rawURL, "http://localhost");
    const target = new URL(requestURL.searchParams.get("url") ?? "");
    if (target.protocol !== "https:" || !allowedHosts.has(target.host)) return null;
    return target;
  } catch {
    return null;
  }
}

function forwardHeaders(headers) {
  const out = {};
  for (const name of ["authorization", "content-type", "x-riido-ai-agent-token"]) {
    if (headers[name]) out[name] = headers[name];
  }
  return out;
}

function readBody(req) {
  return new Promise((resolve, reject) => {
    const chunks = [];
    req.on("data", (chunk) => chunks.push(chunk));
    req.on("end", () => resolve(Buffer.concat(chunks)));
    req.on("error", reject);
  });
}

function send(res, status, body) {
  res.statusCode = status;
  res.end(body);
}
