export function pathParamDefaults(operation) {
  return Object.fromEntries((operation.path_params ?? []).map((name) => [name, ""]));
}

export function resolvePath(path, params) {
  return path.replace(/\{([^}]+)\}/g, (_, name) => encodeURIComponent(params[name] ?? ""));
}

export function shouldSendBody(operation) {
  return Boolean(operation.request) && !["GET", "HEAD"].includes(operation.method);
}

export function buildHeaders(operation, token) {
  const headers = {};
  if (shouldSendBody(operation)) headers["Content-Type"] = "application/json";
  if (!token.trim()) return headers;
  if (operation.auth?.scheme === "bearer") headers.Authorization = `Bearer ${token.trim()}`;
  if (operation.auth?.header) headers[operation.auth.header] = token.trim();
  return headers;
}

export function authLabel(operation) {
  if (operation.auth?.scheme === "bearer") return "Bearer token";
  return operation.auth?.header ?? "Auth token";
}

export function requestURL(baseURL, path, params) {
  const target = baseURL.replace(/\/$/, "") + resolvePath(path, params);
  if (!["localhost", "127.0.0.1"].includes(window.location.hostname)) return target;
  return `/__riido_saas_proxy?url=${encodeURIComponent(target)}`;
}
