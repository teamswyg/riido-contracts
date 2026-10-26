import { useEffect, useState } from "react";
import { jsonText } from "../contractModel.js";
import { authLabel, buildHeaders, pathParamDefaults, queryParamDefaults, requestURL, shouldSendBody } from "../requestModel.js";
import { sampleForRef } from "../schemaSample.js";

export default function LiveRequestPanel({ operation }) {
  const [baseURL, setBaseURL] = useState("https://development.ai-api.riido.io");
  const [token, setToken] = useState("");
  const [params, setParams] = useState(pathParamDefaults(operation));
  const [queryParams, setQueryParams] = useState(queryParamDefaults(operation));
  const [body, setBody] = useState("{}");
  const [result, setResult] = useState("");

  useEffect(() => {
    setParams(pathParamDefaults(operation));
    setQueryParams(queryParamDefaults(operation));
    setBody(jsonText(sampleForRef(operation.request?.ref, operation.schemas)));
    setResult("");
  }, [operation]);

  async function send() {
    const url = requestURL(baseURL, operation.path, params, queryParams);
    const init = { method: operation.method, headers: buildHeaders(operation, token) };
    if (shouldSendBody(operation)) init.body = body;
    const response = await fetch(url, init);
    const text = await response.text();
    setResult(`${response.status} ${response.statusText}\n${text}`);
  }

  return (
    <section className="live-panel">
      <h3>Live SaaS request</h3>
      <div className="env-row">
        <button onClick={() => setBaseURL("https://development.ai-api.riido.io")}>development</button>
        <button onClick={() => setBaseURL("https://testnet.ai-api.riido.io")}>testnet</button>
        <button onClick={() => setBaseURL("https://api.riido.io")}>production</button>
      </div>
      <p className="proxy-note">localhost uses a Riido-only dev proxy, so requests still hit the selected SaaS host.</p>
      <div className="form-row"><label>Base URL<input value={baseURL} onChange={(e) => setBaseURL(e.target.value)} /></label></div>
      <div className="form-row"><label>{authLabel(operation)}<input value={token} onChange={(e) => setToken(e.target.value)} type="password" /></label></div>
      {operation.path_params?.map((name) => (
        <div className="form-row" key={name}><label>{name}<input value={params[name] ?? ""} onChange={(e) => setParams({ ...params, [name]: e.target.value })} /></label></div>
      ))}
      {operation.query_parameters?.map(({ name, type, minimum }) => (
        <div className="form-row" key={name}><label>{name}<input type={type === "integer" ? "number" : "text"} min={minimum} step={type === "integer" ? 1 : undefined} value={queryParams[name] ?? ""} onChange={(e) => setQueryParams({ ...queryParams, [name]: e.target.value })} /></label></div>
      ))}
      {shouldSendBody(operation) && <textarea value={body} onChange={(e) => setBody(e.target.value)} />}
      <button className="send" onClick={send}>Send real request</button>
      {result && <pre className="result">{result}</pre>}
    </section>
  );
}
